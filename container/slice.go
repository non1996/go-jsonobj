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

func SliceConcat[T any](s1, s2 []T) []T {
	var res = make([]T, 0, len(s1)+len(s2))
	res = append(res, s1...)
	res = append(res, s2...)
	return res
}

func SliceConcatM[T any](s1 []T, ss ...[]T) []T {
	var res []T
	res = append(res, s1...)
	for _, s2 := range ss {
		res = append(res, s2...)
	}
	return res
}

func SliceGetLast[T any](s []T) T {
	return s[len(s)-1]
}

func SliceGetLastOr[T any](s []T, d T) T {
	if len(s) == 0 {
		return d
	}
	return s[len(s)-1]
}

func SliceGetFirst[T any](s []T) T {
	return s[0]
}

func SliceGetFirstOr[T any](s []T, d T) T {
	if len(s) == 0 {
		return d
	}
	return s[0]
}

func SliceCopy[T any](s []T) []T {
	if len(s) == 0 {
		return nil
	}

	newSlice := make([]T, len(s))
	for idx, v := range s {
		newSlice[idx] = v
	}

	return newSlice
}
