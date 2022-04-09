package wrapper

import (
	"strconv"
)

type Int32 struct {
	i     int32
	valid bool
}

func (i *Int32) ParseString(s string) (err error) {
	res, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return err
	}
	i.i = int32(res)
	i.valid = true
	return
}

func (i *Int32) ValueOf(v int32) {
	i.i = v
	i.valid = true
}

func (i *Int32) V() int32 {
	if i.valid {
		return i.i
	}
	return 0
}

func (i *Int32) I64() int64 {
	if i.valid {
		return int64(i.i)
	}
	return 0
}

func (i *Int32) Valid() bool {
	return i.valid
}

type Int64 struct {
	i     int64
	valid bool
}

func (i *Int64) ParseString(s string) (err error) {
	res, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err
	}
	i.i = res
	i.valid = true
	return
}

func (i *Int64) ValueOf(v int64) {
	i.i = v
	i.valid = true
}

func (i *Int64) V() int64 {
	if i.valid {
		return i.i
	}
	return 0
}

func (i *Int64) I32() int32 {
	if i.valid {
		return int32(i.i)
	}
	return 0
}

func (i *Int64) Valid() bool {
	return i.valid
}
