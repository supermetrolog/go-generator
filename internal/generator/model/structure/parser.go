package structure

import (
	"errors"
	"go/ast"
	"go/token"
	"log"
)

type parser struct {
	structure structure
}

func New() *parser {
	return &parser{}
}

func (p *parser) Run(node *ast.File, currentNode ast.Node, fset *token.FileSet) bool {
	gen, ok := currentNode.(*ast.GenDecl)

	if !ok {
		return false
	}

	for _, spec := range gen.Specs {

		switch typeSpec := spec.(type) {
		case *ast.TypeSpec:
			switch structType := typeSpec.Type.(type) {
			case *ast.StructType:
				p.StructHandler(typeSpec, structType)
			}
		}

	}
	return true
}

func (p *parser) StructHandler(typeSpec *ast.TypeSpec, structure *ast.StructType) {
	log.Println("STRUCTURE: ", typeSpec.Name.Name)

	p.structure.name = typeSpec.Name.Name

	for _, field := range structure.Fields.List {
		newFields := p.FieldHandler(field)
		err := p.structure.AddField(newFields...)
		if err != nil {
			log.Println(err)
		}
	}

	log.Println(p)
}

func (p *parser) FieldHandler(fieldAST *ast.Field) []*field {

	comment := p.GetFieldComment(fieldAST)
	fieldType, err := p.GetFieldType(fieldAST)
	if err != nil {
		log.Println(err)
		return nil
	}

	var newFields []*field
	for _, name := range fieldAST.Names {
		newField, err := NewStructureField(name.Name, fieldType, "", comment)
		if err != nil {
			log.Println(err)
			continue
		}
		newFields = append(newFields, newField)
	}
	return newFields

}

func (p *parser) GetFieldType(fieldAST *ast.Field) (string, error) {
	var fieldType string
	i, okIndent := fieldAST.Type.(*ast.Ident)
	a, okArray := fieldAST.Type.(*ast.ArrayType)
	if !okIndent && !okArray {
		return "", errors.New("Unknown type")
	}

	if okIndent {
		fieldType = i.Name
	} else if okArray {
		aIdent, ok := a.Elt.(*ast.Ident)
		if ok {
			fieldType = "[]" + aIdent.Name
		}
	}
	return fieldType, nil
}

func (p *parser) GetFieldComment(fieldAST *ast.Field) string {
	var comment string
	if fieldAST.Comment != nil {
		for _, c := range fieldAST.Comment.List {
			comment += c.Text
		}
	}

	return comment
}
