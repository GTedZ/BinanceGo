package validation

import "reflect"

func IsNotDefault(v any) bool {
	if v == nil {
		return false
	}

	val := reflect.ValueOf(v)

	switch val.Kind() {
	case reflect.Slice, reflect.Array, reflect.Map, reflect.String:
		return val.Len() != 0

	case reflect.Ptr, reflect.Interface:
		return !val.IsNil()

	default:
		zero := reflect.Zero(val.Type())
		return !reflect.DeepEqual(v, zero.Interface())
	}
}
