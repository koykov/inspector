package inspector

import "strconv"

func (c *Compiler) writeNodeCopy(_ *node, recv, pname string) error {
	c.wl("var r ", pname)
	c.wl("switch x.(type) {")
	c.wl("case ", pname, ":")
	c.wl("r = x.(", pname, ")")
	c.wl("case *", pname, ":")
	c.wl("r = *x.(*", pname, ")")
	c.wl("case **", pname, ":")
	c.wl("r = **x.(**", pname, ")")
	c.wl("default:")
	c.wl("return nil, inspector.ErrUnsupportedType")
	c.wl("}")

	c.wl("bc:=", recv, ".countBytes(&r)")
	c.wl("var l ", pname)
	c.wl("err:=", recv, ".CopyTo(&r,&l,inspector.NewByteBuffer(bc))")
	c.wl("return &l, err")
	return nil
}

func (c *Compiler) writeNodeCopyTo(_ *node, recv, pname string) error {
	c.wl("var r ", pname)
	c.wl("switch src.(type) {")
	c.wl("case ", pname, ":")
	c.wl("r = src.(", pname, ")")
	c.wl("case *", pname, ":")
	c.wl("r = *src.(*", pname, ")")
	c.wl("case **", pname, ":")
	c.wl("r = **src.(**", pname, ")")
	c.wl("default:")
	c.wl("return inspector.ErrUnsupportedType")
	c.wl("}")

	c.wl("var l *", pname)
	c.wl("switch dst.(type) {")
	c.wl("case ", pname, ":")
	c.wl("return inspector.ErrMustPointerType")
	c.wl("case *", pname, ":")
	c.wl("l = dst.(*", pname, ")")
	c.wl("case **", pname, ":")
	c.wl("l = *dst.(**", pname, ")")
	c.wl("default:")
	c.wl("return inspector.ErrUnsupportedType")
	c.wl("}")

	c.wl("bb:=buf.AcquireBytes()")
	c.wl("var err error")
	c.wl("if bb,err=", recv, ".cpy(bb,l,&r);err!=nil{return err}")
	c.wl("buf.ReleaseBytes(bb)")
	c.wl("return nil")
	return nil
}

func (c *Compiler) writeCopy(node *node, l, r string, depth int) error {
	switch node.typ {
	case typeStruct:
		for _, ch := range node.chld {
			nl := l + "." + ch.name
			nr := r + "." + ch.name
			chPtr := ch.ptr && (ch.typ == typeStruct || ch.typ == typeMap || ch.typ == typeSlice)
			if chPtr {
				c.wl("if ", nr, "!=nil{")
				if ch.ptr && ch.typ == typeStruct {
					c.wl("if ", nl, "==nil{")
					c.wl(nl, "=&", c.fmtT(ch), "{}")
					c.wl("}")
				}
			}
			_ = c.writeCopy(ch, nl, nr, depth+1)
			if chPtr {
				c.wl("}")
			}
		}
	case typeMap:
		ln := "len(" + c.fmtVnb(node, r, depth) + ")"
		c.wl("if ", ln, ">0 {")
		c.wl("if ", l, "==nil{")
		lb1 := "buf" + strconv.Itoa(depth)
		c.wl(lb1, ":=make(", c.fmtT(node), ",", ln, ")")
		c.wl(l, "=", c.fmtP(node, lb1, depth))
		c.wl("}")
		rk := "rk" + strconv.Itoa(depth)
		rv := "rv" + strconv.Itoa(depth)
		c.wl("for ", rk, ",", rv, ":=range ", c.fmtVnb(node, r, depth), "{")
		c.wl("_,_=", rk, ",", rv)
		lk := "lk" + strconv.Itoa(depth)
		c.wl("var ", lk, " ", c.fmtT(node.mapk))
		_ = c.writeCopy(node.mapk, lk, rk, depth+1)
		lv := "lv" + strconv.Itoa(depth)
		c.wl("var ", lv, " ", c.fmtT(node.mapv))
		_ = c.writeCopy(node.mapv, lv, rv, depth+1)
		pfx := ""
		if node.mapv.ptr && node.mapv.typ != typeBasic {
			pfx = "&"
		}
		c.wl(c.fmtVd(node, l, depth), "[", lk, "]=", pfx, lv)
		c.wl("}")
		c.wl("}")
	case typeSlice:
		if node.typn == "[]byte" {
			c.wl("buf,", c.fmtVnb(node, l, depth), "=inspector.Bufferize(buf,", c.fmtVnb(node, r, depth), ")")
		} else {
			c.wl("if len(", c.fmtVnb(node, r, depth), ")>0{")
			lb := "buf" + strconv.Itoa(depth)
			c.wl(lb, ":=", c.fmtVd(node, l, depth))
			c.wl("if ", lb, "==nil {")
			c.wl(lb, "=make(", c.fmtT(node), ",0,len(", c.fmtVnb(node, r, depth), "))")
			c.wl("}")

			ni := "i" + strconv.Itoa(depth)
			c.wl("for ", ni, ":=0; ", ni, "<len(", c.fmtVnb(node, r, depth), "); ", ni, "++{")
			nv := "x" + strconv.Itoa(depth)
			nb := "b" + strconv.Itoa(depth)
			c.wl("var ", nb, " ", c.fmtT(node.slct))
			if node.slct.ptr || c.isBuiltin(node.slct.typn) {
				c.wl(nv, " := ", c.fmtVd(node, r, depth), "[", ni, "]")
			} else {
				c.wl(nv, " := &", c.fmtVd(node, r, depth), "[", ni, "]")
			}
			_ = c.writeCopy(node.slct, nb, nv, depth+1)
			pfx := ""
			if node.slct.ptr && !c.isBuiltin(node.slct.typn) {
				pfx = "&"
			}
			c.wl(lb, "=append(", lb, ",", pfx, nb, ")")
			c.wl("}")
			c.wl(l, "=", c.fmtP(node, lb, depth))
			c.wl("}")
		}
	case typeBasic:
		if node.typu == "string" {
			sname := "c" + strconv.Itoa(c.cntrCpy)
			c.wl("var ", sname, " string")
			c.cntrCpy++
			c.wl("buf,", sname, "=inspector.BufferizeString(buf,string(", c.fmtVnb(node, r, depth), "))")
			pname := c.fmtPtpfx(node.typn)
			c.wl(c.fmtVnb(node, l, depth), "=", pname, node.typn, "(", sname, ")")
		} else {
			c.wl(l, "=", r)
		}
	}
	return nil
}
