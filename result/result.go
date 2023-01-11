package result

import (
	"fmt"

	"github.com/non1996/go-jsonobj/function"
)

type Result[T any] struct {
	val T
	err error
}

func Of[T any](v T, err error) Result[T] {
	return Result[T]{
		val: v,
		err: err,
	}
}

func Ok[T any](v T) Result[T] {
	return Result[T]{
		val: v,
	}
}

func Fail[T any](err error) Result[T] {
	return Result[T]{
		err: err,
	}
}

func (r Result[T]) Value() T {
	return r.val
}

func (r Result[T]) ValueOr(d T) T {
	return function.Ternary(r.err == nil, r.val, d)
}

func (r Result[T]) Must() T {
	if r.err != nil {
		panic(fmt.Errorf("unexpected error in result.Result %+v", r.err))
	}

	return r.val
}

func (r Result[T]) Get() (T, error) {
	return r.val, r.err
}

func (r Result[T]) Err() error {
	return r.err
}

func (r Result[T]) IsOK() bool {
	return r.err == nil
}

func (r Result[T]) IsErr() bool {
	return r.err != nil
}

func (r Result[T]) Then(f func(T)) Result[T] {
	if r.err == nil {
		f(r.val)
	}
	return r
}

func (r Result[T]) Exception(f func(error)) Result[T] {
	if r.err != nil {
		f(r.err)
	}
	return r
}
