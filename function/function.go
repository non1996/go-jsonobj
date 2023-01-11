package function

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
