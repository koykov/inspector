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
		c.pkgName = file.Name.String()
		node, err := c.parseAstExpr(ts.Type, ts.Name, 0)
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

func (c *Compiler) parseAstExpr(expr ast.Expr, id *ast.Ident, depth int) (*node, error) {
	var err error
	node := &node{typ: typeBasic}
	if id != nil {
		node.name = strings.Replace(id.String(), c.pkgDot, "", 1)
		if id.Obj != nil {
			node.name = id.Obj.Name
		}
	}
	if depth == 0 {
		node.typn = id.String()
		node.pkg = c.pkgName
		node.pkgi = c.imp_
	}
	switch expr.(type) {
	case *ast.MapType:
		node.typ = typeMap
		if id != nil {
			node.name = strings.Replace(id.String(), c.pkgDot, "", 1)
		}
		m := expr.(*ast.MapType)
		if node.mapk, err = c.parseAstExpr(m.Key, nil, depth+1); err != nil {
			return nil, err
		}
		if node.mapv, err = c.parseAstExpr(m.Value, nil, depth+1); err != nil {
			return nil, err
		}
		node.hasb = node.mapk.hasb || node.mapv.hasb
		node.hasc = true
		return node, nil
	case *ast.StructType:
		node.typ = typeStruct
		node.typn = node.name
		node.pkg = c.pkgName
		node.pkgi = c.imp_
		s := expr.(*ast.StructType)
		if s.Fields != nil {
			for i := 0; i < len(s.Fields.List); i++ {
				field := s.Fields.List[i]
				var id *ast.Ident
				if len(field.Names) > 0 {
					id = field.Names[0]
				}
				ch, err := c.parseAstExpr(field.Type, id, depth+1)
				if err != nil {
					return nil, err
				}
				ch.name = id.String()
				if len(ch.typn) == 0 {
					ch.typn = c.composeAstTypeName(ch)
				}
				node.chld = append(node.chld, ch)
				node.hasb = node.hasb || ch.hasb
				node.hasc = node.hasc || ch.hasc
			}
		}
		return node, nil
	case *ast.SliceExpr:
		node.typ = typeSlice
		s := expr.(*ast.SliceExpr)
		_ = s
		return node, nil
	case *ast.ArrayType:
		node.typ = typeSlice
		s := expr.(*ast.ArrayType)
		if node.slct, err = c.parseAstExpr(s.Elt, nil, depth+1); err != nil {
			return nil, err
		}
		node.typn = "[]" + node.slct.typn
		node.hasb = node.typn == "[]byte" || node.slct.hasb
		node.hasc = true
		return node, nil
	case *ast.StarExpr:
		s := expr.(*ast.StarExpr)
		if node, err = c.parseAstExpr(s.X, nil, depth+1); err != nil {
			return nil, err
		}
		node.ptr = true
		return node, nil
	case *ast.Ident:
		id := expr.(*ast.Ident)
		node.typn = id.String()
		node.typu = id.String()
		node.hasb = node.typu == "string"
		node.hasc = node.hasb
		if id.Obj != nil {
			if ts, ok := id.Obj.Decl.(*ast.TypeSpec); ok {
				if node, err = c.parseAstExpr(ts.Type, ts.Name, depth+1); err != nil {
					return nil, err
				}
				node.name = ""
				node.typn = ts.Name.String()
				node.pkg = c.pkgName
				node.pkgi = c.imp_
			}
		}
		return node, nil
	default:
		return nil, ErrUnsupportedType
	}
}

func (c *Compiler) composeAstTypeName(node *node) string {
	var typn string
	switch node.typ {
	case typeMap:
		typn = "map["
		if node.mapk.ptr {
			typn += "*"
		}
		typn += node.mapk.typn + "]"
		if node.mapv.ptr {
			typn += "*"
		}
		typn += node.mapv.typn
	}
	return typn
}
