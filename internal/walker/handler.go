package walker

import (
	"go/ast"
	"go/token"
)

type Handler interface {
	Run(node *ast.File, currentNode ast.Node, fset *token.FileSet) bool
}
