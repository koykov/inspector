package inspector

import "errors"

type Inspector interface {
	Get(src, buf interface{}, path ...string) interface{}
	Set(dst, value interface{}, path ...string)
}

type BaseInspector struct {
	Ebuf interface{}
}

func (i *BaseInspector) StrTo(s, typ string) (interface{}, error) {
	if f, ok := conv[typ]; ok {
		return f(s, typ)
	}
	return nil, ErrNoConvFunc
}

type StrToFunc func(s, typ string) (interface{}, error)

var (
	conv = map[string]StrToFunc{}

	ErrNoConvFunc = errors.New("convert function doesn't exists")
)

func RegisterStrToFunc(typ string, f StrToFunc) {
	conv[typ] = f
}

func init() {
	RegisterStrToFunc("int", strToInt)
	RegisterStrToFunc("int8", strToInt)
	RegisterStrToFunc("int16", strToInt)
	RegisterStrToFunc("int32", strToInt)
	RegisterStrToFunc("int64", strToInt)
	RegisterStrToFunc("uint", strToInt)
	RegisterStrToFunc("uint8", strToInt)
	RegisterStrToFunc("uint16", strToInt)
	RegisterStrToFunc("uint32", strToInt)
	RegisterStrToFunc("uint64", strToInt)

	RegisterStrToFunc("float32", strToFloat)
	RegisterStrToFunc("float64", strToFloat)
}
