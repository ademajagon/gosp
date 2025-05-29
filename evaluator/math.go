package evaluator

import (
	"errors"
	"fmt"

	"github.com/ademajagon/gosp/types"
)

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

func mul(args ...types.Expr) (types.Expr, error) {
	if len(args) == 0 {
		return nil, errors.New("* requires at least one argument")
	}

	product := 1.0
	for i, arg := range args {
		num, ok := arg.(types.Number)
		if !ok {
			return nil, fmt.Errorf("argument %d to * is not a number: %v", i, arg)
		}
		product *= float64(num)
	}
	return types.Number(product), nil
}

func div(args ...types.Expr) (types.Expr, error) {
	if len(args) == 0 {
		return nil, errors.New("/ requires at least one argument")
	}

	if len(args) == 1 {
		num, ok := args[0].(types.Number)
		if !ok {
			return nil, fmt.Errorf("argument to / is not a number: %v", args[0])
		}
		if num == 0 {
			return nil, errors.New("division by zero")
		}
		return types.Number(1 / float64(num)), nil
	}

	var result float64
	num, ok := args[0].(types.Number)
	if !ok {
		return nil, fmt.Errorf("first argument to / is not a number: %v", args[0])
	}

	result = float64(num)

	for i, arg := range args[1:] {
		num, ok := arg.(types.Number)
		if !ok {
			return nil, fmt.Errorf("argument %d to / is not a number: %v", i+1, arg)
		}

		if num == 0 {
			return nil, errors.New("division by zero")
		}
		result /= float64(num)
	}
	return types.Number(result), nil
}
