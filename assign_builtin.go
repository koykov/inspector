package inspector

import (
	"regexp"
	"strconv"

	"github.com/koykov/byteconv"
	"github.com/koykov/x2bytes"
)

var (
	// todo optimise string check
	reIsDecInt   = regexp.MustCompile(`^[-+]?[\d]+$`)
	reIsDecUint  = regexp.MustCompile(`^[+]?[\d]+$`)
	reIsDecFloat = regexp.MustCompile(`^[-+]?[\d]*\.?[\d]+([eE][-+]?[\d]+)?$`)
)

// AssignToBytes assigns(converts) source to bytes destination.
func AssignToBytes(dst, src any, buf AccumulativeBuffer) (ok bool) {
	switch dst.(type) {
	case *[]byte:
		switch src.(type) {
		case *[]byte:
			*dst.(*[]byte) = *src.(*[]byte)
			ok = true
		case []byte:
			*dst.(*[]byte) = src.([]byte)
			ok = true
		case *string:
			*dst.(*[]byte) = byteconv.S2B(*src.(*string))
			ok = true
		case string:
			*dst.(*[]byte) = byteconv.S2B(src.(string))
			ok = true
		default:
			var (
				p   = *dst.(*[]byte)
				err error
			)
			// Worst case, try to convert source to bytes.
			if buf == nil {
				p, err = x2bytes.ToBytes(p[:0], src)
				if ok = err == nil; ok {
					*dst.(*[]byte) = p
				}
			} else {
				bb := buf.AcquireBytes()
				offset := len(bb)
				bb, err = x2bytes.ToBytes(bb, src)
				if ok = err == nil; ok {
					*dst.(*[]byte) = bb[offset:]
				}
				buf.ReleaseBytes(bb)
			}
		}
	}
	return
}

// AssignToStr assigns(converts) source to string destination.
func AssignToStr(dst, src any, buf AccumulativeBuffer) (ok bool) {
	switch dst.(type) {
	case *string:
		switch src.(type) {
		case *[]byte:
			*dst.(*string) = byteconv.B2S(*src.(*[]byte))
			ok = true
		case []byte:
			*dst.(*string) = byteconv.B2S(src.([]byte))
			ok = true
		case *string:
			*dst.(*string) = *src.(*string)
			ok = true
		case string:
			*dst.(*string) = src.(string)
			ok = true
		default:
			var (
				p   []byte
				err error
			)
			// Worst case, try to convert source to bytes.
			if buf == nil {
				p = byteconv.S2B(*dst.(*string))
				p, err = x2bytes.ToBytes(p, src)
				if ok = err == nil; ok {
					*dst.(*string) = byteconv.B2S(p)
				}
			} else {
				bb := buf.AcquireBytes()
				offset := len(bb)
				bb, err = x2bytes.ToBytes(bb, src)
				if ok = err == nil; ok {
					*dst.(*string) = byteconv.B2S(bb[offset:])
				}
				buf.ReleaseBytes(bb)
			}
		}
	}
	return
}

// AssignToBool assign(converts) source to bool destination.
func AssignToBool(dst, src any, _ AccumulativeBuffer) (ok bool) {
	switch dst.(type) {
	case *bool:
		switch src.(type) {
		case *bool:
			*dst.(*bool) = *src.(*bool)
			ok = true
		case bool:
			*dst.(*bool) = src.(bool)
			ok = true
		case *[]byte:
			*dst.(*bool) = byteconv.B2S(*src.(*[]byte)) == "true"
			ok = true
		case []byte:
			*dst.(*bool) = byteconv.B2S(src.([]byte)) == "true"
			ok = true
		case *string:
			*dst.(*bool) = *src.(*string) == "true"
			ok = true
		case string:
			*dst.(*bool) = src.(string) == "true"
			ok = true
		case int:
			*dst.(*bool) = src.(int) != 0
			ok = true
		case *int:
			*dst.(*bool) = *src.(*int) != 0
			ok = true
		case int8:
			*dst.(*bool) = src.(int8) != 0
			ok = true
		case *int8:
			*dst.(*bool) = *src.(*int8) != 0
			ok = true
		case int16:
			*dst.(*bool) = src.(int16) != 0
			ok = true
		case *int16:
			*dst.(*bool) = *src.(*int16) != 0
			ok = true
		case int32:
			*dst.(*bool) = src.(int32) != 0
			ok = true
		case *int32:
			*dst.(*bool) = *src.(*int32) != 0
			ok = true
		case int64:
			*dst.(*bool) = src.(int64) != 0
			ok = true
		case *int64:
			*dst.(*bool) = *src.(*int64) != 0
			ok = true
		case uint:
			*dst.(*bool) = src.(uint) != 0
			ok = true
		case *uint:
			*dst.(*bool) = *src.(*uint) != 0
			ok = true
		case uint8:
			*dst.(*bool) = src.(uint8) != 0
			ok = true
		case *uint8:
			*dst.(*bool) = *src.(*uint8) != 0
			ok = true
		case uint16:
			*dst.(*bool) = src.(uint16) != 0
			ok = true
		case *uint16:
			*dst.(*bool) = *src.(*uint16) != 0
			ok = true
		case uint32:
			*dst.(*bool) = src.(uint32) != 0
			ok = true
		case *uint32:
			*dst.(*bool) = *src.(*uint32) != 0
			ok = true
		case uint64:
			*dst.(*bool) = src.(uint64) != 0
			ok = true
		case *uint64:
			*dst.(*bool) = *src.(*uint64) != 0
			ok = true
		case float32:
			*dst.(*bool) = src.(float32) != 0
			ok = true
		case *float32:
			*dst.(*bool) = *src.(*float32) != 0
			ok = true
		case float64:
			*dst.(*bool) = src.(float64) != 0
			ok = true
		case *float64:
			*dst.(*bool) = *src.(*float64) != 0
			ok = true
		}
	}
	return
}

// AssignToInt assigns(converts) source to int destination.
func AssignToInt(dst, src any, _ AccumulativeBuffer) (ok bool) {
	var i int64
	switch src.(type) {
	case int:
		i = int64(src.(int))
		ok = true
	case *int:
		i = int64(*src.(*int))
		ok = true
	case int8:
		i = int64(src.(int8))
		ok = true
	case *int8:
		i = int64(*src.(*int8))
		ok = true
	case int16:
		i = int64(src.(int16))
		ok = true
	case *int16:
		i = int64(*src.(*int16))
		ok = true
	case int32:
		i = int64(src.(int32))
		ok = true
	case *int32:
		i = int64(*src.(*int32))
		ok = true
	case int64:
		i = src.(int64)
		ok = true
	case *int64:
		i = *src.(*int64)
		ok = true
	case []byte:
		i, ok = atoi(byteconv.B2S(src.([]byte)))
	case *[]byte:
		i, ok = atoi(byteconv.B2S(*src.(*[]byte)))
	case string:
		i, ok = atoi(src.(string))
	case *string:
		i, ok = atoi(*src.(*string))
	}
	if ok {
		switch dst.(type) {
		case *int:
			*dst.(*int) = int(i)
		case *int8:
			*dst.(*int8) = int8(i)
		case *int16:
			*dst.(*int16) = int16(i)
		case *int32:
			*dst.(*int32) = int32(i)
		case *int64:
			*dst.(*int64) = i
		default:
			ok = false
		}
	}
	return
}

// AssignToUint assigns(converts) source to unsigned int destination.
func AssignToUint(dst, src any, _ AccumulativeBuffer) (ok bool) {
	var u uint64
	switch src.(type) {
	case uint:
		u = uint64(src.(uint))
		ok = true
	case *uint:
		u = uint64(*src.(*uint))
		ok = true
	case uint8:
		u = uint64(src.(uint8))
		ok = true
	case *uint8:
		u = uint64(*src.(*uint8))
		ok = true
	case uint16:
		u = uint64(src.(uint16))
		ok = true
	case *uint16:
		u = uint64(*src.(*uint16))
		ok = true
	case uint32:
		u = uint64(src.(uint32))
		ok = true
	case *uint32:
		u = uint64(*src.(*uint32))
		ok = true
	case uint64:
		u = src.(uint64)
		ok = true
	case *uint64:
		u = *src.(*uint64)
		ok = true
	case []byte:
		u, ok = atou(byteconv.B2S(src.([]byte)))
	case *[]byte:
		u, ok = atou(byteconv.B2S(*src.(*[]byte)))
	case string:
		u, ok = atou(src.(string))
	case *string:
		u, ok = atou(*src.(*string))
	}
	if ok {
		switch dst.(type) {
		case *uint:
			*dst.(*uint) = uint(u)
		case *uint8:
			*dst.(*uint8) = uint8(u)
		case *uint16:
			*dst.(*uint16) = uint16(u)
		case *uint32:
			*dst.(*uint32) = uint32(u)
		case *uint64:
			*dst.(*uint64) = u
		default:
			ok = false
		}
	}
	return
}

// AssignToFloat assigns(converts) source to float destination.
func AssignToFloat(dst, src any, _ AccumulativeBuffer) (ok bool) {
	var f float64
	switch src.(type) {
	case float32:
		f = float64(src.(float32))
		ok = true
	case *float32:
		f = float64(*src.(*float32))
		ok = true
	case float64:
		f = src.(float64)
		ok = true
	case *float64:
		f = *src.(*float64)
		ok = true
	case []byte:
		f, ok = atof(byteconv.B2S(src.([]byte)))
	case *[]byte:
		f, ok = atof(byteconv.B2S(*src.(*[]byte)))
	case string:
		f, ok = atof(src.(string))
	case *string:
		f, ok = atof(*src.(*string))
	}
	if ok {
		switch dst.(type) {
		case *float32:
			*dst.(*float32) = float32(f)
		case *float64:
			*dst.(*float64) = f
		default:
			ok = false
		}
	}
	return
}

// Check if string contains integer and parse it.
func atoi(s string) (int64, bool) {
	if reIsDecInt.MatchString(s) {
		if i, err := strconv.ParseInt(s, 10, 64); err == nil {
			return i, true
		}
	}
	return 0, false
}

// Check if string contains unsigned integer and parse it.
func atou(s string) (uint64, bool) {
	if reIsDecUint.MatchString(s) {
		if u, err := strconv.ParseUint(s, 10, 64); err == nil {
			return u, true
		}
	}
	return 0, false
}

// Check if string contains float and parse it.
func atof(s string) (float64, bool) {
	if reIsDecFloat.MatchString(s) {
		if f, err := strconv.ParseFloat(s, 64); err == nil {
			return f, true
		}
	}
	return 0, false
}
