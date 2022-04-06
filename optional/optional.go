package optional

import "github.com/non1996/go-jsonobj/function"

type Optional[T any] interface {
	Set(T)
	Get() T
	IsPresent() bool
	IfPresent(function.Consumer[T]) Optional[T]
	Else(func())
}

func New[T any](v T) Optional[T] {
	return &optionalImpl[T]{
		value: v,
		valid: true,
	}
}

func Empty[T any]() Optional[T] {
	return &optionalImpl[T]{}
}

type optionalImpl[T any] struct {
	value T
	valid bool
}

func (o *optionalImpl[T]) Set(v T) {
	o.value = v
	o.valid = true
}

func (o *optionalImpl[T]) Get() T {
	return o.value
}

func (o *optionalImpl[T]) IsPresent() bool {
	return o.valid
}

func (o *optionalImpl[T]) IfPresent(consumer function.Consumer[T]) Optional[T] {
	if o.valid {
		consumer(o.value)
	}
	return o
}

func (o *optionalImpl[T]) Else(action func()) {
	if !o.valid {
		action()
	}
}
