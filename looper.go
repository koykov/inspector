package inspector

// Looper signature.
type Looper interface {
	// RequireKey checks set key requirement.
	RequireKey() bool
	// SetKey sets the key value and inspector to hidden context.
	SetKey(val interface{}, ins Inspector)
	// SetVal sets the value and inspector to context.
	SetVal(val interface{}, ins Inspector)
	// Iterate performs the iteration.
	Iterate() LoopCtl
}
