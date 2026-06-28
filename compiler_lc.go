package inspector

import "strconv"

func (c *Compiler) writeNodeLC(node_ *node, v, fn string, depth int) error {
	if depth == 0 {
		c.wl("*result=0")
	}
	depths := strconv.Itoa(depth)

	requireLenCheck := func(node *node) bool {
		return node.typ == typeStruct || node.typ == typeMap || (node.typ == typeSlice && node.typu != "[]byte")
	}

	if node_.ptr {
		// Value may be nil on pointer types.
		c.wl("if ", v, " == nil { return nil }")
	}
	if depth == 0 && requireLenCheck(node_) {
		c.wl("if len(path) == 0 { return nil }")
	}

	switch node_.typ {
	case typeStruct:
		for _, ch := range node_.chld {
			if (ch.typ == typeBasic && ch.typu != "string") || !ch.hasc {
				continue
			}
			c.wl("if path[", depths, "] == ", `"`, ch.name, `" {`)
			nv := v + "." + ch.name
			chPtr := ch.ptr && (ch.typ == typeStruct || ch.typ == typeMap || ch.typ == typeSlice)
			if chPtr {
				c.wl("if ", nv, "!=nil{")
			}
			_ = c.writeNodeLC(ch, nv, fn, depth+1)
			if chPtr {
				c.wl("}")
			}
			c.wl("}")
		}
	case typeMap:
		origPtr := node_.ptr
		if depth == 0 {
			node_.ptr = true
		}
		if fn == "len" {
			c.wl("if len(path)==", depths, "{")
			c.wl("*result=", fn, "(", c.fmtVnb(node_, v, depth), ")")
			c.wl("return nil")
			c.wl("}")
		}
		if !node_.mapv.hasc {
			return nil
		}
		nv := "x" + strconv.Itoa(depth)
		c.wl("if len(path) < ", strconv.Itoa(depth+1), " { return nil }")
		if node_.mapk.typn == "string" {
			// Key is string, simple case.
			key := c.fmtP(node_.mapk, "path["+depths+"]", depth+1)
			c.wl("if ", nv, ", ok := ", c.fmtV(node_, v), "[", key, "]; ok {")
			c.wl("_ = ", nv)
			if requireLenCheck(node_.mapv) {
				c.wl("if len(path) < ", strconv.Itoa(depth+2), " { return nil }")
			}
			err := c.writeNodeLC(node_.mapv, nv, fn, depth+1)
			if err != nil {
				return err
			}
			c.wl("}")
		} else {
			// Convert path value to the key type and try to find it in the map.
			c.wl("var k ", node_.mapk.typn)
			snippet, imports, err := StrConvSnippet("path["+depths+"]", node_.mapk.typn, node_.mapk.typu, "k")
			c.regImport(imports)
			if err != nil {
				return err
			}
			c.wl(snippet)
			c.wl(nv, " := ", c.fmtV(node_, v), "[", c.fmtP(node_.mapk, "k", depth+1), "]")
			c.wl("_ = ", nv)
			if requireLenCheck(node_.mapv) {
				c.wl("if len(path) < ", strconv.Itoa(depth+2), " { return nil }")
			}
			err = c.writeNodeLC(node_.mapv, nv, fn, depth+1)
			if err != nil {
				return err
			}
		}
		node_.ptr = origPtr
	case typeSlice:
		if node_.typn == "[]byte" {
			c.wl("*result=", fn, "(", c.fmtVnb(node_, v, depth), ")")
			c.wl("return nil")
		} else {
			if !node_.slct.hasc {
				return nil
			}
			c.wl("if len(path)==", depths, "{")
			c.wl("*result=", fn, "(", c.fmtVnb(node_, v, depth), ")")
			c.wl("return nil")
			c.wl("}")
			c.wl("if len(path) < ", strconv.Itoa(depth+1), " { return nil }")

			nv := "x" + strconv.Itoa(depth)
			c.wl("var i int")
			snippet, imports, err := StrConvSnippet("path["+depths+"]", "int", "", "i")
			c.regImport(imports)
			if err != nil {
				return err
			}
			c.wl(snippet)
			c.wl("if len(", c.fmtVnb(node_, v, depth), ") > i {")
			if node_.slct.ptr || c.isBuiltin(node_.slct.typn) {
				c.wl(nv, " := ", c.fmtVd(node_, v, depth), "[i]")
			} else {
				c.wl(nv, " := &", c.fmtVd(node_, v, depth), "[i]")
			}
			c.wl("_ = ", nv)
			if requireLenCheck(node_.slct) {
				c.wl("if len(path) < ", strconv.Itoa(depth+2), " { return nil }")
			}
			err = c.writeNodeLC(node_.slct, nv, fn, depth+1)
			if err != nil {
				return err
			}
			c.wl("}")
		}
	case typeBasic:
		if node_.typu == "string" && fn == "len" {
			c.wl("*result=", fn, "(", c.fmtVnb(node_, v, depth), ")")
			c.wl("return nil")
		}
	}

	return nil
}
