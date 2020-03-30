package inspector

import (
	"errors"
	"strconv"
)

var (
	ErrUnknownIntType   = errors.New("unknown integer type")
	ErrUnknownFloatType = errors.New("unknown float type")
)

func strToInt(s, typ string) (interface{}, error) {
	// i, err := strconv.ParseInt(s, 0, 0)
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	switch typ {
	case "int":
		return int(i), nil
	case "int8":
		return int8(i), nil
	case "int16":
		return int16(i), nil
	case "int32":
		return int32(i), nil
	case "int64":
		return int64(i), nil
	case "uint":
		return uint(i), nil
	case "uint8":
		return uint(i), nil
	case "uint16":
		return uint16(i), nil
	case "uint32":
		return uint32(i), nil
	case "uint64":
		return uint64(i), nil
	default:
		return 0, ErrUnknownIntType
	}
}

func strToFloat(s, typ string) (interface{}, error) {
	f, err := strconv.ParseFloat(s, 0)
	if err != nil {
		return 0, err
	}
	switch typ {
	case "float32":
		return float32(f), nil
	case "float64":
		return f, nil
	default:
		return 0, ErrUnknownFloatType
	}
}
