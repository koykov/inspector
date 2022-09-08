package testobj

type TestObject1 struct {
	IntSlice       []int32
	IntPtrSlice    []*int32
	IntSlicePtr    *[]int32
	IntPtrSlicePtr *[]*int32

	FloatSlice       TestFloatSlice
	FloatPtrSlice    TestFloatPtrSlice
	FloatSlicePtr    *TestFloatSlice
	FloatPtrSlicePtr *TestFloatPtrSlice

	IntStringMap          map[int]string
	IntStringPtrMap       map[int]*string
	IntStringMapPtr       *map[int]string
	IntStringPtrMapPtr    *map[int]*string
	IntPtrStringPtrMapPtr *map[*int]*string
}

type TestFloatSlice []float32
type TestFloatPtrSlice []*float32
