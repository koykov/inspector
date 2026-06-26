package inspector

import (
	"bytes"
	"go/format"
	"go/types"
	"log"
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
	inp           bool
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
		inp:    conf.InPlace,
		wr:     cc.Buf,
		l:      cc.Logger,
		imp:    make([]string, 0),
		cnf:    cc,
	}
	if len(cc.Import) > 0 {
		c.imp_ = cc.Import
		c.pkgDot = cc.Import + "."
	}
	if c.inp {
		c.pkgDot = ""
		c.imp_ = ""
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
		c.dstAbs = c.dst
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
		dstExists = false
	}
	if !dstExists {
		if err := os.MkdirAll(c.dstAbs, 0755); err != nil {
			return ErrDstNotExists
		}
	}

	// Walk over nodes and compile each of them to separate file.
	for _, node := range c.nodes {
		file := strings.ToLower(node.name) + "_ins.go"
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
	if dstExists && !c.nc {
		if err := syscall.Access(c.dstAbs, syscall.O_RDWR); err != nil {
			return err
		}
		if err := os.RemoveAll(c.dstAbs); err != nil {
			return err
		}
		dstExists = false
	}
	if !dstExists {
		if err := os.MkdirAll(c.dstAbs, 0755); err != nil {
			return ErrDstNotExists
		}
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
		if pkg == nil {
			return nil
		}
		c.pkgName = pkg.Pkg.Name()
		err = c.parsePkg(pkg)
		if err != nil {
			return err
		}
	case TargetDirectory:
		if c.l != nil {
			c.l.Print("Parse directory " + c.pkg)
		}
		err = c.parseDir(c.pkg)
	case TargetFile:
		if c.l != nil {
			c.l.Print("Parse file " + c.pkg)
		}
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
		if !c.cnf.Force {
			return err
		}
		log.Println(err)
		fmtSource = source
	}

	return os.WriteFile(filename, fmtSource, 0644)
}

// Write given type including headers and imports.
func (c *Compiler) writeType(node *node) error {
	sfx := "_ins"
	if c.inp {
		sfx = ""
	}
	pkg := c.pkg
	if len(pkg) == 0 && len(c.cnf.File) > 0 {
		pkg = c.cnf.File
	}

	c.wl("// Code generated by inspc. DO NOT EDIT.")
	c.wdl("// source: ", pkg)

	c.wdl("package ", c.pkgName, sfx)

	c.imp = append(c.imp[:0], `"github.com/koykov/inspector"`)
	if !c.inp {
		c.imp = append(c.imp, `"`+c.pkg+`"`)
	}
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
	if c.inp {
		pname = node.typn
	}

	c.wdl("type ", inst, " struct {\ninspector.BaseInspector\n}")

	// Getter methods.
	c.wl("func (", recv, " ", inst, ") TypeName() string {")
	c.wl("return \"", node.typn, "\"")
	c.wdl("}")

	c.wl("func (", recv, " ", inst, ") Instance(ptr bool) any {")
	c.wl("if ptr {return &", pname, "{}}")
	c.wl("return ", pname, "{}")
	c.wdl("}")

	c.wl("func (", recv, " ", inst, ") Get(src any, path ...string) (any, error) {")
	c.wl("var buf any\nerr := " + recv + ".GetTo(src, &buf, path...)\nreturn buf, err")
	c.wdl("}")

	c.wl("func (", recv, " ", inst, ") GetTo(src any, buf *any, path ...string) (err error) {")
	c.wdl(c.getFuncHeaderGetTo(pname))
	err = c.writeNode(node, nil, recv, "x", "", 0, modeGet)
	if err != nil {
		return err
	}
	c.wdl("return }")

	// Compare method.
	c.wl("func (", recv, " ", inst, ") Compare(src any, cond inspector.Op, right string, result *bool, path ...string) (err error) {")
	c.wdl(c.getFuncHeader(pname))
	err = c.writeNode(node, nil, recv, "x", "", 0, modeCmp)
	if err != nil {
		return err
	}
	c.wdl("return }")

	// Loop method.
	c.wl("func (", recv, " ", inst, ") Loop(src any, l inspector.Iterator, buf *[]byte, path ...string) (err error) {")
	c.wdl(c.getFuncHeaderLoop(pname, node.typ))
	err = c.writeNode(node, nil, recv, "x", "", 0, modeLoop)
	if err != nil {
		return err
	}
	c.wdl("return }")

	// Setter methods.
	c.wl("func (", recv, " ", inst, ") SetWithBuffer(dst, value any, buf inspector.AccumulativeBuffer, path ...string) error {")
	c.wdl(c.getFuncHeaderSet(pname))
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
	c.wdl(c.getFuncHeaderEqual(pname))
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
	c.wdl(c.getFuncHeaderLC(pname))
	err = c.writeNodeLC(node, "x", "len", 0)
	if err != nil {
		return err
	}
	c.wdl("return nil}")
	c.wl("func (", recv, " ", inst, ") Capacity(src any, result *int, path ...string) error {")
	c.wdl(c.getFuncHeaderLC(pname))
	err = c.writeNodeLC(node, "x", "cap", 0)
	if err != nil {
		return err
	}
	c.wdl("return nil}")

	c.wl("func (", recv, " ", inst, ") Append(src, value any, path ...string) (any, error) {")
	c.wl("_, _, _ = src, value, path")
	if node.hasc {
		c.wdl(c.getFuncHeaderAppend(pname))
		err = c.writeNodeAppend(node, "x", 0)
	}
	c.wl("return src, nil")
	c.wdl("}")

	// Each method.
	c.wl("func (", recv, " ", inst, ") Each(src any, fn func(i int, field string, value any)) error {")
	c.wdl(c.getFuncHeaderLC(pname))
	err = c.writeNodeEach(node, "x", 0)
	c.wdl("return nil}")

	// Reset methods.
	c.wl("func (", recv, " ", inst, ") Reset(x any, path ...string) error {")
	c.wl("if len(path)==0{return ", recv, ".reset1(x, path...)}else{return ", recv, ".reset2(x, path...)}")
	c.wl("}\n")
	c.wl("func (", recv, " ", inst, ") reset1(x any, path ...string) error {")
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
	err = c.writeNodeResetFull(node, "origin", 0)
	if err != nil {
		return err
	}
	c.wdl("return nil}")
	c.wl("func (", recv, " ", inst, ") reset2(x any, path ...string) error {")
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
