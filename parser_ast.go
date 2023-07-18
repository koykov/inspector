package inspector

import (
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

func (c *Compiler) parseFile(path string) error {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		return err
	}
	return c.parseAstFile(file)
}

func (c *Compiler) parseAstFile(file *ast.File) error {
	ast.Inspect(file, func(n ast.Node) bool {
		ts, ok := n.(*ast.TypeSpec)
		if !ok {
			return true
		}
		node, err := c.parseAstType(ts)
		if err != nil {
			return true
		}
		if node == nil || node.typ == typeBasic || reMap.MatchString(node.typn) || reSlc.MatchString(node.typn) {
			return true
		}
		// Check and skip type if it is already parsed or blacklisted.
		if _, ok := c.uniq[node.name]; ok {
			return true
		}
		if len(c.bl) > 0 {
			if _, ok := c.bl[node.name]; ok {
				return true
			}
		}
		c.uniq[node.name] = true

		c.nodes = append(c.nodes, node)
		return true
	})
	return nil
}

func (c *Compiler) parseAstType(ts *ast.TypeSpec) (*node, error) {
	node := &node{
		typ:  typeBasic,
		typn: strings.Replace(ts.Name.String(), c.pkgDot, "", 1),
		ptr:  false,
	}
	if ts.Type == nil {
		return nil, nil
	}
	switch ts.Type.(type) {
	case *ast.MapType:
		m := ts.Type.(*ast.MapType)
		_ = m
		var err error
		node.typ = typeMap
		node.mapk, err = c.parseAstExpr(m.Key)
		node.mapv, err = c.parseAstExpr(m.Value)
		node.hasb = node.mapk.hasb || node.mapv.hasb
		node.hasc = true
		return node, err
	case *ast.StructType:
		s := ts.Type.(*ast.StructType)
		_ = s
		var err error
		node.typ = typeStruct
		if s.Fields != nil {
			for i := 0; i < len(s.Fields.List); i++ {
				field := s.Fields.List[i]
				ch, err := c.parseAstField(field)
				if err != nil {
					return nil, err
				}
				node.chld = append(node.chld, ch)
				node.hasb = node.hasb || ch.hasb
				node.hasc = node.hasc || ch.hasc
			}
		}
		_ = err
	}
	for i := 0; i < len(ts.TypeParams.List); i++ {
		field := ts.TypeParams.List[i]
		for j := 0; j < len(field.Names); j++ {
			name := field.Names[j]
			println(" *", name.String())
		}
	}
	return node, nil
}

func (c *Compiler) parseAstExpr(expr ast.Expr) (*node, error) {
	id, ok := expr.(*ast.Ident)
	if !ok {
		return nil, nil
	}
	node := &node{
		typ:  typeBasic,
		typn: strings.Replace(id.String(), c.pkgDot, "", 1),
		ptr:  false,
	}
	if id.Obj == nil {
		node.typu = id.String()
		node.hasb = node.typu == "string"
		node.hasc = node.hasb
	}
	return node, nil
}

func (c *Compiler) parseAstField(field *ast.Field) (*node, error) {
	node, err := c.parseAstExpr(field.Type)
	if err != nil {
		return nil, err
	}
	if len(field.Names) > 0 {
		node.name = field.Names[0].String()
	}
	return node, nil
}
