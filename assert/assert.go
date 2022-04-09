package assert

import (
	"fmt"
	"github.com/non1996/go-jsonobj/constraint"
	"github.com/non1996/go-jsonobj/function"
	"github.com/non1996/go-jsonobj/stream"
	"time"
	"unicode/utf8"
)

var errorWrapper function.Operation[error] = func(err error) error { return err }

func setErrorWrapper(wrapper function.Operation[error]) {
	errorWrapper = wrapper
}

func True(cond bool, errMsg ...string) error {
	if !cond {
		return errorWrapper(fmt.Errorf(getErrMsg(errMsg...)))
	}
	return nil
}

func False(cond bool, errMsg ...string) error {
	if cond {
		return errorWrapper(fmt.Errorf(getErrMsg(errMsg...)))
	}
	return nil
}

func TimeValid(t time.Time, errMsg ...string) error {
	return False(t.IsZero(), errMsg...)
}

func Int64EQ(expect, actual int64, errMsg ...string) error {
	return True(actual == expect, errMsg...)
}

func Int64NE(expect, actual int64, errMsg ...string) error {
	return True(actual != expect, errMsg...)
}

func Int64GTE(expect, actual int64, errMsg ...string) error {
	return True(actual >= expect, errMsg...)
}

func Int64GT(expect, actual int64, errMsg ...string) error {
	return True(actual > expect, errMsg...)
}

func Int64LTE(expect, actual int64, errMsg ...string) error {
	return True(actual <= expect, errMsg...)
}

func Int64LT(expect, actual int64, errMsg ...string) error {
	return True(actual < expect, errMsg...)
}

func Int64In(expect []int64, actual int64, errMsg ...string) error {
	return True(find(expect, actual), errMsg...)
}

func StringNotEmpty(s string, errMsg ...string) error {
	return True(len(s) != 0, errMsg...)
}

func StringNoLonger(maxLen int, s string, errMsg ...string) error {
	return True(utf8.RuneCountInString(s) <= maxLen, errMsg...)
}

func StringEqual(expect, actual string, errMsg ...string) error {
	return True(expect == actual, errMsg...)
}

func StringIn(expect []string, actual string, errMsg ...string) error {
	return True(findString(expect, actual), errMsg...)
}

func findString(expect []string, actual string) bool {
	for _, e := range expect {
		if e == actual {
			return true
		}
	}
	return false
}

type Number interface {
	constraint.Int | constraint.Uint | constraint.Float
}

func EQ[T comparable](expect, actual T, errMsg ...string) error {
	return True(actual == expect, errMsg...)
}

func NE[N comparable](expect, actual N, errMsg ...string) error {
	return True(actual != expect, errMsg...)
}

func GTE[N Number](expect, actual N, errMsg ...string) error {
	return True(actual >= expect, errMsg...)
}

func GT[N Number](expect, actual N, errMsg ...string) error {
	return True(actual > expect, errMsg...)
}

func LTE[N Number](expect, actual N, errMsg ...string) error {
	return True(actual <= expect, errMsg...)
}

func LT[N Number](expect, actual N, errMsg ...string) error {
	return True(actual < expect, errMsg...)
}

func In[N comparable](expect []N, actual N, errMsg ...string) error {
	return True(find(expect, actual), errMsg...)
}

func NotEmpty[T any](list []T, errMsg ...string) error {
	return True(len(list) != 0, errMsg...)
}

func NoLonger[T any](maxLen int, list []T, errMsg ...string) error {
	return True(len(list) <= maxLen, errMsg...)
}

func find[T comparable](list []T, expect T) bool {
	return stream.Slice(list).Find(func(v T) bool { return v == expect }).IsPresent()
}

func getErrMsg(msg ...string) string {
	if len(msg) > 0 {
		return msg[0]
	}
	return "assertion failed"
}
