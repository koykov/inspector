package inspector

type StaticInspector struct {
	BaseInspector
}

func (i *StaticInspector) Get(src interface{}, _ ...string) (interface{}, error) {
	return src, nil
}

func (i *StaticInspector) GetTo(src interface{}, buf *interface{}, _ ...string) error {
	*buf = &src
	return nil
}

func (i *StaticInspector) Set(_, _ interface{}, _ ...string) {
	//
}

func (i *StaticInspector) Cmp(_ interface{}, _ Op, _ string, result *bool, _ ...string) error {
	*result = false
	return nil
}

func (i *StaticInspector) Loop(_ interface{}, _ ContextPooler, _ ...string) error {
	return nil
}
