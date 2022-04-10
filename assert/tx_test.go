package assert

import (
	"fmt"
	"testing"
	"time"
)

func Test(t *testing.T) {
	fmt.Println("TimeValid")
	fmt.Println(TimeValid(time.Now(), "success"))
	fmt.Println(TimeValid(time.Time{}, "fail"))

	fmt.Println("EQ")
	fmt.Println(EQ(0, 0, "success"))
	fmt.Println(EQ(0.0, 0.1, "fail"))
	fmt.Println(EQ("a", "a", "success"))

	fmt.Println("NE")
	fmt.Println(NE(0, 1, "success"))
	fmt.Println(NE(0.0, 0.0, "fail"))

	fmt.Println("GTE")
	fmt.Println(GTE(0.0, 0.8, "success"))
	fmt.Println(GTE(0.0, 0.0, "success"))
	fmt.Println(GTE(1.1, 0.8, "fail"))

	fmt.Println("GT")
	fmt.Println(GT(0.0, 0.8, "success"))
	fmt.Println(GT(1.1, 0.8, "fail"))

	fmt.Println("LTE")
	fmt.Println(LTE(0.8, 0.0, "success"))
	fmt.Println(LTE(0.0, 0.0, "success"))
	fmt.Println(LTE(0.1, 1.8, "fail"))

	fmt.Println("LT")
	fmt.Println(LT(0.8, 0.0, "success"))
	fmt.Println(LT(1.1, 1.1, "fail"))

	fmt.Println("IN")
	fmt.Println(In([]int64{1, 2, 3}, 3, "success"))
	fmt.Println(In([]int64{1, 2, 3}, 4, "fail"))

	fmt.Println("NotEmpty")
	fmt.Println(NotEmpty([]int64{1}, "success"))
	fmt.Println(NotEmpty([]int64{}, "fail"))

	fmt.Println("NoLonger")
	fmt.Println(NoLonger(1, []int64{1}, "success"))
	fmt.Println(NoLonger(0, []int64{1, 2, 3}, "fail"))
}

func TestTx(t *testing.T) {
	err := Should(
		TimeValidx(time.Time{}, "fail 1"),
		EQx(0, 0, "success"),
		NotEmptyx([]int64{1}, "success"),
		NoLongerx(1, []int64{1, 2}, "fail 2"),
	)
	fmt.Println(err)
}
