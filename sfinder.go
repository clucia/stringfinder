package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <filename>")
		return
	}

	var accum = make(map[string]int)

	for _, filename := range os.Args[1:] {
		strings := parse(filename)
		for k, v := range strings {
			_, ok := accum[k]
			if ok {
				accum[k] += v
			} else {
				accum[k] = v
			}
		}
	}

	// Print the list of string constants
	for k, v := range accum {
		idx := strings.IndexAny(k, "\t\n=+,")
		if idx >= 0 {
			continue
		}

		fmt.Println(v, "\t", k)
	}
}

func parseYes(node ast.Node, strings *map[string]int) {
	ast.Inspect(node, func(n ast.Node) bool {
		switch trav := n.(type) {
		case *ast.BasicLit:
			if trav.Kind == token.STRING {
				(*strings)[trav.Value]++
			}
		}
		return true
	})
}

func parse(filename string) (strings map[string]int) {
	// Parse the source file
	strings = make(map[string]int)
	fset := token.NewFileSet()

	node, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	ast.Inspect(node, func(n ast.Node) bool {
		switch x := n.(type) {
		case nil:
			return false
		case *ast.ImportSpec:
			return false
		case *ast.Comment:
			return false
		case *ast.CommentGroup:
			return false
		case *ast.FuncType:
			return false
		default:
			_ = x
			parseYes(n, &strings)
		}
		return true
	})
	return strings
}
