package reflectutils

import (
	"reflect"
)

type field struct {
	Level []int8
}

func (this field) get_value(val reflect.Value) (reflect.Value, bool) {
	for _, level := range this.Level {
		// all pointer convert to actual values
		if val.Kind() == reflect.Ptr {
			if val.IsNil() {
				// if pointer is nil then we can't go deeper in sturcture
				return val, false
			}
			val = val.Elem()
		}
		// loop unli we find somthing that isn't structure
		if val.Kind() != reflect.Struct {
			break
		}
		val = val.Field(int(level))
	}

	// if we find pointer and not nil, then return value where pointer is pointing
	if val.Kind() == reflect.Ptr && val.IsNil() == false {
		val = val.Elem()
	}

	return val, true
}
