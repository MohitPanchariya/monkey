package lexer

import (
	"testing"

	"github.com/MohitPanchariya/monkey/token"
)

func TestNextToken(t *testing.T) {
	input := "=+(){},;"

	tests := []struct {
		expectedTokenType token.TokenType
		expectedLiteral   string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedTokenType {
			t.Fatalf("tests[%d] - expectedTokenType wrong. expected = %q, got = %q", i, tt.expectedTokenType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - expectedLiteral wrong. expected = %q, got = %q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
