package testobj

type TestObject1 struct {
	IntSlice       []int32
	IntPtrSlice    []*int32
	IntSlicePtr    *[]int32
	IntPtrSlicePtr *[]*int32
	ByteSlice      []byte
	BytePtrSlice   []*byte
	// ByteSlicePtr    *[]byte
	// BytePtrSlicePtr *[]*byte

	FloatSlice       TestFloatSlice
	FloatPtrSlice    TestFloatPtrSlice
	FloatSlicePtr    *TestFloatSlice
	FloatPtrSlicePtr *TestFloatPtrSlice

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
}

type TestFloatSlice []float32
type TestFloatPtrSlice []*float32

type TestStringFloatMap map[string]float64
type TestStringFloatPtrMap map[string]*float64
type TestStringPtrFloatPtrMap map[*string]*float64
