package inspector

// AssignFn describes signature of assign function.
//
// Assigns src to dst with dynamic typecasting.
// Returns true if typecasting success, false otherwise.
type AssignFn func(dst, src any, buf AccumulativeBuffer) bool

var (
	// Registry of assign functions.
	assignFnRegistry = make([]AssignFn, 0)
)

// RegisterAssignFn registers new assign function to the registry.
func RegisterAssignFn(fn AssignFn) {
	for _, f := range assignFnRegistry {
		if &f == &fn {
			return
		}
	}
	assignFnRegistry = append(assignFnRegistry, fn)
}

// Assign is a generic assign function.
func Assign(dst, src any) (ok bool) {
	return AssignBuf(dst, src, nil)
}

// AssignBuf is a generic buffered assign function.
//
// Walks other registered assign callbacks and try to execute each of them.
// Stops when current callback return true.
func AssignBuf(dst, src any, buf AccumulativeBuffer) (ok bool) {
	for _, fn := range assignFnRegistry {
		if ok = fn(dst, src, buf); ok {
			return
		}
	}
	return
}

var _ = Assign
