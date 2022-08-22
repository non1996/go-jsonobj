package function

import "reflect"

type BiConsumer[T any, U any] func(T, U)
type BiFunction[T any, U any, R any] func(T, U) R
type BiOperation[T any] func(T, T) T
type BiPredicate[T any, U any] func(T, U) bool

type Consumer[T any] func(T)
type Function[T any, R any] func(T) R
type Operation[T any] func(T) T
type Predicate[T any] func(T) bool
type Supplier[T any] func() T
type Comparator[T any] func(T, T) bool

func Noop[T any](t T) T {
	return t
}

// DerefOrDef 对指针解引用，如果是空指针则返回默认值
func DerefOrDef[T any, P *T](ptr P, d T) (v T) {
	if ptr == nil {
		return d
	}
	return *ptr
}

// Ternary 三元表达式
func Ternary[T any](cond bool, v1, v2 T) T {
	if cond {
		return v1
	}
	return v2
}

// TernaryLazy 使用函数闭包提供值的三元表达式
func TernaryLazy[T any](cond bool, supplier1, supplier2 Supplier[T]) T {
	if cond {
		return supplier1()
	}
	return supplier2()
}

// NonNil 判断是否是nil指针
func NonNil[T any](t *T) bool {
	return t != nil
}

// Ref 传入一个值，返回它的指针
func Ref[T any](t T) *T {
	return &t
}

// Indirect 函数接受一个指针，返回指针指向的值（解引用），对于空指针回返回指针指向类型的默认值
func Indirect[T any](t *T) T {
	if t != nil {
		return *t
	}
	return Zero[T]()
}

// IndirectOr 函数和 Indirect 很像，但允许用户自定义遇到空指针时返回的默认值。
func IndirectOr[T any](t *T, d T) T {
	if t != nil {
		return *t
	}
	return d
}

// Zero 函数返回一个类型的默认值 / 零值
func Zero[T any]() T {
	return *new(T)
}

// Type 返回一个类型的指针的 reflect.Type
func Type[T any]() reflect.Type {
	return reflect.TypeOf((*T)(nil))
}
