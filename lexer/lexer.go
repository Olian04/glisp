package lexer

import (
	"bufio"
	"io"
)

type Lexer struct {
	reader *bufio.Reader
	line   int
	column int
	tokens []Token
}

func NewLexer(reader io.Reader) *Lexer {
	return &Lexer{
		reader: bufio.NewReader(reader),
		line:   0,
		column: 0,
		tokens: []Token{},
	}
}

func (l *Lexer) addToken(tok Token) {
	// Update token position
	tok.Line = l.line
	tok.Column = l.column

	// Add token to token list
	l.tokens = append(l.tokens, tok)

	// Update column position
	if tok.IsIdentifier() {
		l.column += len(tok.Identifier)
	} else {
		l.column += 1
	}

	// Update line position when applicable
	if tok.IsNewline() {
		l.line++
		l.column = 0
	}
}

func (l *Lexer) Lex() ([]Token, error) {
	identifierBuf := []rune{}
	commitIdentifier := func() {
		if len(identifierBuf) > 0 {
			l.addToken(Token{
				Type:       TokenTypeIdentifier,
				Identifier: string(identifierBuf[:]),
			})
			identifierBuf = []rune{}
		}
	}

	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				commitIdentifier()
				break
			}
			return nil, err
		}

		if tok := readReserved(r); tok != nil {
			commitIdentifier()
			l.addToken(*tok)
		} else {
			identifierBuf = append(identifierBuf, r)
		}
	}

	return l.tokens, nil
}
