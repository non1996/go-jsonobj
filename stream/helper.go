package stream

import (
	"github.com/non1996/go-jsonobj/function"
	"github.com/non1996/go-jsonobj/optional"
)

func Slice[T any](s []T) Stream[T] {
	return newPipeline[T](&sliceIterator[T]{s: s})
}

func New[T any](iter Iterator[T]) Stream[T] {
	return newPipeline[T](iter)
}

func Map[T1 any, T2 any](s []T1, mapper function.Function[T1, T2]) []T2 {
	if len(s) == 0 {
		return nil
	}
	res := make([]T2, 0, len(s))
	for idx := range s {
		res = append(res, mapper(s[idx]))
	}
	return res
}

func CollectToMap[T any, K comparable, V any](s []T, keyMapper function.Function[T, K], valMapper function.Function[T, V]) map[K]V {
	res := map[K]V{}
	for idx := range s {
		res[keyMapper(s[idx])] = valMapper(s[idx])
	}
	return res
}

func CollectToMapDump[T any, K comparable, V any](
	s []T, keyMapper function.Function[T, K],
	valMapper function.Function[T, V],
) map[K][]V {
	res := map[K][]V{}
	for idx := range s {
		res[keyMapper(s[idx])] = append(res[keyMapper(s[idx])], valMapper(s[idx]))
	}
	return res
}

func Reduce[T any](identity T, s []T, operation function.BiOperation[T]) optional.Optional[T] {
	if len(s) == 0 {
		return optional.Empty[T]()
	}
	for idx := range s {
		identity = operation(identity, s[idx])
	}
	return optional.New(identity)
}

func FlatMap[T1 any, T2 any](s []T1, mapper function.Function[T1, []T2]) []T2 {
	if len(s) == 0 {
		return nil
	}
	res := make([]T2, 0)
	for idx := range s {
		res = append(res, mapper(s[idx])...)
	}
	return res
}

func MapWithError[T1, T2 any](list []T1, mapper func(T1) (T2, error)) (res []T2, err error) {
	if len(list) == 0 {
		return nil, nil
	}
	res = make([]T2, 0, len(list))
	for _, i := range list {
		r, err := mapper(i)
		if err != nil {
			return nil, err
		}
		res = append(res, r)
	}
	return
}

func MapS[T1, T2 any](src Stream[T1], mapper func(T1) T2) Stream[T2] {
	pipe := src.(*pipeline[T1])
	iter := &mappingIterator[T1, T2]{
		upstream: pipe.iter,
		stages: func(in T1) (out T2, nextAction bool, nextElem bool) {
			in, nextAction, nextElem = pipe.stages(in)
			if !nextAction {
				return
			}
			return mapper(in), true, nextElem
		},
	}
	return newPipeline[T2](iter)
}

func CollectToMapS[T any, K comparable, V any](
	s Stream[T],
	keyMapper function.Function[T, K],
	valMapper function.Function[T, V],
) map[K]V {
	res := map[K]V{}
	s.Foreach(func(v T) {
		res[keyMapper(v)] = valMapper(v)
	})
	return res
}
