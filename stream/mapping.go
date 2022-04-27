package stream

import (
	"github.com/non1996/go-jsonobj/function"
)

type mapIterator[IN, OUT any] struct {
	upstream Iterator[IN]
	stages   stage2[IN, OUT]
}

func (i *mapIterator[IN, OUT]) TryAdvance(predicate function.Predicate[OUT]) bool {
	return i.upstream.TryAdvance(func(in IN) bool {
		out, nextAction, nextElem := i.stages(in)
		if !nextAction {
			return nextElem
		}
		return nextElem && predicate(out)
	})
}
