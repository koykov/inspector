package inspector

import (
	"errors"
	"strings"
)

type Inspector interface {
	Get(src interface{}, path ...string) (interface{}, error)
	GetTo(src interface{}, buf *interface{}, path ...string) error
	Set(dst, value interface{}, path ...string)
	Cmp(src interface{}, cond Op, right string, result *bool, path ...string) error
	Loop(src interface{}, ctx ContextPooler, buf *[]byte, path ...string) error
}

type ContextPooler interface {
	Set(key string, val interface{}, ins Inspector)
	Loop()
}

type BaseInspector struct{}

type StrConv struct {
	Snippet string
	Import  []string
}

func StrConvSnippet(s, typ, _var string) (string, []string, error) {
	if sc, ok := convSnippet[typ]; ok {
		snippet := strings.Replace(sc.Snippet, "!{arg}", s, -1)
		snippet = strings.Replace(snippet, "!{var}", _var, -1)
		return snippet, sc.Import, nil
	}
	return "", nil, ErrNoConvFunc
}

var (
	convSnippet = map[string]StrConv{}

	ErrNoConvFunc = errors.New("convert function doesn't exists")
)

func RegisterStrToFunc(typ, snippet string, imports []string) {
	convSnippet[typ] = StrConv{snippet, imports}
}

func init() {
	imp := []string{`"strconv"`}
	RegisterStrToFunc("bool", strToBoolSnippet("bool"), imp)

	RegisterStrToFunc("int", strToIntSnippet("int"), imp)
	RegisterStrToFunc("int8", strToIntSnippet("int8"), imp)
	RegisterStrToFunc("int16", strToIntSnippet("int16"), imp)
	RegisterStrToFunc("int32", strToIntSnippet("int32"), imp)
	RegisterStrToFunc("int64", strToIntSnippet("int64"), imp)

	RegisterStrToFunc("uint", strToUintSnippet("uint"), imp)
	RegisterStrToFunc("uint8", strToUintSnippet("uint8"), imp)
	RegisterStrToFunc("uint16", strToUintSnippet("uint16"), imp)
	RegisterStrToFunc("uint32", strToUintSnippet("uint32"), imp)
	RegisterStrToFunc("uint64", strToUintSnippet("uint64"), imp)

	RegisterStrToFunc("float32", strToFloatSnippet("float32"), imp)
	RegisterStrToFunc("float64", strToFloatSnippet("float64"), imp)

	imp = []string{`"bytes"`, `"github.com/koykov/fastconv"`}
	RegisterStrToFunc("[]byte", strToBytesSnippet("[]byte"), imp)
	RegisterStrToFunc("string", strToStrSnippet("string"), nil)
}
