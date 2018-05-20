package parser

/*
#cgo pkg-config: libgraphqlparser
#cgo CXXFLAGS: -std=c++11
#cgo darwin LDFLAGS: -Wl,-undefined -Wl,dynamic_lookup
#include "c/GraphQLAst.h"
#include "c/GraphQLAstNode.h"
#include "c/GraphQLParser.h"
#include <stdlib.h>
#include "callbacks.h"
*/
import "C"
import (
	"errors"
	"unsafe"
)

func Parse(query string) (*AstNode, error) {
	cQuery := C.CString(query)
	cError := (*C.char)(nil)
	node := C.graphql_parse_string(cQuery, &cError)
	C.free(unsafe.Pointer(cQuery))

	if node == nil {
		err := errors.New(C.GoString(cError))
		C.graphql_error_free(cError)
		return nil, err
	}

	return &AstNode{
		nodeType{"AstNode"},
		node,
	}, nil
}
