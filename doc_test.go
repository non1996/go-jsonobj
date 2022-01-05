package jsonobj

import (
	"encoding/json"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"testing"
)

func TestO(t *testing.T) {
	o := O()

	s := `{
		"a": ["1", "2", "3"],
		"b": {
			"c": [
				{
					"e": "5"
				}
			]
		}
	}`

	err := json.Unmarshal([]byte(s), o)
	if err != nil {
	    t.Fatal(err)
	}
	fmt.Println(o.v)
	a := o.Get("a")
	fmt.Println(a)
	a.Remove("2")
	fmt.Println(a)
	fmt.Println(o)
	a.Set(3, "5")
	fmt.Println(a)
	fmt.Println(o)
	a.Append("5")
	fmt.Println(a)
	fmt.Println(o)

	e := o.GetPath("b", "c", 0)
	fmt.Println(e)
	e.Set("x", 100)

	e.Del("e")
	content, err := json.Marshal(o)
	if err != nil {
	     t.Fatal(err)
	}
	fmt.Println(string(content))
}

func TestDeserialize(t *testing.T) {
	s := `{
		"a": ["1", "2", "3"],
		"b": {
			"c": [
				{
					"e": "5"
				}
			]
		}
	}`

	o := struct {
		B *Doc `json:"b"`
	}{}

	err := json.Unmarshal([]byte(s), &o)
	if err != nil {
	    t.Fatal(err)
	}
	fmt.Println(o.B)
	fmt.Println(o.B.Get("c").Get(0).Get("e").String())
}

func TestHttp(t *testing.T) {
	a := struct {
		A string `json:"a"`
	}{}
	a.A = "https://p3-if-sign.byteimg.com/tos-cn-i-0xlkdnewzn/b48e92cf9f2f4fce93fca7ccfe89ac4d~tplv-0xlkdnewzn-engine-high.webp?x-expires=1644120187&x-signature=uBBzyFQv0Awd2FpQ%2Bkb4dUMyR8k%3D"
	b, _ := jsonAPI.Marshal(a)
	fmt.Println(string(b))
	b2, _ := jsoniter.Marshal(a)
	fmt.Println(string(b2))
	b3, _ := json.Marshal(a)
	fmt.Println(string(b3))
}
