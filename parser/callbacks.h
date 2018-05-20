#include "c/GraphQLAstForEachConcreteType.h"

#define VISITOR_FUNCTION_SIGNATURES(type, snake_type) \
    struct GraphQLAst##type; \
    int visit##type##_cgo(struct GraphQLAst##type *node, void *ptr); \
    int endVisit##type##_cgo(struct GraphQLAst##type *node, void *ptr);

FOR_EACH_CONCRETE_TYPE(VISITOR_FUNCTION_SIGNATURES)
