package inspector

import "strconv"

func (c *Compiler) writeNodeReset(node *node, v string, depth int) error {
	depths := strconv.Itoa(depth)
	depths1 := strconv.Itoa(depth + 1)
	mustLenCheck := node.typ != typeBasic
	if mustLenCheck {
		c.wl("if len(path)>", depths, "{")
	}
	switch node.typ {
	case typeStruct:
		for _, ch := range node.chld {
			nv := v + "." + ch.name
			c.wl("if path[", depths, "]==\"", ch.name, "\"{")
			chPtr := ch.ptr && (ch.typ == typeStruct || ch.typ == typeMap || ch.typ == typeSlice)
			if chPtr {
				c.wl("if ", nv, "!=nil{")
			}
			_ = c.writeNodeReset(ch, nv, depth+1)
			if ch.typ == typeMap {
				c.wl("for k:=range", c.fmtVd(ch, nv, depth+1), "{delete(", c.fmtVd(ch, nv, depth+1), ",k)}")
			}
			if ch.typ == typeSlice {
				c.wl(c.fmtVd(ch, nv, depth+1), "=", c.fmtVd(ch, nv, depth+1), "[:0]")
			}
			if chPtr {
				c.wl("}")
			}
			c.wl("}")
		}
	case typeMap:
		c.wl("if l:=len(", c.fmtVd(node, v, depth), ");l>0{")

		pname := c.fmtPt(node.mapk.typn)
		kv := "k" + depths
		c.wl("var ", kv, " ", pname, node.mapk.typn)
		c.wl("_=", kv)
		if node.mapk.typn == "string" {
			// Key is string, simple case.
			c.wl(kv, "=path["+depths+"]")
		} else {
			// Convert path value to the key type and try to find it in the map.
			c.wl(kv, "=", pname, node.mapk.typn, "(path[", depths, "])")
		}
		nv := "x" + strconv.Itoa(depth)
		c.wl(nv, ":=", c.fmtVd(node, v, depth), "[", c.fmtP(node.mapk, kv, depth+1), "]")
		c.wl("_=", nv)
		_ = c.writeNodeReset(node.mapv, nv, depth+1)
		c.wl(c.fmtVd(node, v, depth), "[", c.fmtP(node.mapk, kv, depth+1), "]", "=", nv)
		c.wl("}")
		c.wl("return nil")
	case typeSlice:
		c.wl("if l:=len(", c.fmtVd(node, v, depth), ");l>0{")
		if node.typn == "[]byte" {
			c.wl(c.fmtVd(node, v, depth), "=", c.fmtVd(node, v, depth), "[:0]")
		} else {
			iv := "i" + depths
			c.wl("var ", iv, " int=-1")
			c.wl("_= ", iv)
			if node.slct.typ != typeBasic {
				c.wl("_=", c.fmtVd(node, v, depth), "[l-1]")
				nv := "x" + strconv.Itoa(depth)

				snippet, imports, err := StrConvSnippet("path["+depths+"]", "int", "", iv)
				c.regImport(imports)
				if err != nil {
					return err
				}
				c.wl(snippet)
				pfx := "&"
				if node.slct.ptr {
					pfx = ""
				}
				c.wl(nv, ":=", pfx, c.fmtVd(node, v, depth), "[", iv, "]")
				_ = c.writeNodeReset(node.slct, nv, depth+1)
				c.wl("if len(path)==", depths1, "{")
				pfx = ""
				if node.slct.ptr {
					pfx = "&"
				}
				c.wl(c.fmtVd(node, v, depth), "[", iv, "]", "=", pfx, c.fmtT(node.slct), "{}")
				c.wl("}")
			}
			c.wl("if ", iv, "==-1{")
			c.wl(c.fmtVd(node, v, depth), "=", c.fmtVd(node, v, depth), "[:0]")
			c.wl("}")
		}
		c.wl("}")
		c.wl("return nil")
	case typeBasic:
		typ := node.typu
		if len(typ) == 0 {
			typ = node.typn
		}
		var def string
		switch typ {
		case "int", "int8", "int16", "int32", "int64",
			"uint", "uint8", "uint16", "uint32", "uint64",
			"float32", "float64", "byte", "rune":
			def = "0"
		case "bool":
			def = "false"
		case "string":
			def = `""`
		}
		if len(def) > 0 {
			c.wl(c.fmtVnb(node, v, depth), "=", def)
		}
	}
	if mustLenCheck {
		c.wl("}")
	}
	return nil
}

func (c *Compiler) writeNodeResetFull(node *node, v string, depth int) error {
	switch node.typ {
	case typeStruct:
		for _, ch := range node.chld {
			nv := v + "." + ch.name
			chPtr := ch.ptr && (ch.typ == typeStruct || ch.typ == typeMap || ch.typ == typeSlice)
			if chPtr {
				c.wl("if ", nv, "!=nil{")
			}
			_ = c.writeNodeResetFull(ch, nv, depth+1)
			if chPtr {
				c.wl("}")
			}
		}
	case typeMap:
		c.wl("if l:=len(", c.fmtVd(node, v, depth), ");l>0{")
		c.wl("for k,_:=range ", c.fmtVd(node, v, depth), "{")
		c.wl("delete(", c.fmtVd(node, v, depth), ",k)")
		c.wl("}")
		c.wl("}")
	case typeSlice:
		c.wl("if l:=len(", c.fmtVd(node, v, depth), ");l>0{")
		if node.typn == "[]byte" {
			c.wl(c.fmtVd(node, v, depth), "=", c.fmtVd(node, v, depth), "[:0]")
		} else {
			if node.slct.typ != typeBasic {
				c.wl("_=", c.fmtVd(node, v, depth), "[l-1]")
				nv := "x" + strconv.Itoa(depth)
				c.wl("for i:=0;i<l;i++{")
				pfx := "&"
				if node.slct.ptr {
					pfx = ""
				}
				c.wl(nv, ":=", pfx, c.fmtVd(node, v, depth), "[i]")
				_ = c.writeNodeResetFull(node.slct, nv, depth+1)
				c.wl("}")
			}
			c.wl(c.fmtVd(node, v, depth), "=", c.fmtVd(node, v, depth), "[:0]")
		}
		c.wl("}")
	case typeBasic:
		typ := node.typu
		if len(typ) == 0 {
			typ = node.typn
		}
		var def string
		switch typ {
		case "int", "int8", "int16", "int32", "int64",
			"uint", "uint8", "uint16", "uint32", "uint64",
			"float32", "float64", "byte", "rune":
			def = "0"
		case "bool":
			def = "false"
		case "string":
			def = `""`
		}
		if len(def) > 0 {
			c.wl(c.fmtVnb(node, v, depth), "=", def)
		}
	}
	return nil
}
