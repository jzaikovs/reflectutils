package reflectutils

import (
	"reflect"
)

var cache_layouts = true

type type_layout struct {
	Fields map[string]field
}

var get_layout_cache = make(map[reflect.Type]type_layout)

func get_layout(typ reflect.Type) type_layout {
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	layout, ok := get_layout_cache[typ]
	if !ok {
		layout = type_layout{Fields: get_fields(typ, nil)}
		if cache_layouts {
			get_layout_cache[typ] = layout
		}
	}
	return layout
}

func get_fields(typ reflect.Type, level []int8) map[string]field {
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	if typ.Kind() != reflect.Struct {
		return nil
	}

	n := typ.NumField()
	m := make(map[string]field)

	for i := 0; i < n; i++ {

		typeField := typ.Field(i)

		if typeField.PkgPath != "" {
			continue
		}

		if typeField.Anonymous {
			levels := make([]int8, len(level))
			copy(levels, level)
			if childs := get_fields(typeField.Type, append(levels, int8(typeField.Index[0]))); childs != nil {
				for k, v := range childs {
					m[k] = v
				}
			}
			continue
		}
		levels := make([]int8, len(level))
		copy(levels, level)
		m[typeField.Name] = field{append(levels, int8(typeField.Index[0]))}
	}
	return m
}
