package inspector

import "github.com/koykov/fastconv"

func AssignToBytes(dst, src interface{}) (ok bool) {
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
			*dst.(*[]byte) = fastconv.S2B(*src.(*string))
			ok = true
		case string:
			*dst.(*[]byte) = fastconv.S2B(src.(string))
			ok = true
		}
	}
	return
}

func AssignToStr(dst, src interface{}) (ok bool) {
	switch dst.(type) {
	case *string:
		switch src.(type) {
		case *[]byte:
			*dst.(*string) = fastconv.B2S(*src.(*[]byte))
			ok = true
		case []byte:
			*dst.(*string) = fastconv.B2S(src.([]byte))
			ok = true
		case *string:
			*dst.(*string) = *src.(*string)
			ok = true
		case string:
			*dst.(*string) = src.(string)
			ok = true
		}
	}
	return
}

func AssignToBool(dst, src interface{}) (ok bool) {
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
			*dst.(*bool) = fastconv.B2S(*src.(*[]byte)) == "true"
			ok = true
		case []byte:
			*dst.(*bool) = fastconv.B2S(src.([]byte)) == "true"
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

func AssignToInt(dst, src interface{}) (ok bool) {
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

func AssignToUint(dst, src interface{}) (ok bool) {
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

func AssignToFloat(dst, src interface{}) (ok bool) {
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
