package main

import (
	"fmt"
	"log"

	"github.com/emwalker/graphqlparser/parser"
)

func main() {
	ast, err := parser.Parse(`
		query {
			organization(externalId: "123") {
				links {
					edges {
						node {
							name
						}
					}
				}
			}
		}
	`)
	defer ast.Release()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(ast.JSON()))
}
