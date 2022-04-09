package wrapper

import (
	"strconv"
)

type Bool struct {
	b     bool
	valid bool
}

func (b *Bool) ParseString(s string) (err error) {
	res, err := strconv.ParseBool(s)
	if err != nil {
		return err
	}
	b.b = res
	b.valid = true
	return
}

func ValueOfBool(b bool) Bool {
	return Bool{b: b, valid: true}
}

func (b *Bool) ValueOf(v bool) {
	b.b = v
	b.valid = true
}

func (b *Bool) Value() bool {
	if b.valid {
		return b.b
	}
	return false
}

func (b *Bool) Valid() bool {
	return b.valid
}
