package lexer

type TokenType string

const (
	TokenTypeIdentifier     TokenType = "identifier"
	TokenTypeSpace          TokenType = "space"
	TokenTypeTab            TokenType = "tab"
	TokenTypeNewline        TokenType = "newline"
	TokenTypeCarriageReturn TokenType = "carriage_return"
	TokenTypeOpenParen      TokenType = "paren_open"
	TokenTypeCloseParen     TokenType = "paren_close"
	TokenTypeSingleQuote    TokenType = "quote_single"
)

type Token struct {
	Line       int
	Column     int
	Type       TokenType
	Identifier string
}

func (t Token) IsNewline() bool {
	return t.Type == TokenTypeNewline
}

func (t Token) IsWhiteSpace() bool {
	return t.Type == TokenTypeSpace || t.Type == TokenTypeTab || t.Type == TokenTypeCarriageReturn || t.Type == TokenTypeNewline
}

func (t Token) IsIdentifier() bool {
	return t.Type == TokenTypeIdentifier
}

func (t Token) IsOpenParen() bool {
	return t.Type == TokenTypeOpenParen
}

func (t Token) IsCloseParen() bool {
	return t.Type == TokenTypeCloseParen
}

func (t Token) IsQuote() bool {
	return t.Type == TokenTypeSingleQuote
}
