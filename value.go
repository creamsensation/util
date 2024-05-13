package util

import "reflect"

func SetValueToReflected(kind reflect.Kind, field reflect.Value, value string) {
	switch kind {
	case reflect.String:
		field.Set(reflect.ValueOf(value))
	case reflect.Bool:
		field.Set(reflect.ValueOf(value == "true"))
	case reflect.Int:
		field.Set(reflect.ValueOf(AssertStringToType[int](value)))
	case reflect.Float32:
		field.Set(reflect.ValueOf(AssertStringToType[float64](value)))
	case reflect.Float64:
		field.Set(reflect.ValueOf(AssertStringToType[float64](value)))
	default:
	}
}
