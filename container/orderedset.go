package container

import (
	util "github.com/non1996/go-jsonobj/utils"
)

type OrderedSet[V SetValue] struct {
	m *OrderedMap[V, placeholder]
}

func NewOrderedSet[V SetValue]() *OrderedSet[V] {
	return &OrderedSet[V]{
		m: NewOrderedMap[V, placeholder](),
	}
}

func (o *OrderedSet[V]) Size() int {
	return o.m.Size()
}

func (o *OrderedSet[V]) Add(val V) {
	o.m.Add(val, placeholder{})
}

func (o *OrderedSet[V]) AddAll(vals []V) {
	for _, v := range vals {
		o.Add(v)
	}
}

func (o *OrderedSet[V]) Exist(val V) bool {
	return o.m.Exist(val)
}

func (o *OrderedSet[V]) Remove(val V) {
	o.m.Remove(val)
}

func (o *OrderedSet[V]) ToList() []V {
	res := make([]V, 0, o.Size())
	for elem := o.m.l.Front(); elem != nil; elem = elem.Next() {
		p := elem.Value.(util.Pair[V, V])
		res = append(res, p.First)
	}
	return res
}
