package stream

import (
	"github.com/non1996/go-jsonobj/function"
	"github.com/non1996/go-jsonobj/optional"
)

type stage[T any] func(T) (v T, nextAction bool, nextElem bool)

type Stream[T any] interface {
	Filter(function.Predicate[T]) Stream[T]
	Peek(function.Consumer[T]) Stream[T]
	Map(function.Operation[T]) Stream[T]
	Skip(int) Stream[T]
	Limit(int) Stream[T]
	Sorted(function.Comparator[T]) Stream[T]
	ToList() []T
	Foreach(function.Consumer[T])
	Count() int
	AnyMatch(function.Predicate[T]) bool
	AllMatch(function.Predicate[T]) bool
	NoneMatch(function.Predicate[T]) bool
	Find(function.Predicate[T]) optional.Optional[T]
	Reduce(T, function.BiOperation[T]) optional.Optional[T]
}

type Iterator[T any] interface {
	Reset()
	TryAdvance(function.Predicate[T]) bool
}
