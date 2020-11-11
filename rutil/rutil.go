package rutil

import (
	"reflect"
	"strings"
)

// DerefValue returns dereferenced value of v
func DerefValue(v reflect.Value) reflect.Value {
	if v.Type().Kind() == reflect.Ptr {
		return v.Elem()
	}
	return v
}

// TypeOfSlicEl returns type of slice elements
func TypeOfSlicEl(t reflect.Type) reflect.Type {
	if t.Kind() == reflect.Slice {
		return t.Elem()
	}
	return t
}

// DerefType returns dereverenced value of t
func DerefType(t reflect.Type) reflect.Type {
	if t.Kind() == reflect.Ptr {
		return t.Elem()
	}
	return t
}

// DerefedTypeOf returns type of dereferenced v
func DerefedTypeOf(v interface{}) reflect.Type {
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Ptr {
		return t.Elem()
	}
	return t
}

// SliceTypeOf returns type of slice element
func SliceTypeOf(v interface{}) reflect.Type {
	t := DerefedTypeOf(v)
	if t.Kind() == reflect.Slice {
		return t.Elem()
	}
	return t
}

func GetInt64FieldValue(structValue interface{}, field string, caseSensitive bool) (int64, bool) {
	v, ok := structValue.(reflect.Value)
	if !ok {
		v = DerefValue(reflect.ValueOf(v))
	}

	if v.Type().Kind() != reflect.Struct {
		return 0, false
	}

	if fld, ok := FindField(v, field, caseSensitive); ok {
		if ok {
			res, ok := fld.Interface().(int64)
			return res, ok
		}
	}

	return 0, false
}

func FindField(v reflect.Value, field string, caseSensitive bool) (reflect.Value, bool) {
	t := v.Type()
	numField := t.NumField()

	if !caseSensitive {
		field = strings.ToLower(field)
	}

	for i := 0; i < numField; i++ {
		fname := t.Field(i).Name

		if !caseSensitive {
			fname = strings.ToLower(fname)
		}

		if fname == field {
			return v.Field(i), true
		}
	}

	return reflect.Value{}, false
}

func GetValue(structValue interface{}, field string, caseSensitive bool) (interface{}, bool) {
	v := DerefValue(reflect.ValueOf(structValue))
	v, ok := FindField(v, field, caseSensitive)
	if !ok {
		return nil, false
	}
	v = DerefValue(v)
	return v.Interface(), true
}
