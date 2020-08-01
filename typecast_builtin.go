package inspector

import "github.com/koykov/fastconv"

func TypeCastToBytes(typ string, val interface{}) (r interface{}, ok bool) {
	var b []byte
	switch val.(type) {
	case *[]byte:
		b = *val.(*[]byte)
		ok = true
	case []byte:
		b = val.([]byte)
		ok = true
	case *string:
		b = fastconv.S2B(*val.(*string))
		ok = true
	case string:
		b = fastconv.S2B(val.(string))
		ok = true
	}
	if ok {
		switch typ {
		case "[]byte":
			r = b
		case "*[]byte":
			r = &b
		default:
			ok = false
		}
	}
	return
}

func TypeCastToStr(typ string, val interface{}) (r interface{}, ok bool) {
	var s string
	switch val.(type) {
	case *[]byte:
		s = fastconv.B2S(*val.(*[]byte))
		ok = true
	case []byte:
		s = fastconv.B2S(val.([]byte))
		ok = true
	case *string:
		s = *val.(*string)
		ok = true
	case string:
		s = val.(string)
		ok = true
	}
	if ok {
		switch typ {
		case "string":
			r = s
		case "*string":
			r = &s
		default:
			ok = false
		}
	}
	return
}

func TypeCastToBool(typ string, val interface{}) (r interface{}, ok bool) {
	var b bool
	switch val.(type) {
	case *bool:
		b = *val.(*bool)
		ok = true
	case bool:
		b = val.(bool)
		ok = true
	case *[]byte:
		b = fastconv.B2S(*val.(*[]byte)) == "true"
		ok = true
	case []byte:
		b = fastconv.B2S(val.([]byte)) == "true"
		ok = true
	case *string:
		b = *val.(*string) == "true"
		ok = true
	case string:
		b = val.(string) == "true"
		ok = true
	case int:
		b = val.(int) != 0
		ok = true
	case *int:
		b = *val.(*int) != 0
		ok = true
	case int8:
		b = val.(int8) != 0
		ok = true
	case *int8:
		b = *val.(*int8) != 0
		ok = true
	case int16:
		b = val.(int16) != 0
		ok = true
	case *int16:
		b = *val.(*int16) != 0
		ok = true
	case int32:
		b = val.(int32) != 0
		ok = true
	case *int32:
		b = *val.(*int32) != 0
		ok = true
	case int64:
		b = val.(int64) != 0
		ok = true
	case *int64:
		b = *val.(*int64) != 0
		ok = true
	case uint:
		b = val.(uint) != 0
		ok = true
	case *uint:
		b = *val.(*uint) != 0
		ok = true
	case uint8:
		b = val.(uint8) != 0
		ok = true
	case *uint8:
		b = *val.(*uint8) != 0
		ok = true
	case uint16:
		b = val.(uint16) != 0
		ok = true
	case *uint16:
		b = *val.(*uint16) != 0
		ok = true
	case uint32:
		b = val.(uint32) != 0
		ok = true
	case *uint32:
		b = *val.(*uint32) != 0
		ok = true
	case uint64:
		b = val.(uint64) != 0
		ok = true
	case *uint64:
		b = *val.(*uint64) != 0
		ok = true
	case float32:
		b = val.(float32) != 0
		ok = true
	case *float32:
		b = *val.(*float32) != 0
		ok = true
	case float64:
		b = val.(float64) != 0
		ok = true
	case *float64:
		b = *val.(*float64) != 0
		ok = true
	}
	if ok {
		switch typ {
		case "bool":
			r = b
		case "*bool":
			r = &b
		default:
			ok = false
		}
	}
	return
}

func TypeCastToInt(typ string, val interface{}) (r interface{}, ok bool) {
	var i int64
	switch val.(type) {
	case int:
		i = int64(val.(int))
		ok = true
	case *int:
		i = int64(*val.(*int))
		ok = true
	case int8:
		i = int64(val.(int8))
		ok = true
	case *int8:
		i = int64(*val.(*int8))
		ok = true
	case int16:
		i = int64(val.(int16))
		ok = true
	case *int16:
		i = int64(*val.(*int16))
		ok = true
	case int32:
		i = int64(val.(int32))
		ok = true
	case *int32:
		i = int64(*val.(*int32))
		ok = true
	case int64:
		i = val.(int64)
		ok = true
	case *int64:
		i = *val.(*int64)
		ok = true
	}
	if ok {
		switch typ {
		case "int":
			r = int(i)
		case "*int":
			c := int(i)
			r = &c
		case "int8":
			r = int8(i)
		case "*int8":
			c := int8(i)
			r = &c
		case "int16":
			r = int16(i)
		case "*int16":
			c := int16(i)
			r = &c
		case "int32":
			r = int32(i)
		case "*int32":
			c := int32(i)
			r = &c
		case "int64":
			r = i
		case "*int64":
			r = &i
		}
	}
	return
}

func TypeCastToUint(typ string, val interface{}) (r interface{}, ok bool) {
	var u uint64
	switch val.(type) {
	case uint:
		u = uint64(val.(uint))
		ok = true
	case *uint:
		u = uint64(*val.(*uint))
		ok = true
	case uint8:
		u = uint64(val.(uint8))
		ok = true
	case *uint8:
		u = uint64(*val.(*uint8))
		ok = true
	case uint16:
		u = uint64(val.(uint16))
		ok = true
	case *uint16:
		u = uint64(*val.(*uint16))
		ok = true
	case uint32:
		u = uint64(val.(uint32))
		ok = true
	case *uint32:
		u = uint64(*val.(*uint32))
		ok = true
	case uint64:
		u = val.(uint64)
		ok = true
	case *uint64:
		u = *val.(*uint64)
		ok = true
	}
	if ok {
		switch typ {
		case "uint":
			r = uint(u)
		case "*uint":
			c := uint(u)
			r = &c
		case "uint8":
			r = uint8(u)
		case "*uint8":
			c := uint8(u)
			r = &c
		case "uint16":
			r = uint16(u)
		case "*uint16":
			c := uint16(u)
			r = &c
		case "uint32":
			r = uint32(u)
		case "*uint32":
			c := uint32(u)
			r = &c
		case "uint64":
			r = u
		case "*uint64":
			r = &u
		}
	}
	return
}

func TypeCastToFloat(typ string, val interface{}) (r interface{}, ok bool) {
	if typ != "float32" && typ != "float64" {
		return
	}
	var f float64
	switch val.(type) {
	case float32:
		f = float64(val.(float32))
		ok = true
	case *float32:
		f = float64(*val.(*float32))
		ok = true
	case float64:
		f = val.(float64)
		ok = true
	case *float64:
		f = *val.(*float64)
		ok = true
	}
	if ok {
		switch typ {
		case "float32":
			r = float32(f)
		case "*float32":
			c := float32(f)
			r = &c
		case "float64":
			r = f
		case "*float64":
			r = &f
		}
	}
	return
}
