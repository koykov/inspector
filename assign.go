package inspector

// Signature of assign function.
//
// Assigns src to dst with dynamic typecasting.
// Returns true if typecasting success, false otherwise.
type AssignFn func(dst, src interface{}, buf AccumulativeBuffer) bool

var (
	// Registry of assign functions.
	assignFnRegistry = make([]AssignFn, 0)
)

// Add new assign function to the registry.
func RegisterAssignFn(fn AssignFn) {
	for _, f := range assignFnRegistry {
		if &f == &fn {
			return
		}
	}
	assignFnRegistry = append(assignFnRegistry, fn)
}

// Generic assign function.
func Assign(dst, src interface{}) (ok bool) {
	return AssignBuf(dst, src, nil)
}

// Generic buffered assign function.
//
// Walks other registered assign callbacks and try to execute each of them.
// Stops when current callback return true.
func AssignBuf(dst, src interface{}, buf AccumulativeBuffer) (ok bool) {
	for _, fn := range assignFnRegistry {
		if ok = fn(dst, src, buf); ok {
			return
		}
	}
	return
}
