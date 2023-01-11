package function

import (
	"reflect"

	"github.com/non1996/go-jsonobj/constraint"
)

// Noop 不进行任何操作，直接返回
func Noop[T any](t T) T {
	return t
}

// DerefOrDef 对指针解引用，如果是空指针则返回默认值
// Deprecated
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

// IsNil 判断是否是nil指针
func IsNil[T any](t *T) bool {
	return t == nil
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

// Add 两个值相加
func Add[T constraint.Number | constraint.Complex | constraint.String](
	x T,
	y T,
) T {
	return x + y
}

// Equal 判断两个值是否相等
func Equal[T comparable](v1, v2 T) bool {
	return v1 == v2
}

// Greater 判断v1是否大于v2
func Greater[T constraint.Orderable](v1, v2 T) bool {
	return v1 > v2
}

// Less 判断v1是否小于v2
func Less[T constraint.Orderable](v1, v2 T) bool {
	return v1 < v2
}

// Max 返回两值中较大的那个
func Max[T constraint.Orderable](v1, v2 T) T {
	if v1 > v2 {
		return v1
	}
	return v2
}

// Min 返回两值中较小的那个
func Min[T constraint.Orderable](v1, v2 T) T {
	if v1 > v2 {
		return v2
	}
	return v1
}

func NonZeroOr[T comparable](v T, d T) T {
	if v == Zero[T]() {
		return d
	}
	return v
}
