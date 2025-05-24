package evaluator

import (
	"errors"
	"fmt"

	"github.com/ademajagon/gosp/types"
)

type Env map[string]func(args []types.Expr) (types.Expr, error)

var DefaultEnv Env

func NewDefaultEnv() Env {

	env := Env{}

	env["+"] = func(args []types.Expr) (types.Expr, error) {
		sum := 0
		for _, arg := range args {
			val, err := Eval(arg, env)
			if err != nil {
				return nil, err
			}
			i, ok := val.(int)
			if !ok {
				return nil, fmt.Errorf("expected int, got %T", val)
			}
			sum += i
		}
		return sum, nil
	}

	env["*"] = func(args []types.Expr) (types.Expr, error) {
		prod := 1
		for _, arg := range args {
			val, err := Eval(arg, env)
			if err != nil {
				return nil, err
			}
			i, ok := val.(int)
			if !ok {
				return nil, fmt.Errorf("expected int, got %T", val)
			}
			prod *= i
		}
		return prod, nil
	}

	return env
}

func Eval(expr types.Expr, env Env) (types.Expr, error) {
	switch e := expr.(type) {
	case int:
		return e, nil
	case string:
		return e, nil
	case []types.Expr:
		if len(e) == 0 {
			return nil, errors.New("empty list")
		}
		fnSym, ok := e[0].(string)
		if !ok {
			return nil, errors.New("first element must be a function name")
		}
		fn, ok := env[fnSym]
		if !ok {
			return nil, fmt.Errorf("unknown function: %s", fnSym)
		}
		return fn(e[1:])
	default:
		return nil, fmt.Errorf("unknown expression type: %T", expr)
	}
}
