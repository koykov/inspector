package inspector

import "errors"

type Inspector interface {
	Get(src interface{}, path ...string) (interface{}, error)
	GetTo(src interface{}, buf *interface{}, path ...string) error
	Set(dst, value interface{}, path ...string)
	Cmp(src interface{}, cond Op, right string, result *bool, path ...string) error
	Loop(src interface{}, l Looper, buf *[]byte, path ...string) error
}

type Looper interface {
	RequireKey() bool
	SetKey(val interface{}, ins Inspector)
	SetVal(val interface{}, ins Inspector)
	Loop()
}

type BaseInspector struct{}

var (
	inspectorRegistry = map[string]Inspector{}

	ErrUnknownInspector = errors.New("unknown inspector")
)

func RegisterInspector(name string, ins Inspector) {
	inspectorRegistry[name] = ins
}

func GetInspector(name string) (Inspector, error) {
	if ins, ok := inspectorRegistry[name]; ok {
		return ins, nil
	}
	return nil, ErrUnknownInspector
}

func init() {
	RegisterInspector("static", &StaticInspector{})
	RegisterInspector("reflect", &ReflectInspector{})

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
