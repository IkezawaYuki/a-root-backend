package util

import "reflect"

func Pointer[T any](value T) *T {
	return &value
}

func IsNil(i interface{}) bool {
	if i == nil {
		return true
	}
	v := reflect.ValueOf(i)
	if v.Kind() == reflect.Ptr {
		return v.IsNil()
	}
	return false
}
