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
}

type TestFloatSlice []float32
type TestFloatPtrSlice []*float32
