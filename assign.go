package inspector

type AssignFn func(dst, src interface{}) bool

var (
	assignFnRegistry = make([]AssignFn, 0)
)

func RegisterAssignFn(fn AssignFn) {
	for _, f := range assignFnRegistry {
		if &f == &fn {
			return
		}
	}
	assignFnRegistry = append(assignFnRegistry, fn)
}

func Assign(dst, src interface{}) (ok bool) {
	for _, fn := range assignFnRegistry {
		if ok = fn(dst, src); ok {
			return
		}
	}
	return
}
