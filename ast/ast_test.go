package ast

import (
	"testing"

	"github.com/MohitPanchariya/monkey/token"
)

func TestString(t *testing.T) {
	expectedString := "let myVar = anotherVar;"
	program := Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}
	if program.String() != expectedString {
		t.Errorf("program.String() wrong.\n Expected = %q\n Got = %q\n", expectedString, program.String())
	}
}
