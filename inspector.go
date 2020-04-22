package inspector

import (
	"errors"
	"strings"
)

type Inspector interface {
	Get(src interface{}, path ...string) (interface{}, error)
	GetTo(src interface{}, buf *interface{}, path ...string) error
	Set(dst, value interface{}, path ...string)
}

type BaseInspector struct{}

func (i *BaseInspector) StrConvSnippet(s, typ, _var string) (string, error) {
	if snippet, ok := convSnippet[typ]; ok {
		snippet := strings.Replace(snippet, "!{arg}", s, -1)
		snippet = strings.Replace(snippet, "!{var}", _var, -1)
	}
	return "", ErrNoConvFunc
}

var (
	convSnippet = map[string]string{}

	ErrNoConvFunc = errors.New("convert function doesn't exists")
)

func RegisterStrToFunc(typ, snippet string) {
	convSnippet[typ] = snippet
}

func init() {
	RegisterStrToFunc("int", strToIntSnippet("int"))
	RegisterStrToFunc("int8", strToIntSnippet("int8"))
	RegisterStrToFunc("int16", strToIntSnippet("int16"))
	RegisterStrToFunc("int32", strToIntSnippet("int32"))
	RegisterStrToFunc("int64", strToIntSnippet("int64"))

	RegisterStrToFunc("uint", strToUintSnippet("uint"))
	RegisterStrToFunc("uint8", strToUintSnippet("uint8"))
	RegisterStrToFunc("uint16", strToUintSnippet("uint16"))
	RegisterStrToFunc("uint32", strToUintSnippet("uint32"))
	RegisterStrToFunc("uint64", strToUintSnippet("uint64"))

	RegisterStrToFunc("float32", strToFloatSnippet("float32"))
	RegisterStrToFunc("float64", strToFloatSnippet("float64"))
}
