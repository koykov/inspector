package inspector

func (c *Compiler) getFuncHeader(pname string) string {
	return `if len(path) == 0 { return }
if src == nil { return }
var x *` + pname + `
_ = x
if p, ok := src.(**` + pname + `); ok { x = *p } else if p, ok := src.(*` + pname + `); ok { x = p } else if v, ok := src.(` + pname + `); ok { x = &v } else { return }`
}

func (c *Compiler) getFuncHeaderGetTo(pname string) string {
	return `if src == nil { return }
var x *` + pname + `
_ = x
if p, ok := src.(**` + pname + `); ok { x = *p } else if p, ok := src.(*` + pname + `); ok { x = p } else if v, ok := src.(` + pname + `); ok { x = &v } else { return }
if len(path) == 0 { *buf = &(*x)
return}`
}

func (c *Compiler) getFuncHeaderSet(pname string) string {
	return `if len(path) == 0 { return nil }
if dst == nil { return nil }
var x *` + pname + `
_ = x
if p, ok := dst.(**` + pname + `); ok { x = *p } else if p, ok := dst.(*` + pname + `); ok { x = p } else if v, ok := dst.(` + pname + `); ok { x = &v } else { return nil }`
}

func (c *Compiler) getFuncHeaderLoop(pname string, nodetyp typ) string {
	var r string
	if nodetyp != typeSlice {
		r += "if len(path) == 0 { return }\n"
	}
	return r + `if src == nil { return }
var x *` + pname + `
_ = x
if p, ok := src.(**` + pname + `); ok { x = *p } else if p, ok := src.(*` + pname + `); ok { x = p } else if v, ok := src.(` + pname + `); ok { x = &v } else { return }`
}

func (c *Compiler) getFuncHeaderEqual(pname string) string {
	return `var (
lx, rx *` + pname + `
leq, req bool
)
_, _, _, _ = lx, rx, leq, req
if lp, ok := l.(**` + pname + `); ok { lx, leq = *lp, true } else if lp, ok := l.(*` + pname + `); ok { lx, leq = lp, true } else if lp, ok := l.(` + pname + `); ok { lx, leq = &lp, true }
if rp, ok := r.(**` + pname + `); ok { rx, req = *rp, true } else if rp, ok := r.(*` + pname + `); ok { rx, req = rp, true } else if rp, ok := r.(` + pname + `); ok { rx, req = &rp, true }
if !leq || !req { return false }
if lx == nil && rx == nil { return true }
if (lx == nil && rx != nil) || (lx != nil && rx == nil) { return false }`
}

func (c *Compiler) getFuncHeaderLC(pname string) string {
	return `if src == nil { return nil }
var x *` + pname + `
_ = x
if p, ok := src.(**` + pname + `); ok { x = *p } else if p, ok := src.(*` + pname + `); ok { x = p } else if v, ok := src.(` + pname + `); ok { x = &v } else { return inspector.ErrUnsupportedType }`
}

func (c *Compiler) getFuncHeaderAppend(pname string) string {
	return `if src == nil { return src, nil }
var x *` + pname + `
_ = x
if p, ok := src.(**` + pname + `); ok { x = *p } else if p, ok := src.(*` + pname + `); ok { x = p } else if v, ok := src.(` + pname + `); ok { x = &v } else { return src, nil }`
}
