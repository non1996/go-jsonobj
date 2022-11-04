package constraint

type Int interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Uint interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type String interface {
	~string
}

type Float interface {
	~float32 | ~float64
}

type Bool interface {
	~bool
}

type Number interface {
	Int | Uint | Float
}
