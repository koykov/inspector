package inspector

import (
	"strconv"
	"strings"
)

func (c *Compiler) writeNodeDEQ(node, parent *node, recv, path, lv, rv string, depth int) error {
	_ = parent
	path = strings.Trim(path, ".")
	if len(path) > 0 && len(node.name) > 0 {
		path += "."
	}
	if depth > 0 {
		path += node.name
	}

	var deqMustSkipByType, deqMustSkipByTypeAndPath bool
	if parent != nil {
		deqMustSkipByType = parent.typ != typeMap && parent.typ != typeSlice
		deqMustSkipByTypeAndPath = len(path) > 0 && parent.typ != typeMap && parent.typ != typeSlice
	}

	if node.ptr {
		c.wl("if (", lv, "==nil && ", rv, "!=nil) || (", lv, "!=nil && ", rv, "==nil) {return false}")
		c.wl("if ", lv, "!=nil && ", rv, "!=nil {")
	}

	switch node.typ {
	case typeStruct:
		if deqMustSkipByTypeAndPath {
			c.wl("if inspector.DEQMustCheck(\"", path, "\",opts){")
		}
		for _, ch := range node.chld {
			nlv, nrv := lv, rv
			isBasic := ch.typ == typeBasic || (ch.typ == typeSlice && ch.typn == "[]byte")
			if !isBasic {
				c.cntrDEQ++
				nlv = "lx" + strconv.Itoa(c.cntrDEQ)
				nrv = "rx" + strconv.Itoa(c.cntrDEQ)
				c.wl(nlv, ":=", lv, ".", ch.name)
				c.wl(nrv, ":=", rv, ".", ch.name)
				c.wl("_,_=", nlv, ",", nrv)
			}
			if err := c.writeNodeDEQ(ch, node, recv, path, nlv, nrv, depth+1); err != nil {
				return err
			}
		}
		if deqMustSkipByTypeAndPath {
			c.wl("}")
		}
	case typeMap:
		nlv := c.fmtVnb(node, lv, depth)
		nrv := c.fmtVnb(node, rv, depth)
		if deqMustSkipByTypeAndPath {
			c.wl("if inspector.DEQMustCheck(\"", path, "\",opts){")
		}
		c.wl("if len(", nlv, ")!=len(", nrv, "){return false}")
		c.wl("for k:=range ", nlv, "{")
		c.cntrDEQ++
		nlv1 := "lx" + strconv.Itoa(c.cntrDEQ)
		nrv1 := "rx" + strconv.Itoa(c.cntrDEQ)
		ok1 := "ok" + strconv.Itoa(c.cntrDEQ)
		c.wl(nlv1, ":=(", nlv, ")[k]")
		c.wl(nrv1, ",", ok1, ":=(", nrv, ")[k]")
		c.wl("_,_,_=", nlv1, ",", nrv1, ",", ok1)
		c.wl("if !", ok1, "{return false}")
		if err := c.writeNodeDEQ(node.mapv, node, recv, path, nlv1, nrv1, depth+1); err != nil {
			return err
		}
		c.wl("}")
		if deqMustSkipByTypeAndPath {
			c.wl("}")
		}
	case typeSlice:
		if node.typn == "[]byte" {
			nlv, nrv := lv+"."+node.name, rv+"."+node.name
			c.wl("if !bytes.Equal(", c.fmtVnb(node, nlv, depth), ",", c.fmtVnb(node, nrv, depth), ") && inspector.DEQMustCheck(\"", path, "\",opts){return false}")
		} else {
			nlv := c.fmtVnb(node, lv, depth)
			nrv := c.fmtVnb(node, rv, depth)
			if deqMustSkipByTypeAndPath {
				c.wl("if inspector.DEQMustCheck(\"", path, "\",opts){")
			}
			c.wl("if len(", nlv, ")!=len(", nrv, "){return false}")
			c.wl("for i:=0;i<len(", nlv, ");i++{")
			c.cntrDEQ++
			nlv1 := "lx" + strconv.Itoa(c.cntrDEQ)
			nrv1 := "rx" + strconv.Itoa(c.cntrDEQ)
			c.wl(nlv1, ":=(", nlv, ")[i]")
			c.wl(nrv1, ":=(", nrv, ")[i]")
			c.wl("_,_=", nlv1, ",", nrv1)
			if err := c.writeNodeDEQ(node.slct, node, recv, path, nlv1, nrv1, depth+1); err != nil {
				return err
			}
			c.wl("}")
			if deqMustSkipByTypeAndPath {
				c.wl("}")
			}
		}
	case typeBasic:
		plv, prv := lv, rv
		if len(node.name) > 0 {
			plv, prv = lv+"."+node.name, rv+"."+node.name
		}
		if node.ptr {
			plv = "*" + plv
			prv = "*" + prv
		}
		if deqMustSkipByType {
			if node.typu == "float32" {
				c.wl("if !inspector.EqualFloat32(", plv, ",", prv, ", opts) && inspector.DEQMustCheck(\"", path, "\",opts){return false}")
			} else if node.typu == "float64" {
				c.wl("if !inspector.EqualFloat64(", plv, ",", prv, ", opts) && inspector.DEQMustCheck(\"", path, "\",opts){return false}")
			} else {
				c.wl("if ", plv, "!=", prv, " && inspector.DEQMustCheck(\"", path, "\",opts){return false}")
			}
		} else {
			c.wl("if ", plv, "!=", prv, "{return false}")
		}
	}

	if node.ptr {
		c.wl("}")
	}

	return nil
}
