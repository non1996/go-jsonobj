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
	Reduce(int64(1), a, func(v1, v2 int64) int64 { return v1 * v2 }).
		IfPresent(func(v int64) { fmt.Println(v) })
}

func TestCollectToMap(t *testing.T) {
	var a = []int64{1, 2, 3, 4, 5, 6}
	res := CollectToMap(a, func(v int64) int64 { return v }, func(v int64) string { return fmt.Sprintf("id-%+v", v) })
	fmt.Println(res)
}

type Service interface {
	Exec() error
}

type Req interface {
	Seriallize() string
}

type BaseService[R Req] struct {
	req R
}

func NewBaseService[T Req](req T) BaseService[T] {
	return BaseService[T]{req: req}
}

type ProjectReq struct {
	ID int64
}

func (p *ProjectReq) Seriallize() string {
	//TODO implement me
	panic("implement me")
}

type ProjectService struct {
	BaseService[*ProjectReq]
}

func NewProjectService() Service {
	return &ProjectService{NewBaseService(&ProjectReq{})}
}

func (s *ProjectService) Exec() error {
	fmt.Println(s.req.ID)
	return nil
}

func TestNewProjectService(t *testing.T) {
	p := NewProjectService()

	p.Exec()
}
