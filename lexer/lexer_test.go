package lexer

import (
	"testing"

	"github.com/milennik/go_interpreter/token"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	l := New(input)

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - wrong format. expected %q, got %q", i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - wrong literal. expected %q, got %q", i, tt.expectedLiteral, tok.Literal)
		}
	}

}
