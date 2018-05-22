package parser

/*
#cgo pkg-config: libgraphqlparser
#cgo CXXFLAGS: -std=c++11
#cgo darwin LDFLAGS: -Wl,-undefined -Wl,dynamic_lookup
#include "c/GraphQLAst.h"
#include "c/GraphQLAstNode.h"
#include "c/GraphQLAstVisitor.h"
#include "c/GraphQLAstToJSON.h"
#include "callbacks.h"
#include <stdlib.h>
#include <string.h>
*/
import "C"
import (
	"unsafe"
)

type Node interface {
	TypeName() string
}

type nodeType struct {
	name string
}

func (n *nodeType) TypeName() string {
	return n.name
}

type AstNode struct {
	nodeType
	node *C.struct_GraphQLAstNode
}

func (n *AstNode) Release() {
	C.graphql_node_free(n.node)
}

func (n *AstNode) Accept(v *ParserVisitor) {
	handle := unsafe.Pointer(newHandle(v))
	defer deleteHandles()
	C.graphql_node_visit(n.node, callbacks(), handle)
}

func (n *AstNode) JSON() []byte {
	cStr := C.graphql_ast_to_json(n.node)
	ptr := unsafe.Pointer(cStr)
	defer C.free(ptr)
	return C.GoBytes(ptr, C.int(C.strlen(cStr)))
}

func newName(node *C.struct_GraphQLAstName) *NameNode {
	return &NameNode{
		nodeType{"NameNode"},
		node,
	}
}

type ArgumentNode struct {
	nodeType
	node *C.struct_GraphQLAstArgument
}

type DocumentNode struct {
	nodeType
	node *C.struct_GraphQLAstDocument
}

type FieldNode struct {
	nodeType
	node *C.struct_GraphQLAstField
}

func (n *FieldNode) Name() *NameNode {
	return newName(C.GraphQLAstField_get_name(n.node))
}

type FragmentDefinitionNode struct {
	nodeType
	node *C.struct_GraphQLAstFragmentDefinition
}

func (n *FragmentDefinitionNode) Name() *NameNode {
	return newName(C.GraphQLAstFragmentDefinition_get_name(n.node))
}

type FragmentSpreadNode struct {
	nodeType
	node *C.struct_GraphQLAstFragmentSpread
}

func (n *FragmentSpreadNode) Name() *NameNode {
	return newName(C.GraphQLAstFragmentSpread_get_name(n.node))
}

type NameNode struct {
	nodeType
	node *C.struct_GraphQLAstName
}

func (n *NameNode) Value() string {
	return C.GoString(C.GraphQLAstName_get_value(n.node))
}

type ObjectFieldNode struct {
	nodeType
	node *C.struct_GraphQLAstObjectField
}

type OperationDefinitionNode struct {
	nodeType
	node *C.struct_GraphQLAstOperationDefinition
}

type SelectionSetNode struct {
	nodeType
	node *C.struct_GraphQLAstSelectionSet
}

type StringValueNode struct {
	nodeType
	node *C.struct_GraphQLAstStringValue
}

func (n *StringValueNode) Value() string {
	return C.GoString(C.GraphQLAstStringValue_get_value(n.node))
}

// MACRO(VariableDefinition, variable_definition) \
// MACRO(InlineFragment, inline_fragment) \
// MACRO(Variable, variable) \
// MACRO(IntValue, int_value) \
// MACRO(FloatValue, float_value) \
// MACRO(BooleanValue, boolean_value) \
// MACRO(NullValue, null_value) \
// MACRO(EnumValue, enum_value) \
// MACRO(ListValue, list_value) \
// MACRO(ObjectValue, object_value) \
// MACRO(NamedType, named_type) \
// MACRO(ListType, list_type) \
// MACRO(NonNullType, non_null_type) \
// MACRO(SchemaDefinition, schema_definition) \
// MACRO(OperationTypeDefinition, operation_type_definition) \
// MACRO(ScalarTypeDefinition, scalar_type_definition) \
// MACRO(ObjectTypeDefinition, object_type_definition) \
// MACRO(FieldDefinition, field_definition) \
// MACRO(InputValueDefinition, input_value_definition) \
// MACRO(InterfaceTypeDefinition, interface_type_definition) \
// MACRO(UnionTypeDefinition, union_type_definition) \
// MACRO(EnumTypeDefinition, enum_type_definition) \
// MACRO(EnumValueDefinition, enum_value_definition) \
// MACRO(InputObjectTypeDefinition, input_object_type_definition) \
// MACRO(TypeExtensionDefinition, type_extension_definition) \
// MACRO(DirectiveDefinition, directive_definition)

func callbacks() *C.struct_GraphQLAstVisitorCallbacks {
	return &C.struct_GraphQLAstVisitorCallbacks{
		end_visit_argument:             (C.end_visit_argument_func)(C.endVisitArgument_cgo),
		end_visit_directive:            (C.end_visit_directive_func)(C.endVisitDirective_cgo),
		end_visit_document:             (C.end_visit_document_func)(C.endVisitDocument_cgo),
		end_visit_field:                (C.end_visit_field_func)(C.endVisitField_cgo),
		end_visit_fragment_definition:  (C.end_visit_fragment_definition_func)(C.endVisitFragmentDefinition_cgo),
		end_visit_fragment_spread:      (C.end_visit_fragment_spread_func)(C.endVisitFragmentSpread_cgo),
		end_visit_name:                 (C.end_visit_name_func)(C.endVisitName_cgo),
		end_visit_operation_definition: (C.end_visit_operation_definition_func)(C.endVisitOperationDefinition_cgo),
		end_visit_selection_set:        (C.end_visit_selection_set_func)(C.endVisitSelectionSet_cgo),
		visit_argument:                 (C.visit_argument_func)(C.visitArgument_cgo),
		visit_document:                 (C.visit_document_func)(C.visitDocument_cgo),
		visit_field:                    (C.visit_field_func)(C.visitField_cgo),
		visit_fragment_definition:      (C.visit_fragment_definition_func)(C.visitFragmentDefinition_cgo),
		visit_fragment_spread:          (C.visit_fragment_spread_func)(C.visitFragmentSpread_cgo),
		visit_name:                     (C.visit_name_func)(C.visitName_cgo),
		visit_object_field:             (C.visit_object_field_func)(C.visitObjectField_cgo),
		visit_operation_definition:     (C.visit_operation_definition_func)(C.visitOperationDefinition_cgo),
		visit_selection_set:            (C.visit_selection_set_func)(C.visitSelectionSet_cgo),
		visit_string_value:             (C.visit_string_value_func)(C.visitStringValue_cgo),
	}
}
