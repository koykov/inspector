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

func (n node) write(w ByteStringWriter) error {
	_, _ = w.WriteString(`<?xml version="1.0" encoding="UTF-8"?>`)
	_ = w.WriteByte('\n')
	return n.write_(w, "node", &n, 0)
}

func (n node) write_(w ByteStringWriter, tag string, node *node, depth int) error {
	pad(w, depth)
	_ = w.WriteByte('<')
	_, _ = w.WriteString(tag)
	_, _ = w.WriteString(` type="`)
	_, _ = w.WriteString(node.typ.String())
	_, _ = w.WriteString(`"`)
	if len(node.name) > 0 {
		_, _ = w.WriteString(` name="`)
		_, _ = w.WriteString(node.name)
		_, _ = w.WriteString(`"`)
	}
	if len(node.typn) > 0 {
		_, _ = w.WriteString(` typeName="`)
		_, _ = w.WriteString(node.typn)
		_, _ = w.WriteString(`"`)
	}
	if len(node.typu) > 0 {
		_, _ = w.WriteString(` underlyingName="`)
		_, _ = w.WriteString(node.typu)
		_, _ = w.WriteString(`"`)
	}
	if len(node.pkg) > 0 {
		_, _ = w.WriteString(` package="`)
		_, _ = w.WriteString(node.pkg)
		_, _ = w.WriteString(`"`)
	}
	if len(node.pkgi) > 0 {
		_, _ = w.WriteString(` packageImport="`)
		_, _ = w.WriteString(node.pkgi)
		_, _ = w.WriteString(`"`)
	}
	if node.ptr {
		_, _ = w.WriteString(` pointer="true"`)
	}
	if node.hasb {
		_, _ = w.WriteString(` hasBytes="true"`)
	}
	if node.hasc {
		_, _ = w.WriteString(` hasLC="true"`)
	}

	hasNL := node.mapk != nil || node.mapv != nil || node.slct != nil || len(node.chld) > 0
	if hasNL {
		_, _ = w.WriteString(">")
		_ = w.WriteByte('\n')
	} else {
		_, _ = w.WriteString("/>")

	}

	if node.mapk != nil {
		_ = n.write_(w, "mapKey", node.mapk, depth+1)
	}
	if node.mapv != nil {
		_ = n.write_(w, "mapValue", node.mapv, depth+1)
	}
	if node.slct != nil {
		_ = n.write_(w, "slice", node.slct, depth+1)
	}
	if len(node.chld) > 0 {
		pad(w, depth+1)
		_, _ = w.WriteString("<nodes>\n")
		for i := 0; i < len(node.chld); i++ {
			child := node.chld[i]
			_ = n.write_(w, "node", child, depth+2)
		}
		pad(w, depth+1)
		_, _ = w.WriteString("</nodes>\n")
	}

	if hasNL {
		pad(w, depth)
		_, _ = w.WriteString("</")
		_, _ = w.WriteString(tag)
		_, _ = w.WriteString(">\n")
	} else {
		_ = w.WriteByte('\n')
	}
	return nil
}

func pad(w ByteStringWriter, c int) {
	for i := 0; i < c; i++ {
		_ = w.WriteByte('\t')
	}
}
