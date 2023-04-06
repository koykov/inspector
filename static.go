package inspector

// Built-in inspector for simple types, like int, unit, ...

import (
	"bytes"
	"encoding/json"
	"math"
	"strconv"

	"github.com/koykov/fastconv"
)

type StaticInspector struct {
	BaseInspector
}

func (i StaticInspector) TypeName() string {
	return "static"
}

func (i StaticInspector) Get(src any, _ ...string) (any, error) {
	return src, nil
}

func (i StaticInspector) GetTo(src any, buf *any, _ ...string) error {
	*buf = src
	return nil
}

func (i StaticInspector) Set(_, _ any, _ ...string) error {
	return nil
}

func (i StaticInspector) SetWB(_, _ any, _ AccumulativeBuffer, _ ...string) error {
	return nil
}

func (i StaticInspector) Cmp(src any, cond Op, right string, result *bool, _ ...string) error {
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

func (i StaticInspector) cmpInt(left int64, cond Op, right int64) bool {
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

func (i StaticInspector) cmpUint(left uint64, cond Op, right uint64) bool {
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

func (i StaticInspector) cmpFloat(left float64, cond Op, right float64) bool {
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

func (i StaticInspector) cmpStr(left string, cond Op, right string) bool {
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

func (i StaticInspector) cmpBytes(left []byte, cond Op, right []byte) bool {
	switch cond {
	case OpEq:
		return bytes.Equal(left, right)
	case OpNq:
		return !bytes.Equal(left, right)
	}
	return false
}

func (i StaticInspector) Loop(_ any, _ Looper, _ *[]byte, _ ...string) error {
	return nil
}

func (i StaticInspector) DeepEqual(l, r any) bool {
	return i.DeepEqualWithOptions(l, r, nil)
}

func (i StaticInspector) DeepEqualWithOptions(l, r any, _ *DEQOptions) bool {
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

func (i StaticInspector) Unmarshal(p []byte, typ Encoding) (any, error) {
	var x any
	switch typ {
	case EncodingJSON:
		err := json.Unmarshal(p, &x)
		return x, err
	default:
		return nil, ErrUnknownEncodingType
	}
}

func (i StaticInspector) Copy(x any) (dst any, err error) {
	switch x.(type) {
	case bool:
		dst = x.(bool)
	case *bool:
		dst = *x.(*bool)
	case int:
		dst = x.(int)
	case *int:
		dst = *x.(*int)
	case int8:
		dst = x.(int8)
	case *int8:
		dst = *x.(*int8)
	case int16:
		dst = x.(int16)
	case *int16:
		dst = *x.(*int16)
	case int32:
		dst = x.(int32)
	case *int32:
		dst = *x.(*int32)
	case int64:
		dst = x.(int64)
	case *int64:
		dst = *x.(*int64)
	case uint:
		dst = x.(uint)
	case *uint:
		dst = *x.(*uint)
	case uint8:
		dst = x.(uint8)
	case *uint8:
		dst = *x.(*uint8)
	case uint16:
		dst = x.(uint16)
	case *uint16:
		dst = *x.(*uint16)
	case uint32:
		dst = x.(uint32)
	case *uint32:
		dst = *x.(*uint32)
	case uint64:
		dst = x.(uint64)
	case *uint64:
		dst = *x.(*uint64)
	case float32:
		dst = x.(float32)
	case *float32:
		dst = *x.(*float32)
	case float64:
		dst = x.(float64)
	case *float64:
		dst = *x.(*float64)
	case []byte:
		origin := x.([]byte)
		cpy := append([]byte(nil), origin...)
		dst = cpy
	case *[]byte:
		origin := *x.(*[]byte)
		cpy := append([]byte(nil), origin...)
		dst = cpy
	case string:
		origin := x.(string)
		cpy := append([]byte(nil), origin...)
		dst = fastconv.B2S(cpy)
	case *string:
		origin := *x.(*string)
		cpy := append([]byte(nil), origin...)
		dst = fastconv.B2S(cpy)
	default:
		return nil, ErrUnsupportedType
	}
	return
}

func (i StaticInspector) CopyTo(src, dst any, buf AccumulativeBuffer) error {
	switch src.(type) {
	case bool:
		if _, ok := dst.(*bool); !ok {
			return ErrMustPointerType
		}
		*dst.(*bool) = src.(bool)
	case *bool:
		if _, ok := dst.(*bool); !ok {
			return ErrMustPointerType
		}
		*dst.(*bool) = *src.(*bool)
	case int:
		if _, ok := dst.(*int); !ok {
			return ErrMustPointerType
		}
		*dst.(*int) = src.(int)
	case *int:
		if _, ok := dst.(*int); !ok {
			return ErrMustPointerType
		}
		*dst.(*int) = *src.(*int)
	case int8:
		if _, ok := dst.(*int8); !ok {
			return ErrMustPointerType
		}
		*dst.(*int8) = src.(int8)
	case *int8:
		if _, ok := dst.(*int8); !ok {
			return ErrMustPointerType
		}
		*dst.(*int8) = *src.(*int8)
	case int16:
		if _, ok := dst.(*int16); !ok {
			return ErrMustPointerType
		}
		*dst.(*int16) = src.(int16)
	case *int16:
		if _, ok := dst.(*int16); !ok {
			return ErrMustPointerType
		}
		*dst.(*int16) = *src.(*int16)
	case int32:
		if _, ok := dst.(*int32); !ok {
			return ErrMustPointerType
		}
		*dst.(*int32) = src.(int32)
	case *int32:
		if _, ok := dst.(*int32); !ok {
			return ErrMustPointerType
		}
		*dst.(*int32) = *src.(*int32)
	case int64:
		if _, ok := dst.(*int64); !ok {
			return ErrMustPointerType
		}
		*dst.(*int64) = src.(int64)
	case *int64:
		if _, ok := dst.(*int64); !ok {
			return ErrMustPointerType
		}
		*dst.(*int64) = *src.(*int64)
	case uint:
		if _, ok := dst.(*uint); !ok {
			return ErrMustPointerType
		}
		*dst.(*uint) = src.(uint)
	case *uint:
		if _, ok := dst.(*uint); !ok {
			return ErrMustPointerType
		}
		*dst.(*uint) = *src.(*uint)
	case uint8:
		if _, ok := dst.(*uint8); !ok {
			return ErrMustPointerType
		}
		*dst.(*uint8) = src.(uint8)
	case *uint8:
		if _, ok := dst.(*uint8); !ok {
			return ErrMustPointerType
		}
		*dst.(*uint8) = *src.(*uint8)
	case uint16:
		if _, ok := dst.(*uint16); !ok {
			return ErrMustPointerType
		}
		*dst.(*uint16) = src.(uint16)
	case *uint16:
		if _, ok := dst.(*uint16); !ok {
			return ErrMustPointerType
		}
		*dst.(*uint16) = *src.(*uint16)
	case uint32:
		if _, ok := dst.(*uint32); !ok {
			return ErrMustPointerType
		}
		*dst.(*uint32) = src.(uint32)
	case *uint32:
		if _, ok := dst.(*uint32); !ok {
			return ErrMustPointerType
		}
		*dst.(*uint32) = *src.(*uint32)
	case uint64:
		if _, ok := dst.(*uint64); !ok {
			return ErrMustPointerType
		}
		*dst.(*uint64) = src.(uint64)
	case *uint64:
		if _, ok := dst.(*uint64); !ok {
			return ErrMustPointerType
		}
		*dst.(*uint64) = *src.(*uint64)
	case float32:
		if _, ok := dst.(*float32); !ok {
			return ErrMustPointerType
		}
		*dst.(*float32) = src.(float32)
	case *float32:
		if _, ok := dst.(*float32); !ok {
			return ErrMustPointerType
		}
		*dst.(*float32) = *src.(*float32)
	case float64:
		if _, ok := dst.(*float64); !ok {
			return ErrMustPointerType
		}
		*dst.(*float64) = src.(float64)
	case *float64:
		if _, ok := dst.(*float64); !ok {
			return ErrMustPointerType
		}
		*dst.(*float64) = *src.(*float64)
	case []byte:
		p, ok := dst.(*[]byte)
		if !ok {
			return ErrMustPointerType
		}
		origin := src.([]byte)
		buf1 := buf.AcquireBytes()
		off := len(buf1)
		buf1 = append(buf1, origin...)
		buf.ReleaseBytes(buf1)
		*p = buf1[off:]
	case *[]byte:
		p, ok := dst.(*[]byte)
		if !ok {
			return ErrMustPointerType
		}
		origin := src.([]byte)
		buf1 := buf.AcquireBytes()
		off := len(buf1)
		buf1 = append(buf1, origin...)
		buf.ReleaseBytes(buf1)
		*p = buf1[off:]
	case string:
		p, ok := dst.(*string)
		if !ok {
			return ErrMustPointerType
		}
		origin := src.(string)
		buf1 := buf.AcquireBytes()
		off := len(buf1)
		buf1 = append(buf1, origin...)
		buf.ReleaseBytes(buf1)
		*p = fastconv.B2S(buf1[off:])
	case *string:
		p, ok := dst.(*string)
		if !ok {
			return ErrMustPointerType
		}
		origin := *src.(*string)
		buf1 := buf.AcquireBytes()
		off := len(buf1)
		buf1 = append(buf1, origin...)
		buf.ReleaseBytes(buf1)
		*p = fastconv.B2S(buf1[off:])
	default:
		return ErrUnsupportedType
	}
	return nil
}

func (i StaticInspector) Reset(x any) error {
	switch x.(type) {
	case bool:
		x = false
	case *bool:
		*x.(*bool) = false
	case int:
		x = 0
	case *int:
		*x.(*int) = 0
	case *int8:
		*x.(*int8) = 0
	case int8:
		x = 0
	case *int16:
		*x.(*int16) = 0
	case int16:
		x = 0
	case *int32:
		*x.(*int32) = 0
	case int32:
		x = 0
	case *int64:
		*x.(*int64) = 0
	case int64:
		x = 0
	case *uint:
		*x.(*uint) = 0
	case uint:
		x = 0
	case *uint8:
		*x.(*uint8) = 0
	case uint8:
		x = 0
	case *uint16:
		*x.(*uint16) = 0
	case uint16:
		x = 0
	case *uint32:
		*x.(*uint32) = 0
	case uint32:
		x = 0
	case *uint64:
		*x.(*uint64) = 0
	case uint64:
		x = 0
	case *float32:
		*x.(*float32) = 0
	case float32:
		x = 0
	case *float64:
		*x.(*float64) = 0
	case float64:
		x = 0
	case []byte:
		p := x.([]byte)
		x = p[:0]
	case *[]byte:
		p := *x.(*[]byte)
		p = p[:0]
		x = &p
	case string:
		x = ""
	case *string:
		var s string
		x = &s
	}
	return nil
}

func (i StaticInspector) indBool(x any) (bool, bool) {
	switch x.(type) {
	case bool:
		return x.(bool), true
	case *bool:
		return *x.(*bool), true
	}
	return false, false
}

func (i StaticInspector) indInt(x any) (int64, bool) {
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

func (i StaticInspector) indUint(x any) (uint64, bool) {
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

func (i StaticInspector) indFloat(x any) (float64, bool) {
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

func (i StaticInspector) indString(x any) (string, bool) {
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

func (i StaticInspector) indBytes(x any) ([]byte, bool) {
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

func (i StaticInspector) eqlf64(a, b float64) bool {
	return math.Abs(a-b) <= FloatPrecision
}
