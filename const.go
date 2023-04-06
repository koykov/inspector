package inspector

// Op describes operation type.
type Op int

// LoopCtl describes loop control type.
type LoopCtl int

const (
	OpUnk Op = 0
	OpEq  Op = 1
	OpNq  Op = 2
	OpGt  Op = 3
	OpGtq Op = 4
	OpLt  Op = 5
	OpLtq Op = 6
	OpInc Op = 7
	OpDec Op = 8

	LoopCtlNone = 0
	LoopCtlBrk  = 1
	LoopCtlCnt  = 2

	Nil = "nil"

	FloatPrecision = 1e-3
)

var _, _, _, _ = OpUnk, OpInc, OpDec, LoopCtlNone
