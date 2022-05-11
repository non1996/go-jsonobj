package stream

import (
	"fmt"
	"testing"
)

func TestSliceStream(t *testing.T) {
	var a = []int64{1, 2, 3, 5, 5, 4, 3}
	res1 := Slice(a).
		Limit(5).
		Skip(3).
		Peek(func(v int64) { fmt.Println("peek", v) }).
		Map(func(v int64) int64 { return v * 2 }).
		ToList()
	fmt.Println(res1)

	fmt.Println(Slice(a).AllMatch(func(v int64) bool { return v > 4 }))
	fmt.Println(Slice(a).AllMatch(func(v int64) bool { return v < 10 }))
	fmt.Println(Slice(a).AnyMatch(func(v int64) bool { return v == 5 }))
	fmt.Println(Slice(a).AnyMatch(func(v int64) bool { return v == 6 }))
	fmt.Println(Slice(a).NoneMatch(func(v int64) bool { return v > 10 }))
	fmt.Println(Slice(a).NoneMatch(func(v int64) bool { return v < 10 }))
	fmt.Println(Slice(a).NoneMatch(func(v int64) bool { return v < 3 }))
	fmt.Println(Slice(a).Limit(5).Count())
	fmt.Println(Slice(a).Sorted(intcomp).Limit(5).Count())

	fmt.Println("filter", Slice(a).Filter(func(v int64) bool { return v > 3 }).Count())

	Slice(a).Find(func(v int64) bool { return v > 3 }).
		IfPresent(func(v int64) { fmt.Println("find", v) })
	Slice(a).Find(func(v int64) bool { return v > 30 }).
		IfPresent(func(v int64) { fmt.Println("find", v) }).
		Else(func() { fmt.Println("not found") })
}

func TestMap(t *testing.T) {
	var a = []int64{1, 2, 3, 4, 5, 6}
	var b = Map(a, func(v int64) string { return fmt.Sprintf("%d", v) })
	fmt.Println(b)
}

func TestReduce(t *testing.T) {
	var a = []int64{1, 2, 3, 4, 5, 6}
	Slice(a).Limit(4).Reduce(int64(1), func(v1, v2 int64) int64 { return v1 * v2 }).
		IfPresent(func(v int64) { fmt.Println(v) })
}

func TestCollectToMap(t *testing.T) {
	var a = []int64{1, 2, 3, 4, 5, 6}
	res := CollectToMap(a, func(v int64) int64 { return v }, func(v int64) string { return fmt.Sprintf("id-%+v", v) })
	fmt.Println(res)
}

func TestSorted(t *testing.T) {
	var a = []int64{1, 2, 3, 9, 5, 6, 7}
	after := Slice(a).
		Skip(1).
		Limit(5).
		Filter(func(i int64) bool { return i > 5 }).
		Sorted(func(i1, i2 int64) bool { return i1 > i2 }).
		Limit(3).
		ToList()
	fmt.Println(after)
}

func TestMapping(t *testing.T) {
	var a = []int64{1, 2, 3, 9, 5, 6, 7}

	m := MapS(Slice(a).Sorted(intcomp).Limit(4), func(i int64) string {
		return fmt.Sprintf("xxx-%d", i)
	}).ToList()
	fmt.Println(m)
}

func intcomp(i, j int64) bool {
	return i < j
}
