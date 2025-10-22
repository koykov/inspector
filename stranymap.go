package inspector

import (
	"encoding/json"
)

type StringAnyMapInspector struct {
	BaseInspector
}

func (i StringAnyMapInspector) TypeName() string {
	return "map[string]any"
}

func (i StringAnyMapInspector) Instance(ptr bool) any {
	if ptr {
		return &map[string]any{}
	}
	return map[string]any{}
}

func (i StringAnyMapInspector) Get(src any, path ...string) (any, error) {
	var buf any
	err := i.GetTo(src, &buf, path...)
	return buf, err
}

func (i StringAnyMapInspector) GetTo(src any, buf *any, path ...string) error {
	if len(path) == 0 {
		*buf = src
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
	return i.GetTo(x, buf, path[1:]...)
}

func (i StringAnyMapInspector) Set(dst, value any, path ...string) error {
	var buf ByteBuffer
	return i.SetWithBuffer(dst, value, &buf, path...)
}

func (i StringAnyMapInspector) SetWithBuffer(dst, value any, buf AccumulativeBuffer, path ...string) (err error) {
	if len(path) == 0 {
		return
	}
	var buf_ map[string]any
	if err = i.indir1(&buf_, dst); err != nil || buf_ == nil {
		return err
	}
	if len(path) > 1 {
		x, ok := buf_[path[0]]
		if !ok {
			x = make(map[string]any)
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
			buf_[path[0]] = buf.Bufferize(x)
		case *[]byte:
			buf_[path[0]] = buf.Bufferize(*x)
		default:
			buf_[path[0]] = value
		}
	}
	return
}

func (i StringAnyMapInspector) Compare(src any, cond Op, right string, result *bool, path ...string) (err error) {
	if len(path) == 0 {
		return
	}
	var buf_ map[string]any
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

func (i StringAnyMapInspector) Loop(src any, it Iterator, buf *[]byte, path ...string) error {
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

func (i StringAnyMapInspector) DeepEqual(l, r any) bool {
	return i.DeepEqualWithOptions(l, r, nil)
}

func (i StringAnyMapInspector) DeepEqualWithOptions(l, r any, opts *DEQOptions) bool {
	_, _, _ = l, r, opts
	// todo implement me
	return false
}

func (i StringAnyMapInspector) Unmarshal(p []byte, typ Encoding) (any, error) {
	var x map[string]any
	switch typ {
	case EncodingJSON:
		err := json.Unmarshal(p, &x)
		return x, err
	default:
		return nil, ErrUnknownEncodingType
	}
}

func (i StringAnyMapInspector) Copy(x any) (dst any, err error) {
	var buf ByteBuffer
	x_ := make(map[string]any)
	err = i.CopyTo(x, &x_, &buf)
	dst = x_
	return
}

func (i StringAnyMapInspector) CopyTo(src, dst any, buf AccumulativeBuffer) (err error) {
	var msrc, mdst map[string]any
	if err = i.indir1(&msrc, src); err != nil || msrc == nil {
		return
	}
	if err = i.indir2(&mdst, dst); err != nil || mdst == nil {
		return
	}
	for k := range mdst {
		delete(mdst, k)
	}

	return i.cpy(&mdst, msrc, buf)
}

func (i StringAnyMapInspector) Length(x any, result *int, path ...string) error {
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

func (i StringAnyMapInspector) Capacity(x any, result *int, path ...string) error {
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

func (i StringAnyMapInspector) Reset(x any, path ...string) error {
	var m map[string]any
	if err := i.indir2(&m, x); err != nil || x == nil {
		return nil
	}
	if len(path) == 0 {
		for k := range m {
			delete(m, k)
		}
		return nil
	}
	if v, ok := m[path[0]]; ok {
		if err := i.Reset(v, path[1:]...); err != nil {
			return err
		}
		m[path[0]] = v
	}
	return nil
}

func (i StringAnyMapInspector) indir(val any) (buf map[string]any, err error) {
	err = i.indir1(&buf, val)
	return
}

func (i StringAnyMapInspector) indir1(dst *map[string]any, val any) error {
	switch x := val.(type) {
	case map[string]any:
		*dst = x
	case *map[string]any:
		*dst = *x
	case **map[string]any:
		*dst = *(*x)
	default:
		return ErrUnsupportedType
	}
	return nil
}

func (i StringAnyMapInspector) indir2(dst *map[string]any, val any) error {
	switch x := val.(type) {
	case map[string]any:
		return ErrMustPointerType
	case *map[string]any:
		*dst = *x
	case **map[string]any:
		*dst = *(*x)
	default:
		return ErrUnsupportedType
	}
	return nil
}

func (i StringAnyMapInspector) cpy(dst *map[string]any, src map[string]any, buf AccumulativeBuffer) error {
	for k := range src {
		x := src[k]
		m, err := i.indir(x)
		if err == nil {
			m1 := make(map[string]any, len(m))
			if err = i.cpy(&m1, m, buf); err != nil {
				return err
			}
			switch x.(type) {
			case map[string]any:
				(*dst)[k] = m1
			case *map[string]any:
				(*dst)[k] = &m1
			case **map[string]any:
				m2 := &m1
				(*dst)[k] = &m2
			}
		} else {
			switch x1 := x.(type) {
			case string:
				(*dst)[k] = buf.BufferizeString(x1)
			case *string:
				(*dst)[k] = buf.BufferizeString(*x1)
			case []byte:
				(*dst)[k] = buf.Bufferize(x1)
			case *[]byte:
				(*dst)[k] = buf.Bufferize(*x1)
			default:
				// todo check possible pointer type of x?
				(*dst)[k] = x
			}
		}
	}
	return nil
}
