package lexer

import (
	"github.com/RealConrad/monkey-interpreter/srcs/token"
)

type Lexer struct {
	input		string // string we checking for
	position	int // current position in input string
	readPostion	int // current reading position in input string (after current char)
	ch			byte // the char we testing
}

func (lexer *Lexer) NextToken() token.Token {
	var tok token.Token

	lexer.readChar()
	switch lexer.ch {
		case '=':
			tok = newToken(token.ASSIGN, lexer.ch)
		case ';':
			tok = newToken(token.SEMICOLON, lexer.ch)
		case '(':
			tok = newToken(token.LPAREN, lexer.ch)
		case ')':
			tok = newToken(token.RPAREN, lexer.ch)
		case ',':
			tok = newToken(token.COMMA, lexer.ch)
		case '+':
			tok = newToken(token.PLUS, lexer.ch)
		case '{':
			tok = newToken(token.LBRACE, lexer.ch)
		case '}':
			tok = newToken(token.RBRACE, lexer.ch)
		case 0:
			tok.Literal = ""
			tok.Type = token.EOF
	}
	return tok
}

func New(input string) *Lexer {
	lex := &Lexer{input: input}
	return lex
}

func (lexer *Lexer) readChar() {
	if lexer.readPostion >= len(lexer.input) {
		lexer.ch = 0
	} else {
		lexer.ch = lexer.input[lexer.readPostion]
	}
	lexer.position = lexer.readPostion
	lexer.readPostion++;
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
