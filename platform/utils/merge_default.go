package utils

import (
	"reflect"
)

// deprecated after v1.1.0
// MergeStructs is a function that merges a map of string/any into a struct
// The map should match with the struct fields
// T is a JSON struct
func MergeStructs[T any](dataMap map[string]any, w T) T {
	for k, v := range dataMap {
		reflect.ValueOf(&w).Elem().FieldByName(k).Set(reflect.ValueOf(v))
	}
	return w
}
