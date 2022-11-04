package container

func NilToEmpty[T any](s []T) []T {
	if s != nil {
		return s
	}
	return make([]T, 0)
}

func EmptySlice[T any]() []T {
	return make([]T, 0)
}

func NilSlice[T any]() []T {
	return ([]T)(nil)
}

func SingletonSlice[T any](t T) []T {
	s := make([]T, 1, 1)
	s[0] = t
	return s
}
