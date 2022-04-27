package container

func NilToEmpty[T any](s []T) []T {
	if s != nil {
		return s
	}
	return make([]T, 0)
}
