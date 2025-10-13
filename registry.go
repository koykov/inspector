package inspector

var (
	// Global registry of all inspectors.
	inspectorRegistry = map[string]Inspector{}

	_, _ = GetInspector, EachInspector
)

// RegisterInspector saves inspector to the registry.
func RegisterInspector(name string, ins Inspector) {
	inspectorRegistry[name] = ins
}

// GetInspector returns inspector from the registry.
func GetInspector(name string) (Inspector, error) {
	if ins, ok := inspectorRegistry[name]; ok {
		return ins, nil
	}
	return nil, ErrUnknownInspector
}

// EachInspector visits all registered inspectors.
func EachInspector(fn func(name string, ins Inspector)) {
	for name, ins := range inspectorRegistry {
		fn(name, ins)
	}
}
