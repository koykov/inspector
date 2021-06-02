package inspector

// Reflect inspector. Strongly recommend to not use it. Need only for performance/allocation comparison.

import (
	"fmt"
	"reflect"
	"strconv"
)

type ReflectInspector struct {
	BaseInspector
}

func (i *ReflectInspector) Get(src interface{}, path ...string) (interface{}, error) {
	var (
		r interface{}
		c int
		k interface{}
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

func (i *ReflectInspector) GetTo(src interface{}, buf *interface{}, path ...string) error {
	var err error
	*buf, err = i.Get(src, path...)
	return err
}

func (i *ReflectInspector) Cmp(_ interface{}, _ Op, _ string, _ *bool, _ ...string) error {
	// Empty method, I'm too lazy to implement it now.
	return nil
}

func (i *ReflectInspector) Set(_, _ interface{}, _ ...string) error {
	// Empty method, there is no way to update data using reflection.
	return nil
}

func (i *ReflectInspector) SetWB(_, _ interface{}, _ AccumulativeBuffer, _ ...string) error {
	// Empty method, there is no way to update data using reflection.
	return nil
}

func (i *ReflectInspector) Loop(_ interface{}, _ Looper, _ *[]byte, _ ...string) (err error) {
	// Empty method. todo implement it
	return nil
}

func (i *ReflectInspector) inspect(node interface{}, key string) interface{} {
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
	}
	return nil
}
