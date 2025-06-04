package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ademajagon/gosp/evaluator"
	"github.com/ademajagon/gosp/parser"
	"github.com/ademajagon/gosp/scanner"
)

func main() {

	if len(os.Args) > 1 {
		code := strings.Join(os.Args[1:], " ")
		execute(code)
		return
	}

	startREPL()
}

func execute(code string) {
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
		fmt.Printf("Eval error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Result: %v\n", result)
}

func startREPL() {
	inputScanner := bufio.NewScanner(os.Stdin)
	env := evaluator.NewDefaultEnv()

	fmt.Println("Welcome to GoSP interpreter!")
	fmt.Println("Enter expressions to evaluate, or type 'exit' to quit.")

	for {
		fmt.Println("gosp> ")
		if !inputScanner.Scan() {
			break
		}

		input := strings.TrimSpace(inputScanner.Text())
		if input == "" {
			continue
		}
		if strings.EqualFold(input, "exit") {
			break
		}

		tokens := scanner.Scan(input)
		p := parser.New(tokens)
		ast, err := p.Parse()
		if err != nil {
			fmt.Printf("Parse error: %v\n", err)
			continue
		}

		result, err := evaluator.Eval(ast, env)
		if err != nil {
			fmt.Printf("Eval error: %v\n", err)
			continue
		}

		fmt.Printf("=> %v\n", result)
	}

	if err := inputScanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
	}
}
