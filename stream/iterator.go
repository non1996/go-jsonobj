package stream

import "github.com/non1996/go-jsonobj/function"

// slice迭代器
type sliceIterator[T any] struct {
	s   []T
	idx int
}

func (s *sliceIterator[T]) TryAdvance(predicate function.Predicate[T]) bool {
	if s.idx >= len(s.s) {
		return false
	}
	next := predicate(s.s[s.idx])
	s.idx++
	return next
}

type mappingIterator[IN, OUT any] struct {
	upstream Iterator[IN]
	stages   stage2[IN, OUT]
}

func (i *mappingIterator[IN, OUT]) TryAdvance(predicate function.Predicate[OUT]) bool {
	return i.upstream.TryAdvance(func(in IN) bool {
		out, nextAction, nextElem := i.stages(in)
		if !nextAction {
			return nextElem
		}
		return nextElem && predicate(out)
	})
}
