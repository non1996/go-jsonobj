package stream

import (
	"github.com/non1996/go-jsonobj/function"
	"github.com/non1996/go-jsonobj/optional"
	"sort"
)

type streamImpl[T any] struct {
	iter   Iterator[T]
	stages []stage[T]
}

func (s *streamImpl[T]) Filter(predicate function.Predicate[T]) Stream[T] {
	s.stages = append(s.stages, func(v T) (res T, nextAction bool, nextElem bool) {
		return v, predicate(v), true
	})
	return s
}

func (s *streamImpl[T]) Peek(consumer function.Consumer[T]) Stream[T] {
	s.stages = append(s.stages, func(v T) (res T, nextAction bool, nextElem bool) {
		consumer(v)
		return v, true, true
	})
	return s
}

func (s *streamImpl[T]) Map(operation function.Operation[T]) Stream[T] {
	s.stages = append(s.stages, func(v T) (res T, nextAction bool, nextElem bool) {
		return operation(v), true, true
	})
	return s
}

func (s *streamImpl[T]) Skip(n int) Stream[T] {
	var count int
	s.stages = append(s.stages, func(v T) (res T, nextAction bool, nextElem bool) {
		count++
		return v, count > n, true
	})
	return s
}

func (s *streamImpl[T]) Limit(n int) Stream[T] {
	var count int
	s.stages = append(s.stages, func(v T) (res T, nextAction bool, nextElem bool) {
		count++
		return v, true, count < n
	})
	return s
}

func (s *streamImpl[T]) Sorted(comparator function.Comparator[T]) Stream[T] {
	arr := s.ToList()
	sort.Slice(arr, func(i, j int) bool { return comparator(arr[i], arr[j]) })
	return Slice(arr)
}

func (s *streamImpl[T]) ToList() (res []T) {
	s.advanceEach(func(v T) bool {
		res = append(res, v)
		return true
	})
	return res
}

func (s *streamImpl[T]) Foreach(consumer function.Consumer[T]) {
	s.advanceEach(func(v T) bool {
		consumer(v)
		return true
	})
}

func (s *streamImpl[T]) Count() (count int) {
	s.advanceEach(func(T) bool {
		count++
		return true
	})
	return
}

func (s *streamImpl[T]) AnyMatch(predicate function.Predicate[T]) (match bool) {
	s.advanceEach(func(v T) bool {
		match = predicate(v)
		return !match
	})
	return match
}

func (s *streamImpl[T]) AllMatch(predicate function.Predicate[T]) (allMatch bool) {
	allMatch = true
	s.advanceEach(func(v T) bool {
		allMatch = allMatch && predicate(v)
		return allMatch
	})
	return allMatch
}

func (s *streamImpl[T]) NoneMatch(predicate function.Predicate[T]) (nonMatch bool) {
	nonMatch = true
	s.advanceEach(func(v T) bool {
		if predicate(v) {
			nonMatch = false
			return false
		}
		return true
	})
	return nonMatch
}

func (s *streamImpl[T]) Find(predicate function.Predicate[T]) (o optional.Optional[T]) {
	o = optional.Empty[T]()
	s.advanceEach(func(v T) bool {
		if predicate(v) {
			o.Set(v)
			return false
		}
		return true
	})
	return o
}

func (s *streamImpl[T]) advanceEach(advancer function.Predicate[T]) {
	s.iter.Reset()
	for s.iter.TryAdvance(func(v T) bool {
		var nextAction bool
		var nextElem = true
		for _, stage := range s.stages {
			var tmp bool
			v, nextAction, tmp = stage(v)
			nextElem = nextElem && tmp
			if !nextAction {
				return nextElem
			}
		}
		return advancer(v) && nextElem
	}) {
	}
}
