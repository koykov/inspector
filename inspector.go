package inspector

// Inspector signature.
type Inspector interface {
	// TypeName returns name of underlying type.
	TypeName() string
	// Get returns value from src according path.
	Get(src interface{}, path ...string) (interface{}, error)
	// GetTo writes value from src to buf according path.
	GetTo(src interface{}, buf *interface{}, path ...string) error
	// Set sets value to dst according path.
	Set(dst, value interface{}, path ...string) error
	// SetWB is a buffered version of Set().
	SetWB(dst, value interface{}, buf AccumulativeBuffer, path ...string) error
	// Cmp compares value according path with right using cond.
	// Result will be present in result.
	Cmp(src interface{}, cond Op, right string, result *bool, path ...string) error
	// Loop iterates in src value taking by path using l looper.
	Loop(src interface{}, l Looper, buf *[]byte, path ...string) error
	// DeepEqual compares l and r.
	DeepEqual(l, r interface{}) bool
	// DeepEqualWithOptions compares l and r corresponding options.
	DeepEqualWithOptions(l, r interface{}, options *DEQOptions) bool
	// Parse parses encoded data according encoding type.
	Parse([]byte, Encoding) (interface{}, error)
}

// Looper signature.
type Looper interface {
	// RequireKey checks set key requirement.
	RequireKey() bool
	// SetKey sets the key value and inspector to hidden context.
	SetKey(val interface{}, ins Inspector)
	// SetVal sets the value and inspector to context.
	SetVal(val interface{}, ins Inspector)
	// Iterate performs the iteration.
	Iterate() LoopCtl
}

// AccumulativeBuffer describes buffer signature.
// Collects data during assign functions work.
type AccumulativeBuffer interface {
	// AcquireBytes returns more space to use.
	AcquireBytes() []byte
	// ReleaseBytes returns space to the buffer.
	ReleaseBytes([]byte)
	// Reset all accumulated data.
	Reset()
}

// BaseInspector describes base struct.
type BaseInspector struct{}

var (
	// Global registry of all inspectors.
	inspectorRegistry = map[string]Inspector{}

	_ = GetInspector
)

// RegisterInspector saves inspector to the registry.
func RegisterInspector(name string, ins Inspector) {
	inspectorRegistry[name] = ins
}

// GetInspector returns inspector from the registry.
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

	// Register functions to typecast to built-in types.
	RegisterAssignFn(AssignToBytes)
	RegisterAssignFn(AssignToStr)
	RegisterAssignFn(AssignToBool)
	RegisterAssignFn(AssignToInt)
	RegisterAssignFn(AssignToUint)
	RegisterAssignFn(AssignToFloat)
}
