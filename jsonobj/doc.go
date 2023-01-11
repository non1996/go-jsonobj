package jsonobj

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/modern-go/reflect2"
)

type Key interface{}

// Doc json obj/arr/value的包装
type Doc struct {
	parent    *Doc        // 该doc的parent
	parentKey Key         // 该doc在parent中的key
	v         interface{} // 值
}

func O(kvs ...interface{}) (d *Doc) {
	assertKVNum(len(kvs))
	m := make(map[string]interface{})
	for i := 0; i < len(kvs); i += 2 {
		m[kvs[i].(string)] = assertValueType(kvs[i+1])
	}
	return &Doc{v: m}
}

func NewObj(data []byte) (d *Doc, err error) {
	m := make(map[string]interface{})
	err = jsonAPI.Unmarshal(data, &m)
	if err != nil {
		return nil, err
	}
	return &Doc{v: m}, nil
}

func MustNewObj(data []byte) (d *Doc) {
	d, err := NewObj(data)
	if err != nil {
		d = O()
	}
	return d
}

func A(vs ...interface{}) (d *Doc) {
	a := make([]interface{}, 0, len(vs))
	for _, v := range vs {
		a = append(a, assertValueType(v))
	}
	return &Doc{v: a}
}

func NewArr(data []byte) (d *Doc, err error) {
	var a []interface{}
	err = jsonAPI.Unmarshal(data, &a)
	if err != nil {
		return nil, err
	}
	return &Doc{v: a}, nil
}

func MustNewArr(data []byte) (d *Doc) {
	d, err := NewArr(data)
	if err != nil {
		d = A()
	}
	return d
}

func (d *Doc) IsNil() bool {
	return reflect2.IsNil(d.v)
}

func (d *Doc) Interface() interface{} {
	return d.v
}

func (d *Doc) MarshalJSON() ([]byte, error) {
	return jsonAPI.Marshal(d.v)
}

func (d *Doc) UnmarshalJSON(p []byte) error {
	err := jsonAPI.Unmarshal(p, &d.v)
	if err != nil {
		return err
	}
	d.setBack()
	return nil
}

// Check map 独有，检查 key 是否存在
func (d *Doc) Check(key string) bool {
	m, err := d.CastMap()
	if err != nil {
		return false
	}
	_, exist := m[key]
	return exist
}

// Keys map 独有，获取 map 中所有的 key
func (d *Doc) Keys() []string {
	m, err := d.CastMap()
	if err != nil {
		return nil
	}
	res := make([]string, 0, len(m))
	for k := range m {
		res = append(res, k)
	}
	return res
}

func (d *Doc) Replace(value interface{}) {
	d.v = assertValueType(value)
	d.setBack()
}

func (d *Doc) Set(key Key, value interface{}) *Doc {
	key = assertKeyType(key)
	d.preAlloc(key)
	switch k := key.(type) {
	case int64:
		d.setArr(k, assertValueType(value))
	case string:
		d.setMap(k, assertValueType(value))
	}
	d.setBack()
	return d
}

func (d *Doc) SetPath(path []Key, value interface{}) {
	if len(path) == 0 {
		return
	}

	prefix := path[:len(path)-1]
	final := path[len(path)-1]
	iter := d
	for _, p := range prefix {
		p = assertKeyType(p)
		if iter.IsNil() {
			iter.preAlloc(p)
		}
		iter = iter.Get(p)
	}
	iter.Set(assertKeyType(final), value)
}

func (d *Doc) Get(key Key) *Doc {
	key = assertKeyType(key)
	switch k := key.(type) {
	case int64:
		return d.getArr(k)
	case string:
		return d.getMap(k)
	default:
		panic(errInvalidKey(key))
	}
}

func (d *Doc) GetPath(path ...Key) *Doc {
	iter := d
	for _, k := range path {
		k = assertKeyType(k)
		if iter.IsNil() {
			return &Doc{v: nil}
		}
		iter = iter.Get(k)
	}
	return iter
}

func (d *Doc) Del(key Key) {
	key = assertKeyType(key)
	switch k := key.(type) {
	case int64:
		d.delArr(k)
	case string:
		d.delMap(k)
	default:
		panic(errInvalidKey(key))
	}
	d.setBack()
}

// Add Array 类型独有，往数组中添加，确保不重复
func (d *Doc) Add(value interface{}) {
	value = assertValueType(value)
	a, err := d.CastArr()
	if err != nil {
		return
	}
	for _, i := range a {
		if reflect.DeepEqual(i, value) {
			return
		}
	}
	a = append(a, value)
	d.v = a
	d.setBack()
}

// Append Array 类型独有，往数组中添加
func (d *Doc) Append(value interface{}) {
	value = assertValueType(value)
	a, err := d.CastArr()
	if err != nil {
		return
	}
	a = append(a, value)
	d.v = a
	d.setBack()
}

func (d *Doc) Insert(index int64, value interface{}) {
	value = assertValueType(value)
	a, err := d.CastArr()
	if err != nil {
		return
	}
	if index < 0 || index > int64(len(a)) {
		return
	}
	anew := make([]interface{}, 0, len(a)+1)
	anew = append(anew, a[:index]...)
	anew = append(anew, value)
	anew = append(anew, a[index:]...)
	d.v = anew
	d.setBack()
}

func (d *Doc) PopBack() (res interface{}) {
	a, err := d.CastArr()
	if err != nil {
		return nil
	}
	if len(a) == 0 {
		return nil
	}
	res = a[len(a)-1]
	d.v = a[:len(a)-1]
	d.setBack()
	return
}

// Remove Array 类型独有，从数组中删除值等于value的第一项，目前只能对string生效
func (d *Doc) Remove(value interface{}) int64 {
	a, err := d.CastArr()
	if err != nil {
		return -1
	}
	var i int
	for ; i < len(a); i++ {
		if reflect.DeepEqual(a[i], value) {
			break
		}
	}
	if i == len(a) {
		return -1
	}
	anew := make([]interface{}, 0, len(a))
	anew = append(anew, a[:i]...)
	anew = append(anew, a[i+1:]...)
	d.v = anew
	d.setBack()
	return int64(i)
}

func (d *Doc) RemoveAll(value interface{}) {
	a, err := d.CastArr()
	if err != nil {
		return
	}
	anew := make([]interface{}, 0, len(a))
	for _, v := range a {
		if !reflect.DeepEqual(v, value) {
			anew = append(anew, v)
		}
	}
	d.v = anew
	d.setBack()
}

func (d *Doc) Len() int {
	a, err := d.CastArr()
	if err == nil {
		return len(a)
	}
	m, err := d.CastMap()
	if err == nil {
		return len(m)
	}
	return 0
}

func (d *Doc) preAlloc(key Key) {
	if !d.IsNil() {
		return
	}
	switch key.(type) {
	case string:
		d.v = make(map[string]interface{})
	case int64:
		d.v = make([]interface{}, 0)
	}
}

func (d *Doc) setArr(key int64, value interface{}) {
	a, err := d.CastArr()
	if err != nil {
		return
	}
	for key >= int64(len(a)) {
		a = append(a, nil)
	}
	a[key] = value
	d.v = a
}

func (d *Doc) setMap(key string, value interface{}) {
	m, err := d.CastMap()
	if err != nil {
		return
	}
	m[key] = value
}

func (d *Doc) delArr(key int64) {
	a, err := d.CastArr()
	if err != nil {
		return
	}
	if key >= int64(len(a)) || key < 0 {
		return
	}
	newArr := make([]interface{}, 0, len(a)-1)
	newArr = append(newArr, a[:key]...)
	newArr = append(newArr, a[key+1:]...)
	d.v = newArr
}

func (d *Doc) delMap(key string) {
	m, err := d.CastMap()
	if err != nil {
		return
	}
	delete(m, key)
}

func (d *Doc) getArr(key int64) *Doc {
	a, err := d.CastArr()
	if err != nil {
		return &Doc{v: nil}
	}
	if key < 0 || key >= int64(len(a)) {
		return &Doc{v: nil}
	}
	return &Doc{parent: d, parentKey: key, v: a[key]}
}

func (d *Doc) getMap(key string) *Doc {
	m, err := d.CastMap()
	if err != nil {
		return &Doc{v: nil}
	}
	return &Doc{parent: d, parentKey: key, v: m[key]}
}

func (d *Doc) setBack() {
	if d.parent != nil {
		d.parent.Set(d.parentKey, d)
	}
}

func (d *Doc) CastMap() (map[string]interface{}, error) {
	m, ok := d.v.(map[string]interface{})
	if !ok {
		return nil, errCastFailed("map[string]interface{}")
	}
	return m, nil
}

func (d *Doc) CastArr() ([]interface{}, error) {
	a, ok := d.v.([]interface{})
	if !ok {
		return nil, errCastFailed("[]interface{}")
	}
	return a, nil
}

func (d *Doc) CastInt() (i int64, err error) {
	switch v2 := d.v.(type) {
	case json.Number:
		i, err = v2.Int64()
	case float32, float64:
		i = int64(reflect.ValueOf(v2).Float())
	case int, int8, int16, int32, int64:
		i = reflect.ValueOf(v2).Int()
	case uint, uint8, uint16, uint32, uint64:
		i = int64(reflect.ValueOf(v2).Uint())
	default:
		err = errCastFailed("int")
	}
	return i, err
}

func (d *Doc) CastFloat() (f float64, err error) {
	switch v2 := d.v.(type) {
	case json.Number:
		f, err = v2.Float64()
	case float32, float64:
		f = reflect.ValueOf(v2).Float()
	case int, int8, int16, int32, int64:
		f = float64(reflect.ValueOf(v2).Int())
	case uint, uint8, uint16, uint32, uint64:
		f = float64(reflect.ValueOf(v2).Uint())
	default:
		err = errCastFailed("float")
	}
	return f, err
}

func (d *Doc) CastBool() (bool, error) {
	if b, ok := d.v.(bool); ok {
		return b, nil
	}
	return false, errCastFailed("bool")
}

func (d *Doc) CastString() (string, error) {
	if s, ok := d.v.(string); ok {
		return s, nil
	}
	return "", errCastFailed("string")
}

func (d *Doc) Map(df ...map[string]interface{}) map[string]interface{} {
	o, err := d.CastMap()
	if err != nil && len(df) > 0 {
		o = df[0]
	}
	return o
}

func (d *Doc) Arr(df ...[]interface{}) []interface{} {
	a, err := d.CastArr()
	if err != nil && len(df) > 0 {
		a = df[0]
	}
	return a
}

func (d *Doc) Int(df ...int64) int64 {
	i, err := d.CastInt()
	if err != nil && len(df) > 0 {
		i = df[0]
	}
	return i
}

func (d *Doc) Float(df ...float64) float64 {
	f, err := d.CastFloat()
	if err != nil && len(df) > 0 {
		f = df[0]
	}
	return f
}

func (d *Doc) Bool(df ...bool) bool {
	b, err := d.CastBool()
	if err != nil && len(df) > 0 {
		b = df[0]
	}
	return b
}

func (d *Doc) String(df ...string) string {
	s, err := d.CastString()
	if err != nil && len(df) > 0 {
		s = df[0]
	}
	return s
}

func (d *Doc) CastStrArr() (res []string, err error) {
	a, err := d.CastArr()
	if err != nil {
		return nil, err
	}
	res = make([]string, 0, len(a))
	for _, s := range a {
		if s2, ok := s.(string); ok {
			res = append(res, s2)
		}
	}
	return res, nil
}

func (d *Doc) StrArr(df ...[]string) (res []string) {
	a, err := d.CastStrArr()
	if err != nil && len(df) != 0 {
		a = df[0]
	}
	return a
}

func (d *Doc) ToString() string {
	b, _ := json.Marshal(d)
	return string(b)
}

func assertKVNum(num int) {
	if num&1 != 0 {
		panic("expect even num params")
	}
}

func assertKeyType(k interface{}) Key {
	switch k.(type) {
	case int8, int16, int32, int64, int,
		uint, uint8, uint16, uint32, uint64:
		return reflect.ValueOf(k).Int()
	case string:
		return k
	default:
		panic(errInvalidKey(k))
	}
}

func assertValueType(v interface{}) interface{} {
	switch v2 := v.(type) {
	case int8, int16, int32, int64, int,
		uint, uint8, uint16, uint32, uint64,
		float32, float64,
		bool, string, nil:
		return v
	case *Doc:
		return v2.v
	default:
		panic(errInvalidValue(v))
	}
}

func errCastFailed(name string) error {
	return fmt.Errorf("type assertion to %s failed", name)
}

func errInvalidKey(key Key) error {
	return fmt.Errorf("invalid json key type: %s", reflect.TypeOf(key).Name())
}

func errInvalidValue(value interface{}) error {
	return fmt.Errorf("invalid json value type: %s", reflect.TypeOf(value).Name())
}
