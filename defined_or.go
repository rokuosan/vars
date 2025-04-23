package vars

import "reflect"

func DefinedOr[T any](p *T, v T) T {
	if p == nil {
		return v
	}

	rv := reflect.ValueOf(*p)
	if !rv.IsValid() {
		return v
	}
	switch rv.Kind() {
	case reflect.Slice, reflect.Map, reflect.Ptr, reflect.Chan, reflect.Func:
		if rv.IsNil() {
			return v
		}
	case reflect.Invalid:
		return v
	}

	return *p
}
