#include "c/GraphQLAstForEachConcreteType.h"
#include "callbacks.h"

#define VISITOR_FUNCTIONS(type, snake_type) \
    int visit##type(struct GraphQLAst##type *node, void *ptr); \
    int visit##type##_cgo(struct GraphQLAst##type *node, void *ptr) { \
      return visit##type(node, ptr); \
    } \
    int endVisit##type(struct GraphQLAst##type *node, void *ptr); \
    int endVisit##type##_cgo(struct GraphQLAst##type *node, void *ptr) { \
      return endVisit##type(node, ptr); \
    }

FOR_EACH_CONCRETE_TYPE(VISITOR_FUNCTIONS)
