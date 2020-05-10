package inspector

import (
	"go/format"
	"go/types"
	"strconv"
	"strings"

	"golang.org/x/tools/go/loader"
)

type typ int
type mode int

const (
	typeStruct typ = iota
	typeMap
	typeSlice
	typeBasic

	modeGet mode = iota
	modeCmp
	modeLoop
)

type ByteStringWriter interface {
	Write(p []byte) (n int, err error)
	WriteByte(c byte) error
	WriteString(s string) (n int, err error)
	Bytes() []byte
	String() string
	Reset()
}

type Compiler struct {
	pkg     string
	pkgDot  string
	pkgName string
	outDir  string
	uniq    map[string]bool
	nodes   []*node
	wr      ByteStringWriter
	imp     []string
	err     error
}

type node struct {
	typ  typ
	typn string
	name string
	pkg  string
	pkgi string
	ptr  bool
	chld []*node
	mapk *node
	mapv *node
	slct *node
}

func NewCompiler(pkg, outDir string, w ByteStringWriter) *Compiler {
	c := Compiler{
		pkg:    pkg,
		pkgDot: pkg + ".",
		outDir: outDir,
		uniq:   make(map[string]bool),
		wr:     w,
		imp:    make([]string, 0),
	}
	return &c
}

func (c *Compiler) String() string {
	return ""
}

func (c *Compiler) Compile() error {
	var conf loader.Config
	conf.Import(c.pkg)
	prog, err := conf.Load()
	if err != nil {
		return err
	}

	pkg := prog.Package(c.pkg)
	c.pkgName = pkg.Pkg.Name()
	err = c.parsePkg(pkg)
	if err != nil {
		return err
	}

	err = c.write()
	if err != nil {
		return err
	}

	source := c.wr.Bytes()
	fmtSource, err := format.Source(source)
	if err != nil {
		return err
	}
	c.wr.Reset()
	_, _ = c.wr.Write(fmtSource)

	return nil
}

func (c *Compiler) parsePkg(pkg *loader.PackageInfo) error {
	for _, scope := range pkg.Info.Scopes {
		if parent := scope.Parent(); parent != nil {
			for _, name := range parent.Names() {
				o := parent.Lookup(name)
				t := o.Type()
				node, err := c.parseType(t)
				if err != nil {
					return err
				}
				if node.typ == typeStruct {
					node.typn = o.Name()
				}
				node.pkg = o.Pkg().Name()
				node.name = o.Name()
				c.nodes = append(c.nodes, node)
			}
		}
	}

	return nil
}

func (c *Compiler) parseType(t types.Type) (*node, error) {
	node := &node{
		typ:  typeBasic,
		typn: strings.Replace(t.String(), c.pkgDot, "", 1),
		ptr:  false,
	}
	if n, ok := t.(*types.Named); ok {
		node.typn = n.Obj().Name()
		node.pkg = n.Obj().Pkg().Name()
		node.pkgi = n.Obj().Pkg().Path()
	}

	u := t.Underlying()
	if p, ok := u.(*types.Pointer); ok {
		e := p.Elem()
		u = e.Underlying()
		unode, err := c.parseType(u)
		if err != nil {
			return node, err
		}
		// unode.typn = strings.TrimLeft(node.typn, "*")
		unode.ptr = true
		if n, ok := e.(*types.Named); ok {
			unode.typn = n.Obj().Name()
			unode.pkg = n.Obj().Pkg().Name()
			unode.pkgi = n.Obj().Pkg().Path()
		}
		return unode, err
	}

	if s, ok := u.(*types.Struct); ok {
		node.typ = typeStruct
		for i := 0; i < s.NumFields(); i++ {
			f := s.Field(i)
			ch, err := c.parseType(f.Type())
			if err != nil {
				return node, err
			}
			ch.name = f.Name()
			if ch.ptr {
				ch.typn = strings.Replace(f.Type().String(), c.pkgDot, "", 1)
			}
			node.chld = append(node.chld, ch)
		}
		return node, nil
	}

	if m, ok := u.(*types.Map); ok {
		var err error
		node.typ = typeMap
		node.mapk, err = c.parseType(m.Key())
		node.mapv, err = c.parseType(m.Elem())
		return node, err
	}

	if s, ok := u.(*types.Slice); ok {
		var err error
		node.typ = typeSlice
		node.slct, err = c.parseType(s.Elem())
		return node, err
	}

	return node, nil
}

func (c *Compiler) write() error {
	c.wl("// Code generated by inspc. DO NOT EDIT.")
	c.wdl("// source: ", c.pkg)

	c.wdl("package ", c.pkgName, "_ins")

	c.imp = append(c.imp, `"github.com/koykov/inspector"`, `"`+c.pkg+`"`)
	c.wdl("import (\n!{import}\n)")

	c.wl("func init() {")
	for _, node := range c.nodes {
		c.wl(`inspector.RegisterInspector("`, node.name, `", &`, node.name, `Inspector{})`)
	}
	c.wdl("}")

	for i, node := range c.nodes {
		err := c.writeRootNode(node, i)
		if err != nil {
			return err
		}
	}
	c.r("!{import}", strings.Join(c.imp, "\n"))

	return nil
}

func (c *Compiler) regImport(imports []string) {
	for _, i := range imports {
		found := false
		for _, j := range c.imp {
			if i == j {
				found = true
				break
			}
		}
		if !found {
			c.imp = append(c.imp, i)
		}
	}
}

func (c *Compiler) writeRootNode(node *node, idx int) (err error) {
	inst := node.name + "Inspector"
	recv := "i" + strconv.Itoa(idx)
	pname := c.pkgName + "." + node.typn

	c.wdl("type ", inst, " struct {\ninspector.BaseInspector\n}")

	// Common checks and type cast.
	funcHeader := `if len(path) == 0 { return }
if src == nil { return }
var x *` + pname + `
_ = x
if p, ok := src.(*` + pname + `); ok { x = p } else if v, ok := src.(` + pname + `); ok { x = &v } else { return }`

	// Getter methods.
	c.wl("func (", recv, " *", inst, ") Get(src interface{}, path ...string) (interface{}, error) {")
	c.wl("var buf interface{}\nerr := " + recv + ".GetTo(src, &buf, path...)\nreturn buf, err")
	c.wdl("}")

	c.wl("func (", recv, " *", inst, ") GetTo(src interface{}, buf *interface{}, path ...string) (err error) {")
	c.wdl(funcHeader)
	err = c.writeNode(node, recv, "x", 0, modeGet)
	if err != nil {
		return err
	}
	c.wdl("return }")

	// Compare method.
	c.wl("func (", recv, " *", inst, ") Cmp(src interface{}, cond inspector.Op, right string, result *bool, path ...string) (err error) {")
	c.wdl(funcHeader)
	err = c.writeNode(node, recv, "x", 0, modeCmp)
	if err != nil {
		return err
	}
	c.wdl("return }")

	// Loop method.
	c.wl("func (", recv, " *", inst, ") Loop(src interface{}, l inspector.Looper, buf *[]byte, path ...string) (err error) {")
	c.wdl(funcHeader)
	err = c.writeNode(node, recv, "x", 0, modeLoop)
	if err != nil {
		return err
	}
	c.wdl("return }")

	// Setter method.
	c.wl("func (", recv, " *", inst, ") Set(dst, value interface{}, path ...string) {")
	// ...
	c.wdl("}")

	return c.err
}

func (c *Compiler) writeNode(node *node, recv, v string, depth int, mode mode) error {
	depths := strconv.Itoa(depth)

	// requireLenCheck := node.typ != typeBasic
	requireLenCheck := node.typ != typeBasic && !(mode == modeLoop && (node.typ == typeMap || (node.typ == typeSlice && node.typn != "[]byte")))

	if requireLenCheck {
		c.wl("if len(path) > ", depths, " {")
	}
	if node.ptr {
		c.wl("if ", v, " == nil { return }")
	}
	switch node.typ {
	case typeStruct:
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
				}
			} else {
				nv := "x" + strconv.Itoa(depth)
				c.wl(nv, " := ", v, ".", ch.name)
				c.wl("_ = ", nv)
				err := c.writeNode(ch, recv, nv, depth+1, mode)
				if err != nil {
					return err
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
			c.wl("for k := range ", c.fmtV(node, v), " {")
			c.wl("if l.RequireKey() {")
			switch node.mapk.typn {
			case "string", "[]byte":
				c.wl("*buf = append((*buf)[:0], k...)")
			case "bool":
				c.wl(`if k { *buf = append((*buf)[:0], "true"...) } else { *buf = append(*buf[:0], "false"...) }`)
			case "int", "int8", "int16", "int32", "int64":
				c.wl("*buf = strconv.AppendInt((*buf)[:0], int64(k), 10)")
			case "uint", "uint8", "uint16", "uint32", "uint64":
				c.wl("*buf = strconv.AppendUint((*buf)[:0], uint64(k), 10)")
			case "float32", "float64":
				c.wl("*buf = strconv.AppendFloat((*buf)[:0], float6464(k), 'f', -1, 64)")
			default:
				c.regImport([]string{`"github.com/koykov/cbytealg"`})
				c.wl("*buf, err = cbytealg.AnyToBytes(*buf[:0], k)")
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
			c.wl("l.Loop()")
			c.wl("}")
			c.wl("return")
		default:
			nv := "x" + strconv.Itoa(depth)
			if node.mapk.typn == "string" {
				c.wl("if ", nv, ", ok := ", c.fmtV(node, v), "[path[", depths, "]]; ok {")
				c.wl("_ = ", nv)
				err := c.writeNode(node.mapv, recv, nv, depth+1, mode)
				if err != nil {
					return err
				}
				c.wl("}")
			} else {
				c.wl("var k ", node.mapk.typn)
				snippet, imports, err := StrConvSnippet("path["+depths+"]", node.mapk.typn, "k")
				c.regImport(imports)
				if err != nil {
					return err
				}
				c.wl(snippet)
				c.wl(nv, " := ", c.fmtV(node, v), "[k]")
				c.wl("_ = ", nv)
				err = c.writeNode(node.mapv, recv, nv, depth+1, mode)
				if err != nil {
					return err
				}
			}
		}
		node.ptr = origPtr
	case typeSlice:
		if node.typn == "[]byte" {
			switch mode {
			case modeGet:
				c.wl("*buf = &", v)
				c.wl("return")
			case modeCmp:
				c.writeCmp(node, v)
				c.wl("return")
			}
		}
		switch mode {
		case modeLoop:
			c.wl("for k := range ", c.fmtV(node, v), " {")
			c.wl("if l.RequireKey() {")
			c.wl("*buf = strconv.AppendInt((*buf)[:0], int64(k), 10)")
			c.wl("l.SetKey(buf, &inspector.StaticInspector{})")
			c.wl("}")
			c.wl("l.SetVal(&", c.fmtV(node, v), "[k], &TestHistoryInspector{})")
			c.wl("l.Loop()")
			c.wl("}")
			c.wl("return")
		default:
			nv := "x" + strconv.Itoa(depth)
			c.wl("var i int")
			snippet, imports, err := StrConvSnippet("path["+depths+"]", "int", "i")
			c.regImport(imports)
			if err != nil {
				return err
			}
			c.wl(snippet)
			c.wl("if len(", v, ") > i {")
			c.wl(nv, " := &", v, "[i]")
			c.wl("_ = ", nv)
			err = c.writeNode(node.slct, recv, nv, depth+1, mode)
			if err != nil {
				return err
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
		}
	}
	if requireLenCheck {
		c.wl("}")
	}
	if (node.typ == typeMap || node.typ == typeSlice) && mode == modeGet {
		c.wl("*buf = ", v)
	}

	return c.err
}

func (c *Compiler) writeCmp(left *node, leftVar string) {
	c.wl("var rightExact ", left.typn)
	snippet, imports, err := StrConvSnippet("right", left.typn, "rightExact")
	c.regImport(imports)
	if err != nil {
		c.err = err
		return
	}
	c.wl(snippet)

	switch left.typn {
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

func (c *Compiler) fmtV(node *node, v string) string {
	if node.ptr {
		return "(*" + v + ")"
	}
	return "(" + v + ")"
}

func (c *Compiler) w(s ...string) {
	_, err := c.wr.WriteString(strings.Join(s, ""))
	if err != nil && c.err == nil {
		c.err = err
	}
}

func (c *Compiler) wl(s ...string) {
	s = append(s, "\n")
	c.w(s...)
}

func (c *Compiler) wdl(s ...string) {
	s = append(s, "\n\n")
	c.w(s...)
}

func (c *Compiler) r(old, new string) {
	s := c.wr.String()
	s = strings.Replace(s, old, new, -1)
	c.wr.Reset()
	_, _ = c.wr.WriteString(s)
}
