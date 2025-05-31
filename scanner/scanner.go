package scanner

import (
	"strings"
	"unicode"

	"github.com/ademajagon/gosp/types"
)

func Scan(input string) []types.Token {
	var tokens []types.Token
	var current strings.Builder

	addToken := func() {
		if current.Len() > 0 {
			tokens = append(tokens, types.Token(current.String()))
			current.Reset()
		}
	}

	for _, r := range input {
		switch {
		case unicode.IsSpace(r):
			addToken()
		case strings.ContainsRune("()", r):
			addToken()
			tokens = append(tokens, types.Token(string(r)))
		default:
			current.WriteRune(r)
		}
	}

	addToken()

	return tokens
}
