package inspector

import (
	"errors"
)

func marshal(w ByteStringWriter, nodes []*node) error {
	if len(nodes) == 0 {
		return errors.New("no nodes provided")
	}
	_, _ = w.WriteString(`<?xml version="1.0" encoding="UTF-8"?>`)
	_, _ = w.WriteString("\n<nodes>\n")
	for i := 0; i < len(nodes); i++ {
		node := nodes[i]
		if err := marshalNode(w, "node", node, 1); err != nil {
			return err
		}
	}
	_, _ = w.WriteString("</nodes>\n")
	return nil
}

func marshalNode(w ByteStringWriter, tag string, node *node, depth int) error {
	pad(w, depth)
	_ = w.WriteByte('<')
	_, _ = w.WriteString(tag)
	_, _ = w.WriteString(` type="`)
	_, _ = w.WriteString(node.typ.String())
	_, _ = w.WriteString(`"`)
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
	if len(node.name) > 0 {
		_, _ = w.WriteString(` name="`)
		_, _ = w.WriteString(node.name)
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
	_, _ = w.WriteString(">")
	_, _ = w.WriteString("</")
	_, _ = w.WriteString(tag)
	_, _ = w.WriteString(">\n")
	return nil
}

func pad(w ByteStringWriter, c int) {
	for i := 0; i < c; i++ {
		_ = w.WriteByte('\t')
	}
}
