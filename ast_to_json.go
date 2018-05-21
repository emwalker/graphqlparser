package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/emwalker/graphqlparser/parser"
)

func main() {
	var input bytes.Buffer
	io.Copy(&input, os.Stdin)
	ast, err := parser.Parse(input.String())
	defer ast.Release()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(ast.JSON()))
}
