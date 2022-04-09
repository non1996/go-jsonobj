package container

func Keys[K comparable, V any](m map[K]V) (keys []K) {
	if len(m) == 0 {
		return nil
	}
	keys = make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func Values[K comparable, V any](m map[K]V) (values []V) {
	if len(m) == 0 {
		return nil
	}
	values = make([]V, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}
