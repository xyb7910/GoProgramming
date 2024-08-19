package reflect

import (
	"errors"
	"reflect"
)

func IterateFields(val any) (map[string]any, error) {
	// 判断是否为nil
	if val == nil {
		return nil, errors.New("不能为nil")
	}

	// 获取类型
	tpe := reflect.TypeOf(val)
	// 获取值
	refVal := reflect.ValueOf(val)

	// 判断是否为指针，并且处理指针
	for tpe.Kind() == reflect.Ptr {
		tpe = tpe.Elem()
		refVal = refVal.Elem()
	}

	// 确保是结构体
	if tpe.Kind() != reflect.Struct {
		return nil, errors.New("传入的值必须是结构体类型")
	}

	// 结构体取值
	numField := tpe.NumField()
	res := make(map[string]any, numField)
	for i := 0; i < numField; i++ {
		fdType := tpe.Field(i)
		if fdType.IsExported() {
			res[fdType.Name] = refVal.Field(i).Interface()
		} else {
			res[fdType.Name] = reflect.Zero(fdType.Type).Interface()
		}
	}
	return res, nil
}

func SetField(entity any, fieldName string, value any) error {
	val := reflect.ValueOf(entity)
	typ := val.Type()

	// 只能是一级指针
	if typ.Kind() != reflect.Ptr || typ.Elem().Kind() != reflect.Struct {
		return errors.New("非法类型")
	}

	typ = typ.Elem()
	val = val.Elem()

	fd := val.FieldByName(fieldName)
	if _, found := typ.FieldByName(fieldName); !found {
		return errors.New("字段不存在")
	}
	if !fd.CanSet() {
		return errors.New("字段不可设置")
	}
	fd.Set(reflect.ValueOf(value))
	return nil
}
