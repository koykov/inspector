package inspector

// BaseInspector describes base struct.
type BaseInspector struct{}

func (i BaseInspector) TypeName() string {
	return "not implement!"
}

func (i BaseInspector) Instance(_ bool) any { return nil }

func (i BaseInspector) Get(_ any, _ ...string) (any, error) {
	return nil, ErrNotImplement
}

func (i BaseInspector) GetTo(_ any, _ *any, _ ...string) error {
	return ErrNotImplement
}

func (i BaseInspector) Set(_, _ any, _ ...string) error {
	return ErrNotImplement
}

func (i BaseInspector) SetWithBuffer(_, _ any, _ AccumulativeBuffer, _ ...string) error {
	return ErrNotImplement
}

func (i BaseInspector) Compare(_ any, _ Op, _ string, _ *bool, _ ...string) error {
	return ErrNotImplement
}

func (i BaseInspector) Loop(_ any, _ Iterator, _ *[]byte, _ ...string) error {
	return ErrNotImplement
}

func (i BaseInspector) DeepEqual(l, r any) bool {
	return i.DeepEqualWithOptions(l, r, nil)
}

func (i BaseInspector) DeepEqualWithOptions(_, _ any, _ *DEQOptions) bool {
	return false
}

func (i BaseInspector) Unmarshal(_ []byte, _ Encoding) (any, error) {
	return nil, ErrNotImplement
}

func (i BaseInspector) Copy(_ any) (dst any, err error) {
	return nil, ErrNotImplement
}

func (i BaseInspector) CopyTo(_, _ any, _ AccumulativeBuffer) error {
	return ErrNotImplement
}

func (i BaseInspector) Length(_ any, _ *int, _ ...string) error {
	return ErrNotImplement
}

func (i BaseInspector) Capacity(_ any, _ *int, _ ...string) error {
	return ErrNotImplement
}

func (i BaseInspector) Reset(_ any) error {
	return ErrNotImplement
}
