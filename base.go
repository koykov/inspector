package inspector

// BaseInspector describes base struct.
type BaseInspector struct{}

// DEQMustCheck makes a decision must field (by given path) checks or not during DeepEqual().
func (i BaseInspector) DEQMustCheck(path string, options *DEQOptions) bool {
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
