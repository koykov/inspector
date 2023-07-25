package inspector

import (
	"bytes"
	"go/format"
	"go/types"
	"os"
	"regexp"
	"strconv"
	"strings"
	"syscall"

	"golang.org/x/tools/go/loader"
)

type Target int

// Internal node type.
type typ int

// Node compile mode.
type mode int

const (
	TargetPackage Target = iota
	TargetDirectory
	TargetFile
)
const (
	// Possible node's types.
	typeStruct typ = iota
	typeMap
	typeSlice
	typeBasic
)
const (
	// Possible compile modes.
	modeGet mode = iota
	modeSet
	modeCmp
	modeLoop
)

// ByteStringWriter is an extended writer interface.
type ByteStringWriter interface {
	Write(p []byte) (n int, err error)
	WriteByte(c byte) error
	WriteString(s string) (n int, err error)
	Bytes() []byte
	String() string
	Reset()
}

// Logger interface.
type Logger interface {
	Print(v ...any)
	Println(v ...any)
}

type Compiler struct {
	trg           Target
	pkg           string
	pkgDot        string
	pkgName       string
	imp_          string
	dst           string
	dstAbs        string
	bl            map[string]struct{}
	uniq          map[string]struct{}
	nodes         []*node
	imp           []string
	cntr, cntrDEQ int
	l             Logger
	wr            ByteStringWriter
	nc            bool

	cnf *Config
	err error
}

var (
	// Check map.
	reMap = regexp.MustCompile(`map\[[^]]+].*`)
	// Check slice.
	reSlc = regexp.MustCompile(`\[].*`)
	// Check if first symbol is in upper case.
	reUp = regexp.MustCompile(`^[A-Z]+.*`)
	// Check vendor directory.
	reVnd = regexp.MustCompile(`.*/vendor/(.*)`)
)

func NewCompiler(conf *Config) (*Compiler, error) {
	cc := conf.Copy()
	var source string
	switch conf.Target {
	case TargetPackage:
		source = conf.Package
	case TargetDirectory:
		source = conf.Directory
	case TargetFile:
		source = conf.File
	default:
		return nil, ErrUnknownTarget
	}
	c := Compiler{
		trg:    cc.Target,
		pkg:    source,
		pkgDot: source + ".",
		dst:    cc.Destination,
		bl:     cc.BlackList,
		uniq:   make(map[string]struct{}),
		nc:     cc.NoClean || len(conf.XML) > 0,
		wr:     cc.Buf,
		l:      cc.Logger,
		imp:    make([]string, 0),
		cnf:    cc,
	}
	if len(cc.Import) > 0 {
		c.imp_ = cc.Import
		c.pkgDot = cc.Import + "."
	}
	return &c, nil
}

func (c *Compiler) Compile() error {
	if c.cnf == nil {
		return ErrNoConfig
	}
	if err := c.parse(); err != nil {
		return err
	}
	// Prepare destination.
	if c.l != nil {
		c.l.Print("Prepare destination dir " + c.dst)
	}
	ps := string(os.PathSeparator)
	switch c.trg {
	case TargetPackage:
		gopath := os.Getenv("GOPATH")
		if len(gopath) == 0 {
			return ErrNoGOPATH
		}
		c.dstAbs = os.Getenv("GOPATH") + ps + "src" + ps + c.dst
	default:
		wd, err := os.Getwd()
		if err != nil {
			return err
		}
		c.dstAbs = wd + ps + c.dst
		c.pkg = c.imp_
	}

	dstExists := true
	if _, err := os.Stat(c.dstAbs); os.IsNotExist(err) {
		dstExists = false
	}
	if dstExists && !c.nc {
		if err := syscall.Access(c.dstAbs, syscall.O_RDWR); err != nil {
			return err
		}
		if err := os.RemoveAll(c.dstAbs); err != nil {
			return err
		}
	}
	if !dstExists {
		if err := os.MkdirAll(c.dstAbs, 0755); err != nil {
			return ErrDstNotExists
		}
	}

	// Walk over nodes and compile each of them to separate file.
	for _, node := range c.nodes {
		file := strings.ToLower(node.name) + ".ins.go"
		if c.l != nil {
			c.l.Print("Compiling ", node.name, "Inspector to "+file)
		}
		c.wr.Reset()
		if err := c.writeType(node); err != nil {
			return err
		}
		if err := c.writeFile(c.dstAbs + ps + file); err != nil {
			return err
		}
	}

	return nil
}

func (c *Compiler) WriteXML() error {
	if err := c.parse(); err != nil {
		return err
	}
	// Prepare destination.
	ps := string(os.PathSeparator)
	if len(c.dstAbs) == 0 {
		wd, err := os.Getwd()
		if err != nil {
			return err
		}
		c.dstAbs = wd + ps + c.cnf.XML
	}
	if c.l != nil {
		c.l.Print("Prepare destination dir " + c.dst)
	}
	dstExists := true
	if _, err := os.Stat(c.dstAbs); os.IsNotExist(err) {
		dstExists = false
	}
	if dstExists {
		if err := syscall.Access(c.dstAbs, syscall.O_RDWR); err != nil {
			return err
		}
		if err := os.RemoveAll(c.dstAbs); err != nil {
			return err
		}
	}
	if err := os.MkdirAll(c.dstAbs, 0755); err != nil {
		return ErrDstNotExists
	}

	// Walk over nodes and compile each of them to separate file.
	var buf bytes.Buffer
	for _, node := range c.nodes {
		file := strings.ToLower(node.name) + ".xml"
		if c.l != nil {
			c.l.Print("Writing ", node.name, " type to ", file)
		}
		buf.Reset()
		if err := node.write(&buf); err != nil {
			return err
		}
		if err := os.WriteFile(c.dstAbs+ps+file, buf.Bytes(), 0644); err != nil {
			return err
		}
		c.cntr++
	}
	return nil
}

func (c *Compiler) parse() error {
	var err error
	switch c.trg {
	case TargetPackage:
		if c.l != nil {
			c.l.Print("Parse package " + c.pkg)
		}
		// Try import the package.
		var conf loader.Config
		conf.Import(c.pkg)
		prog, err := conf.Load()
		if err != nil {
			return err
		}

		// Parse the package to nodes list.
		pkg := prog.Package(c.pkg)
		c.pkgName = pkg.Pkg.Name()
		err = c.parsePkg(pkg)
		if err != nil {
			return err
		}
	case TargetDirectory:
		if c.l != nil {
			c.l.Print("Parse directory " + c.pkg)
		}
		// todo implement me
	case TargetFile:
		err = c.parseFile(c.pkg)
	}
	return err
}

// GetTotal returns the total number of compiled types.
func (c *Compiler) GetTotal() int {
	return c.cntr
}

// Format compiled inspector code using format package and write it to the given file.
func (c *Compiler) writeFile(filename string) error {
	source := c.wr.Bytes()
	fmtSource, err := format.Source(source)
	if err != nil {
		// return err
		return os.WriteFile(filename, source, 0644)
	}

	return os.WriteFile(filename, fmtSource, 0644)
}

// Write given type including headers and imports.
func (c *Compiler) writeType(node *node) error {
	c.wl("// Code generated by inspc. DO NOT EDIT.")
	c.wdl("// source: ", c.pkg)

	c.wdl("package ", c.pkgName, "_ins")

	c.imp = append(c.imp[:0], `"github.com/koykov/inspector"`, `"`+c.pkg+`"`)
	c.wdl("import (\n!{import}\n)")

	c.wl("func init() {")
	c.wl(`inspector.RegisterInspector("`, node.name, `", `, node.name, `Inspector{})`)
	c.wdl("}")

	err := c.writeRootNode(node)
	if err != nil {
		return err
	}
	c.cntr++
	c.r("!{import}", strings.Join(c.imp, "\n"))

	return nil
}

// Register imports for output file.
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

// Write methods of the inspector required by Inspector interface type.
func (c *Compiler) writeRootNode(node *node) (err error) {
	inst := node.name + "Inspector"
	recv := "i" + strconv.Itoa(c.cntr)
	pname := c.pkgName + "." + node.typn

	c.wdl("type ", inst, " struct {\ninspector.BaseInspector\n}")

	// Common checks and type cast.
	funcHeader := `if len(path) == 0 { return }
if src == nil { return }
var x *` + pname + `
_ = x
if p, ok := src.(**` + pname + `); ok { x = *p } else if p, ok := src.(*` + pname + `); ok { x = p } else if v, ok := src.(` + pname + `); ok { x = &v } else { return }`
	// Custom header for Set() method.
	funcHeaderSet := `if len(path) == 0 { return nil }
if dst == nil { return nil }
var x *` + pname + `
_ = x
if p, ok := dst.(**` + pname + `); ok { x = *p } else if p, ok := dst.(*` + pname + `); ok { x = p } else if v, ok := dst.(` + pname + `); ok { x = &v } else { return nil }`
	// Custom header for GetTo() method.
	funcHeaderGetTo := `if src == nil { return }
var x *` + pname + `
_ = x
if p, ok := src.(**` + pname + `); ok { x = *p } else if p, ok := src.(*` + pname + `); ok { x = p } else if v, ok := src.(` + pname + `); ok { x = &v } else { return }
if len(path) == 0 { *buf = &(*x)
return}`
	// Header for Loop() method.
	funcHeaderLoop := ""
	if node.typ != typeSlice {
		funcHeaderLoop += "if len(path) == 0 { return }\n"
	}
	funcHeaderLoop += `if src == nil { return }
var x *` + pname + `
_ = x
if p, ok := src.(**` + pname + `); ok { x = *p } else if p, ok := src.(*` + pname + `); ok { x = p } else if v, ok := src.(` + pname + `); ok { x = &v } else { return }`
	// Header for DeepEqual() method.
	funcHeaderEqual := `var (
lx, rx *` + pname + `
leq, req bool
)
_, _, _, _ = lx, rx, leq, req
if lp, ok := l.(**` + pname + `); ok { lx, leq = *lp, true } else if lp, ok := l.(*` + pname + `); ok { lx, leq = lp, true } else if lp, ok := l.(` + pname + `); ok { lx, leq = &lp, true }
if rp, ok := r.(**` + pname + `); ok { rx, req = *rp, true } else if rp, ok := r.(*` + pname + `); ok { rx, req = rp, true } else if rp, ok := r.(` + pname + `); ok { rx, req = &rp, true }
if !leq || !req { return false }
if lx == nil && rx == nil { return true }
if (lx == nil && rx != nil) || (lx != nil && rx == nil) { return false }
`
	// Header for Length()/Capacity() methods.
	funcHeaderLC := `if src == nil { return nil }
var x *` + pname + `
_ = x
if p, ok := src.(**` + pname + `); ok { x = *p } else if p, ok := src.(*` + pname + `); ok { x = p } else if v, ok := src.(` + pname + `); ok { x = &v } else { return inspector.ErrUnsupportedType }`

	// Getter methods.
	c.wl("func (", recv, " ", inst, ") TypeName() string {")
	c.wl("return \"", node.typn, "\"")
	c.wdl("}")

	c.wl("func (", recv, " ", inst, ") Get(src any, path ...string) (any, error) {")
	c.wl("var buf any\nerr := " + recv + ".GetTo(src, &buf, path...)\nreturn buf, err")
	c.wdl("}")

	c.wl("func (", recv, " ", inst, ") GetTo(src any, buf *any, path ...string) (err error) {")
	c.wdl(funcHeaderGetTo)
	err = c.writeNode(node, nil, recv, "x", "", 0, modeGet)
	if err != nil {
		return err
	}
	c.wdl("return }")

	// Compare method.
	c.wl("func (", recv, " ", inst, ") Compare(src any, cond inspector.Op, right string, result *bool, path ...string) (err error) {")
	c.wdl(funcHeader)
	err = c.writeNode(node, nil, recv, "x", "", 0, modeCmp)
	if err != nil {
		return err
	}
	c.wdl("return }")

	// Loop method.
	c.wl("func (", recv, " ", inst, ") Loop(src any, l inspector.Iterator, buf *[]byte, path ...string) (err error) {")
	c.wdl(funcHeaderLoop)
	err = c.writeNode(node, nil, recv, "x", "", 0, modeLoop)
	if err != nil {
		return err
	}
	c.wdl("return }")

	// Setter methods.
	c.wl("func (", recv, " ", inst, ") SetWithBuffer(dst, value any, buf inspector.AccumulativeBuffer, path ...string) error {")
	c.wdl(funcHeaderSet)
	err = c.writeNode(node, nil, recv, "x", "", 0, modeSet)
	if err != nil {
		return err
	}
	c.wdl("return nil }")
	c.wl("func (", recv, " ", inst, ") Set(dst, value any, path ...string) error {")
	c.wl("return ", recv, ".SetWithBuffer(dst, value, nil, path...)")
	c.wdl("}")

	// DeepEqual methods.
	c.cntrDEQ = 0
	c.wl("func (", recv, " ", inst, ") DeepEqual(l, r any) bool {")
	c.wl("return ", recv, ".DeepEqualWithOptions(l, r, nil)")
	c.wdl("}")
	c.wl("func (", recv, " ", inst, ") DeepEqualWithOptions(l, r any, opts *inspector.DEQOptions) bool {")
	c.wdl(funcHeaderEqual)
	err = c.writeNodeDEQ(node, nil, recv, "", "lx", "rx", 0)
	if err != nil {
		return err
	}
	c.wdl("return true }")

	// Encoding methods.
	c.wl("func (", recv, " ", inst, ") Unmarshal(p []byte, typ inspector.Encoding) (any, error) {")
	err = c.writeNodeParse(node, pname)
	if err != nil {
		return err
	}
	c.wdl("}")

	// Copy methods.
	c.wl("func (", recv, " ", inst, ") Copy(x any) (any, error) {")
	err = c.writeNodeCopy(node, recv, pname)
	if err != nil {
		return err
	}
	c.wdl("}")
	c.wl("func (", recv, " ", inst, ") CopyTo(src, dst any, buf inspector.AccumulativeBuffer) error {")
	err = c.writeNodeCopyTo(node, recv, pname)
	if err != nil {
		return err
	}
	c.wdl("}")
	c.wl("func (", recv, " ", inst, ") countBytes(x *", pname, ") (c int) {")
	err = c.writeCountBytes(node, "x", 0)
	if err != nil {
		return err
	}
	c.wdl("return c}")
	c.wl("func (", recv, " ", inst, ") cpy(buf []byte,l, r *", pname, ") ([]byte, error) {")
	err = c.writeCopy(node, "l", "r", 0)
	if err != nil {
		return err
	}
	c.wdl("return buf,nil}")

	// Len/cap methods.
	c.wl("func (", recv, " ", inst, ") Length(src any, result *int, path ...string) error {")
	c.wdl(funcHeaderLC)
	err = c.writeNodeLC(node, "x", "len", 0)
	if err != nil {
		return err
	}
	c.wdl("return nil}")
	c.wl("func (", recv, " ", inst, ") Capacity(src any, result *int, path ...string) error {")
	c.wdl(funcHeaderLC)
	err = c.writeNodeLC(node, "x", "cap", 0)
	if err != nil {
		return err
	}
	c.wdl("return nil}")

	// Reset methods.
	c.wl("func (", recv, " ", inst, ") Reset(x any) error {")
	c.wl("var origin *", pname)
	c.wl("_=origin")
	c.wl("switch x.(type) {")
	c.wl("case ", pname, ":")
	c.wl("return inspector.ErrMustPointerType")
	c.wl("case *", pname, ":")
	c.wl("origin = x.(*", pname, ")")
	c.wl("case **", pname, ":")
	c.wl("origin = *x.(**", pname, ")")
	c.wl("default:")
	c.wl("return inspector.ErrUnsupportedType")
	c.wl("}")
	err = c.writeNodeReset(node, "origin", 0)
	if err != nil {
		return err
	}
	c.wdl("return nil}")

	return c.err
}

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
					if mode == modeSet {
						c.wl("inspector.AssignBuf(", pfx, v, ".", ch.name, ", value, buf)")
					}
					c.wl("return nil")
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
				c.wl("*buf, err = x2bytes.AnyToBytes(*buf[:0], k)")
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
				c.wl("var k ", node.mapk.typn)
				snippet, imports, err := StrConvSnippet("path["+depths+"]", node.mapk.typn, node.mapk.typu, "k")
				c.regImport(imports)
				if err != nil {
					return err
				}
				c.wl(snippet)
				c.wl(nv, " := ", c.fmtV(node, v), "[", c.fmtP(node.mapk, "k", depth+1), "]")
				c.wl("_ = ", nv)
				err = c.writeNode(node.mapv, node, recv, nv, "", depth+1, mode)
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
				if mode == modeSet {
					c.wl("inspector.AssignBuf(", pfx, v, ", value, buf)")
				}
				c.wl("return nil")
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
			if mode == modeSet {
				c.wl("inspector.AssignBuf(", pfx, v, ", value, buf)")
			}
			if parent.typ != typeMap {
				c.wl("return nil")
			}
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

func (c *Compiler) writeNodeParse(_ *node, pname string) error {
	c.regImport([]string{`"encoding/json"`})
	c.wl("var x ", pname)
	c.wl(`switch typ {
case inspector.EncodingJSON:
err := json.Unmarshal(p, &x)
return &x, err
default:
return nil, inspector.ErrUnknownEncodingType
}`)
	return nil
}

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
			c.wl("buf,", c.fmtVnb(node, l, depth), "=inspector.BufferizeString(buf,", c.fmtVnb(node, r, depth), ")")
		} else {
			c.wl(l, "=", r)
		}
	}
	return nil
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

func (c *Compiler) writeNodeReset(node *node, v string, depth int) error {
	switch node.typ {
	case typeStruct:
		for _, ch := range node.chld {
			nv := v + "." + ch.name
			chPtr := ch.ptr && (ch.typ == typeStruct || ch.typ == typeMap || ch.typ == typeSlice)
			if chPtr {
				c.wl("if ", nv, "!=nil{")
			}
			_ = c.writeNodeReset(ch, nv, depth+1)
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
				_ = c.writeNodeReset(node.slct, nv, depth+1)
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

// Check if given type is a builtin type.
func (c *Compiler) isBuiltin(typ string) bool {
	switch typ {
	case "int", "int8", "int16", "int32", "int64",
		"uint", "uint8", "uint16", "uint32", "uint64",
		"float32", "float64",
		"string", "[]byte", "byte", "bool":
		return true
	default:
		return false
	}
}

// Check if given object has an exportable type.
func (c *Compiler) isExported(o types.Object) bool {
	if !o.Exported() {
		return false
	}
	if _, ok := o.(*types.Var); ok {
		return false
	}
	t := o.Type()
	if n, ok := t.(*types.Named); ok {
		if !reUp.MatchString(n.Obj().Name()) {
			return false
		}
	}
	u := t.Underlying()
	if p, ok := u.(*types.Pointer); ok {
		_ = p
		return true
	}
	return true
}

// Remove vendor path from the package.
func (c *Compiler) clearPkg(pkgi string) string {
	if m := reVnd.FindStringSubmatch(pkgi); m != nil {
		return m[1]
	}
	return pkgi
}

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
func (c *Compiler) fmtT(node *node) string {
	switch node.typ {
	case typeStruct:
		return node.pkg + "." + strings.Trim(node.typn, "*")
	case typeMap:
		if strings.Contains(node.typn, "map[") {
			s := "map["
			if node.mapk.ptr {
				s += "*"
			}
			if len(node.mapk.pkg) > 0 {
				s += node.mapk.pkg + "."
			}
			s += node.mapk.typn
			s += "]"
			if node.mapv.ptr {
				s += "*"
			}
			if len(node.mapv.pkg) > 0 {
				s += node.mapv.pkg + "."
			}
			s += node.mapv.typn
			return s
		} else {
			return node.pkg + "." + strings.Trim(node.typn, "*")
		}
	case typeSlice:
		s := node.slct.typn
		if !c.isBuiltin(s) {
			s = node.slct.pkg + "." + strings.Trim(node.slct.typn, "*")
			if node.slct.ptr {
				s = "*" + s
			}
		}
		if strings.Index(node.typn, "[]") != -1 {
			pfx := ""
			if node.slct.ptr {
				pfx = "*"
			}
			s = "[]" + pfx + strings.TrimLeft(s, "*")
		} else {
			s = node.pkg + "." + strings.Trim(node.typn, "*")
		}
		return s
	case typeBasic:
		pfx := ""
		if node.ptr {
			pfx = "*"
		}
		return pfx + node.typn
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

func (t typ) String() string {
	switch t {
	case typeBasic:
		return "basic"
	case typeStruct:
		return "struct"
	case typeMap:
		return "map"
	case typeSlice:
		return "slice"
	default:
		return "unknown"
	}
}
