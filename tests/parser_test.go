package tests

import (
	"testing"

	"github.com/Olian04/glisp/lexer"
	"github.com/Olian04/glisp/parser"
	"github.com/Olian04/glisp/tests/assert"
)

func TestParser(t *testing.T) {
	p := parser.Parser{}
	expressions, errs := p.Parse([]lexer.Token{
		{Type: lexer.TokenTypeOpenParen},
		{Type: lexer.TokenTypeIdentifier, Identifier: "+"},
		{Type: lexer.TokenTypeIdentifier, Identifier: "1"},
		{Type: lexer.TokenTypeIdentifier, Identifier: "2"},
		{Type: lexer.TokenTypeCloseParen},
	})

	assert.Slice(t, errs, nil)
	assert.Slice(t, expressions, []parser.Expression{
		{Type: parser.ExpressionTypeCall, Statements: []parser.Statement{
			{Expression: &parser.Expression{
				Type: parser.ExpressionTypeCall,
				Statements: []parser.Statement{
					{Atom: &parser.Atom{Type: parser.AtomTypeIdentifier, Identifier: "+"}},
					{Atom: &parser.Atom{Type: parser.AtomTypeNumber, NumberValue: 1}},
					{Atom: &parser.Atom{Type: parser.AtomTypeNumber, NumberValue: 2}},
				},
			},
			},
		}},
	})
}
