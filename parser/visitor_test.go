package parser_test

import (
	"reflect"
	"testing"

	"github.com/emwalker/graphqlparser/parser"
)

type visitor struct {
	labels        *[]string
	parserVisitor *parser.ParserVisitor
}

func newVisitor(handlers parser.Handlers) *visitor {
	visitor := &visitor{
		labels: &[]string{},
	}
	visitor.parserVisitor = parser.NewVisitor(visitor, handlers)
	return visitor
}

func (t *visitor) Label(label string) {
	*t.labels = append(*t.labels, label)
}

func (t *visitor) Visit(node *parser.AstNode) {
	node.Accept(t.parserVisitor)
}

func makeLabel(prefix string) parser.Handler {
	return func(v parser.Visitor, node parser.Node) bool {
		v.(*visitor).Label(prefix + node.TypeName())
		return true
	}
}

func TestVisitor(t *testing.T) {
	ast, err := parser.Parse(`query QueryRoot { organization(ext: "1") { name } }`)
	if err != nil {
		t.Fail()
	}

	visitor := newVisitor(parser.Handlers{
		"Document":               makeLabel(""),
		"EndDocument":            makeLabel("End"),
		"EndField":               makeLabel("End"),
		"EndName":                makeLabel("End"),
		"EndOperationDefinition": makeLabel("End"),
		"Field":                  makeLabel(""),
		"Name":                   makeLabel(""),
	})

	visitor.Visit(ast)
	defer ast.Release()

	expected := []string{
		"Document",
		"Name",
		"EndName",
		"Field",
		"Name",
		"EndName",
		"Name",
		"EndName",
		"Field",
		"Name",
		"EndName",
		"EndField",
		"EndField",
		"EndOperationDefinition",
		"EndDocument",
	}

	if !reflect.DeepEqual(*visitor.labels, expected) {
		t.Errorf("visitor test failed: %#v", *visitor.labels)
	}
}

func TestFragmentSpread(t *testing.T) {
	ast, err := parser.Parse(`
		query QueryRoot {
			organization(ext: "1") {
				...Links_organization
			}
		}

		fragment Links_organization on Organization {
			name
		}
	`)

	if err != nil {
		t.Fail()
	}

	visitor := newVisitor(parser.Handlers{
		"EndFragmentDefinition": makeLabel("End"),
		"EndFragmentSpread":     makeLabel("End"),
		"FragmentDefinition":    makeLabel(""),
		"FragmentSpread":        makeLabel(""),
	})

	visitor.Visit(ast)
	defer ast.Release()

	expected := []string{
		"FragmentSpread",
		"EndFragmentSpread",
		"FragmentDefinition",
		"EndFragmentDefinition",
	}

	if !reflect.DeepEqual(*visitor.labels, expected) {
		t.Errorf("visitor test failed: %#v", *visitor.labels)
	}
}
