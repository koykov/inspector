package inspector

import (
	"encoding/json"

	"github.com/koykov/byteconv"
)

type StringStringMapInspector struct {
	BaseInspector
}

func (i StringStringMapInspector) TypeName() string {
	return "map[string]string"
}

func (i StringStringMapInspector) Instance(ptr bool) any {
	if ptr {
		return &map[string]string{}
	}
	return map[string]string{}
}

func (i StringStringMapInspector) Get(src any, path ...string) (any, error) {
	var buf any
	err := i.GetTo(src, &buf, path...)
	return buf, err
}

func (i StringStringMapInspector) GetTo(src any, buf *any, path ...string) error {
	if len(path) == 0 {
		*buf = src
		return nil
	}
	m, err := i.indir(src)
	if err != nil {
		return err
	}
	x, ok := m[path[0]]
	if ok {
		*buf = &x
		return nil
	}
	return nil
}

func (i StringStringMapInspector) Set(dst, value any, path ...string) error {
	var buf ByteBuffer
	return i.SetWithBuffer(dst, value, &buf, path...)
}

func (i StringStringMapInspector) SetWithBuffer(dst, value any, buf AccumulativeBuffer, path ...string) (err error) {
	if len(path) == 0 {
		return
	}
	var buf_ map[string]string
	if err = i.indir1(&buf_, dst); err != nil || buf_ == nil {
		return err
	}
	if len(path) > 1 {
		x, ok := buf_[path[0]]
		if !ok {
			return
		}
		err = i.SetWithBuffer(x, value, buf, path[1:]...)
		buf_[path[0]] = x
	} else {
		switch x := value.(type) {
		case string:
			buf_[path[0]] = buf.BufferizeString(x)
		case *string:
			buf_[path[0]] = buf.BufferizeString(*x)
		case []byte:
			buf_[path[0]] = byteconv.B2S(buf.Bufferize(x))
		case *[]byte:
			buf_[path[0]] = byteconv.B2S(buf.Bufferize(*x))
		default:
			buf_[path[0]], err = buf.BufferizeAnyString(value)
		}
	}
	return
}

func (i StringStringMapInspector) Compare(src any, cond Op, right string, result *bool, path ...string) (err error) {
	if len(path) == 0 {
		return
	}
	var buf_ map[string]string
	if err = i.indir1(&buf_, src); err != nil || buf_ == nil {
		return err
	}
	x, ok := buf_[path[0]]
	if !ok {
		return
	}
	if len(path) > 1 {
		err = i.Compare(x, cond, right, result, path[1:]...)
	} else {
		si := StaticInspector{}
		err = si.Compare(x, cond, right, result)
	}
	return
}

func (i StringStringMapInspector) Loop(src any, it Iterator, buf *[]byte, path ...string) error {
	if len(path) == 0 {
		m, err := i.indir(src)
		if err != nil {
			return err
		}
		for k := range m {
			if it.RequireKey() {
				it.SetKey(&k, StaticInspector{})
			}
			it.SetVal(m[k], StaticInspector{})
			ctl := it.Iterate()
			if ctl == LoopCtlBrk {
				break
			}
			if ctl == LoopCtlCnt {
				continue
			}
		}
		return nil
	}

	m, err := i.indir(src)
	if err != nil {
		return err
	}
	x, ok := m[path[0]]
	if !ok {
		return nil
	}
	return i.Loop(x, it, buf, path[1:]...)
}

func (i StringStringMapInspector) DeepEqual(l, r any) bool {
	return i.DeepEqualWithOptions(l, r, nil)
}

func (i StringStringMapInspector) DeepEqualWithOptions(l, r any, opts *DEQOptions) bool {
	var ml, mr map[string]string
	if err := i.indir1(&ml, l); err != nil || ml == nil {
		return false
	}
	if err := i.indir2(&mr, r); err != nil || mr == nil {
		return false
	}
	if len(ml) != len(mr) {
		return false
	}

	for k, v := range ml {
		if mr[k] != v {
			return false
		}
	}
	for k, v := range mr {
		if ml[k] != v {
			return false
		}
	}

	return false
}

func (i StringStringMapInspector) Unmarshal(p []byte, typ Encoding) (any, error) {
	var x map[string]string
	switch typ {
	case EncodingJSON:
		err := json.Unmarshal(p, &x)
		return x, err
	default:
		return nil, ErrUnknownEncodingType
	}
}

func (i StringStringMapInspector) Copy(x any) (dst any, err error) {
	var buf ByteBuffer
	x_ := make(map[string]string)
	err = i.CopyTo(x, &x_, &buf)
	dst = x_
	return
}

func (i StringStringMapInspector) CopyTo(src, dst any, buf AccumulativeBuffer) (err error) {
	var msrc, mdst map[string]string
	if err = i.indir1(&msrc, src); err != nil || msrc == nil {
		return
	}
	if err = i.indir2(&mdst, dst); err != nil || mdst == nil {
		return
	}
	for k := range mdst {
		delete(mdst, k)
	}

	for k, v := range msrc {
		mdst[k] = v
	}
	return
}

func (i StringStringMapInspector) Each(src any, fn func(i int, field string, value any)) error {
	m, err := i.indir(src)
	if err != nil {
		return err
	}
	var c int
	for k, v := range m {
		fn(c, k, v)
		c++
	}
	return nil
}

func (i StringStringMapInspector) Length(x any, result *int, path ...string) error {
	if len(path) == 0 {
		m, err := i.indir(x)
		if err == nil {
			*result = len(m)
			return nil
		}
		switch x1 := x.(type) {
		case string:
			*result = len(x1)
		case *string:
			*result = len(*x1)
		case []byte:
			*result = len(x1)
		case *[]byte:
			*result = len(*x1)
		}
		return nil
	}
	m, err := i.indir(x)
	if err != nil {
		return err
	}
	x1, ok := m[path[0]]
	if !ok {
		return nil
	}
	return i.Length(x1, result, path[1:]...)
}

func (i StringStringMapInspector) Capacity(x any, result *int, path ...string) error {
	if len(path) == 0 {
		switch x1 := x.(type) {
		case []byte:
			*result = cap(x1)
		case *[]byte:
			*result = cap(*x1)
		}
		return nil
	}
	m, err := i.indir(x)
	if err != nil {
		return err
	}
	x1, ok := m[path[0]]
	if !ok {
		return nil
	}
	return i.Length(x1, result, path[1:]...)
}

func (i StringStringMapInspector) Reset(x any, path ...string) error {
	var m map[string]string
	if err := i.indir1(&m, x); err != nil || x == nil {
		return nil
	}
	if len(path) == 0 {
		for k := range m {
			delete(m, k)
		}
		return nil
	}
	if v, ok := m[path[0]]; ok {
		if len(path) == 1 {
			var m1 map[string]string
			if err := i.indir1(&m1, v); err == nil {
				for k := range m1 {
					delete(m1, k)
				}
				return nil
			}
			delete(m, path[0])
			return nil
		}
		if err := i.Reset(v, path[1:]...); err != nil {
			return err
		}
		m[path[0]] = v
	}
	return nil
}

func (i StringStringMapInspector) indir(val any) (buf map[string]string, err error) {
	err = i.indir1(&buf, val)
	return
}

func (i StringStringMapInspector) indir1(dst *map[string]string, val any) error {
	switch x := val.(type) {
	case map[string]string:
		*dst = x
	case *map[string]string:
		*dst = *x
	case **map[string]string:
		*dst = *(*x)
	default:
		return ErrUnsupportedType
	}
	return nil
}

func (i StringStringMapInspector) indir2(dst *map[string]string, val any) error {
	switch x := val.(type) {
	case map[string]string:
		return ErrMustPointerType
	case *map[string]string:
		*dst = *x
	case **map[string]string:
		*dst = *(*x)
	default:
		return ErrUnsupportedType
	}
	return nil
}
