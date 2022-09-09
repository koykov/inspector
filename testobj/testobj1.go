package testobj

type TestObject1 struct {
	IntSlice        []int32
	IntPtrSlice     []*int32
	IntSlicePtr     *[]int32
	IntPtrSlicePtr  *[]*int32
	ByteSlice       []byte
	BytePtrSlice    []*byte
	ByteSlicePtr    *[]byte
	BytePtrSlicePtr *[]*byte

	FloatSlice       TestFloatSlice
	FloatPtrSlice    TestFloatPtrSlice
	FloatSlicePtr    *TestFloatSlice
	FloatPtrSlicePtr *TestFloatPtrSlice

	StructSlice       []TestStruct
	StructPtrSlice    []*TestStruct
	StructSlicePtr    *[]TestStruct
	StructPtrSlicePtr *[]*TestStruct

	IntStringMap          map[int]string
	IntStringPtrMap       map[int]*string
	IntStringMapPtr       *map[int]string
	IntStringPtrMapPtr    *map[int]*string
	IntPtrStringPtrMapPtr *map[*int]*string

	StringFloatMap          TestStringFloatMap
	StringFloatPtrMap       TestStringFloatPtrMap
	StringFloatMapPtr       *TestStringFloatMap
	StringFloatPtrMapPtr    *TestStringFloatPtrMap
	StringPtrFloatPtrMapPtr *TestStringPtrFloatPtrMap

	FloatStructMap          map[float64]TestStruct
	FloatStructPtrMap       map[float64]*TestStruct
	FloatPtrStructMap       map[*float64]TestStruct
	FloatPtrStructPtrMap    map[*float64]*TestStruct
	FloatPtrStructPtrMapPtr *map[*float64]*TestStruct
}

type TestFloatSlice []float32
type TestFloatPtrSlice []*float32

type TestStringFloatMap map[string]float64
type TestStringFloatPtrMap map[string]*float64
type TestStringPtrFloatPtrMap map[*string]*float64

type TestStruct struct {
	A   byte
	S   string
	B   []byte
	I   int
	I8  int8
	I16 int16
	I32 int32
	I64 int64
	U   uint
	U8  uint8
	U16 uint16
	U32 uint32
	U64 uint64
	F   float32
	D   float64
}
