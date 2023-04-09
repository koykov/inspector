package inspector

import "strconv"

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
	_, _, _ = dst, value, path
	// todo: implement me
	return nil
}

func (i StringsInspector) SetWithBuffer(dst, value any, buf AccumulativeBuffer, path ...string) error {
	_, _, _, _ = dst, value, buf, path
	// todo: implement me
	return nil
}

func (i StringsInspector) Compare(src any, cond Op, right string, result *bool, path ...string) error {
	_, _, _, _, _ = src, cond, right, result, path
	// todo: implement me
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
