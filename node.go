package inspector

// Node type.
// It's an internal representation of Go type.
type node struct {
	// Type flag, see constants above.
	typ typ
	// Type name as string.
	typn string
	// Underlying type name.
	typu string
	// If node has a parent, eg struct or map, this field contains the name or key.
	name string
	// Package name, requires for import suffixes.
	pkg string
	// Package path, requires for imports.
	pkgi string
	// Is node passed as pointer in parent or not.
	ptr bool
	// List of child nodes.
	chld []*node
	// If is a map, that node contains information about key's type.
	mapk *node
	// If is a map, that node contains information about value's type.
	mapv *node
	// If is a slice, that node contains information about value's type.
	slct *node
	// Flag if node contains bytes or string inside.
	hasb bool
	// Flag if node contains string/bytes/slice/map inside.
	hasc bool
}
