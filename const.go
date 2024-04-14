package inspector

// Op describes operation type.
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

// LoopCtl describes loop control type.
type LoopCtl int

const (
	LoopCtlNone LoopCtl = iota
	LoopCtlBrk
	LoopCtlCnt
)

const Nil = "nil"
const FloatPrecision = 1e-3

var _, _, _, _ = OpUnk, OpInc, OpDec, LoopCtlNone
