package container

import (
	util "github.com/non1996/go-jsonobj/utils"
)

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

func MapKeys[K comparable, V any](m map[K]V) (keys []K) {
	if len(m) == 0 {
		return nil
	}
	keys = make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func MapValues[K comparable, V any](m map[K]V) (values []V) {
	if len(m) == 0 {
		return nil
	}
	values = make([]V, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

func MapEntries[K comparable, V any](m map[K]V) (entries []util.Pair[K, V]) {
	if len(m) == 0 {
		return nil
	}
	entries = make([]util.Pair[K, V], 0, len(m))
	for k, v := range m {
		entries = append(entries, util.NewPair(k, v))
	}
	return entries
}

func MapToSlice[K comparable, V any, T any](m map[K]V, mapping func(K, V) T) []T {
	res := make([]T, 0, len(m))
	for k, v := range m {
		res = append(res, mapping(k, v))
	}
	return res
}

func MapSize[K comparable, V any](m map[K]V) int64 {
	return int64(len(m))
}

func MapIsEmpty[K comparable, V any](m map[K]V) bool {
	return len(m) == 0
}

func MapContainsKey[K comparable, V any](m map[K]V, k K) bool {
	_, exist := m[k]
	return exist
}

func MapPutAll[K comparable, V any](m, other map[K]V) {
	for k, v := range other {
		m[k] = v
	}
}

func MapClear[K comparable, V any](m map[K]V) {
	keys := MapKeys(m)
	for _, k := range keys {
		delete(m, k)
	}
}

func MapExist[K comparable, V any](m map[K]V, key K) bool {
	_, exist := m[key]
	return exist
}

func MapGet[K comparable, V any](m map[K]V, key K) V {
	return m[key]
}

func MapGetOrDefault[K comparable, V any](m map[K]V, key K, defaultValue V) V {
	v, exist := m[key]
	if exist {
		return v
	}
	return defaultValue
}

func MapForEach[K comparable, V any](m map[K]V, fn func(K, V)) {
	for k, v := range m {
		fn(k, v)
	}
}

func MapReplaceAll[K comparable, V any](m map[K]V, function func(K, V) V) {
	keys := MapKeys(m)
	for _, key := range keys {
		m[key] = function(key, m[key])
	}
}

func MapPutIfAbsent[K comparable, V any](m map[K]V, k K, v V) {
	if !MapContainsKey(m, k) {
		m[k] = v
	}
}

func MapReplace[K comparable, V any](m map[K]V, k K, v V) {
	if MapContainsKey(m, k) {
		m[k] = v
	}
}

func MapComputeIfAbsent[K comparable, V any](m map[K]V, k K, mapping func(K) (V, bool)) {
	if !MapContainsKey(m, k) {
		newValue, toAdd := mapping(k)
		if toAdd {
			m[k] = newValue
		}
	}
}

func MapComputeIfPresent[K comparable, V any](m map[K]V, k K, mapping func(K) (V, bool)) {
	if MapContainsKey(m, k) {
		newValue, toAdd := mapping(k)
		if toAdd {
			m[k] = newValue
		} else {
			delete(m, k)
		}
	}
}

func MapCompute[K comparable, V any](m map[K]V, k K, mapping func(K, V, bool) (V, bool)) (V, bool) {
	oldValue, exist := m[k]
	newValue, toAdd := mapping(k, oldValue, exist)

	if !toAdd {
		if exist {
			delete(m, k)
			return newValue, false
		} else {
			return newValue, false
		}
	} else {
		m[k] = newValue
		return newValue, true
	}
}

func MapCopy[K comparable, V any](m map[K]V) map[K]V {
	newMap := make(map[K]V, len(m))

	for k, v := range m {
		newMap[k] = v
	}

	return newMap
}

func MapMerge[K comparable, V any](m1, m2 map[K]V) map[K]V {
	newMap := make(map[K]V, len(m1)+len(m2))

	for k, v := range m1 {
		newMap[k] = v
	}

	for k, v := range m2 {
		newMap[k] = v
	}

	return newMap
}
