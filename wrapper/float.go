package wrapper

import (
	"strconv"
)

type Float32 struct {
	f     float32
	valid bool
}

func (f *Float32) ParseString(s string) (err error) {
	res, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return err
	}
	f.f = float32(res)
	f.valid = true
	return
}

func ValueOfFloat32(f float32) Float32 {
	return Float32{f: f, valid: true}
}

func (f *Float32) ValueOf(v float32) {
	f.f = v
	f.valid = true
}

func (f *Float32) Value() float32 {
	if f.valid {
		return f.f
	}
	return 0
}

func (f *Float32) Valid() bool {
	return f.valid
}

type Float64 struct {
	f     float64
	valid bool
}

func (f *Float64) ParseString(s string) (err error) {
	res, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return err
	}
	f.f = res
	f.valid = true
	return
}

func ValueOfFloat64(f float64) Float64 {
	return Float64{f: f, valid: true}
}

func (f *Float64) ValueOf(v float64) {
	f.f = v
	f.valid = true
}

func (f *Float64) Value() float64 {
	if f.valid {
		return f.f
	}
	return 0
}

func (f *Float64) Valid() bool {
	return f.valid
}
