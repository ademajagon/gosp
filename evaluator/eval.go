package evaluator

import (
	"errors"
	"fmt"

	"github.com/ademajagon/gosp/types"
)

type Environment struct {
	parent *Environment
	vars   map[types.Symbol]types.Expr
}

func NewDefaultEnv() *Environment {
	env := &Environment{
		vars: make(map[types.Symbol]types.Expr),
	}

	env.vars["+"] = types.Function(add)
	env.vars["-"] = types.Function(sub)

	return env
}

func (e *Environment) Lookup(sym types.Symbol) (types.Expr, bool) {
	val, ok := e.vars[sym]
	if !ok && e.parent != nil {
		return e.parent.Lookup(sym)
	}
	return val, ok
}

func Eval(expr types.Expr, env *Environment) (types.Expr, error) {
	switch e := expr.(type) {
	case types.Number:
		return e, nil
	case types.Symbol:
		if val, ok := env.Lookup(e); ok {
			return val, nil
		}
		return nil, fmt.Errorf("undefined symbol: %s", e)
	default:
		return nil, fmt.Errorf("cannot evaluate %T", expr)
	}
}

func add(args ...types.Expr) (types.Expr, error) {
	if len(args) == 0 {
		return nil, errors.New("+ requires at least one argument")
	}

	var sum float64
	for i, arg := range args {
		num, ok := arg.(types.Number)
		if !ok {
			return nil, fmt.Errorf("argument %d to + is not a number: %v", i, arg)
		}
		sum += float64(num)
	}
	return types.Number(sum), nil
}

func sub(args ...types.Expr) (types.Expr, error) {
	if len(args) == 0 {
		return nil, errors.New("- requires at least one argument")
	}

	if len(args) == 1 {
		num, ok := args[0].(types.Number)
		if !ok {
			return nil, fmt.Errorf("argument to - is not a number: %v", args[0])
		}
		return types.Number(-float64(num)), nil
	}

	var result float64
	num, ok := args[0].(types.Number)
	if !ok {
		return nil, fmt.Errorf("first argument to - is not a number: %v", args[0])
	}
	result = float64(num)

	for i, arg := range args[1:] {
		num, ok := arg.(types.Number)
		if !ok {
			return nil, fmt.Errorf("argument %d to - is not a number: %v", i+1, arg)
		}
		result -= float64(num)
	}
	return types.Number(result), nil
}
