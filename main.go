package main

import (
	"fmt"

	"github.com/ademajagon/gosp/evaluator"
	"github.com/ademajagon/gosp/parser"
	"github.com/ademajagon/gosp/scanner"
)

func main() {
	code := "(+ 1 (* 2 3))"
	tokens := scanner.Scan(code)

	p := parser.New(tokens)
	ast, err := p.Parse()
	if err != nil {
		panic(err)
	}

	env := evaluator.NewDefaultEnv()
	result, err := evaluator.Eval(ast, env)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Result: %v\n", result)
}
