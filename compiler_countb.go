package inspector

import "strconv"

func (c *Compiler) writeCountBytes(node *node, v string, depth int) error {
	if !node.hasb {
		return nil
	}
	switch node.typ {
	case typeStruct:
		for _, ch := range node.chld {
			if !ch.hasb {
				continue
			}
			nv := v + "." + ch.name
			chPtr := ch.ptr && (ch.typ == typeStruct || ch.typ == typeMap || ch.typ == typeSlice)
			if chPtr {
				c.wl("if ", nv, "!=nil{")
			}
			_ = c.writeCountBytes(ch, nv, depth+1)
			if chPtr {
				c.wl("}")
			}
		}
	case typeMap:
		nk := "k" + strconv.Itoa(depth)
		nx := "v" + strconv.Itoa(depth)
		c.wl("for ", nk, ", ", nx, " := range ", c.fmtVnb(node, v, depth), "{")
		c.wl("_,_=", nk, ",", nx)
		_ = c.writeCountBytes(node.mapk, nk, depth+1)
		_ = c.writeCountBytes(node.mapv, nx, depth+1)
		c.wl("}")
	case typeSlice:
		if node.typn == "[]byte" {
			c.wl("c+=len(", c.fmtVnb(node, v, depth), ")")
		} else {
			ni := "i" + strconv.Itoa(depth)
			c.wl("for ", ni, ":=0; ", ni, "<len(", c.fmtVnb(node, v, depth), "); ", ni, "++{")
			nv := "x" + strconv.Itoa(depth)
			if node.slct.ptr || c.isBuiltin(node.slct.typn) {
				c.wl(nv, " := ", c.fmtVd(node, v, depth), "[", ni, "]")
			} else {
				c.wl(nv, " := &", c.fmtVd(node, v, depth), "[", ni, "]")
			}
			_ = c.writeCountBytes(node.slct, nv, depth+1)
			c.wl("}")
		}
	case typeBasic:
		if node.typu == "string" {
			c.wl("c+=len(", c.fmtVnb(node, v, depth), ")")
		}
	}
	return nil
}
