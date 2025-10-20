package inspector

// Inspector signature.
type Inspector interface {
	// TypeName returns name of underlying type.
	TypeName() string
	// Instance creates new instance (or pointer to instance) of underlying type.
	Instance(ptr bool) any
	// Get returns value from src according path.
	Get(src any, path ...string) (any, error)
	// GetTo writes value from src to buf according path.
	GetTo(src any, buf *any, path ...string) error
	// Set sets value to dst according path.
	Set(dst, value any, path ...string) error
	// SetWithBuffer is a buffered version of Set().
	SetWithBuffer(dst, value any, buf AccumulativeBuffer, path ...string) error
	// Compare applies condition cond to value in src by path and right.
	// Result will set to result buffer.
	Compare(src any, cond Op, right string, result *bool, path ...string) error
	// Loop iterates in src value taking by path using iterator iter.
	Loop(src any, iter Iterator, buf *[]byte, path ...string) error
	// DeepEqual compares l and r.
	DeepEqual(l, r any) bool
	// DeepEqualWithOptions compares l and r corresponding options.
	DeepEqualWithOptions(l, r any, options *DEQOptions) bool
	// Unmarshal parses encoded data according encoding type.
	Unmarshal(p []byte, typ Encoding) (any, error)
	// Copy makes a copy of x.
	Copy(x any) (any, error)
	// CopyTo makes a copy of src to dst using buffer.
	CopyTo(src, dst any, buf AccumulativeBuffer) error
	// Append adds value to the end of dst (dst must be a slice).
	Append(src, value any, path ...string) (any, error)
	// Length puts length of field value by path in src to the buffer.
	Length(src any, result *int, path ...string) error
	// Capacity puts capacity of field value by path in src to the buffer.
	Capacity(src any, result *int, path ...string) error
	// Reset resets x.
	Reset(x any) error
}

func init() {
	// Register inspectors known by default.
	RegisterInspector("static", StaticInspector{})
	RegisterInspector("[]string", StringsInspector{})
	RegisterInspector("strings", StringsInspector{})
	RegisterInspector("[][]byte", StringsInspector{})
	RegisterInspector("bytes", StringsInspector{})
	RegisterInspector("stringAnyMap", StringAnyMapInspector{})
	RegisterInspector("map[string]any", StringAnyMapInspector{})
	RegisterInspector("reflect", ReflectInspector{})

	// Register snippets to convert string to built-in types.
	imp := []string{`"strconv"`}
	RegisterStrToXFn("bool", strToBoolSnippet("bool"), imp)

	RegisterStrToXFn("int", strToIntSnippet("int"), imp)
	RegisterStrToXFn("int8", strToIntSnippet("int8"), imp)
	RegisterStrToXFn("int16", strToIntSnippet("int16"), imp)
	RegisterStrToXFn("int32", strToIntSnippet("int32"), imp)
	RegisterStrToXFn("int64", strToIntSnippet("int64"), imp)

	RegisterStrToXFn("uint", strToUintSnippet("uint"), imp)
	RegisterStrToXFn("uint8", strToUintSnippet("uint8"), imp)
	RegisterStrToXFn("uint16", strToUintSnippet("uint16"), imp)
	RegisterStrToXFn("uint32", strToUintSnippet("uint32"), imp)
	RegisterStrToXFn("uint64", strToUintSnippet("uint64"), imp)

	RegisterStrToXFn("float32", strToFloatSnippet("float32"), imp)
	RegisterStrToXFn("float64", strToFloatSnippet("float64"), imp)

	imp = []string{`"bytes"`, `"github.com/koykov/byteconv"`}
	RegisterStrToXFn("[]byte", strToBytesSnippet("[]byte"), imp)
	RegisterStrToXFn("string", strToStrSnippet("string"), nil)
	RegisterStrToXFn("byte", strToByteSnippet("byte"), []string{`"github.com/koykov/byteconv"`})

	// Register functions to typecast to built-in types.
	RegisterAssignFn(AssignToBytes)
	RegisterAssignFn(AssignToStr)
	RegisterAssignFn(AssignToBool)
	RegisterAssignFn(AssignToInt)
	RegisterAssignFn(AssignToUint)
	RegisterAssignFn(AssignToFloat)
}
