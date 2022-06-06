package function

import (
	"go/ast"
	"go/token"
	"log"
)

type function struct {
	comments []*ast.CommentGroup
}

func New() *function {
	return &function{}
}

func (f *function) Run(node *ast.File, currentNode ast.Node, fset *token.FileSet) bool {

	c, ok := currentNode.(*ast.CommentGroup)
	if ok {
		f.comments = append(f.comments, c)
		log.Println("Node CG: ", c.List[0])
	}
	fn, ok := currentNode.(*ast.FuncDecl)
	if ok {
		log.Println("Func doc text: ", fn.Doc.Text())
		if fn.Name.IsExported() {
			log.Printf("exported function declaration without documentation found on line %d: \n\t%s\n", fset.Position(fn.Pos()).Line, fn.Name.Name)
			log.Println("POSITION", fn.Pos(), fset.Position(fn.Pos()))

			comment := &ast.Comment{
				Text:  "// FUCK THE POLICE 1",
				Slash: fn.Pos() - 1,
			}
			comment2 := &ast.Comment{
				Text:  "// FUCK THE POLICE 2",
				Slash: fn.Pos() - 1,
			}
			comment3 := &ast.Comment{
				Text:  "// FUCK THE POLICE 3",
				Slash: fn.Pos() - 1,
			}
			comment4 := &ast.Comment{
				Text:  "// FUCK THE POLICE 4",
				Slash: fn.Pos() - 1,
			}
			cg := &ast.CommentGroup{
				List: []*ast.Comment{comment, comment2, comment3, comment4},
			}
			if fn.Doc.Text() != "" {
				cg.List = append(cg.List, fn.Doc.List...)
			}
			fn.Doc = cg
		}

	}
	node.Comments = f.comments
	return true
}
