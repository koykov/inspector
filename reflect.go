package inspector

// Reflect inspector. Strongly recommend to not use it. Need only for performance/allocation comparison.

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"

	"github.com/koykov/byteconv"
)

type ReflectInspector struct {
	BaseInspector
}

func (i ReflectInspector) TypeName() string {
	return "reflect"
}

func (i ReflectInspector) Get(src any, path ...string) (any, error) {
	var (
		r any
		c int
		k any
	)
	r = src
	for c, k = range path {
		r = i.inspect(r, k.(string))
	}
	if c < len(path)-1 {
		r = nil
	}
	return r, nil
}

func (i ReflectInspector) GetTo(src any, buf *any, path ...string) error {
	var err error
	*buf, err = i.Get(src, path...)
	return err
}

func (i ReflectInspector) Compare(_ any, _ Op, _ string, _ *bool, _ ...string) error {
	// Empty method, I'm too lazy to implement it now.
	return nil
}

func (i ReflectInspector) Set(_, _ any, _ ...string) error {
	// Empty method, there is no way to update data using reflection.
	return nil
}

func (i ReflectInspector) SetWithBuffer(_, _ any, _ AccumulativeBuffer, _ ...string) error {
	// Empty method, there is no way to update data using reflection.
	return nil
}

func (i ReflectInspector) Loop(_ any, _ Iterator, _ *[]byte, _ ...string) (err error) {
	// Empty method. todo implement it
	return nil
}

func (i ReflectInspector) DeepEqual(l, r any) bool {
	return reflect.DeepEqual(l, r)
}

func (i ReflectInspector) DeepEqualWithOptions(l, r any, _ *DEQOptions) bool {
	return reflect.DeepEqual(l, r)
}

func (i ReflectInspector) inspect(node any, key string) any {
	v := reflect.ValueOf(node)
	switch v.Kind() {
	case reflect.Ptr:
		if elem := v.Elem(); elem.IsValid() && elem.CanInterface() {
			node = elem.Interface()
			return i.inspect(node, key)
		}
	case reflect.Map:
		kv := reflect.ValueOf(key)
		_ = kv
		for _, f := range v.MapKeys() {
			fv := f.Interface()
			if fvs, ok := fv.(string); ok {
				if fvs == key {
					mv := v.MapIndex(f)
					if mv.IsValid() && mv.CanInterface() {
						return mv.Interface()
					}
					return nil
				}
			}
			if fmt.Sprintf("%v", fv) == key {
				mv := v.MapIndex(f)
				if mv.IsValid() && mv.CanInterface() {
					return mv.Interface()
				}
			}
		}
	case reflect.Struct:
		f := v.FieldByName(key)
		if f.IsValid() && f.CanInterface() {
			return f.Interface()
		}
	case reflect.Slice:
		if bytes, ok := node.([]byte); ok {
			return bytes
		}
		idx, err := strconv.Atoi(key)
		if err != nil {
			return nil
		}
		sv := v.Index(idx)
		if sv.IsValid() && sv.CanInterface() {
			return sv.Interface()
		}
	default:
		// noop
	}
	return nil
}

func (i ReflectInspector) Unmarshal(p []byte, typ Encoding) (any, error) {
	var x any
	switch typ {
	case EncodingJSON:
		err := json.Unmarshal(p, &x)
		return x, err
	default:
		return nil, ErrUnknownEncodingType
	}
}

func (i ReflectInspector) Copy(x any) (any, error) {
	return x, nil
}

func (i ReflectInspector) CopyTo(_, _ any, _ AccumulativeBuffer) error {
	return nil
}

func (i ReflectInspector) Each(src any, fn func(i int, field string, value any)) error {
	v := reflect.ValueOf(src)
	if v.Kind() == reflect.Ptr {
		if elem := v.Elem(); elem.IsValid() && elem.CanInterface() {
			return i.Each(elem.Interface(), fn)
		}
		return ErrUnsupportedType
	}
	switch v.Kind() {
	case reflect.Map:
		var c int
		for _, f := range v.MapKeys() {
			fv := f.Interface()
			var fvs string
			switch x := fv.(type) {
			case string:
				fvs = x
			case *string:
				fvs = *x
			case []byte:
				fvs = byteconv.B2S(x)
			case *[]byte:
				fvs = byteconv.B2S(*x)
			default:
				fvs = fmt.Sprintf("%v", fv)
			}
			mv := v.MapIndex(f)
			if mv.IsValid() && mv.CanInterface() {
				fn(c, fvs, mv.Interface())
			}
		}
	case reflect.Struct:
		var c int
		for k, v1 := range v.Fields() {
			if v1.IsValid() && v1.CanInterface() {
				fn(c, k.Name, v1.Interface())
				c++
			}
		}
	default:
		// noop
	}
	return nil
}

func (i ReflectInspector) Length(_ any, _ *int, _ ...string) error { return nil }

func (i ReflectInspector) Capacity(_ any, _ *int, _ ...string) error { return nil }

func (i ReflectInspector) Reset(_ any, _ ...string) error { return nil }
