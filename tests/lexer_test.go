package tests

import (
	"strings"
	"testing"

	"github.com/Olian04/glisp/lexer"
	"github.com/Olian04/glisp/tests/assert"
)

func removePosition(tokens []lexer.Token) []lexer.Token {
	out := []lexer.Token{}
	for _, token := range tokens {
		out = append(out, lexer.Token{
			Type:       token.Type,
			Identifier: token.Identifier,
		})
	}
	return out
}

func TestBasicExpression(t *testing.T) {
	lex := lexer.NewLexer(strings.NewReader("(+ 123 45 6)"))
	tokens, err := lex.Lex()
	assert.NoError(t, err)
	assert.Slice(t, tokens, []lexer.Token{
		{Type: lexer.TokenTypeOpenParen, Line: 0, Column: 0},
		{Type: lexer.TokenTypeIdentifier, Identifier: "+", Line: 0, Column: 1},
		{Type: lexer.TokenTypeSpace, Line: 0, Column: 2},
		{Type: lexer.TokenTypeIdentifier, Identifier: "123", Line: 0, Column: 3},
		{Type: lexer.TokenTypeSpace, Line: 0, Column: 6},
		{Type: lexer.TokenTypeIdentifier, Identifier: "45", Line: 0, Column: 7},
		{Type: lexer.TokenTypeSpace, Line: 0, Column: 9},
		{Type: lexer.TokenTypeIdentifier, Identifier: "6", Line: 0, Column: 10},
		{Type: lexer.TokenTypeCloseParen, Line: 0, Column: 11},
	})
}

func TestQuotedExpression(t *testing.T) {
	lex := lexer.NewLexer(strings.NewReader("'(* 123 45 6)"))
	tokens, err := lex.Lex()
	assert.NoError(t, err)
	assert.Slice(t, tokens, []lexer.Token{
		{Type: lexer.TokenTypeSingleQuote, Line: 0, Column: 0},
		{Type: lexer.TokenTypeOpenParen, Line: 0, Column: 1},
		{Type: lexer.TokenTypeIdentifier, Identifier: "*", Line: 0, Column: 2},
		{Type: lexer.TokenTypeSpace, Line: 0, Column: 3},
		{Type: lexer.TokenTypeIdentifier, Identifier: "123", Line: 0, Column: 4},
		{Type: lexer.TokenTypeSpace, Line: 0, Column: 7},
		{Type: lexer.TokenTypeIdentifier, Identifier: "45", Line: 0, Column: 8},
		{Type: lexer.TokenTypeSpace, Line: 0, Column: 10},
		{Type: lexer.TokenTypeIdentifier, Identifier: "6", Line: 0, Column: 11},
		{Type: lexer.TokenTypeCloseParen, Line: 0, Column: 12},
	})
}

func TestNewline(t *testing.T) {
	lex := lexer.NewLexer(strings.NewReader("(+ 123 45 6\n78 90)"))
	tokens, err := lex.Lex()
	assert.NoError(t, err)
	assert.Slice(t, tokens, []lexer.Token{
		{Type: lexer.TokenTypeOpenParen, Line: 0, Column: 0},
		{Type: lexer.TokenTypeIdentifier, Identifier: "+", Line: 0, Column: 1},
		{Type: lexer.TokenTypeSpace, Line: 0, Column: 2},
		{Type: lexer.TokenTypeIdentifier, Identifier: "123", Line: 0, Column: 3},
		{Type: lexer.TokenTypeSpace, Line: 0, Column: 6},
		{Type: lexer.TokenTypeIdentifier, Identifier: "45", Line: 0, Column: 7},
		{Type: lexer.TokenTypeSpace, Line: 0, Column: 9},
		{Type: lexer.TokenTypeIdentifier, Identifier: "6", Line: 0, Column: 10},
		{Type: lexer.TokenTypeNewline, Line: 0, Column: 11},
		{Type: lexer.TokenTypeIdentifier, Identifier: "78", Line: 1, Column: 0},
		{Type: lexer.TokenTypeSpace, Line: 1, Column: 2},
		{Type: lexer.TokenTypeIdentifier, Identifier: "90", Line: 1, Column: 3},
		{Type: lexer.TokenTypeCloseParen, Line: 1, Column: 5},
	})
}

func TestNestedExpressions(t *testing.T) {
	lex := lexer.NewLexer(strings.NewReader("(+ (- 4 2) (* 5 6))"))
	tokens, err := lex.Lex()
	assert.NoError(t, err)
	assert.Slice(t, removePosition(tokens), []lexer.Token{
		{Type: lexer.TokenTypeOpenParen},
		{Type: lexer.TokenTypeIdentifier, Identifier: "+"},
		{Type: lexer.TokenTypeSpace},
		{Type: lexer.TokenTypeOpenParen},
		{Type: lexer.TokenTypeIdentifier, Identifier: "-"},
		{Type: lexer.TokenTypeSpace},
		{Type: lexer.TokenTypeIdentifier, Identifier: "4"},
		{Type: lexer.TokenTypeSpace},
		{Type: lexer.TokenTypeIdentifier, Identifier: "2"},
		{Type: lexer.TokenTypeCloseParen},
		{Type: lexer.TokenTypeSpace},
		{Type: lexer.TokenTypeOpenParen},
		{Type: lexer.TokenTypeIdentifier, Identifier: "*"},
		{Type: lexer.TokenTypeSpace},
		{Type: lexer.TokenTypeIdentifier, Identifier: "5"},
		{Type: lexer.TokenTypeSpace},
		{Type: lexer.TokenTypeIdentifier, Identifier: "6"},
		{Type: lexer.TokenTypeCloseParen},
		{Type: lexer.TokenTypeCloseParen},
	})
}
