package inspector

// Built-in inspector for simple types, like int, unit, ...

import (
	"bytes"
	"encoding/json"
	"math"
	"strconv"

	"github.com/koykov/byteconv"
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

func (i StaticInspector) Set(dst, value any, path ...string) error {
	var buf ByteBuffer
	return i.SetWithBuffer(dst, value, &buf, path...)
}

func (i StaticInspector) SetWithBuffer(dst, value any, buf AccumulativeBuffer, path ...string) error {
	switch lx := dst.(type) {
	case bool, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, string, []byte:
		return ErrMustPointerType
	case *bool:
		if rx, ok := i.indBool(value); ok {
			*lx = rx
			return nil
		}
	case *int:
		if rx, ok := i.indInt(value); ok {
			*lx = int(rx)
			return nil
		}
	case *int8:
		if rx, ok := i.indInt(value); ok {
			*lx = int8(rx)
			return nil
		}
	case *int16:
		if rx, ok := i.indInt(value); ok {
			*lx = int16(rx)
			return nil
		}
	case *int32:
		if rx, ok := i.indInt(value); ok {
			*lx = int32(rx)
			return nil
		}
	case *int64:
		if rx, ok := i.indInt(value); ok {
			*lx = rx
			return nil
		}
	case *uint:
		if rx, ok := i.indUint(value); ok {
			*lx = uint(rx)
			return nil
		}
	case *uint8:
		if rx, ok := i.indUint(value); ok {
			*lx = uint8(rx)
			return nil
		}
	case *uint16:
		if rx, ok := i.indUint(value); ok {
			*lx = uint16(rx)
			return nil
		}
	case *uint32:
		if rx, ok := i.indUint(value); ok {
			*lx = uint32(rx)
			return nil
		}
	case *uint64:
		if rx, ok := i.indUint(value); ok {
			*lx = rx
			return nil
		}
	case *float32:
		if rx, ok := i.indFloat(value); ok {
			*lx = float32(rx)
			return nil
		}
	case *float64:
		if rx, ok := i.indFloat(value); ok {
			*lx = float64(rx)
			return nil
		}
	case *[]byte:
		if rx, ok := i.indBytes(value); ok {
			*lx = buf.Bufferize(rx)
			return nil
		}
	case *string:
		if rx, ok := i.indString(value); ok {
			*lx = buf.BufferizeString(rx)
			return nil
		}
	}
	return nil
}

func (i StaticInspector) Compare(src any, cond Op, right string, result *bool, _ ...string) error {
	switch x := src.(type) {
	case int:
		if r, err := strconv.ParseInt(right, 0, 0); err == nil {
			*result = i.cmpInt(int64(x), cond, r)
		}
	case *int:
		if r, err := strconv.ParseInt(right, 0, 0); err == nil {
			*result = i.cmpInt(int64(*x), cond, r)
		}
	case int8:
		if r, err := strconv.ParseInt(right, 0, 0); err == nil {
			*result = i.cmpInt(int64(x), cond, r)
		}
	case *int8:
		if r, err := strconv.ParseInt(right, 0, 0); err == nil {
			*result = i.cmpInt(int64(*x), cond, r)
		}
	case int16:
		if r, err := strconv.ParseInt(right, 0, 0); err == nil {
			*result = i.cmpInt(int64(x), cond, r)
		}
	case *int16:
		if r, err := strconv.ParseInt(right, 0, 0); err == nil {
			*result = i.cmpInt(int64(*x), cond, r)
		}
	case int32:
		if r, err := strconv.ParseInt(right, 0, 0); err == nil {
			*result = i.cmpInt(int64(x), cond, r)
		}
	case *int32:
		if r, err := strconv.ParseInt(right, 0, 0); err == nil {
			*result = i.cmpInt(int64(*x), cond, r)
		}
	case int64:
		if r, err := strconv.ParseInt(right, 0, 0); err == nil {
			*result = i.cmpInt(x, cond, r)
		}
	case *int64:
		if r, err := strconv.ParseInt(right, 0, 0); err == nil {
			*result = i.cmpInt(*x, cond, r)
		}
	case uint:
		if r, err := strconv.ParseUint(right, 0, 0); err == nil {
			*result = i.cmpUint(uint64(x), cond, r)
		}
	case *uint:
		if r, err := strconv.ParseUint(right, 0, 0); err == nil {
			*result = i.cmpUint(uint64(*x), cond, r)
		}
	case uint8:
		if r, err := strconv.ParseUint(right, 0, 0); err == nil {
			*result = i.cmpUint(uint64(x), cond, r)
		}
	case *uint8:
		if r, err := strconv.ParseUint(right, 0, 0); err == nil {
			*result = i.cmpUint(uint64(*x), cond, r)
		}
	case uint16:
		if r, err := strconv.ParseUint(right, 0, 0); err == nil {
			*result = i.cmpUint(uint64(x), cond, r)
		}
	case *uint16:
		if r, err := strconv.ParseUint(right, 0, 0); err == nil {
			*result = i.cmpUint(uint64(*x), cond, r)
		}
	case uint32:
		if r, err := strconv.ParseUint(right, 0, 0); err == nil {
			*result = i.cmpUint(uint64(x), cond, r)
		}
	case *uint32:
		if r, err := strconv.ParseUint(right, 0, 0); err == nil {
			*result = i.cmpUint(uint64(*x), cond, r)
		}
	case uint64:
		if r, err := strconv.ParseUint(right, 0, 0); err == nil {
			*result = i.cmpUint(x, cond, r)
		}
	case *uint64:
		if r, err := strconv.ParseUint(right, 0, 0); err == nil {
			*result = i.cmpUint(*x, cond, r)
		}
	case float32:
		if r, err := strconv.ParseFloat(right, 0); err == nil {
			*result = i.cmpFloat(float64(x), cond, r)
		}
	case *float32:
		if r, err := strconv.ParseFloat(right, 0); err == nil {
			*result = i.cmpFloat(float64(*x), cond, r)
		}
	case float64:
		if r, err := strconv.ParseFloat(right, 0); err == nil {
			*result = i.cmpFloat(x, cond, r)
		}
	case *float64:
		if r, err := strconv.ParseFloat(right, 0); err == nil {
			*result = i.cmpFloat(*x, cond, r)
		}
	case bool:
		if r, err := strconv.ParseBool(right); err == nil {
			*result = i.cmpBool(x, cond, r)
		}
	case *bool:
		if r, err := strconv.ParseBool(right); err == nil {
			*result = i.cmpBool(*x, cond, r)
		}
	case []byte:
		*result = i.cmpBytes(x, cond, byteconv.S2B(right))
	case *[]byte:
		*result = i.cmpBytes(*x, cond, byteconv.S2B(right))
	case string:
		*result = i.cmpStr(x, cond, right)
	case *string:
		*result = i.cmpStr(*x, cond, right)
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
	default:
		return false
	}
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
	default:
		return false
	}
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
	default:
		return false
	}
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
	default:
		return false
	}
}

func (i StaticInspector) cmpBytes(left []byte, cond Op, right []byte) bool {
	switch cond {
	case OpEq:
		return bytes.Equal(left, right)
	case OpNq:
		return !bytes.Equal(left, right)
	default:
		return false
	}
}

func (i StaticInspector) cmpBool(left bool, cond Op, right bool) bool {
	switch cond {
	case OpEq:
		return left == right
	case OpNq:
		return left != right
	default:
		return false
	}
}

func (i StaticInspector) Loop(_ any, _ Iterator, _ *[]byte, _ ...string) error {
	return nil
}

func (i StaticInspector) DeepEqual(l, r any) bool {
	return i.DeepEqualWithOptions(l, r, nil)
}

func (i StaticInspector) DeepEqualWithOptions(l, r any, _ *DEQOptions) bool {
	switch lx := l.(type) {
	case bool:
		if rx, ok := i.indBool(r); ok {
			return rx == lx
		}
	case *bool:
		if rx, ok := i.indBool(r); ok {
			return rx == *lx
		}
	case int:
		if rx, ok := i.indInt(r); ok {
			return rx == int64(lx)
		}
	case *int:
		if rx, ok := i.indInt(r); ok {
			return rx == int64(*lx)
		}
	case int8:
		if rx, ok := i.indInt(r); ok {
			return rx == int64(lx)
		}
	case *int8:
		if rx, ok := i.indInt(r); ok {
			return rx == int64(*lx)
		}
	case int16:
		if rx, ok := i.indInt(r); ok {
			return rx == int64(lx)
		}
	case *int16:
		if rx, ok := i.indInt(r); ok {
			return rx == int64(*lx)
		}
	case int32:
		if rx, ok := i.indInt(r); ok {
			return rx == int64(lx)
		}
	case *int32:
		if rx, ok := i.indInt(r); ok {
			return rx == int64(*lx)
		}
	case int64:
		if rx, ok := i.indInt(r); ok {
			return rx == lx
		}
	case *int64:
		if rx, ok := i.indInt(r); ok {
			return rx == *lx
		}
	case uint:
		if rx, ok := i.indUint(r); ok {
			return rx == uint64(lx)
		}
	case *uint:
		if rx, ok := i.indUint(r); ok {
			return rx == uint64(*lx)
		}
	case uint8:
		if rx, ok := i.indUint(r); ok {
			return rx == uint64(lx)
		}
	case *uint8:
		if rx, ok := i.indUint(r); ok {
			return rx == uint64(*lx)
		}
	case uint16:
		if rx, ok := i.indUint(r); ok {
			return rx == uint64(lx)
		}
	case *uint16:
		if rx, ok := i.indUint(r); ok {
			return rx == uint64(*lx)
		}
	case uint32:
		if rx, ok := i.indUint(r); ok {
			return rx == uint64(lx)
		}
	case *uint32:
		if rx, ok := i.indUint(r); ok {
			return rx == uint64(*lx)
		}
	case uint64:
		if rx, ok := i.indUint(r); ok {
			return rx == lx
		}
	case *uint64:
		if rx, ok := i.indUint(r); ok {
			return rx == *lx
		}
	case float32:
		if rx, ok := i.indFloat(r); ok {
			return i.eqlf64(rx, float64(lx))
		}
	case *float32:
		if rx, ok := i.indFloat(r); ok {
			return i.eqlf64(rx, float64(*lx))
		}
	case float64:
		if rx, ok := i.indFloat(r); ok {
			return i.eqlf64(rx, lx)
		}
	case *float64:
		if rx, ok := i.indFloat(r); ok {
			return i.eqlf64(rx, *lx)
		}
	case []byte:
		if rx, ok := i.indBytes(r); ok {
			return bytes.Equal(rx, lx)
		}
	case *[]byte:
		if rx, ok := i.indBytes(r); ok {
			return bytes.Equal(rx, *lx)
		}
	case string:
		if rx, ok := i.indString(r); ok {
			return rx == lx
		}
	case *string:
		if rx, ok := i.indString(r); ok {
			return rx == *lx
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

func (i StaticInspector) Copy(val any) (dst any, err error) {
	switch x := val.(type) {
	case bool:
		dst = x
	case *bool:
		dst = *x
	case int:
		dst = x
	case *int:
		dst = *x
	case int8:
		dst = x
	case *int8:
		dst = *x
	case int16:
		dst = x
	case *int16:
		dst = *x
	case int32:
		dst = x
	case *int32:
		dst = *x
	case int64:
		dst = x
	case *int64:
		dst = *x
	case uint:
		dst = x
	case *uint:
		dst = *x
	case uint8:
		dst = x
	case *uint8:
		dst = *x
	case uint16:
		dst = x
	case *uint16:
		dst = *x
	case uint32:
		dst = x
	case *uint32:
		dst = *x
	case uint64:
		dst = x
	case *uint64:
		dst = *x
	case float32:
		dst = x
	case *float32:
		dst = *x
	case float64:
		dst = x
	case *float64:
		dst = *x
	case []byte:
		origin := x
		cpy := append([]byte(nil), origin...)
		dst = cpy
	case *[]byte:
		origin := *x
		cpy := append([]byte(nil), origin...)
		dst = cpy
	case string:
		origin := x
		cpy := append([]byte(nil), origin...)
		dst = byteconv.B2S(cpy)
	case *string:
		origin := *x
		cpy := append([]byte(nil), origin...)
		dst = byteconv.B2S(cpy)
	default:
		return nil, ErrUnsupportedType
	}
	return
}

func (i StaticInspector) CopyTo(src, dst any, buf AccumulativeBuffer) error {
	switch x := src.(type) {
	case bool:
		if _, ok := dst.(*bool); !ok {
			return ErrMustPointerType
		}
		*dst.(*bool) = x
	case *bool:
		if _, ok := dst.(*bool); !ok {
			return ErrMustPointerType
		}
		*dst.(*bool) = *x
	case int:
		if _, ok := dst.(*int); !ok {
			return ErrMustPointerType
		}
		*dst.(*int) = x
	case *int:
		if _, ok := dst.(*int); !ok {
			return ErrMustPointerType
		}
		*dst.(*int) = *x
	case int8:
		if _, ok := dst.(*int8); !ok {
			return ErrMustPointerType
		}
		*dst.(*int8) = x
	case *int8:
		if _, ok := dst.(*int8); !ok {
			return ErrMustPointerType
		}
		*dst.(*int8) = *x
	case int16:
		if _, ok := dst.(*int16); !ok {
			return ErrMustPointerType
		}
		*dst.(*int16) = x
	case *int16:
		if _, ok := dst.(*int16); !ok {
			return ErrMustPointerType
		}
		*dst.(*int16) = *x
	case int32:
		if _, ok := dst.(*int32); !ok {
			return ErrMustPointerType
		}
		*dst.(*int32) = x
	case *int32:
		if _, ok := dst.(*int32); !ok {
			return ErrMustPointerType
		}
		*dst.(*int32) = *x
	case int64:
		if _, ok := dst.(*int64); !ok {
			return ErrMustPointerType
		}
		*dst.(*int64) = x
	case *int64:
		if _, ok := dst.(*int64); !ok {
			return ErrMustPointerType
		}
		*dst.(*int64) = *x
	case uint:
		if _, ok := dst.(*uint); !ok {
			return ErrMustPointerType
		}
		*dst.(*uint) = x
	case *uint:
		if _, ok := dst.(*uint); !ok {
			return ErrMustPointerType
		}
		*dst.(*uint) = *x
	case uint8:
		if _, ok := dst.(*uint8); !ok {
			return ErrMustPointerType
		}
		*dst.(*uint8) = x
	case *uint8:
		if _, ok := dst.(*uint8); !ok {
			return ErrMustPointerType
		}
		*dst.(*uint8) = *x
	case uint16:
		if _, ok := dst.(*uint16); !ok {
			return ErrMustPointerType
		}
		*dst.(*uint16) = x
	case *uint16:
		if _, ok := dst.(*uint16); !ok {
			return ErrMustPointerType
		}
		*dst.(*uint16) = *x
	case uint32:
		if _, ok := dst.(*uint32); !ok {
			return ErrMustPointerType
		}
		*dst.(*uint32) = x
	case *uint32:
		if _, ok := dst.(*uint32); !ok {
			return ErrMustPointerType
		}
		*dst.(*uint32) = *x
	case uint64:
		if _, ok := dst.(*uint64); !ok {
			return ErrMustPointerType
		}
		*dst.(*uint64) = x
	case *uint64:
		if _, ok := dst.(*uint64); !ok {
			return ErrMustPointerType
		}
		*dst.(*uint64) = *x
	case float32:
		if _, ok := dst.(*float32); !ok {
			return ErrMustPointerType
		}
		*dst.(*float32) = x
	case *float32:
		if _, ok := dst.(*float32); !ok {
			return ErrMustPointerType
		}
		*dst.(*float32) = *x
	case float64:
		if _, ok := dst.(*float64); !ok {
			return ErrMustPointerType
		}
		*dst.(*float64) = x
	case *float64:
		if _, ok := dst.(*float64); !ok {
			return ErrMustPointerType
		}
		*dst.(*float64) = *x
	case []byte:
		p, ok := dst.(*[]byte)
		if !ok {
			return ErrMustPointerType
		}
		*p = buf.Bufferize(x)
	case *[]byte:
		p, ok := dst.(*[]byte)
		if !ok {
			return ErrMustPointerType
		}
		*p = buf.Bufferize(*x)
	case string:
		p, ok := dst.(*string)
		if !ok {
			return ErrMustPointerType
		}
		*p = buf.BufferizeString(x)
	case *string:
		p, ok := dst.(*string)
		if !ok {
			return ErrMustPointerType
		}
		*p = buf.BufferizeString(*x)
	default:
		return ErrUnsupportedType
	}
	return nil
}

func (i StaticInspector) Length(x any, result *int, _ ...string) error {
	l, _ := i.lc(x)
	*result = l
	return nil
}

func (i StaticInspector) Capacity(x any, result *int, _ ...string) error {
	_, c := i.lc(x)
	*result = c
	return nil
}

func (i StaticInspector) Reset(val any) error {
	switch x := val.(type) {
	case bool:
		x = false
	case *bool:
		*x = false
	case int:
		x = 0
	case *int:
		*x = 0
	case *int8:
		*x = 0
	case int8:
		x = 0
	case *int16:
		*x = 0
	case int16:
		x = 0
	case *int32:
		*x = 0
	case int32:
		x = 0
	case *int64:
		*x = 0
	case int64:
		x = 0
	case *uint:
		*x = 0
	case uint:
		x = 0
	case *uint8:
		*x = 0
	case uint8:
		x = 0
	case *uint16:
		*x = 0
	case uint16:
		x = 0
	case *uint32:
		*x = 0
	case uint32:
		x = 0
	case *uint64:
		*x = 0
	case uint64:
		x = 0
	case *float32:
		*x = 0
	case float32:
		x = 0
	case *float64:
		*x = 0
	case float64:
		x = 0
	case []byte:
		p := x
		val = p[:0]
	case *[]byte:
		p := *x
		p = p[:0]
		val = &p
	case string:
		x = ""
	case *string:
		var s string
		x = &s
	}
	return nil
}

func (i StaticInspector) indBool(val any) (bool, bool) {
	switch x := val.(type) {
	case bool:
		return x, true
	case *bool:
		return *x, true
	}
	return false, false
}

func (i StaticInspector) indInt(val any) (int64, bool) {
	switch x := val.(type) {
	case int:
		return int64(x), true
	case *int:
		return int64(*x), true
	case int8:
		return int64(x), true
	case *int8:
		return int64(*x), true
	case int16:
		return int64(x), true
	case *int16:
		return int64(*x), true
	case int32:
		return int64(x), true
	case *int32:
		return int64(*x), true
	case int64:
		return x, true
	case *int64:
		return *x, true
	case float32:
		return int64(x), true
	case *float32:
		return int64(*x), true
	case float64:
		return int64(x), true
	case *float64:
		return int64(*x), true
	}
	return 0, false
}

func (i StaticInspector) indUint(val any) (uint64, bool) {
	switch x := val.(type) {
	case uint:
		return uint64(x), true
	case *uint:
		return uint64(*x), true
	case uint8:
		return uint64(x), true
	case *uint8:
		return uint64(*x), true
	case uint16:
		return uint64(x), true
	case *uint16:
		return uint64(*x), true
	case uint32:
		return uint64(x), true
	case *uint32:
		return uint64(*x), true
	case uint64:
		return x, true
	case *uint64:
		return *x, true
	case float32:
		return uint64(x), true
	case *float32:
		return uint64(*x), true
	case float64:
		return uint64(x), true
	case *float64:
		return uint64(*x), true
	}
	return 0, false
}

func (i StaticInspector) indFloat(val any) (float64, bool) {
	switch x := val.(type) {
	case float32:
		return float64(x), true
	case *float32:
		return float64(*x), true
	case float64:
		return x, true
	case *float64:
		return *x, true
	case int:
		return float64(x), true
	case *int:
		return float64(*x), true
	case int8:
		return float64(x), true
	case *int8:
		return float64(*x), true
	case int16:
		return float64(x), true
	case *int16:
		return float64(*x), true
	case int32:
		return float64(x), true
	case *int32:
		return float64(*x), true
	case int64:
		return float64(x), true
	case *int64:
		return float64(*x), true
	case uint:
		return float64(x), true
	case *uint:
		return float64(*x), true
	case uint8:
		return float64(x), true
	case *uint8:
		return float64(*x), true
	case uint16:
		return float64(x), true
	case *uint16:
		return float64(*x), true
	case uint32:
		return float64(x), true
	case *uint32:
		return float64(*x), true
	case uint64:
		return float64(x), true
	case *uint64:
		return float64(*x), true
	}
	return 0, false
}

func (i StaticInspector) indString(val any) (string, bool) {
	switch x := val.(type) {
	case string:
		return x, true
	case *string:
		return *x, true
	}
	if b, ok := i.indBytes(val); ok {
		return byteconv.B2S(b), true
	}
	return "", false
}

func (i StaticInspector) indBytes(val any) ([]byte, bool) {
	switch x := val.(type) {
	case []byte:
		return x, true
	case *[]byte:
		return *x, true
	}
	if s, ok := i.indString(val); ok {
		return byteconv.S2B(s), true
	}
	return nil, false
}

func (i StaticInspector) eqlf64(a, b float64) bool {
	return math.Abs(a-b) <= FloatPrecision
}

func (i StaticInspector) lc(val any) (int, int) {
	switch x := val.(type) {
	case []byte:
		p := x
		return len(p), cap(p)
	case *[]byte:
		p := *x
		return len(p), cap(p)
	case string:
		s := x
		return len(s), len(s)
	case *string:
		s := *x
		return len(s), len(s)
	default:
		return 0, 0
	}
}
