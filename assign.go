package inspector

// Signature of assign function.
//
// Assigns src to dst with dynamic typecasting.
// Returns true if typecasting success, false otherwise.
type AssignFn func(dst, src interface{}) bool

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
//
// Walks other registered assign callbacks and try to execute each of them.
// Stops when current callback return true.
func Assign(dst, src interface{}) (ok bool) {
	for _, fn := range assignFnRegistry {
		if ok = fn(dst, src); ok {
			return
		}
	}
	return
}
