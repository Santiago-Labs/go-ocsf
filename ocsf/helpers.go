package ocsf

import "reflect"

type hasObservable interface {
	Observable() (*int, string)
}

func PresentObservablesOf(obj any) [][]any {
	if obj == nil {
		return nil
	}

	var out [][]any
	var walk func(reflect.Value)

	walk = func(v reflect.Value) {
		if !v.IsValid() || v.Kind() == reflect.Ptr && v.IsNil() {
			return
		}

		if typeId, name := asObservable(v); typeId != nil {
			out = append(out, []any{typeId, name})
		}

		switch v.Kind() {
		case reflect.Ptr:
			if v.IsNil() {
				return
			}
			walk(v.Elem())

		case reflect.Struct:
			for i := 0; i < v.NumField(); i++ {
				walk(v.Field(i))
			}

		case reflect.Slice, reflect.Array:
			if v.Len() == 0 {
				return
			}
			for i := 0; i < v.Len(); i++ {
				walk(v.Index(i))
			}
		}
	}

	rv := reflect.ValueOf(obj)
	if rv.Kind() == reflect.Ptr && rv.Elem().Kind() == reflect.Struct {
		rv = rv.Elem()
	}
	walk(rv)
	return out
}

func asObservable(v reflect.Value) (*int, string) {
	if v.IsValid() && v.CanInterface() {
		if h, ok := v.Interface().(hasObservable); ok {
			if typeId, name := h.Observable(); typeId != nil {
				return typeId, name
			}
		}
	}
	if v.CanAddr() {
		av := v.Addr()
		if av.CanInterface() {
			if h, ok := av.Interface().(hasObservable); ok {
				if typeId, name := h.Observable(); typeId != nil {
					return typeId, name
				}
			}
		}
	}
	return nil, ""
}
