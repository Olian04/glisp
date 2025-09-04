package lexer

func readReserved(r rune) *Token {
	if tok, ok := spaceRunes[r]; ok {
		return &Token{
			Type: tok,
		}
	}
	if tok, ok := syntaxRunes[r]; ok {
		return &Token{
			Type: tok,
		}
	}
	return nil
}

var spaceRunes = map[rune]TokenType{
	' ':  TokenTypeSpace,
	'\t': TokenTypeTab,
	'\n': TokenTypeNewline,
	'\r': TokenTypeCarriageReturn,
}

var syntaxRunes = map[rune]TokenType{
	'\'': TokenTypeSingleQuote,
	'(':  TokenTypeOpenParen,
	')':  TokenTypeCloseParen,
}
