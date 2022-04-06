package stream

import "github.com/non1996/go-jsonobj/function"

// slice迭代器
type sliceIterator[T any] struct {
	s   []T
	idx int
}

func (s *sliceIterator[T]) Reset() {
	s.idx = 0
}

func (s *sliceIterator[T]) TryAdvance(predicate function.Predicate[T]) bool {
	if s.idx >= len(s.s) {
		return false
	}
	next := predicate(s.s[s.idx])
	s.idx++
	return next
}
