package stream

import (
	"github.com/non1996/go-jsonobj/function"
	"github.com/non1996/go-jsonobj/optional"
	"sort"
)

type pipeline[T any] struct {
	iter   Iterator[T]
	stages stage[T]
}

func newPipeline[T any](iter Iterator[T]) *pipeline[T] {
	return &pipeline[T]{
		iter: iter,
		stages: func(in T) (out T, nextAction bool, nextElem bool) {
			return in, true, true
		},
	}
}

func (p *pipeline[T]) Filter(predicate function.Predicate[T]) Stream[T] {
	return p.mergeStage(func(in T) (out T, nextAction bool, nextElem bool) {
		return in, predicate(in), true
	})
}

func (p *pipeline[T]) Peek(consumer function.Consumer[T]) Stream[T] {
	return p.mergeStage(func(in T) (out T, nextAction bool, nextElem bool) {
		consumer(in)
		return in, true, true
	})
}

func (p *pipeline[T]) Map(operation function.Operation[T]) Stream[T] {
	return p.mergeStage(func(in T) (out T, nextAction bool, nextElem bool) {
		return operation(in), true, true
	})
}

func (p *pipeline[T]) Skip(n int) Stream[T] {
	var count int
	return p.mergeStage(func(in T) (out T, nextAction bool, nextElem bool) {
		count++
		return in, count > n, true
	})
}

func (p *pipeline[T]) Limit(n int) Stream[T] {
	var count int
	return p.mergeStage(func(in T) (out T, nextAction bool, nextElem bool) {
		count++
		return in, true, count < n
	})
}

func (p *pipeline[T]) Sorted(comparator function.Comparator[T]) Stream[T] {
	arr := p.ToList()
	sort.Slice(arr, func(i, j int) bool { return comparator(arr[i], arr[j]) })
	return Slice(arr)
}

func (p *pipeline[T]) ToList() (res []T) {
	p.advanceEach(func(v T) bool {
		res = append(res, v)
		return true
	})
	return res
}

func (p *pipeline[T]) Foreach(consumer function.Consumer[T]) {
	p.advanceEach(func(v T) bool {
		consumer(v)
		return true
	})
}

func (p *pipeline[T]) Count() (count int) {
	p.advanceEach(func(T) bool {
		count++
		return true
	})
	return
}

func (p *pipeline[T]) AnyMatch(predicate function.Predicate[T]) (match bool) {
	p.advanceEach(func(v T) bool {
		match = predicate(v)
		return !match
	})
	return match
}

func (p *pipeline[T]) AllMatch(predicate function.Predicate[T]) (allMatch bool) {
	allMatch = true
	p.advanceEach(func(v T) bool {
		allMatch = allMatch && predicate(v)
		return allMatch
	})
	return allMatch
}

func (p *pipeline[T]) NoneMatch(predicate function.Predicate[T]) (nonMatch bool) {
	nonMatch = true
	p.advanceEach(func(v T) bool {
		if predicate(v) {
			nonMatch = false
			return false
		}
		return true
	})
	return nonMatch
}

func (p *pipeline[T]) Find(predicate function.Predicate[T]) (o optional.Optional[T]) {
	o = optional.Empty[T]()
	p.advanceEach(func(v T) bool {
		if predicate(v) {
			o.Set(v)
			return false
		}
		return true
	})
	return o
}

func (p *pipeline[T]) Reduce(identity T, operation function.BiOperation[T]) optional.Optional[T] {
	var do bool
	p.advanceEach(func(v T) bool {
		do = true
		identity = operation(identity, v)
		return true
	})
	if !do {
		return optional.Empty[T]()
	}
	return optional.New(identity)
}

func (p *pipeline[T]) advanceEach(advancer function.Predicate[T]) {
	p.mergeStage(func(in T) (out T, nextAction bool, nextElem bool) {
		return out, nextAction, advancer(in)
	})

	for p.iter.TryAdvance(func(v T) bool {
		_, _, nextElem := p.stages(v)
		return nextElem
	}) {
	}
}

func (p *pipeline[T]) mergeStage(downstream stage[T]) Stream[T] {
	upstream := p.stages
	p.stages = func(in T) (out T, nextAction bool, nextElem bool) {
		out, nextAction, nextElem = upstream(in)
		if !nextAction {
			return
		}
		out2, nextAction2, nextElem2 := downstream(out)
		return out2, nextAction2, nextElem && nextElem2
	}
	return p
}
