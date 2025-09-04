package parser

import (
	"github.com/Olian04/glisp/lexer"
)

// Used for multi threaded parsing
func getNextExpression(tokens []lexer.Token) (expressionTokens []lexer.Token, remainingTokens []lexer.Token, err *ParseError) {
	if !tokens[0].IsOpenParen() {
		return nil, nil, &ParseError{
			Message: "expected open parenthesis",
			Line:    tokens[0].Line,
			Column:  tokens[0].Column,
		}
	}

	openExpressions := 1

	for i, token := range tokens {
		if token.IsOpenParen() {
			openExpressions++
		}
		if token.IsCloseParen() {
			openExpressions--
		}
		if openExpressions == 0 {
			return tokens[:i], tokens[i+1:], nil
		}
	}

	return nil, nil, &ParseError{
		Message: "no matching closing parenthesis found",
		Line:    tokens[0].Line,
		Column:  tokens[0].Column,
	}
}
