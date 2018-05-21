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
	EndFragmentDefinition  = "EndFragmentDefinition"
	EndFragmentSpread      = "EndFragmentSpread"
	EndName                = "EndName"
	EndObjectField         = "EndObjectField"
	EndOperationDefinition = "EndOperationDefinition"
	EndSelectionSet        = "EndSelectionSet"
	Field                  = "Field"
	FragmentDefinition     = "FragmentDefinition"
	FragmentSpread         = "FragmentSpread"
	Name                   = "Name"
	ObjectField            = "ObjectField"
	OperationDefinition    = "OperationDefinition"
	SelectionSet           = "SelectionSet"
	StringValue            = "StringValue"
)

var (
	passthroughHandler = func(v Visitor, n Node) bool {
		return true
	}

	defaultHandlers = Handlers{
		Argument:               passthroughHandler,
		Directive:              passthroughHandler,
		Document:               passthroughHandler,
		EndArgument:            passthroughHandler,
		EndDirective:           passthroughHandler,
		EndDocument:            passthroughHandler,
		EndFragmentDefinition:  passthroughHandler,
		EndFragmentSpread:      passthroughHandler,
		EndName:                passthroughHandler,
		EndObjectField:         passthroughHandler,
		EndOperationDefinition: passthroughHandler,
		EndSelectionSet:        passthroughHandler,
		Field:                  passthroughHandler,
		Name:                   passthroughHandler,
		ObjectField:            passthroughHandler,
		OperationDefinition:    passthroughHandler,
		SelectionSet:           passthroughHandler,
		StringValue:            passthroughHandler,
	}
)

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
