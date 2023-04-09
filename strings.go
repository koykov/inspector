package inspector

import (
	"strconv"

	"github.com/koykov/fastconv"
)

// StringsInspector is a built-in inspector for []string/[][]byte types.
type StringsInspector struct{}

func (i StringsInspector) TypeName() string {
	return "strings"
}

func (i StringsInspector) Get(src any, path ...string) (any, error) {
	var x any
	err := i.GetTo(src, &x, path...)
	return x, err
}

func (i StringsInspector) GetTo(src any, buf *any, path ...string) error {
	if len(path) != 1 {
		return nil
	}
	var (
		ss []string
		pp [][]byte
	)
	switch src.(type) {
	case []string:
		ss = src.([]string)
	case *[]string:
		ss = *(src.(*[]string))
	case [][]byte:
		pp = src.([][]byte)
	case *[][]byte:
		pp = *(src.(*[][]byte))
	default:
		return nil
	}
	idx, err := strconv.Atoi(path[0])
	if err != nil {
		return err
	}
	if idx < 0 {
		return nil
	}
	switch {
	case len(ss) > 0 && idx < len(ss):
		*buf = &ss[idx]
	case len(pp) > 0 && idx < len(pp):
		*buf = &pp[idx]
	}
	return nil
}

func (i StringsInspector) Set(dst, value any, path ...string) error {
	var buf ByteBuffer
	return i.SetWithBuffer(dst, value, &buf, path...)
}

func (i StringsInspector) SetWithBuffer(dst, value any, buf AccumulativeBuffer, path ...string) error {
	if len(path) != 1 {
		return nil
	}
	var (
		ss []string
		pp [][]byte
	)
	switch dst.(type) {
	case []string:
		ss = dst.([]string)
	case *[]string:
		ss = *(dst.(*[]string))
	case [][]byte:
		pp = dst.([][]byte)
	case *[][]byte:
		pp = *(dst.(*[][]byte))
	default:
		return nil
	}
	idx, err := strconv.Atoi(path[0])
	if err != nil {
		return err
	}
	if idx < 0 {
		return nil
	}
	var p []byte
	switch {
	case len(ss) > 0 && idx < len(ss):
		switch value.(type) {
		case string:
			p = fastconv.S2B(value.(string))
		case *string:
			p = fastconv.S2B(*value.(*string))
		}
		if len(p) > 0 {
			bp := buf.AcquireBytes()
			off := len(bp)
			bp = append(bp, p...)
			buf.ReleaseBytes(bp)
			ss[idx] = fastconv.B2S(bp[off:])
		}
	case len(pp) > 0 && idx < len(pp):
		switch value.(type) {
		case []byte:
			p = value.([]byte)
		case *[]byte:
			p = *value.(*[]byte)
		}
		if len(p) > 0 {
			bp := buf.AcquireBytes()
			off := len(bp)
			bp = append(bp, p...)
			buf.ReleaseBytes(bp)
			pp[idx] = bp[off:]
		}
	}
	return nil
}

func (i StringsInspector) Compare(src any, cond Op, right string, result *bool, path ...string) error {
	if len(path) != 1 {
		return nil
	}
	var (
		ss []string
		pp [][]byte
	)
	switch src.(type) {
	case []string:
		ss = src.([]string)
	case *[]string:
		ss = *(src.(*[]string))
	case [][]byte:
		pp = src.([][]byte)
	case *[][]byte:
		pp = *(src.(*[][]byte))
	default:
		return nil
	}
	idx, err := strconv.Atoi(path[0])
	if err != nil {
		return err
	}
	if idx < 0 {
		return nil
	}
	var s string
	switch {
	case len(ss) > 0 && idx < len(ss):
		s = ss[idx]
	case len(pp) > 0 && idx < len(pp):
		s = fastconv.B2S(pp[idx])
	}
	switch cond {
	case OpNq:
		*result = s != right
	case OpEq:
		*result = s == right
	case OpGt:
		*result = s > right
	case OpGtq:
		*result = s >= right
	case OpLt:
		*result = s < right
	case OpLtq:
		*result = s <= right
	}
	return nil
}

func (i StringsInspector) Loop(src any, l Iterator, buf *[]byte, path ...string) error {
	_, _, _, _ = src, l, buf, path
	// todo: implement me
	return nil
}

func (i StringsInspector) DeepEqual(l, r any) bool {
	_, _ = l, r
	// todo: implement me
	return false
}

func (i StringsInspector) DeepEqualWithOptions(l, r any, options *DEQOptions) bool {
	_, _, _ = l, r, options
	// todo: implement me
	return false
}

func (i StringsInspector) Unmarshal(p []byte, typ Encoding) (any, error) {
	_, _ = p, typ
	// todo: implement me
	return nil, nil
}

func (i StringsInspector) Copy(x any) (any, error) {
	_ = x
	// todo: implement me
	return nil, nil
}

func (i StringsInspector) CopyTo(src, dst any, buf AccumulativeBuffer) error {
	_, _, _ = src, dst, buf
	// todo: implement me
	return nil
}

func (i StringsInspector) Length(src any, result *int, path ...string) error {
	_, _, _ = src, result, path
	// todo: implement me
	return nil
}

func (i StringsInspector) Capacity(src any, result *int, path ...string) error {
	_, _, _ = src, result, path
	// todo: implement me
	return nil
}

func (i StringsInspector) Reset(x any) error {
	_ = x
	// todo: implement me
	return nil
}
