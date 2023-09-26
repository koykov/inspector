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
	_, _ = src, path
	return nil, nil
}

func (i StringAnyMapInspector) GetTo(src any, buf *any, path ...string) error {
	_, _, _ = src, buf, path
	return nil
}

func (i StringAnyMapInspector) Set(dst, value any, path ...string) error {
	_, _, _ = dst, value, path
	return nil
}

func (i StringAnyMapInspector) SetWithBuffer(dst, value any, buf AccumulativeBuffer, path ...string) error {
	_, _, _, _ = dst, value, buf, path
	return nil
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
