package parser

import (
	"fmt"
)

type ParseError struct {
	Message string
	Line    int
	Column  int
}

func (e ParseError) Error() string {
	return fmt.Sprintf("%s at line %d, column %d", e.Message, e.Line, e.Column)
}
