package evaluator

import (
	"errors"
	"fmt"

	"github.com/ademajagon/gosp/global"
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
	env.vars["*"] = types.Function(mul)
	env.vars["/"] = types.Function(div)

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
	case types.List:
		if len(e) == 0 {
			return nil, errors.New("empty list")
		}
		return evalList(e, env)
	default:
		return nil, fmt.Errorf("cannot evaluate %T", expr)
	}
}

func evalList(list types.List, env *Environment) (types.Expr, error) {
	global.Log("Evaluating list:", list)

	fnExpr, err := Eval(list[0], env)
	if err != nil {
		return nil, err
	}

	fn, ok := fnExpr.(types.Function)
	if !ok {
		return nil, fmt.Errorf("first list element must be a function, got %T", fnExpr)
	}

	args := make([]types.Expr, 0, len(list)-1)
	for _, arg := range list[1:] {
		global.Log("Evaluating argument:", arg)
		evalArg, err := Eval(arg, env)
		if err != nil {
			return nil, err
		}
		global.Log("Evaluated to:", evalArg)
		args = append(args, evalArg)
	}

	global.Log("Calling function with arguments:", args)

	return fn(args...)
}
