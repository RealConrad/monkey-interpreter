package lexer

import (
	"testing"

	"github.com/RealConrad/monkey-interpreter/srcs/token"
)

type Lexer struct {
	input		string // string we checking for
	position	int // current position in input string
	readPostion	int // current reading position in input string (after current char)
	ch			byte // the char we testing
}

func New(input string) *Lexer {
	lex := &Lexer{input: input}
	return lex
}

func TestNextToken(t *testing.T) {
	input := "=+(){};"

	tests := []struct {
		expectedType	token.TokenType
		expectedLiteral	string
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

	for i, tokenType := range tests {
		tok := l.NextToken()

		if tok.Type != tokenType.expectedType {
			t.Fatal("tests[%d] - token type is wrong. expected=%q, got=%q", i, tokenType.expectedLiteral, tok.Type)
		}
		if tok.Literal != tokenType.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tokenType.expectedLiteral, tok.Literal)
		}
	}
}
