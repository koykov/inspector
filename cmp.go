package inspector

// Must be sync with cbytetpl.Op
// Copied to avoid dependency with that package.
type Op int

const (
	OpUnk Op = iota
	OpEq
	OpNq
	OpGt
	OpGtq
	OpLt
	OpLtq
	OpInc
	OpDec
)
