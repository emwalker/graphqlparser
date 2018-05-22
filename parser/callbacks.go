package parser

/*
#cgo pkg-config: libgraphqlparser
#cgo CXXFLAGS: -std=c++11
#cgo darwin LDFLAGS: -Wl,-undefined -Wl,dynamic_lookup
#include "c/GraphQLAst.h"
#include "callbacks.h"
*/
import "C"
import (
	"unsafe"
)

func parserVisitor(handle unsafe.Pointer) *ParserVisitor {
	return lookupHandle(uintptr(handle)).(*ParserVisitor)
}

func shouldContinue(result bool) int {
	if result {
		return -1
	}
	return 0
}

//export endVisitArgument
func endVisitArgument(node *C.struct_GraphQLAstArgument, handle unsafe.Pointer) int {
	res := parserVisitor(handle).Dispatch(EndArgument, &ArgumentNode{
		nodeType: nodeType{Argument},
		node:     node,
	})
	return shouldContinue(res)
}

//export endVisitDocument
func endVisitDocument(node *C.struct_GraphQLAstDocument, handle unsafe.Pointer) int {
	res := parserVisitor(handle).Dispatch(EndDocument, &DocumentNode{
		nodeType: nodeType{Document},
		node:     node,
	})
	return shouldContinue(res)
}

//export endVisitField
func endVisitField(node *C.struct_GraphQLAstField, handle unsafe.Pointer) int {
	res := parserVisitor(handle).Dispatch(EndField, &FieldNode{
		nodeType: nodeType{Field},
		node:     node,
	})
	return shouldContinue(res)
}

//export endVisitFragmentDefinition
func endVisitFragmentDefinition(node *C.struct_GraphQLAstFragmentDefinition, handle unsafe.Pointer) int {
	res := parserVisitor(handle).Dispatch(EndFragmentDefinition, &FragmentDefinitionNode{
		nodeType: nodeType{FragmentDefinition},
		node:     node,
	})
	return shouldContinue(res)
}

//export endVisitFragmentSpread
func endVisitFragmentSpread(node *C.struct_GraphQLAstFragmentSpread, handle unsafe.Pointer) int {
	res := parserVisitor(handle).Dispatch(EndFragmentSpread, &FragmentSpreadNode{
		nodeType: nodeType{FragmentSpread},
		node:     node,
	})
	return shouldContinue(res)
}

//export endVisitName
func endVisitName(node *C.struct_GraphQLAstName, handle unsafe.Pointer) int {
	res := parserVisitor(handle).Dispatch(EndName, &NameNode{
		nodeType: nodeType{Name},
		node:     node,
	})
	return shouldContinue(res)
}

//export endVisitObjectField
func endVisitObjectField(node *C.struct_GraphQLAstObjectField, handle unsafe.Pointer) int {
	res := parserVisitor(handle).Dispatch(EndObjectField, &ObjectFieldNode{
		nodeType: nodeType{ObjectField},
		node:     node,
	})
	return shouldContinue(res)
}

//export endVisitOperationDefinition
func endVisitOperationDefinition(node *C.struct_GraphQLAstOperationDefinition, handle unsafe.Pointer) int {
	res := parserVisitor(handle).Dispatch(EndOperationDefinition, &OperationDefinitionNode{
		nodeType: nodeType{OperationDefinition},
		node:     node,
	})
	return shouldContinue(res)
}

//export endVisitSelectionSet
func endVisitSelectionSet(node *C.struct_GraphQLAstSelectionSet, handle unsafe.Pointer) int {
	res := parserVisitor(handle).Dispatch(EndSelectionSet, &SelectionSetNode{
		nodeType: nodeType{SelectionSet},
		node:     node,
	})
	return shouldContinue(res)
}

//export visitArgument
func visitArgument(node *C.struct_GraphQLAstArgument, handle unsafe.Pointer) int {
	res := parserVisitor(handle).Dispatch(Argument, &ArgumentNode{
		nodeType: nodeType{Argument},
		node:     node,
	})
	return shouldContinue(res)
}

//export visitDocument
func visitDocument(node *C.struct_GraphQLAstDocument, handle unsafe.Pointer) int {
	res := parserVisitor(handle).Dispatch(Document, &DocumentNode{
		nodeType: nodeType{Document},
		node:     node,
	})
	return shouldContinue(res)
}

//export visitField
func visitField(node *C.struct_GraphQLAstField, handle unsafe.Pointer) int {
	res := parserVisitor(handle).Dispatch(Field, &FieldNode{
		nodeType: nodeType{Field},
		node:     node,
	})
	return shouldContinue(res)
}

//export visitFragmentDefinition
func visitFragmentDefinition(node *C.struct_GraphQLAstFragmentDefinition, handle unsafe.Pointer) int {
	res := parserVisitor(handle).Dispatch(FragmentDefinition, &FragmentDefinitionNode{
		nodeType: nodeType{FragmentDefinition},
		node:     node,
	})
	return shouldContinue(res)
}

//export visitFragmentSpread
func visitFragmentSpread(node *C.struct_GraphQLAstFragmentSpread, handle unsafe.Pointer) int {
	res := parserVisitor(handle).Dispatch(FragmentSpread, &FragmentSpreadNode{
		nodeType: nodeType{FragmentSpread},
		node:     node,
	})
	return shouldContinue(res)
}

//export visitName
func visitName(node *C.struct_GraphQLAstName, handle unsafe.Pointer) int {
	res := parserVisitor(handle).Dispatch(Name, &NameNode{
		nodeType: nodeType{Name},
		node:     node,
	})
	return shouldContinue(res)
}

//export visitObjectField
func visitObjectField(node *C.struct_GraphQLAstObjectField, handle unsafe.Pointer) int {
	res := parserVisitor(handle).Dispatch(ObjectField, &ObjectFieldNode{
		nodeType: nodeType{ObjectField},
		node:     node,
	})
	return shouldContinue(res)
}

//export visitOperationDefinition
func visitOperationDefinition(node *C.struct_GraphQLAstOperationDefinition, handle unsafe.Pointer) int {
	res := parserVisitor(handle).Dispatch(OperationDefinition, &OperationDefinitionNode{
		nodeType: nodeType{OperationDefinition},
		node:     node,
	})
	return shouldContinue(res)
}

//export visitSelectionSet
func visitSelectionSet(node *C.struct_GraphQLAstSelectionSet, handle unsafe.Pointer) int {
	res := parserVisitor(handle).Dispatch(SelectionSet, &SelectionSetNode{
		nodeType: nodeType{SelectionSet},
		node:     node,
	})
	return shouldContinue(res)
}

//export visitStringValue
func visitStringValue(node *C.struct_GraphQLAstStringValue, handle unsafe.Pointer) int {
	res := parserVisitor(handle).Dispatch(StringValue, &StringValueNode{
		nodeType: nodeType{StringValue},
		node:     node,
	})
	return shouldContinue(res)
}
