package inspector

import (
	"strconv"

	"github.com/koykov/byteconv"
	"github.com/koykov/x2bytes"
)

// AssignToBytes assigns(converts) source to bytes destination.
func AssignToBytes(dst, src any, buf AccumulativeBuffer) (ok bool) {
	switch lx := dst.(type) {
	case *[]byte:
		switch rx := src.(type) {
		case *[]byte:
			*lx = *rx
			ok = true
		case []byte:
			*lx = rx
			ok = true
		case *string:
			*lx = byteconv.S2B(*rx)
			ok = true
		case string:
			*lx = byteconv.S2B(rx)
			ok = true
		default:
			var (
				p   = *lx
				err error
			)
			// Worst case, try to convert source to bytes.
			if buf == nil {
				p, err = x2bytes.ToBytes(p[:0], src)
				if ok = err == nil; ok {
					*lx = p
				}
			} else {
				bb := buf.AcquireBytes()
				offset := len(bb)
				bb, err = x2bytes.ToBytes(bb, src)
				if ok = err == nil; ok {
					*lx = bb[offset:]
				}
				buf.ReleaseBytes(bb)
			}
		}
	}
	return
}

// AssignToStr assigns(converts) source to string destination.
func AssignToStr(dst, src any, buf AccumulativeBuffer) (ok bool) {
	switch lx := dst.(type) {
	case *string:
		switch rx := src.(type) {
		case *[]byte:
			*lx = byteconv.B2S(*rx)
			ok = true
		case []byte:
			*lx = byteconv.B2S(rx)
			ok = true
		case *string:
			*lx = *rx
			ok = true
		case string:
			*lx = rx
			ok = true
		default:
			var (
				p   []byte
				err error
			)
			// Worst case, try to convert source to bytes.
			if buf == nil {
				p = byteconv.S2B(*lx)
				p, err = x2bytes.ToBytes(p, src)
				if ok = err == nil; ok {
					*lx = byteconv.B2S(p)
				}
			} else {
				bb := buf.AcquireBytes()
				offset := len(bb)
				bb, err = x2bytes.ToBytes(bb, src)
				if ok = err == nil; ok {
					*lx = byteconv.B2S(bb[offset:])
				}
				buf.ReleaseBytes(bb)
			}
		}
	}
	return
}

// AssignToBool assign(converts) source to bool destination.
func AssignToBool(dst, src any, _ AccumulativeBuffer) (ok bool) {
	switch lx := dst.(type) {
	case *bool:
		switch rx := src.(type) {
		case *bool:
			*lx = *rx
			ok = true
		case bool:
			*lx = rx
			ok = true
		case *[]byte:
			*lx = byteconv.B2S(*rx) == "true"
			ok = true
		case []byte:
			*lx = byteconv.B2S(rx) == "true"
			ok = true
		case *string:
			*lx = *rx == "true"
			ok = true
		case string:
			*lx = rx == "true"
			ok = true
		case int:
			*lx = rx != 0
			ok = true
		case *int:
			*lx = *rx != 0
			ok = true
		case int8:
			*lx = rx != 0
			ok = true
		case *int8:
			*lx = *rx != 0
			ok = true
		case int16:
			*lx = rx != 0
			ok = true
		case *int16:
			*lx = *rx != 0
			ok = true
		case int32:
			*lx = rx != 0
			ok = true
		case *int32:
			*lx = *rx != 0
			ok = true
		case int64:
			*lx = rx != 0
			ok = true
		case *int64:
			*lx = *rx != 0
			ok = true
		case uint:
			*lx = rx != 0
			ok = true
		case *uint:
			*lx = *rx != 0
			ok = true
		case uint8:
			*lx = rx != 0
			ok = true
		case *uint8:
			*lx = *rx != 0
			ok = true
		case uint16:
			*lx = rx != 0
			ok = true
		case *uint16:
			*lx = *rx != 0
			ok = true
		case uint32:
			*lx = rx != 0
			ok = true
		case *uint32:
			*lx = *rx != 0
			ok = true
		case uint64:
			*lx = rx != 0
			ok = true
		case *uint64:
			*lx = *rx != 0
			ok = true
		case float32:
			*lx = rx != 0
			ok = true
		case *float32:
			*lx = *rx != 0
			ok = true
		case float64:
			*lx = rx != 0
			ok = true
		case *float64:
			*lx = *rx != 0
			ok = true
		}
	}
	return
}

// AssignToInt assigns(converts) source to int destination.
func AssignToInt(dst, src any, _ AccumulativeBuffer) (ok bool) {
	var i int64
	switch x := src.(type) {
	case int:
		i = int64(x)
		ok = true
	case *int:
		i = int64(*x)
		ok = true
	case int8:
		i = int64(x)
		ok = true
	case *int8:
		i = int64(*x)
		ok = true
	case int16:
		i = int64(x)
		ok = true
	case *int16:
		i = int64(*x)
		ok = true
	case int32:
		i = int64(x)
		ok = true
	case *int32:
		i = int64(*x)
		ok = true
	case int64:
		i = x
		ok = true
	case *int64:
		i = *x
		ok = true
	case uint:
		i = int64(x)
		ok = true
	case *uint:
		i = int64(*x)
		ok = true
	case uint8:
		i = int64(x)
		ok = true
	case *uint8:
		i = int64(*x)
		ok = true
	case uint16:
		i = int64(x)
		ok = true
	case *uint16:
		i = int64(*x)
		ok = true
	case uint32:
		i = int64(x)
		ok = true
	case *uint32:
		i = int64(*x)
		ok = true
	case uint64:
		i = int64(x)
		ok = true
	case *uint64:
		i = int64(*x)
		ok = true
	case float32:
		i = int64(x)
		ok = true
	case *float32:
		i = int64(*x)
		ok = true
	case float64:
		i = int64(x)
		ok = true
	case *float64:
		i = int64(*x)
		ok = true
	case []byte:
		i, ok = atoi(byteconv.B2S(x))
	case *[]byte:
		i, ok = atoi(byteconv.B2S(*x))
	case string:
		i, ok = atoi(x)
	case *string:
		i, ok = atoi(*x)
	}
	if ok {
		switch x := dst.(type) {
		case *int:
			*x = int(i)
		case *int8:
			*x = int8(i)
		case *int16:
			*x = int16(i)
		case *int32:
			*x = int32(i)
		case *int64:
			*x = i
		default:
			ok = false
		}
	}
	return
}

// AssignToUint assigns(converts) source to unsigned int destination.
func AssignToUint(dst, src any, _ AccumulativeBuffer) (ok bool) {
	var u uint64
	switch x := src.(type) {
	case uint:
		u = uint64(x)
		ok = true
	case *uint:
		u = uint64(*x)
		ok = true
	case uint8:
		u = uint64(x)
		ok = true
	case *uint8:
		u = uint64(*x)
		ok = true
	case uint16:
		u = uint64(x)
		ok = true
	case *uint16:
		u = uint64(*x)
		ok = true
	case uint32:
		u = uint64(x)
		ok = true
	case *uint32:
		u = uint64(*x)
		ok = true
	case uint64:
		u = x
		ok = true
	case *uint64:
		u = *x
		ok = true
	case int:
		u = uint64(x)
		ok = true
	case *int:
		u = uint64(*x)
		ok = true
	case int8:
		u = uint64(x)
		ok = true
	case *int8:
		u = uint64(*x)
		ok = true
	case int16:
		u = uint64(x)
		ok = true
	case *int16:
		u = uint64(*x)
		ok = true
	case int32:
		u = uint64(x)
		ok = true
	case *int32:
		u = uint64(*x)
		ok = true
	case int64:
		u = uint64(x)
		ok = true
	case *int64:
		u = uint64(*x)
		ok = true
	case float32:
		u = uint64(x)
		ok = true
	case *float32:
		u = uint64(*x)
		ok = true
	case float64:
		u = uint64(x)
		ok = true
	case *float64:
		u = uint64(*x)
		ok = true
	case []byte:
		u, ok = atou(byteconv.B2S(x))
	case *[]byte:
		u, ok = atou(byteconv.B2S(*x))
	case string:
		u, ok = atou(x)
	case *string:
		u, ok = atou(*x)
	}
	if ok {
		switch x := dst.(type) {
		case *uint:
			*x = uint(u)
		case *uint8:
			*x = uint8(u)
		case *uint16:
			*x = uint16(u)
		case *uint32:
			*x = uint32(u)
		case *uint64:
			*x = u
		default:
			ok = false
		}
	}
	return
}

// AssignToFloat assigns(converts) source to float destination.
func AssignToFloat(dst, src any, _ AccumulativeBuffer) (ok bool) {
	var f float64
	switch x := src.(type) {
	case uint:
		f = float64(x)
		ok = true
	case *uint:
		f = float64(*x)
		ok = true
	case uint8:
		f = float64(x)
		ok = true
	case *uint8:
		f = float64(*x)
		ok = true
	case uint16:
		f = float64(x)
		ok = true
	case *uint16:
		f = float64(*x)
		ok = true
	case uint32:
		f = float64(x)
		ok = true
	case *uint32:
		f = float64(*x)
		ok = true
	case uint64:
		f = float64(x)
		ok = true
	case *uint64:
		f = float64(*x)
		ok = true
	case int:
		f = float64(x)
		ok = true
	case *int:
		f = float64(*x)
		ok = true
	case int8:
		f = float64(x)
		ok = true
	case *int8:
		f = float64(*x)
		ok = true
	case int16:
		f = float64(x)
		ok = true
	case *int16:
		f = float64(*x)
		ok = true
	case int32:
		f = float64(x)
		ok = true
	case *int32:
		f = float64(*x)
		ok = true
	case int64:
		f = float64(x)
		ok = true
	case *int64:
		f = float64(*x)
		ok = true
	case float32:
		f = float64(x)
		ok = true
	case *float32:
		f = float64(*x)
		ok = true
	case float64:
		f = x
		ok = true
	case *float64:
		f = *x
		ok = true
	case []byte:
		f, ok = atof(byteconv.B2S(x))
	case *[]byte:
		f, ok = atof(byteconv.B2S(*x))
	case string:
		f, ok = atof(x)
	case *string:
		f, ok = atof(*x)
	}
	if ok {
		switch x := dst.(type) {
		case *float32:
			*x = float32(f)
		case *float64:
			*x = f
		default:
			ok = false
		}
	}
	return
}

// Check if string contains integer and parse it.
func atoi(s string) (int64, bool) {
	if i, err := strconv.ParseInt(s, 10, 64); err == nil {
		return i, true
	}
	return 0, false
}

// Check if string contains unsigned integer and parse it.
func atou(s string) (uint64, bool) {
	if u, err := strconv.ParseUint(s, 10, 64); err == nil {
		return u, true
	}
	return 0, false
}

// Check if string contains float and parse it.
func atof(s string) (float64, bool) {
	if f, err := strconv.ParseFloat(s, 64); err == nil {
		return f, true
	}
	return 0, false
}
