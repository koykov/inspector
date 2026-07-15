package inspector

import (
	"strconv"
	"strings"
)

// Write body of the methods according given mode.
func (c *Compiler) writeNode(node, parent *node, recv, v, vsrc string, depth int, mode mode) error {
	depths := strconv.Itoa(depth)

	// Flag to check length of the path array.
	requireLenCheck := node.typ != typeBasic && !(mode == modeLoop && (node.typ == typeMap || (node.typ == typeSlice && node.typn != "[]byte")))

	if requireLenCheck {
		c.wl("if len(path) > ", depths, " {")
	}
	if node.ptr {
		// Value may be nil on pointer types.
		c.wl("if ", v, " == nil { ", c.fmtR(mode, "nil"), " }")
	}

	switch node.typ {
	case typeStruct:
		// Walk over fields.
		for _, ch := range node.chld {
			isBasic := ch.typ == typeBasic || (ch.typ == typeSlice && ch.typn == "[]byte")
			if isBasic && mode == modeLoop {
				continue
			}
			c.wl("if path[", depths, "] == ", `"`, ch.name, `" {`)
			if isBasic {
				switch mode {
				case modeGet:
					c.wl("*buf = &", v, ".", ch.name)
					c.wl("return")
				case modeCmp:
					c.writeCmp(ch, v+"."+ch.name)
					c.wl("return")
				case modeSet:
					pfx := ""
					if !c.isBuiltin(ch.typn) && !c.isBuiltin(ch.typu) {
						pfx = ch.pkg + "."
					}
					if !ch.ptr {
						pfx = "&" + pfx
					}
					c.wl("inspector.AssignBuf(", pfx, v, ".", ch.name, ", value, buf)")
					c.wl("return nil")
				default:
					// noop
				}
			} else {
				nv := "x" + strconv.Itoa(depth)
				vsrc := v + "." + ch.name
				pfx := ""
				nvPtr := false
				if !ch.ptr && ch.typ != typeMap && ch.typ != typeSlice {
					pfx = "&"
					nvPtr = true
				}
				c.wl(nv, " := ", pfx, vsrc)
				if mode == modeSet && (ch.typ == typeStruct || ch.typ == typeMap || ch.typ == typeSlice) {
					nilChk := ch.ptr || ch.typ == typeMap || ch.typ == typeSlice
					typ := c.fmtT(ch)
					pfx := "*"
					if ch.ptr || nvPtr {
						pfx = ""
					}
					c.ensureImports(typ)
					c.wl("if uvalue, ok := value.(*", typ, "); ok {")
					c.wl(nv, " = ", pfx, "uvalue")
					c.wl("}")

					pfx = ""
					if ch.ptr {
						pfx = "&"
					}
					if nilChk {
						c.wl("if ", nv, " == nil {")
						switch ch.typ {
						case typeStruct:
							c.wl(nv, " = ", pfx, typ, "{}")
						case typeMap:
							c.wl("z := make(", typ, ")")
							c.wl(nv, " = ", pfx, "z")
						case typeSlice:
							c.wl("z := make(", typ, ", 0)")
							c.wl(nv, " = ", pfx, "z")
						default:
							// noop
						}
						c.wl(v+"."+ch.name, " = ", nv)
						c.wl("}")
					}
				}
				c.wl("_ = ", nv)
				if mode == modeCmp && ch.ptr {
					c.writeCmp(ch, nv)
				}
				err := c.writeNode(ch, node, recv, nv, vsrc, depth+1, mode)
				if err != nil {
					return err
				}
				if mode == modeGet {
					c.wl("return")
				}
				if mode == modeSet {
					pfx := ""
					if nvPtr && !ch.ptr {
						pfx = "*"
					}
					c.wl(vsrc, " = ", pfx, nv)
				}
			}
			c.wl("}")
		}
	case typeMap:
		// todo find a way to get element from the map without allocation (see runtime.mapaccess1 and runtime.mapaccess2).
		origPtr := node.ptr
		if depth == 0 {
			node.ptr = true
		}
		switch mode {
		case modeLoop:
			// Loop magic.
			c.wl("for k := range ", c.fmtV(node, v), " {")
			c.wl("if l.RequireKey() {")
			switch node.mapk.typn {
			case "string", "[]byte":
				c.wl("*buf = append((*buf)[:0], ", c.fmtVnb(node.mapk, "k", depth+1), "...)")
			case "bool":
				c.wl(`if k { *buf = append((*buf)[:0], "true"...) } else { *buf = append(*buf[:0], "false"...) }`)
			case "int", "int8", "int16", "int32", "int64":
				c.wl("*buf = strconv.AppendInt((*buf)[:0], int64(", c.fmtVnb(node.mapk, "k", depth+1), "), 10)")
			case "uint", "uint8", "uint16", "uint32", "uint64":
				c.wl("*buf = strconv.AppendUint((*buf)[:0], uint64(", c.fmtVnb(node.mapk, "k", depth+1), "), 10)")
			case "float32", "float64":
				c.wl("*buf = strconv.AppendFloat((*buf)[:0], float64(", c.fmtVnb(node.mapk, "k", depth+1), "), 'f', -1, 64)")
			default:
				c.regImport([]string{`"github.com/koykov/x2bytes"`})
				c.wl("*buf, err = x2bytes.ToBytes((*buf)[:0], k)")
				c.wl("if err != nil { return }")
			}
			c.wl("l.SetKey(buf, &inspector.StaticInspector{})")
			c.wl("}")
			insName := "inspector.StaticInspector"
			if node.mapv.typ == typeStruct {
				insName = ""
				if node.mapv.pkg != c.pkgName {
					insName = node.mapv.pkg + "_ins."
					c.regImport([]string{`"` + node.mapv.pkgi + `_ins"`})
				}
				insName += node.mapv.typn + "Inspector"
			}
			c.wl("l.SetVal(", c.fmtV(node, v), "[k], &", insName, "{})")
			c.wl("ctl := l.Iterate()")
			c.wl("if ctl == inspector.LoopCtlBrk { break }")
			c.wl("if ctl == inspector.LoopCtlCnt { continue }")
			c.wl("}")
			c.wl("return")
		default:
			nv := "x" + strconv.Itoa(depth)
			if node.mapk.typn == "string" {
				// Key is string, simple case.
				key := c.fmtP(node.mapk, "path["+depths+"]", depth+1)
				if mode == modeSet {
					c.wl(nv, " := ", c.fmtV(node, v), "[", key, "]")
				} else {
					c.wl("if ", nv, ", ok := ", c.fmtV(node, v), "[", key, "]; ok {")
				}
				c.wl("_ = ", nv)
				err := c.writeNode(node.mapv, node, recv, nv, "", depth+1, mode)
				if err != nil {
					return err
				}
				if mode == modeSet {
					c.wl(c.fmtV(node, v), "[", key, "] = ", nv)
					c.wl("return nil")
				}
				if mode != modeSet {
					c.wl("}")
				}
			} else {
				// Convert path value to the key type and try to find it in the map.
				c.wl("var k ", c.pkgName, ".", node.mapk.typn)
				pname := ""
				if node.mapk.typn != "string" {
					pname = c.pkgName + "."
				}
				c.wl("k=", pname, node.mapk.typn, "(path[", depths, "])")
				c.wl(nv, " := ", c.fmtV(node, v), "[", c.fmtP(node.mapk, "k", depth+1), "]")
				c.wl("_ = ", nv)
				err := c.writeNode(node.mapv, node, recv, nv, "", depth+1, mode)
				if mode == modeSet {
					c.wl(c.fmtV(node, v), "[", c.fmtP(node.mapk, "k", depth+1), "] = ", nv)
					c.wl("return nil")
				}
				if err != nil {
					return err
				}
			}
		}
		node.ptr = origPtr
	case typeSlice:
		// Quick check of []byte type.
		if node.typn == "[]byte" {
			switch mode {
			case modeGet:
				c.wl("*buf = &", v)
				c.wl("return")
			case modeCmp:
				c.writeCmp(node, v)
				c.wl("return")
			case modeSet:
				pfx := ""
				if !node.ptr {
					pfx = "&"
				}
				c.wl("inspector.AssignBuf(", pfx, v, ", value, buf)")
				c.wl("return nil")
			default:
				// noop
			}
		}
		switch mode {
		case modeLoop:
			// Loop magic.
			c.wl("for k := range ", c.fmtVd(node, v, depth), " {")
			c.wl("if l.RequireKey() {")
			c.wl("*buf = strconv.AppendInt((*buf)[:0], int64(k), 10)")
			c.wl("l.SetKey(buf, &inspector.StaticInspector{})")
			c.wl("}")
			insName := "inspector.StaticInspector"
			if node.slct.typ == typeStruct {
				insName = ""
				if node.slct.pkg != c.pkgName {
					insName = node.slct.pkg + "_ins."
					c.regImport([]string{`"` + node.slct.pkgi + `_ins"`})
				}
				insName += node.slct.typn + "Inspector"
			}
			c.wl("l.SetVal(&", c.fmtVd(node, v, depth), "[k], &", insName, "{})")
			c.wl("ctl := l.Iterate()")
			c.wl("if ctl == inspector.LoopCtlBrk { break }")
			c.wl("if ctl == inspector.LoopCtlCnt { continue }")
			c.wl("}")
			c.wl("return")
		default:
			// Convert path value to the int index and try to find value in the slice using it.
			nv := "x" + strconv.Itoa(depth)
			c.wl("var i int")
			snippet, imports, err := StrConvSnippet("path["+depths+"]", "int", "", "i")
			c.regImport(imports)
			if err != nil {
				return err
			}
			c.wl(snippet)
			c.wl("if len(", c.fmtVnb(node, v, depth), ") > i {")
			if node.slct.ptr || c.isBuiltin(node.slct.typn) {
				c.wl(nv, " := ", c.fmtVd(node, v, depth), "[i]")
			} else {
				c.wl(nv, " := &", c.fmtVd(node, v, depth), "[i]")
			}
			c.wl("_ = ", nv)
			err = c.writeNode(node.slct, node, recv, nv, "", depth+1, mode)
			if err != nil {
				return err
			}
			if mode == modeSet {
				pfx := ""
				if !node.slct.ptr && !c.isBuiltin(node.slct.typn) {
					pfx = "*"
				}
				c.wl(c.fmtVd(node, v, depth), "[i] = ", pfx, nv)
				c.wl("return nil")
			}
			c.wl("}")
		}
	case typeBasic:
		switch mode {
		case modeGet:
			c.wl("*buf = &", v)
			c.wl("return")
		case modeCmp:
			c.writeCmp(node, v)
			c.wl("return")
		case modeSet:
			pfx := ""
			if !c.isBuiltin(node.typn) {
				pfx = node.pkg + "."
			}
			if !node.ptr {
				pfx = "&" + pfx
			}
			c.wl("inspector.AssignBuf(", pfx, v, ", value, buf)")
			if parent != nil && parent.typ != typeMap {
				c.wl("return nil")
			}
		default:
			// noop
		}
	}
	if requireLenCheck {
		c.wl("}")
	}
	// Special case to take value by pointer to avoid allocation.
	if (node.typ == typeStruct || node.typ == typeMap || node.typ == typeSlice) && mode == modeGet && v != "x" {
		pfx := ""
		if parent != nil && parent.typ != typeSlice {
			pfx = "&"
		}
		if len(vsrc) > 0 {
			c.wl("*buf = ", pfx, vsrc)
		} else {
			c.wl("*buf = ", pfx, v)
		}
	}

	return c.err
}

// Write comparison code of the node.
func (c *Compiler) writeCmp(left *node, leftVar string) {
	// Node is pointer, check nil case.
	if left.ptr {
		c.wl("if right==inspector.Nil {")
		c.wl("if cond == inspector.OpEq {")
		c.wl("*result = ", leftVar, " == nil")
		c.wl("} else {")
		c.wl("*result = ", leftVar, " != nil")
		c.wl("}")
		c.wl("return\n}")
		return
	}

	// Get type name as a string.
	bi := c.isBuiltin(left.typn)
	pname := left.typn
	if !bi {
		pname = c.pkgName + "." + pname
	}
	// Get the exact value of the right variable.
	c.wl("var rightExact ", pname)
	snippet, imports, err := StrConvSnippet("right", left.typn, left.typu, "rightExact")
	c.regImport(imports)
	if err != nil {
		c.err = err
		return
	}
	if !bi {
		snippet = strings.Replace(snippet, left.typu, pname, -1)
	}
	c.wl(snippet)

	switch left.typn {
	// []byte and bool types allows only equal and non-equal comparison operations.
	case "[]byte":
		c.wl("if cond == inspector.OpEq {")
		c.wl("*result = bytes.Equal(", leftVar, ", rightExact)")
		c.wl("} else {")
		c.wl("*result = !bytes.Equal(", leftVar, ", rightExact)")
		c.wl("}")
	case "bool":
		c.wl("if cond == inspector.OpEq {")
		c.wl("*result = ", leftVar, " == rightExact")
		c.wl("} else {")
		c.wl("*result = ", leftVar, " != rightExact")
		c.wl("}")
	default:
		// All other cases allows all type sof the comparison.
		c.wl("switch cond {")
		c.wl("case inspector.OpEq:")
		c.wl("*result = ", leftVar, " == rightExact")
		c.wl("case inspector.OpNq:")
		c.wl("*result = ", leftVar, " != rightExact")
		c.wl("case inspector.OpGt:")
		c.wl("*result = ", leftVar, " > rightExact")
		c.wl("case inspector.OpGtq:")
		c.wl("*result = ", leftVar, " >= rightExact")
		c.wl("case inspector.OpLt:")
		c.wl("*result = ", leftVar, " < rightExact")
		c.wl("case inspector.OpLtq:")
		c.wl("*result = ", leftVar, " <= rightExact")
		c.wl("}")
	}
}

func (c *Compiler) ensureImports(typ string) {
	switch typ {
	case "time.Time":
		c.regImport([]string{`"time"`})
	}
}
