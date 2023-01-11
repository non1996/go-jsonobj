package container

import (
	"container/list"

	util "github.com/non1996/go-jsonobj/utils"
)

type OrderedMap[K comparable, V any] struct {
	l   *list.List
	idx map[K]*list.Element
}

func NewOrderedMap[K comparable, V any]() *OrderedMap[K, V] {
	return &OrderedMap[K, V]{
		l:   list.New(),
		idx: map[K]*list.Element{},
	}
}

func (o *OrderedMap[K, V]) Size() int {
	return o.l.Len()
}

func (o *OrderedMap[K, V]) Add(key K, val V) {
	if elem, exist := o.idx[key]; exist {
		elem.Value = util.NewPair(key, val)
	} else {
		o.idx[key] = o.l.PushBack(util.NewPair(key, val))
	}
}

func (o *OrderedMap[K, V]) Exist(key K) bool {
	_, exist := o.idx[key]
	return exist
}

func (o *OrderedMap[K, V]) Get(key K) (val V) {
	if elem, exist := o.idx[key]; exist {
		return elem.Value.(util.Pair[K, V]).Second
	}
	return
}

func (o *OrderedMap[K, V]) Remove(key K) {
	if elem, exist := o.idx[key]; exist {
		o.l.Remove(elem)
		delete(o.idx, key)
	}
}

func (o *OrderedMap[K, V]) Keys() []K {
	keys := make([]K, 0, len(o.idx))
	for k := range o.idx {
		keys = append(keys, k)
	}
	return keys
}

func (o *OrderedMap[K, V]) ForeachErr(action func(K, V) error) error {
	for elem := o.l.Front(); elem != nil; elem = elem.Next() {
		p := elem.Value.(util.Pair[K, V])
		err := action(p.First, p.Second)
		if err != nil {
			return err
		}
	}
	return nil
}
