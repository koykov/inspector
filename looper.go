package inspector

// Looper signature.
type Looper interface {
	// RequireKey checks set key requirement.
	RequireKey() bool
	// SetKey sets the key value and inspector to hidden context.
	SetKey(val any, ins Inspector)
	// SetVal sets the value and inspector to context.
	SetVal(val any, ins Inspector)
	// Iterate performs the iteration.
	Iterate() LoopCtl
}
