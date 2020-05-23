package inspector

type Op int
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
)
