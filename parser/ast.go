package parser

type Expression struct {
	Type       ExpressionType
	Statements []Statement
}

type Statement struct {
	Atom       *Atom
	Expression *Expression
}

type Atom struct {
	Type         AtomType
	Identifier   string
	StringValue  string
	NumberValue  float64
	BooleanValue bool
}

type AtomType string

const (
	AtomTypeIdentifier AtomType = "identifier"
	AtomTypeString     AtomType = "string"
	AtomTypeNumber     AtomType = "number"
	AtomTypeBoolean    AtomType = "boolean"
)

type ExpressionType string

const (
	ExpressionTypeCall  ExpressionType = "call"
	ExpressionTypeQuote ExpressionType = "quote"
)
