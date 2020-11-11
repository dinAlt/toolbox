package slices

import (
	"fmt"
	"reflect"
)

func StrContains(slice []string, value string, other ...string) bool {
	if slice == nil {
		return false
	}
	for i := range slice {
		if slice[i] == value {
			return true
		}
		for ii := range other {
			if slice[i] == other[ii] {
				return true
			}
		}
	}
	return false
}

func StrDropEmpty(slice []string) []string {
	if len(slice) == 0 {
		return nil
	}
	cur, next := 0, 1
	for next < len(slice) {
		switch {
		case slice[cur] != "":
			cur++
			if next <= cur {
				next++
			}
		case slice[next] != "":
			slice[cur] = slice[next]
			slice[next] = ""
			next++
		default:
			next++
		}
	}
	if slice[cur] == "" {
		return slice[:cur]
	}
	return slice[:cur+1]

}

func FindValue(slice interface{}, f func(interface{}) bool) (int, error) {
	v := reflect.ValueOf(slice)
	if v.Type().Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Type().Kind() != reflect.Slice {
		return -1, fmt.Errorf("FindValue(): \"slice\" value should be slice or *slice, not %v", v.Kind())
	}
	for i := 0; i < v.Len(); i++ {
		elem := v.Index(i).Interface()
		if f(elem) {
			return i, nil
		}
	}
	return -1, nil
}
