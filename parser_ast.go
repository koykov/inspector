package inspector

import (
	"go/ast"
	"go/parser"
	"go/token"
)

func (c *Compiler) parseFile(path string) error {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		return err
	}
	ast.Inspect(file, func(node ast.Node) bool {
		ts, ok := node.(*ast.TypeSpec)
		if !ok {
			return true
		}
		println(ts.Name.String())
		if ts.TypeParams == nil {
			return true
		}
		for i := 0; i < len(ts.TypeParams.List); i++ {
			field := ts.TypeParams.List[i]
			for j := 0; j < len(field.Names); j++ {
				name := field.Names[j]
				println(" *", name.String())
			}
		}
		return true
	})
	return nil
}
