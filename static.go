package inspector

// Built-in inspector for simple types, like int, unit, ...

import (
	"bytes"
	"encoding/json"
	"strconv"

	"github.com/koykov/fastconv"
)

type StaticInspector struct {
	BaseInspector
}

func (i *StaticInspector) TypeName() string {
	return "static"
}

func (i *StaticInspector) Get(src interface{}, _ ...string) (interface{}, error) {
	return src, nil
}

func (i *StaticInspector) GetTo(src interface{}, buf *interface{}, _ ...string) error {
	*buf = src
	return nil
}

func (i *StaticInspector) Set(_, _ interface{}, _ ...string) error {
	return nil
}

func (i *StaticInspector) SetWB(_, _ interface{}, _ AccumulativeBuffer, _ ...string) error {
	return nil
}

func (i *StaticInspector) Cmp(src interface{}, cond Op, right string, result *bool, _ ...string) error {
	switch src.(type) {
	case int:
		if r, err := strconv.ParseInt(right, 0, 0); err == nil {
			*result = i.cmpInt(int64(src.(int)), cond, r)
		}
	case *int:
		if r, err := strconv.ParseInt(right, 0, 0); err == nil {
			*result = i.cmpInt(int64(*src.(*int)), cond, r)
		}
	case int8:
		if r, err := strconv.ParseInt(right, 0, 0); err == nil {
			*result = i.cmpInt(int64(src.(int8)), cond, r)
		}
	case *int8:
		if r, err := strconv.ParseInt(right, 0, 0); err == nil {
			*result = i.cmpInt(int64(*src.(*int8)), cond, r)
		}
	case int16:
		if r, err := strconv.ParseInt(right, 0, 0); err == nil {
			*result = i.cmpInt(int64(src.(int16)), cond, r)
		}
	case *int16:
		if r, err := strconv.ParseInt(right, 0, 0); err == nil {
			*result = i.cmpInt(int64(*src.(*int16)), cond, r)
		}
	case int32:
		if r, err := strconv.ParseInt(right, 0, 0); err == nil {
			*result = i.cmpInt(int64(src.(int32)), cond, r)
		}
	case *int32:
		if r, err := strconv.ParseInt(right, 0, 0); err == nil {
			*result = i.cmpInt(int64(*src.(*int32)), cond, r)
		}
	case int64:
		if r, err := strconv.ParseInt(right, 0, 0); err == nil {
			*result = i.cmpInt(src.(int64), cond, r)
		}
	case *int64:
		if r, err := strconv.ParseInt(right, 0, 0); err == nil {
			*result = i.cmpInt(*src.(*int64), cond, r)
		}
	case uint:
		if r, err := strconv.ParseUint(right, 0, 0); err == nil {
			*result = i.cmpUint(uint64(src.(uint)), cond, r)
		}
	case *uint:
		if r, err := strconv.ParseUint(right, 0, 0); err == nil {
			*result = i.cmpUint(uint64(*src.(*uint)), cond, r)
		}
	case uint8:
		if r, err := strconv.ParseUint(right, 0, 0); err == nil {
			*result = i.cmpUint(uint64(src.(uint8)), cond, r)
		}
	case *uint8:
		if r, err := strconv.ParseUint(right, 0, 0); err == nil {
			*result = i.cmpUint(uint64(*src.(*uint8)), cond, r)
		}
	case uint16:
		if r, err := strconv.ParseUint(right, 0, 0); err == nil {
			*result = i.cmpUint(uint64(src.(uint16)), cond, r)
		}
	case *uint16:
		if r, err := strconv.ParseUint(right, 0, 0); err == nil {
			*result = i.cmpUint(uint64(*src.(*uint16)), cond, r)
		}
	case uint32:
		if r, err := strconv.ParseUint(right, 0, 0); err == nil {
			*result = i.cmpUint(uint64(src.(uint32)), cond, r)
		}
	case *uint32:
		if r, err := strconv.ParseUint(right, 0, 0); err == nil {
			*result = i.cmpUint(uint64(*src.(*uint32)), cond, r)
		}
	case uint64:
		if r, err := strconv.ParseUint(right, 0, 0); err == nil {
			*result = i.cmpUint(src.(uint64), cond, r)
		}
	case *uint64:
		if r, err := strconv.ParseUint(right, 0, 0); err == nil {
			*result = i.cmpUint(*src.(*uint64), cond, r)
		}
	case float32:
		if r, err := strconv.ParseFloat(right, 0); err == nil {
			*result = i.cmpFloat(float64(src.(float32)), cond, r)
		}
	case *float32:
		if r, err := strconv.ParseFloat(right, 0); err == nil {
			*result = i.cmpFloat(float64(*src.(*float32)), cond, r)
		}
	case float64:
		if r, err := strconv.ParseFloat(right, 0); err == nil {
			*result = i.cmpFloat(src.(float64), cond, r)
		}
	case *float64:
		if r, err := strconv.ParseFloat(right, 0); err == nil {
			*result = i.cmpFloat(*src.(*float64), cond, r)
		}
	case []byte:
		*result = i.cmpBytes(src.([]byte), cond, fastconv.S2B(right))
	case *[]byte:
		*result = i.cmpBytes(*src.(*[]byte), cond, fastconv.S2B(right))
	case string:
		*result = i.cmpStr(src.(string), cond, right)
	case *string:
		*result = i.cmpStr(*src.(*string), cond, right)
	default:
		*result = false
	}
	return nil
}

func (i *StaticInspector) cmpInt(left int64, cond Op, right int64) bool {
	switch cond {
	case OpEq:
		return left == right
	case OpNq:
		return left != right
	case OpGt:
		return left > right
	case OpGtq:
		return left >= right
	case OpLt:
		return left < right
	case OpLtq:
		return left <= right
	}
	return false
}

func (i *StaticInspector) cmpUint(left uint64, cond Op, right uint64) bool {
	switch cond {
	case OpEq:
		return left == right
	case OpNq:
		return left != right
	case OpGt:
		return left > right
	case OpGtq:
		return left >= right
	case OpLt:
		return left < right
	case OpLtq:
		return left <= right
	}
	return false
}

func (i *StaticInspector) cmpFloat(left float64, cond Op, right float64) bool {
	switch cond {
	case OpEq:
		return left == right
	case OpNq:
		return left != right
	case OpGt:
		return left > right
	case OpGtq:
		return left >= right
	case OpLt:
		return left < right
	case OpLtq:
		return left <= right
	}
	return false
}

func (i *StaticInspector) cmpStr(left string, cond Op, right string) bool {
	switch cond {
	case OpEq:
		return left == right
	case OpNq:
		return left != right
	case OpGt:
		return left > right
	case OpGtq:
		return left >= right
	case OpLt:
		return left < right
	case OpLtq:
		return left <= right
	}
	return false
}

func (i *StaticInspector) cmpBytes(left []byte, cond Op, right []byte) bool {
	switch cond {
	case OpEq:
		return bytes.Equal(left, right)
	case OpNq:
		return !bytes.Equal(left, right)
	}
	return false
}

func (i *StaticInspector) Loop(_ interface{}, _ Looper, _ *[]byte, _ ...string) error {
	return nil
}

func (i *StaticInspector) DeepEqual(l, r interface{}) bool {
	return i.DeepEqualWithOptions(l, r, nil)
}

func (i *StaticInspector) DeepEqualWithOptions(l, r interface{}, _ *DEQOptions) bool {
	switch l.(type) {
	case int:
		if rx, ok := r.(int); ok {
			return rx == l.(int)
		}
	case *int:
		if rx, ok := r.(*int); ok {
			return *rx == *l.(*int)
		}
	case int8:
		if rx, ok := r.(int8); ok {
			return rx == l.(int8)
		}
	case *int8:
		if rx, ok := r.(*int8); ok {
			return *rx == *l.(*int8)
		}
	case int16:
		if rx, ok := r.(int16); ok {
			return rx == l.(int16)
		}
	case *int16:
		if rx, ok := r.(*int16); ok {
			return *rx == *l.(*int16)
		}
	case int32:
		if rx, ok := r.(int32); ok {
			return rx == l.(int32)
		}
	case *int32:
		if rx, ok := r.(*int32); ok {
			return *rx == *l.(*int32)
		}
	case int64:
		if rx, ok := r.(int64); ok {
			return rx == l.(int64)
		}
	case *int64:
		if rx, ok := r.(*int64); ok {
			return *rx == *l.(*int64)
		}
	case uint:
		if rx, ok := r.(uint); ok {
			return rx == l.(uint)
		}
	case *uint:
		if rx, ok := r.(*uint); ok {
			return *rx == *l.(*uint)
		}
	case uint8:
		if rx, ok := r.(uint8); ok {
			return rx == l.(uint8)
		}
	case *uint8:
		if rx, ok := r.(*uint8); ok {
			return *rx == *l.(*uint8)
		}
	case uint16:
		if rx, ok := r.(uint16); ok {
			return rx == l.(uint16)
		}
	case *uint16:
		if rx, ok := r.(*uint16); ok {
			return *rx == *l.(*uint16)
		}
	case uint32:
		if rx, ok := r.(uint32); ok {
			return rx == l.(uint32)
		}
	case *uint32:
		if rx, ok := r.(*uint32); ok {
			return *rx == *l.(*uint32)
		}
	case uint64:
		if rx, ok := r.(uint64); ok {
			return rx == l.(uint64)
		}
	case *uint64:
		if rx, ok := r.(*uint64); ok {
			return *rx == *l.(*uint64)
		}
	case float32:
		if rx, ok := r.(float32); ok {
			return rx == l.(float32)
		}
	case *float32:
		if rx, ok := r.(*float32); ok {
			return *rx == *l.(*float32)
		}
	case float64:
		if rx, ok := r.(float64); ok {
			return rx == l.(float64)
		}
	case *float64:
		if rx, ok := r.(*float64); ok {
			return *rx == *l.(*float64)
		}
	case []byte:
		if rx, ok := r.([]byte); ok {
			return bytes.Equal(rx, l.([]byte))
		}
	case *[]byte:
		if rx, ok := r.(*[]byte); ok {
			return bytes.Equal(*rx, *l.(*[]byte))
		}
	case string:
		if rx, ok := r.(string); ok {
			return rx == l.(string)
		}
	case *string:
		if rx, ok := r.(*string); ok {
			return *rx == *l.(*string)
		}
	}
	return false
}

func (i *StaticInspector) Unmarshal(p []byte, typ Encoding) (interface{}, error) {
	var x interface{}
	switch typ {
	case EncodingJSON:
		err := json.Unmarshal(p, &x)
		return x, err
	default:
		return nil, ErrUnknownEncodingType
	}
}
