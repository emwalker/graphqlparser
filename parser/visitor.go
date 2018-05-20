package parser

import (
	"log"
)

type Handler func(Visitor, Node) bool

type Visitor interface{}

type Handlers map[string]Handler

type ParserVisitor struct {
	visitor  Visitor
	handlers *Handlers
}

const (
	Argument               = "Argument"
	Directive              = "Directive"
	Document               = "Document"
	EndArgument            = "EndArgument"
	EndDirective           = "EndDirective"
	EndDocument            = "EndDocument"
	EndName                = "EndName"
	EndObjectField         = "EndObjectField"
	EndOperationDefinition = "EndOperationDefinition"
	EndSelectionSet        = "EndSelectionSet"
	Field                  = "Field"
	Name                   = "Name"
	ObjectField            = "ObjectField"
	OperationDefinition    = "OperationDefinition"
	SelectionSet           = "SelectionSet"
	StringValue            = "StringValue"
)

var defaultHandlers = Handlers{
	Argument: func(v Visitor, n Node) bool {
		return true
	},

	Directive: func(v Visitor, n Node) bool {
		return true
	},

	Document: func(v Visitor, n Node) bool {
		return true
	},

	EndDirective: func(v Visitor, n Node) bool {
		return true
	},

	EndDocument: func(v Visitor, n Node) bool {
		return true
	},

	EndName: func(v Visitor, n Node) bool {
		return true
	},

	EndObjectField: func(v Visitor, n Node) bool {
		return true
	},

	EndOperationDefinition: func(v Visitor, n Node) bool {
		return true
	},

	EndSelectionSet: func(v Visitor, n Node) bool {
		return true
	},

	Field: func(v Visitor, n Node) bool {
		return true
	},

	Name: func(v Visitor, n Node) bool {
		return true
	},

	ObjectField: func(v Visitor, n Node) bool {
		return true
	},

	OperationDefinition: func(v Visitor, n Node) bool {
		return true
	},

	SelectionSet: func(v Visitor, n Node) bool {
		return true
	},

	StringValue: func(v Visitor, n Node) bool {
		return true
	},
}

func (v *ParserVisitor) Dispatch(eventName string, node Node) bool {
	accept, ok := (*v.handlers)[eventName]
	if !ok {
		log.Fatal("unknown event:", eventName)
	}
	return accept(v.visitor, node)
}

func NewVisitor(v Visitor, handlers *Handlers) *ParserVisitor {
	merged := make(Handlers)

	for eventName, handler := range *handlers {
		merged[eventName] = handler
	}

	for eventName, defaultHandler := range defaultHandlers {
		if _, ok := merged[eventName]; !ok {
			merged[eventName] = defaultHandler
		}
	}

	return &ParserVisitor{v, &merged}
}
