package reflectutils

import (
	"reflect"
)

func type_of(obj interface{}) reflect.Type {
	typ := value_of(obj).Type()
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	return typ
}

func value_of(obj interface{}) reflect.Value {
	val := reflect.ValueOf(obj)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	return val
}

func ToMap(obj interface{}) map[string]interface{} {
	m := make(map[string]interface{})

	val := value_of(obj)
	typ := val.Type()

	if typ.Kind() == reflect.Map {
		for _, k := range val.MapKeys() {
			m[k.String()] = val.MapIndex(k).Interface()
		}
		return m
	}

	layout := get_layout(typ)

	for name, field := range layout.Fields {
		if val, ok := field.get_value(val); ok {
			m[name] = val.Interface()
		}
	}

	return m
}

func IsSlice(obj interface{}) bool {
	return type_of(obj).Kind() == reflect.Slice
}

func Foreach(obj interface{}, fn func(int, interface{}) bool) {
	val := value_of(obj)
	typ := val.Type()

	if typ.Kind() == reflect.Slice {
		n := val.Len()
		for i := 0; i < n; i++ {
			if !fn(i, val.Index(i).Interface()) {
				break
			}
		}
	}
}
