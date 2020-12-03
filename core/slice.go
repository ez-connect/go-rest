package core

import (
	"reflect"
)

func IndexOf(values []string, value string) int {
	for i, v := range values {
		if v == value {
			return i
		}
	}

	return -1
}

func IndexOfWithField(values interface{}, value interface{}, key string) int {
	rv := reflect.ValueOf(values)
	if rv.Kind() != reflect.Slice || rv.Len() <= 0 {
		return -1
	}

	sv := value
	if reflect.ValueOf(value).Kind() == reflect.Ptr {
		sv = reflect.ValueOf(value).Elem().Interface()
	}

	elementType := reflect.TypeOf(values).Elem()
	if elementType.Kind() == reflect.Ptr {
		elementType = elementType.Elem()
	}

	st := reflect.TypeOf(sv)
	if key != "" && elementType.Kind() == reflect.Struct {
		t, ok := elementType.FieldByName(key)
		if ok && ((t.Type.Kind() == reflect.Ptr && st != t.Type.Elem()) || (t.Type.Kind() != reflect.Ptr && st != t.Type)) {
			return -1
		}
	} else if elementType != reflect.TypeOf(sv) {
		return -1
	}

	if key != "" && elementType.Kind() == reflect.Struct {
		for i := 0; i < rv.Len(); i++ {
			t := rv.Index(i)
			field := t.FieldByName(key)
			if !field.IsValid() {
				return -1
			}

			if field.Kind() == reflect.Ptr {
				if field.Elem().Interface() == sv {
					return i
				}
			} else if field.Interface() == sv {
				return i
			}
		}
	} else {
		for i := 0; i < rv.Len(); i++ {
			t := rv.Index(i)

			if t.Kind() == reflect.Ptr {
				if t.Elem().Interface() == sv {
					return i
				}
			} else if t.Interface() == sv {
				return i
			}
		}
	}

	return -1
}

func IndexOfWithFunction(values interface{}, value interface{}, compareFunction func(interface{}, interface{}) bool) int {
	rv := reflect.ValueOf(values)
	if rv.Kind() != reflect.Slice || rv.Len() <= 0 {
		return -1
	}
	for i := 0; i < rv.Len(); i++ {
		t := rv.Index(i).Interface()
		if compareFunction(t, value) {
			return i
		}
	}
	return -1
}
