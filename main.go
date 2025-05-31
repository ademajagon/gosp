package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/ademajagon/gosp/evaluator"
	"github.com/ademajagon/gosp/parser"
	"github.com/ademajagon/gosp/scanner"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: gosp <lisp-expression>")
		return
	}

	code := strings.Join(os.Args[1:], " ")
	tokens := scanner.Scan(code)

	p := parser.New(tokens)
	ast, err := p.Parse()
	if err != nil {
		fmt.Printf("Parse error: %v\n", err)
		os.Exit(1)
	}

	env := evaluator.NewDefaultEnv()
	result, err := evaluator.Eval(ast, env)
	if err != nil {
		fmt.Printf("Evaluation error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Result: %v\n", result)
}
