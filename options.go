package inspector

// DEQOptions describes DeepEqual options.
type DEQOptions struct {
	// Float fields comparison precision.
	Precision float64
	// List of fields to exclude from check.
	Exclude map[string]struct{}
	// Limit fields to check.
	Filter map[string]struct{}
}

// DEQMustCheck makes a decision must field (by given path) checks or not during DeepEqual().
func DEQMustCheck(path string, options *DEQOptions) bool {
	if options == nil {
		return true
	}
	if len(options.Exclude) > 0 {
		_, ok := options.Exclude[path]
		return !ok
	}
	if len(options.Filter) > 0 {
		_, ok := options.Filter[path]
		return ok
	}
	return true
}
