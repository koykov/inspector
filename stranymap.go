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

func (i StringAnyMapInspector) Compare(src any, cond Op, right string, result *bool, path ...string) error {
	_, _, _, _, _ = src, cond, right, result, path
	return nil
}

func (i StringAnyMapInspector) Loop(src any, l Iterator, buf *[]byte, path ...string) error {
	_, _, _, _ = src, l, buf, path
	return nil
}

func (i StringAnyMapInspector) DeepEqual(l, r any) bool {
	return i.DeepEqualWithOptions(l, r, nil)
}

func (i StringAnyMapInspector) DeepEqualWithOptions(l, r any, opts *DEQOptions) bool {
	_, _, _ = l, r, opts
	return false
}

func (i StringAnyMapInspector) Unmarshal(p []byte, typ Encoding) (any, error) {
	var x any
	switch typ {
	case EncodingJSON:
		err := json.Unmarshal(p, &x)
		return x, err
	default:
		return nil, ErrUnknownEncodingType
	}
}

func (i StringAnyMapInspector) Copy(x any) (dst any, err error) {
	_ = x
	return
}

func (i StringAnyMapInspector) CopyTo(src, dst any, buf AccumulativeBuffer) error {
	_, _, _ = src, dst, buf
	return nil
}

func (i StringAnyMapInspector) Length(x any, result *int, path ...string) error {
	_, _, _ = x, result, path
	return nil
}

func (i StringAnyMapInspector) Capacity(x any, result *int, path ...string) error {
	_, _, _ = x, result, path
	return nil
}

func (i StringAnyMapInspector) Reset(x any) error {
	_ = x
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
