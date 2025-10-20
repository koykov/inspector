package inspector

import (
	"go/types"
	"strings"

	"golang.org/x/tools/go/loader"
)

// Parse package.
func (c *Compiler) parsePkg(pkg *loader.PackageInfo) error {
	for _, scope := range pkg.Info.Scopes {
		// Only scopes without parents available.
		if parent := scope.Parent(); parent != nil {
			for _, name := range parent.Names() {
				// Get the object representation of scope.
				o := parent.Lookup(name)
				if !c.isExported(o) {
					continue
				}
				// Get type of object and parse it.
				t := o.Type()
				node, err := c.parsePkgType(t)
				if err != nil {
					return err
				}
				if node == nil || node.typ == typeBasic || reMap.MatchString(node.typn) || reSlc.MatchString(node.typn) {
					continue
				}
				if node.typ == typeStruct {
					// Make type name of struct node actual.
					node.typn = o.Name()
				}
				node.pkg = o.Pkg().Name()
				node.name = o.Name()
				// Check and skip type if it is already parsed or blacklisted.
				if _, ok := c.uniq[node.name]; ok {
					continue
				}
				if len(c.bl) > 0 {
					if _, ok := c.bl[node.name]; ok {
						continue
					}
				}
				c.uniq[node.name] = struct{}{}

				c.nodes = append(c.nodes, node)
			}
		}
	}

	return nil
}

// Main parse method.
func (c *Compiler) parsePkgType(t types.Type) (*node, error) {
	// Prepare and fill up default node.
	node := &node{
		typ:  typeBasic,
		typn: strings.Replace(t.String(), c.pkgDot, "", 1),
		ptr:  false,
	}
	if n, ok := t.(*types.Named); ok {
		node.typn = n.Obj().Name()
		if n.Obj().Pkg() == nil {
			return nil, nil
		}
		node.pkg = n.Obj().Pkg().Name()
		node.pkgi = c.clearPkg(n.Obj().Pkg().Path())
	}

	// Get the underlying type.
	u := t.Underlying()
	// Common skips considering by underlying type.
	if _, ok := u.(*types.Interface); ok {
		return nil, nil
	}
	if _, ok := u.(*types.Signature); ok {
		return nil, nil
	}

	if p, ok := u.(*types.Pointer); ok {
		// Dereference pointer type and go inside to parse.
		e := p.Elem()
		u = e.Underlying()
		unode, err := c.parsePkgType(u)
		if err != nil {
			return node, err
		}
		if unode == nil {
			return nil, nil
		}
		unode.ptr = true
		if n, ok := e.(*types.Named); ok {
			unode.typn = n.Obj().Name()
			unode.pkg = n.Obj().Pkg().Name()
			unode.pkgi = c.clearPkg(n.Obj().Pkg().Path())
		}
		return unode, err
	}

	if s, ok := u.(*types.Struct); ok {
		// Walk over fields in struct and parse each of them as separate node.
		node.typ = typeStruct
		for i := 0; i < s.NumFields(); i++ {
			f := s.Field(i)
			if !f.Exported() {
				continue
			}
			ch, err := c.parsePkgType(f.Type())
			if err != nil {
				return node, err
			}
			if ch == nil {
				continue
			}
			ch.name = f.Name()
			if ch.ptr {
				ch.typn = strings.Replace(f.Type().String(), c.pkgDot, "", 1)
				ch.typn = strings.Replace(ch.typn, "*", "", 1)
			}
			node.chld = append(node.chld, ch)
			node.hasb = node.hasb || ch.hasb
			node.hasc = node.hasc || ch.hasc
			node.hasa = node.hasa || ch.hasa
		}
		return node, nil
	}

	if m, ok := u.(*types.Map); ok {
		// Just parse key and value types of map.
		var err error
		node.typ = typeMap
		node.mapk, err = c.parsePkgType(m.Key())
		node.mapv, err = c.parsePkgType(m.Elem())
		node.hasb = node.mapk.hasb || node.mapv.hasb
		node.hasa = node.mapk.hasa || node.mapv.hasa
		node.hasc = true
		return node, err
	}

	if s, ok := u.(*types.Slice); ok {
		// Just parse value type of slice.
		var err error
		node.typ = typeSlice
		node.slct, err = c.parsePkgType(s.Elem())
		node.hasb = node.typn == "[]byte" || node.slct.hasb
		node.hasa = true
		node.hasc = true
		return node, err
	}

	if b, ok := u.(*types.Basic); ok {
		// Fill up the underlying type of basic node.
		node.typu = b.Name()
		node.hasb = node.typu == "string"
		node.hasc = node.hasb
	}

	return node, nil
}
