package inspector

import (
	"go/types"
	"strings"

	"golang.org/x/tools/go/loader"
)

type typ int

const (
	typeStruct typ = iota
	typeMap
	typeSlice
	typeBasic
)

type Compiler struct {
	pkg     string
	pkgDot  string
	pkgName string
	outDir  string
	uniq    map[string]bool
	nodes   []*node
}

type node struct {
	typ  typ
	typn string
	name string
	ptr  bool
	chld []*node
	mapk *node
	mapv *node
	slct *node
}

func NewCompiler(pkg, outDir string) *Compiler {
	c := Compiler{
		pkg:    pkg,
		pkgDot: pkg + ".",
		outDir: outDir,
		uniq:   make(map[string]bool),
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

	u := t.Underlying()
	if p, ok := u.(*types.Pointer); ok {
		u = p.Elem().Underlying()
		node, err := c.parseType(u)
		if err != nil {
			return node, err
		}
		node.ptr = true
		return node, err
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
				ch.typn = "*" + ch.name
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
	return nil
}
