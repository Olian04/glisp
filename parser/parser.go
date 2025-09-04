package parser

import (
	"sync"

	"github.com/Olian04/glisp/lexer"
)

type Parser struct{}

func (p *Parser) Parse(tokens []lexer.Token) ([]Expression, []*ParseError) {
	out := make(chan Expression)
	errOut := make(chan *ParseError)

	wg := sync.WaitGroup{}

	for len(tokens) > 0 {
		expressionTokens, remainingTokens, err := getNextExpression(tokens)
		if err != nil {
			return nil, []*ParseError{err}
		}
		wg.Go(func() {
			expression, err := parseExpression(expressionTokens)
			if err != nil {
				errOut <- err
				return
			}
			out <- *expression
		})
		tokens = remainingTokens
	}

	wg.Wait()
	close(out)
	close(errOut)

	errs := []*ParseError{}
	for err := range errOut {
		errs = append(errs, err)
	}

	expressions := []Expression{}
	for expression := range out {
		expressions = append(expressions, expression)
	}

	return expressions, errs
}

func parseExpression(tokens []lexer.Token) (*Expression, *ParseError) {
	if len(tokens) == 0 {
		return nil, &ParseError{Message: "no tokens", Line: 0, Column: 0}
	}

	if !tokens[0].IsOpenParen() {
		return nil, &ParseError{Message: "expected open parenthesis", Line: tokens[0].Line, Column: tokens[0].Column}
	}

	if !tokens[len(tokens)-1].IsCloseParen() {
		return nil, &ParseError{Message: "expected close parenthesis", Line: tokens[len(tokens)-1].Line, Column: tokens[len(tokens)-1].Column}
	}

	return &Expression{
		Type:       ExpressionTypeCall,
		Statements: []Statement{},
	}, nil
}
