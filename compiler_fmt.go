package inspector

import "strings"

// Bounds variable according node options.
func (c *Compiler) fmtV(node *node, v string) string {
	if node.ptr {
		return "(*" + v + ")"
	}
	return "(" + v + ")"
}

// Bounds variable according node options and depth.
func (c *Compiler) fmtVd(node *node, v string, depth int) string {
	if node.ptr || depth == 0 {
		return "(*" + v + ")"
	}
	return "(" + v + ")"
}

// No brackets version of fmtV().
func (c *Compiler) fmtVnb(node *node, v string, depth int) string {
	if node.ptr || depth == 0 {
		return "*" + v
	}
	return v
}

// Take pointer of v.
func (c *Compiler) fmtP(node *node, v string, depth int) string {
	if node.ptr || depth == 0 {
		return "&" + v
	}
	return v
}

// Format type to use in new()/make() functions.
func (c *Compiler) fmtT(node_ *node) string {
	return c.fmtTfs(node_, false)
}

func (c *Compiler) fmtTfs(node_ *node, forceStructPtr bool) string {
	pfx := func(node_ *node) string {
		if c.inp || len(node_.pkg) == 0 {
			return ""
		}
		return node_.pkg + "."
	}
	switch node_.typ {
	case typeStruct:
		var pfx_ string
		if node_.ptr {
			pfx_ = "*"
		}
		if forceStructPtr {
			return pfx_ + pfx(node_) + node_.typn
		}
		return pfx(node_) + strings.Trim(node_.typn, "*")
	case typeMap:
		if strings.Contains(node_.typn, "map[") {
			s := "map["
			if node_.mapk.ptr {
				s += "*"
			}
			if len(node_.mapk.pkg) > 0 {
				s += node_.mapk.pkg + "."
			}
			s += node_.mapk.typn
			s += "]"
			if node_.mapv.ptr {
				s += "*"
			}
			if len(node_.mapv.pkg) > 0 && !c.inp {
				s += node_.mapv.pkg + "."
			}
			s += node_.mapv.typn
			return s
		} else {
			return pfx(node_) + strings.Trim(node_.typn, "*")
		}
	case typeSlice:
		s := node_.slct.typn
		if !c.isBuiltin(s) {
			s = pfx(node_.slct) + strings.Trim(node_.slct.typn, "*")
			if node_.slct.ptr {
				s = "*" + s
			}
		}
		if strings.Index(node_.typn, "[]") != -1 {
			pfx_ := ""
			if node_.slct.ptr {
				pfx_ = "*"
			}
			s = "[]" + pfx_ + strings.TrimLeft(s, "*")
		} else {
			s = pfx(node_) + strings.Trim(node_.typn, "*")
		}
		return s
	case typeBasic:
		pfx_ := ""
		if node_.ptr {
			pfx_ = "*"
		}
		return pfx_ + pfx(node_) + node_.typn
	}
	return ""
}

func (c *Compiler) fmtR(mode mode, err string) string {
	if mode == modeSet {
		return "return " + err
	} else {
		return "return"
	}
}

// Parent type prefix format.
func (c *Compiler) fmtPtpfx(typ string) string {
	if !c.isBuiltin(typ) {
		return c.pkgName + "."
	}
	return ""
}

// Shorthand write method.
func (c *Compiler) w(s ...string) {
	_, err := c.wr.WriteString(strings.Join(s, ""))
	if err != nil && c.err == nil {
		c.err = err
	}
}

// Shorthand write new line helper.
func (c *Compiler) wl(s ...string) {
	s = append(s, "\n")
	c.w(s...)
}

// Shorthand write new double lines helper.
func (c *Compiler) wdl(s ...string) {
	s = append(s, "\n\n")
	c.w(s...)
}

// Shorthand replace helper.
func (c *Compiler) r(old, new string) {
	s := c.wr.String()
	s = strings.Replace(s, old, new, -1)
	c.wr.Reset()
	_, _ = c.wr.WriteString(s)
}
