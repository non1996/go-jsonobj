package assert

import (
	"fmt"
	"time"
	"unicode/utf8"

	"github.com/non1996/go-jsonobj/constraint"
	"github.com/non1996/go-jsonobj/function"
	"github.com/non1996/go-jsonobj/stream"
)

var errorWrapper function.Operation[error] = func(err error) error { return err }

func SetErrorWrapper(wrapper function.Operation[error]) {
	errorWrapper = wrapper
}

func True(cond bool, errMsg ...string) error {
	return assert(cond, func() string { return "expect condition is true but false" }, errMsg...)
}

func Truex(cond bool, errMsg ...string) func() error {
	return func() error {
		return True(cond, errMsg...)
	}
}

func False(cond bool, errMsg ...string) error {
	return assert(cond, func() string { return "expect condition is false but true" }, errMsg...)
}

func Falsex(cond bool, errMsg ...string) func() error {
	return func() error {
		return False(cond, errMsg...)
	}
}

func TimeValid(t time.Time, errMsg ...string) error {
	return assert(!t.IsZero(), func() string { return "invalid time" }, errMsg...)
}

func TimeValidx(t time.Time, errMsg ...string) func() error {
	return func() error {
		return TimeValid(t, errMsg...)
	}
}

func EQ[T comparable](expect, actual T, errMsg ...string) error {
	return assert(actual == expect, func() string {
		return fmt.Sprintf("expect %+v but %+v", expect, actual)
	}, errMsg...)
}

func EQx[T comparable](expect, actual T, errMsg ...string) func() error {
	return func() error {
		return EQ(expect, actual, errMsg...)
	}
}

func NE[T comparable](expect, actual T, errMsg ...string) error {
	return assert(actual != expect, func() string {
		return fmt.Sprintf("expect not %+v but it is", expect)
	}, errMsg...)
}

func NEx[T comparable](expect, actual T, errMsg ...string) func() error {
	return func() error {
		return NE(expect, actual, errMsg...)
	}
}

func GTE[N constraint.Number](expect, actual N, errMsg ...string) error {
	return assert(actual >= expect, func() string {
		return fmt.Sprintf("expect greater than or equal to %+v, but %+v", expect, actual)
	}, errMsg...)
}

func GTEx[N constraint.Number](expect, actual N, errMsg ...string) func() error {
	return func() error {
		return GTE(expect, actual, errMsg...)
	}
}

func GT[N constraint.Number](expect, actual N, errMsg ...string) error {
	return assert(actual > expect, func() string {
		return fmt.Sprintf("expect greater than %+v, but %+v", expect, actual)
	}, errMsg...)
}

func GTx[N constraint.Number](expect, actual N, errMsg ...string) func() error {
	return func() error {
		return GT(expect, actual, errMsg...)
	}
}

func LTE[N constraint.Number](expect, actual N, errMsg ...string) error {
	return assert(actual <= expect, func() string {
		return fmt.Sprintf("expect less than or equal to %+v, but %+v", expect, actual)
	}, errMsg...)
}

func LTEx[N constraint.Number](expect, actual N, errMsg ...string) func() error {
	return func() error {
		return LTE(expect, actual, errMsg...)
	}
}

func LT[N constraint.Number](expect, actual N, errMsg ...string) error {
	return assert(actual < expect, func() string {
		return fmt.Sprintf("expect less than %+v, but %+v", expect, actual)
	}, errMsg...)
}

func LTx[N constraint.Number](expect, actual N, errMsg ...string) func() error {
	return func() error {
		return LT(expect, actual, errMsg...)
	}
}

func In[T comparable](expect []T, actual T, errMsg ...string) error {
	return assert(find(expect, actual), func() string {
		return fmt.Sprintf("expect value in %+v, but %+v", expect, actual)
	}, errMsg...)
}

func Inx[T comparable](expect []T, actual T, errMsg ...string) func() error {
	return func() error {
		return In(expect, actual, errMsg...)
	}
}

func Between[T constraint.Number](min, max T, actual T, errMsg ...string) error {
	return assert(actual >= min && actual < max, func() string {
		return fmt.Sprintf("expect value between %+v and %+v, but %+v", min, max, actual)
	}, errMsg...)
}

func Betweenx[T constraint.Number](min, max T, actual T, errMsg ...string) func() error {
	return func() error {
		return Between(min, max, actual, errMsg...)
	}
}

func NotEmpty[T any](list []T, errMsg ...string) error {
	return assert(len(list) != 0, func() string {
		return "should not be empty list"
	}, errMsg...)
}

func NotEmptyx[T any](list []T, errMsg ...string) func() error {
	return func() error {
		return NotEmpty(list, errMsg...)
	}
}

func NoLonger[T any](maxLen int, list []T, errMsg ...string) error {
	return assert(len(list) <= maxLen, func() string {
		return fmt.Sprintf("should not be longer than %d", maxLen)
	}, errMsg...)
}

func NoLongerx[T any](maxLen int, list []T, errMsg ...string) func() error {
	return func() error {
		return NoLonger(maxLen, list, errMsg...)
	}
}

func StringNoLonger(maxLen int, s string, errMsg ...string) error {
	return assert(utf8.RuneCountInString(s) <= maxLen, func() string {
		return fmt.Sprintf("should not be longer than %d", maxLen)
	}, errMsg...)
}

func StringNoLongerx(maxLen int, s string, errMsg ...string) func() error {
	return func() error {
		return StringNoLonger(maxLen, s, errMsg...)
	}
}

func StringNotEmpty(s string, errMsg ...string) error {
	return assert(len(s) != 0, func() string {
		return fmt.Sprintf("should not be empty")
	}, errMsg...)
}

func StringNotEmptyx(s string, errMsg ...string) func() error {
	return func() error {
		return StringNotEmpty(s, errMsg...)
	}
}

func NotNil[T any](v *T, errMsg ...string) error {
	return assert(v != nil, func() string {
		return "should not be nil"
	}, errMsg...)
}

func NotNilx[T any](v *T, errMsg ...string) func() error {
	return func() error {
		return NotNil(v, errMsg...)
	}
}

func find[T comparable](list []T, expect T) bool {
	return stream.Slice(list).Find(func(v T) bool { return v == expect }).IsPresent()
}

func assert(cond bool, hint func() string, errMsg ...string) (err error) {
	if !cond {
		return errorWrapper(fmt.Errorf("[assertion] %s, %s", getErrMsg(errMsg...), hint()))
	}
	return nil
}

func getErrMsg(msg ...string) string {
	if len(msg) > 0 {
		return msg[0]
	}
	return "failed"
}
