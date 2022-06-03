package generator

import (
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"log"
	"os"
	"super-gen/internal/generator/model/structure"
	"super-gen/internal/walker"
)

type generator struct {
	WritePath string
}

func New(writePath string) *generator {
	return &generator{
		WritePath: writePath,
	}
}

func (gen *generator) Run() {

	fileName := "model.go.super"

	filePath := fmt.Sprintf("test-go-files/%s", fileName)
	walker, err := walker.NewWalker(filePath)
	if err != nil {
		log.Fatal(err)
	}
	// funcGenerator := functype.New()
	// genTypeGenerator := gentype.New()
	walker.RegisterHandler(structure.New())

	fset, node := walker.Run()
	gen.Save(fset, node, fileName)
}

func (gen *generator) Save(fset *token.FileSet, node *ast.File, fileName string) error {
	f, err := os.Create(fmt.Sprintf("%s/gen-%s", gen.WritePath, fileName))
	if err != nil {
		return err
	}
	defer f.Close()
	if err := printer.Fprint(f, fset, node); err != nil {
		return err
	}

	return nil
}
