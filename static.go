package inspector

import (
	"bytes"
	"strconv"

	"github.com/koykov/fastconv"
)

type StaticInspector struct {
	BaseInspector
}

func (i *StaticInspector) Get(src interface{}, _ ...string) (interface{}, error) {
	return src, nil
}

func (i *StaticInspector) GetTo(src interface{}, buf *interface{}, _ ...string) error {
	*buf = src
	return nil
}

func (i *StaticInspector) Set(_, _ interface{}, _ ...string) {
	//
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
