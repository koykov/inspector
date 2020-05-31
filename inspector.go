package inspector

import "errors"

// Inspector signature.
type Inspector interface {
	// Get value from src according path.
	Get(src interface{}, path ...string) (interface{}, error)
	// Get value from src to buf according path.
	GetTo(src interface{}, buf *interface{}, path ...string) error
	// Set value ti dst according path.
	Set(dst, value interface{}, path ...string)
	// Compare value according path with right using cond.
	// Result will be present in result.
	Cmp(src interface{}, cond Op, right string, result *bool, path ...string) error
	// Iterate in src value taking by path using l looper.
	Loop(src interface{}, l Looper, buf *[]byte, path ...string) error
}

// Looper signature.
type Looper interface {
	// Need to set key.
	RequireKey() bool
	// Set the key value and inspector to hidden context.
	SetKey(val interface{}, ins Inspector)
	// Set the value and inspector to context.
	SetVal(val interface{}, ins Inspector)
	// Main method to perform the iteration.
	Iterate() LoopCtl
}

// Need for possible further improvements.
type BaseInspector struct{}

var (
	// Global registry of all inspectors.
	inspectorRegistry = map[string]Inspector{}

	ErrUnknownInspector = errors.New("unknown inspector")
)

// Save inspector to the registry.
func RegisterInspector(name string, ins Inspector) {
	inspectorRegistry[name] = ins
}

// Get inspector from the registry.
func GetInspector(name string) (Inspector, error) {
	if ins, ok := inspectorRegistry[name]; ok {
		return ins, nil
	}
	return nil, ErrUnknownInspector
}

func init() {
	// Register inspectors known by default.
	RegisterInspector("static", &StaticInspector{})
	RegisterInspector("reflect", &ReflectInspector{})

	// Register snippets to convert string to built-in types.
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
