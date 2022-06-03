package walker

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
)

type walker struct {
	filename string
	fset     *token.FileSet
	Node     *ast.File
	handlers []Handler
}

func NewWalker(filename string) (*walker, error) {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}
	return &walker{
		filename: filename,
		fset:     fset,
		Node:     node,
	}, nil
}

func (walker *walker) Run() (*token.FileSet, *ast.File) {
	walker.traversalAST()
	return walker.fset, walker.Node
}

func (walker *walker) RegisterHandler(handler Handler) {
	walker.handlers = append(walker.handlers, handler)
}

func (walker *walker) traversalAST() {
	log.Println("Traversal AST")
	ast.Inspect(walker.Node, func(n ast.Node) bool {
		for _, handler := range walker.handlers {
			handler.Run(walker.Node, n, walker.fset)
		}
		return true
	})
}

// func modelFuncGenerate() func(node *ast.File, currentNode ast.Node, fset *token.FileSet) {

// }
