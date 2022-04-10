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

func SetErrorWrapper(wrapper function.Operation[error]) {
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

func TimeValidx(t time.Time, errMsg ...string) func() error {
	return func() error {
		return TimeValid(t, errMsg...)
	}
}

type Number interface {
	constraint.Int | constraint.Uint | constraint.Float
}

func EQ[T comparable](expect, actual T, errMsg ...string) error {
	return True(actual == expect, errMsg...)
}

func EQx[T comparable](expect, actual T, errMsg ...string) func() error {
	return func() error {
		return EQ(expect, actual, errMsg...)
	}
}

func NE[T comparable](expect, actual T, errMsg ...string) error {
	return True(actual != expect, errMsg...)
}

func NEx[T comparable](expect, actual T, errMsg ...string) func() error {
	return func() error {
		return NE(expect, actual, errMsg...)
	}
}

func GTE[N Number](expect, actual N, errMsg ...string) error {
	return True(actual >= expect, errMsg...)
}

func GTEx[N Number](expect, actual N, errMsg ...string) func() error {
	return func() error {
		return GTE(expect, actual, errMsg...)
	}
}

func GT[N Number](expect, actual N, errMsg ...string) error {
	return True(actual > expect, errMsg...)
}

func GTx[N Number](expect, actual N, errMsg ...string) func() error {
	return func() error {
		return GT(expect, actual, errMsg...)
	}
}

func LTE[N Number](expect, actual N, errMsg ...string) error {
	return True(actual <= expect, errMsg...)
}

func LTEx[N Number](expect, actual N, errMsg ...string) func() error {
	return func() error {
		return LTE(expect, actual, errMsg...)
	}
}

func LT[N Number](expect, actual N, errMsg ...string) error {
	return True(actual < expect, errMsg...)
}

func LTx[N Number](expect, actual N, errMsg ...string) func() error {
	return func() error {
		return LT(expect, actual, errMsg...)
	}
}

func In[T comparable](expect []T, actual T, errMsg ...string) error {
	return True(find(expect, actual), errMsg...)
}

func Inx[T comparable](expect []T, actual T, errMsg ...string) func() error {
	return func() error {
		return In(expect, actual, errMsg...)
	}
}

func NotEmpty[T any](list []T, errMsg ...string) error {
	return True(len(list) != 0, errMsg...)
}

func NotEmptyx[T any](list []T, errMsg ...string) func() error {
	return func() error {
		return NotEmpty(list, errMsg...)
	}
}

func NoLonger[T any](maxLen int, list []T, errMsg ...string) error {
	return True(len(list) <= maxLen, errMsg...)
}

func NoLongerx[T any](maxLen int, list []T, errMsg ...string) func() error {
	return func() error {
		return NoLonger(maxLen, list, errMsg...)
	}
}

func StringNoLonger(maxLen int, s string, errMsg ...string) error {
	return True(utf8.RuneCountInString(s) <= maxLen, errMsg...)
}

func StringNoLongerx(maxLen int, s string, errMsg ...string) func() error {
	return func() error {
		return StringNoLonger(maxLen, s, errMsg...)
	}
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
