package parser

import (
	"errors"
	"strconv"

	"github.com/ademajagon/gosp/types"
)

type Parser struct {
	tokens []types.Token
	pos    int
}

func New(tokens []types.Token) *Parser {
	return &Parser{tokens: tokens}
}

func (p *Parser) parseToken() (types.Expr, error) {
	if p.pos >= len(p.tokens) {
		return nil, errors.New("unexpected end of input")
	}

	tok := p.tokens[p.pos]
	p.pos++

	switch tok {
	case "(":
		var list types.List
		for {
			if p.pos >= len(p.tokens) {
				return nil, errors.New("missing ')'")
			}
			if p.tokens[p.pos] == ")" {
				p.pos++
				break
			}

			elem, err := p.parseToken()
			if err != nil {
				return nil, err
			}
			list = append(list, elem)
		}
		return list, nil
	case ")":
		return nil, errors.New("unexpected ')'")
	default:
		if i, err := strconv.Atoi(string(tok)); err == nil {
			return types.Number(i), nil
		}
		return types.Symbol(tok), nil
	}
}

func (p *Parser) Parse() (types.Expr, error) {
	return p.parseToken()
}
