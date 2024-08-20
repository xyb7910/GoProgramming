package reflect

import (
	"errors"
	"reflect"
)

// Iterate 遍历数组，切片，或者字符串
func Iterate(input any) ([]any, error) {
	val := reflect.ValueOf(input)
	typ := val.Type()
	kind := typ.Kind()

	if kind != reflect.Array && kind != reflect.Slice && kind != reflect.String {
		return nil, errors.New("input must be array, slice or string")
	}
	res := make([]any, 0, val.Len())
	for i := 0; i < val.Len(); i++ {
		ele := val.Index(i)
		if kind == reflect.String {
			res = append(res, string(ele.Interface().(uint8)))
		} else {
			res = append(res, ele.Interface())
		}
	}
	return res, nil
}

// IterateMapV1 遍历map
func IterateMapV1(input any) ([]any, []any, error) {
	val := reflect.ValueOf(input)
	if val.Kind() != reflect.Map {
		return nil, nil, errors.New("input must be map")
	}
	l := val.Len()
	keys := make([]any, 0, l)
	values := make([]any, 0, l)
	itr := val.MapRange()
	for itr.Next() {
		keys = append(keys, itr.Key().Interface())
		values = append(values, itr.Value().Interface())
	}
	return keys, values, nil
}

// IterateMapV2 遍历map
func IterateMapV2(input any) ([]any, []any, error) {
	val := reflect.ValueOf(input)
	if val.Kind() != reflect.Map {
		return nil, nil, errors.New("input must be map")
	}
	l := val.Len()
	keys := make([]any, 0, l)
	values := make([]any, 0, l)
	for _, key := range val.MapKeys() {
		keys = append(keys, key.Interface())
		v := val.MapIndex(key)
		values = append(values, v.Interface())
	}
	return keys, values, nil
}
