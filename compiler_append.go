package inspector

import (
	"strconv"
	"strings"
)

func (c *Compiler) writeNodeAppend(node_ *node, v string, depth int) error {
	if !node_.hasa {
		return nil
	}
	depths := strconv.Itoa(depth)

	if node_.ptr && node_.typ != typeSlice {
		// Value may be nil on pointer types.
		c.wl("if ", v, " == nil { ", v, "=new(", c.fmtT(node_), ") }")
	}

	if node_.typ == typeSlice {
		t := c.fmtTfs(node_.slct, true)
		var pfx string
		if !node_.slct.ptr {
			pfx = "*"
		}
		c.wl("if len(path)==", depths, "{")
		c.wl("var raw ", pfx, t)
		c.wl("var ok bool")
		c.wl("switch y:=value.(type){")
		if node_.slct.ptr {
			c.wl("case ", t, ": raw=y; ok=true")
			c.wl("case *", t, ": raw=*y; ok=true")
		} else {
			c.wl("case ", t, ": raw=&y; ok=true")
			c.wl("case *", t, ": raw=y; ok=true")
		}
		c.wl("}")
		c.wl("if ok {")
		c.wl(c.fmtVnb(node_, v, depth), "=append(", c.fmtVnb(node_, v, depth), ",", pfx, "raw)")
		pfx = ""
		if !node_.ptr {
			pfx = "&"
		}
		c.wl("return ", pfx, v, ", nil")
		c.wl("}}")
	}

	switch node_.typ {
	case typeStruct:
		for _, ch := range node_.chld {
			if !ch.hasa {
				continue
			}
			c.wl("if path[", depths, "] == ", `"`, ch.name, `" {`)
			nv := v + "." + ch.name
			_ = c.writeNodeAppend(ch, nv, depth+1)
			c.wl("}")
		}
	case typeMap:
		origPtr := node_.ptr
		if depth == 0 {
			node_.ptr = true
		}
		if !node_.mapv.hasa {
			return nil
		}
		nv := "x" + strconv.Itoa(depth)
		c.wl("if len(path) < ", strconv.Itoa(depth+1), " { return src, nil }")
		if node_.mapk.typn == "string" {
			// Key is string, simple case.
			key := c.fmtP(node_.mapk, "path["+depths+"]", depth+1)
			c.wl("if ", nv, ", ok := ", c.fmtV(node_, v), "[", key, "]; ok {")
			c.wl("_ = ", nv)
			err := c.writeNodeAppend(node_.mapv, nv, depth+1)
			if err != nil {
				return err
			}
			c.wl(c.fmtV(node_, v), "[", key, "]=", nv)
			c.wl("}")
		} else {
			// Convert path value to the key type and try to find it in the map.
			c.wl("var k ", node_.mapk.typn)
			snippet, imports, err := StrConvSnippet("path["+depths+"]", node_.mapk.typn, node_.mapk.typu, "k")
			snippet = strings.ReplaceAll(snippet, "return err", "return src, err") // dirty way(((
			c.regImport(imports)
			if err != nil {
				return err
			}
			c.wl(snippet)
			c.wl(nv, " := ", c.fmtV(node_, v), "[", c.fmtP(node_.mapk, "k", depth+1), "]")
			c.wl("_ = ", nv)
			err = c.writeNodeAppend(node_.mapv, nv, depth+1)
			if err != nil {
				return err
			}
		}
		node_.ptr = origPtr
	case typeSlice:
		if !node_.slct.hasa {
			return nil
		}
		c.wl("if len(path) < ", strconv.Itoa(depth+1), " { return src, nil }")

		nv := "x" + strconv.Itoa(depth)
		c.wl("var i int")
		snippet, imports, err := StrConvSnippet("path["+depths+"]", "int", "", "i")
		snippet = strings.ReplaceAll(snippet, "return err", "return src, err") // dirty way(((
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
		err = c.writeNodeAppend(node_.slct, nv, depth+1)
		if err != nil {
			return err
		}
		c.wl("}")
	case typeBasic:
		// do nothing
	}

	return nil
}
