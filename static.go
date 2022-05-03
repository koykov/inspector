package inspector

// Built-in inspector for simple types, like int, unit, ...

import (
	"bytes"
	"encoding/json"
	"math"
	"strconv"

	"github.com/koykov/fastconv"
)

const (
	eqlf64 = 1e-3
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
	case bool:
		if rx, ok := i.indBool(r); ok {
			return rx == l.(bool)
		}
	case *bool:
		if rx, ok := i.indBool(r); ok {
			return rx == *l.(*bool)
		}
	case int:
		if rx, ok := i.indInt(r); ok {
			return rx == int64(l.(int))
		}
	case *int:
		if rx, ok := i.indInt(r); ok {
			return rx == int64(*l.(*int))
		}
	case int8:
		if rx, ok := i.indInt(r); ok {
			return rx == int64(l.(int8))
		}
	case *int8:
		if rx, ok := i.indInt(r); ok {
			return rx == int64(*l.(*int8))
		}
	case int16:
		if rx, ok := i.indInt(r); ok {
			return rx == int64(l.(int16))
		}
	case *int16:
		if rx, ok := i.indInt(r); ok {
			return rx == int64(*l.(*int16))
		}
	case int32:
		if rx, ok := i.indInt(r); ok {
			return rx == int64(l.(int32))
		}
	case *int32:
		if rx, ok := i.indInt(r); ok {
			return rx == int64(*l.(*int32))
		}
	case int64:
		if rx, ok := i.indInt(r); ok {
			return rx == l.(int64)
		}
	case *int64:
		if rx, ok := i.indInt(r); ok {
			return rx == *l.(*int64)
		}
	case uint:
		if rx, ok := i.indUint(r); ok {
			return rx == uint64(l.(uint))
		}
	case *uint:
		if rx, ok := i.indUint(r); ok {
			return rx == uint64(*l.(*uint))
		}
	case uint8:
		if rx, ok := i.indUint(r); ok {
			return rx == uint64(l.(uint8))
		}
	case *uint8:
		if rx, ok := i.indUint(r); ok {
			return rx == uint64(*l.(*uint8))
		}
	case uint16:
		if rx, ok := i.indUint(r); ok {
			return rx == uint64(l.(uint16))
		}
	case *uint16:
		if rx, ok := i.indUint(r); ok {
			return rx == uint64(*l.(*uint16))
		}
	case uint32:
		if rx, ok := i.indUint(r); ok {
			return rx == uint64(l.(uint32))
		}
	case *uint32:
		if rx, ok := i.indUint(r); ok {
			return rx == uint64(*l.(*uint32))
		}
	case uint64:
		if rx, ok := i.indUint(r); ok {
			return rx == l.(uint64)
		}
	case *uint64:
		if rx, ok := i.indUint(r); ok {
			return rx == *l.(*uint64)
		}
	case float32:
		if rx, ok := i.indFloat(r); ok {
			return i.eqlf64(rx, float64(l.(float32)))
		}
	case *float32:
		if rx, ok := i.indFloat(r); ok {
			return i.eqlf64(rx, float64(*l.(*float32)))
		}
	case float64:
		if rx, ok := i.indFloat(r); ok {
			return i.eqlf64(rx, l.(float64))
		}
	case *float64:
		if rx, ok := i.indFloat(r); ok {
			return i.eqlf64(rx, *l.(*float64))
		}
	case []byte:
		if rx, ok := i.indBytes(r); ok {
			return bytes.Equal(rx, l.([]byte))
		}
	case *[]byte:
		if rx, ok := i.indBytes(r); ok {
			return bytes.Equal(rx, *l.(*[]byte))
		}
	case string:
		if rx, ok := i.indString(r); ok {
			return rx == l.(string)
		}
	case *string:
		if rx, ok := i.indString(r); ok {
			return rx == *l.(*string)
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

func (i *StaticInspector) Copy(x interface{}) (interface{}, error) {
	var t interface{}
	switch x.(type) {
	case *bool:
		t = *x.(*bool)
	case *int:
		t = *x.(*int)
	case *int8:
		t = *x.(*int8)
	case *int16:
		t = *x.(*int16)
	case *int32:
		t = *x.(*int32)
	case *int64:
		t = *x.(*int64)
	case *uint:
		t = *x.(*uint)
	case *uint8:
		t = *x.(*uint8)
	case *uint16:
		t = *x.(*uint16)
	case *uint32:
		t = *x.(*uint32)
	case *uint64:
		t = *x.(*uint64)
	case *float32:
		t = *x.(*float32)
	case *float64:
		t = *x.(*float64)
	case []byte:
		p := x.([]byte)
		t = append([]byte(nil), p...)
	case *[]byte:
		p := *x.(*[]byte)
		t = append([]byte(nil), p...)
	case *string:
		t = *x.(*string)
	default:
		t = x
	}
	return t, nil
}

func (i *StaticInspector) indBool(x interface{}) (bool, bool) {
	switch x.(type) {
	case bool:
		return x.(bool), true
	case *bool:
		return *x.(*bool), true
	}
	return false, false
}

func (i *StaticInspector) indInt(x interface{}) (int64, bool) {
	switch x.(type) {
	case int:
		return int64(x.(int)), true
	case *int:
		return int64(*x.(*int)), true
	case int8:
		return int64(x.(int8)), true
	case *int8:
		return int64(*x.(*int8)), true
	case int16:
		return int64(x.(int16)), true
	case *int16:
		return int64(*x.(*int16)), true
	case int32:
		return int64(x.(int32)), true
	case *int32:
		return int64(*x.(*int32)), true
	case int64:
		return x.(int64), true
	case *int64:
		return *x.(*int64), true
	case float32:
		return int64(x.(float32)), true
	case *float32:
		return int64(*x.(*float32)), true
	case float64:
		return int64(x.(float64)), true
	case *float64:
		return int64(*x.(*float64)), true
	}
	return 0, false
}

func (i *StaticInspector) indUint(x interface{}) (uint64, bool) {
	switch x.(type) {
	case uint:
		return uint64(x.(uint)), true
	case *uint:
		return uint64(*x.(*uint)), true
	case uint8:
		return uint64(x.(uint8)), true
	case *uint8:
		return uint64(*x.(*uint8)), true
	case uint16:
		return uint64(x.(uint16)), true
	case *uint16:
		return uint64(*x.(*uint16)), true
	case uint32:
		return uint64(x.(uint32)), true
	case *uint32:
		return uint64(*x.(*uint32)), true
	case uint64:
		return x.(uint64), true
	case *uint64:
		return *x.(*uint64), true
	case float32:
		return uint64(x.(float32)), true
	case *float32:
		return uint64(*x.(*float32)), true
	case float64:
		return uint64(x.(float64)), true
	case *float64:
		return uint64(*x.(*float64)), true
	}
	return 0, false
}

func (i *StaticInspector) indFloat(x interface{}) (float64, bool) {
	switch x.(type) {
	case float32:
		return float64(x.(float32)), true
	case *float32:
		return float64(*x.(*float32)), true
	case float64:
		return x.(float64), true
	case *float64:
		return *x.(*float64), true
	default:
		switch x.(type) {
		case int:
			return float64(x.(int)), true
		case *int:
			return float64(*x.(*int)), true
		case int8:
			return float64(x.(int8)), true
		case *int8:
			return float64(*x.(*int8)), true
		case int16:
			return float64(x.(int16)), true
		case *int16:
			return float64(*x.(*int16)), true
		case int32:
			return float64(x.(int32)), true
		case *int32:
			return float64(*x.(*int32)), true
		case int64:
			return float64(x.(int64)), true
		case *int64:
			return float64(*x.(*int64)), true
		case uint:
			return float64(x.(uint)), true
		case *uint:
			return float64(*x.(*uint)), true
		case uint8:
			return float64(x.(uint8)), true
		case *uint8:
			return float64(*x.(*uint8)), true
		case uint16:
			return float64(x.(uint16)), true
		case *uint16:
			return float64(*x.(*uint16)), true
		case uint32:
			return float64(x.(uint32)), true
		case *uint32:
			return float64(*x.(*uint32)), true
		case uint64:
			return float64(x.(uint64)), true
		case *uint64:
			return float64(*x.(*uint64)), true
		}
	}
	return 0, false
}

func (i *StaticInspector) indString(x interface{}) (string, bool) {
	switch x.(type) {
	case string:
		return x.(string), true
	case *string:
		return *x.(*string), true
	}
	if b, ok := i.indBytes(x); ok {
		return fastconv.B2S(b), true
	}
	return "", false
}

func (i *StaticInspector) indBytes(x interface{}) ([]byte, bool) {
	switch x.(type) {
	case []byte:
		return x.([]byte), true
	case *[]byte:
		return *x.(*[]byte), true
	}
	if s, ok := i.indString(x); ok {
		return fastconv.S2B(s), true
	}
	return nil, false
}

func (i *StaticInspector) eqlf64(a, b float64) bool {
	return math.Abs(a-b) <= eqlf64
}
