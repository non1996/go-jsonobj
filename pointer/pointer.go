package pointer

func NonNil[T any](t *T) bool {
	return t != nil
}
