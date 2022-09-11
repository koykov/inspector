package inspector

var (
	// Global registry of all inspectors.
	inspectorRegistry = map[string]Inspector{}

	_ = GetInspector
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
