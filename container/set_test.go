package container

import (
	"fmt"
	"testing"
)

func TestSet(t *testing.T) {
	set := NewSet[int]()
	set.Add(1)
	set.Add(2)
	set.Add(2)
	set.Add(3)

	fmt.Println(set.Contains(1))
	fmt.Println(set.Contains(2))
	fmt.Println(set.Contains(3))
	fmt.Println(set.Contains(4))
	fmt.Println(set.ToList())
}
