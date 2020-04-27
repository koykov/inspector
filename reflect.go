package inspector

import (
	"fmt"
	"reflect"
	"strconv"
	"unsafe"
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

func (i *ReflectInspector) GetUnsafe(src interface{}, path ...string) (unsafe.Pointer, uint, error) {
	val, err := i.Get(src, path...)
	return unsafe.Pointer(&val), 0, err
}

func (i *ReflectInspector) Set(dst, value interface{}, path ...string) {
	// Empty method, there is no way to update data using reflection.
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
