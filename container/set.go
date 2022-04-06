package container

import (
	"encoding/json"
)

type SetValue interface {
	comparable
}

type placeholder struct{}

type Set[V SetValue] map[V]placeholder

func NewSet[V SetValue](list ...V) Set[V] {
	set := Set[V]{}
	for _, l := range list {
		set.Add(l)
	}
	return set
}

func (s Set[V]) Size() int64 {
	return int64(len((map[V]placeholder)(s)))
}

func (s Set[V]) Add(i V) bool {
	if s.Contains(i) {
		return false
	}
	s[i] = placeholder{}
	return true
}

func (s Set[V]) Remove(i V) {
	delete(s, i)
}

func (s Set[V]) Contains(i V) bool {
	_, exist := s[i]
	return exist
}

// IsSubset 判断是不是另一个set的subset
func (s Set[V]) IsSubset(other Set[V]) bool {
	if other == nil {
		return true
	}
	if s.Size() > other.Size() {
		return false
	}
	for elem := range s {
		if !other.Contains(elem) {
			return false
		}
	}
	return true
}

// IsSuperSet 判断是不是另一个set的superset
func (s Set[V]) IsSuperSet(other Set[V]) bool {
	return other != nil && other.IsSubset(s)
}

// Equal 判断集合是否相等
func (s Set[V]) Equal(other Set[V]) bool {
	if s.Size() != other.Size() {
		return false
	}
	for elem := range s {
		if !other.Contains(elem) {
			return false
		}
	}
	return true
}

// Union 并集
func (s Set[V]) Union(other Set[V]) Set[V] {
	unioned := Set[V]{}
	for elem := range s {
		unioned.Add(elem)
	}
	for elem := range other {
		unioned.Add(elem)
	}
	return unioned
}

// Intersect 交集
func (s Set[V]) Intersect(other Set[V]) Set[V] {
	intersection := Set[V]{}
	// loop over smaller Set
	var small, large Set[V]
	if s.Size() < other.Size() {
		small, large = s, other
	} else {
		small, large = other, s
	}
	for elem := range small {
		if large.Contains(elem) {
			intersection.Add(elem)
		}
	}
	return intersection
}

// Difference 差集
func (s Set[V]) Difference(other Set[V]) Set[V] {
	difference := Set[V]{}
	for elem := range s {
		if !other.Contains(elem) {
			difference.Add(elem)
		}
	}
	return difference
}

func (s Set[V]) ToList() []V {
	list := make([]V, 0, s.Size())
	for elem := range s {
		list = append(list, elem)
	}
	return list
}

func (s Set[V]) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.ToList())
}

func (s Set[V]) UnmarshalJSON(b []byte) error {
	var res []V
	err := json.Unmarshal(b, &res)
	if err != nil {
		return err
	}
	for _, i := range res {
		s.Add(i)
	}
	return nil
}
