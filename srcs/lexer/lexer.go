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

	lexer.skipWhitespace()
	switch lexer.ch {
		case '=':
			if lexer.peekChar() == '=' {
				ch := lexer.ch
				lexer.readChar()
				tok = token.Token{Type: token.EQ, Literal: string(ch) + string(lexer.ch)}
			} else {
				tok = newToken(token.ASSIGN, lexer.ch)
			}
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
		case '-':
			tok = newToken(token.MINUS, lexer.ch)
		case '!':
			if lexer.peekChar() == '=' {
				ch := lexer.ch
				lexer.readChar()
				tok = token.Token{Type: token.NOT_EQ, Literal: string(ch) + string(lexer.ch)}
			} else {
				tok = newToken(token.BANG, lexer.ch)
			}
		case '/':
			tok = newToken(token.SLASH, lexer.ch)
		case '*':
			tok = newToken(token.ASTERISK, lexer.ch)
		case '<':
			tok = newToken(token.LT, lexer.ch)
		case '>':
			tok = newToken(token.GT, lexer.ch)
		case '{':
			tok = newToken(token.LBRACE, lexer.ch)
		case '}':
			tok = newToken(token.RBRACE, lexer.ch)
		case 0:
			tok.Literal = ""
			tok.Type = token.EOF
		default:
			if isLetter(lexer.ch) {
				tok.Literal = lexer.readIdentifier()
					tok.Type = token.LookupIdent(tok.Literal)
					return tok
				} else if isDigit(lexer.ch) {
					tok.Type = token.INT
					tok.Literal = lexer.readNumber()
					return tok
				} else {
					tok = newToken(token.ILLEGAL, lexer.ch)
			}
	}
	lexer.readChar()
	return tok
}

func New(input string) *Lexer {
	lexer := &Lexer{input: input}
	lexer.readChar()
	return lexer
}

func (lexer *Lexer) peekChar() byte {
	if lexer.readPostion >= len(lexer.input) {
		return 0
	} else {
		return lexer.input[lexer.readPostion]
	}
}

func (lexer *Lexer) skipWhitespace() {
	for lexer.ch == ' ' || lexer.ch == '\t' || lexer.ch == '\n' || lexer.ch == '\r' {
		lexer.readChar()
	}
}	

func (lexer *Lexer) readIdentifier() string {
	position := lexer.position
	for isLetter(lexer.ch) {
		lexer.readChar()
	}
	return lexer.input[position:lexer.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' || ch == '!' || ch == '?'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (lexer *Lexer) readNumber() string {
	position := lexer.position
	for isDigit(lexer.ch) {
		lexer.readChar()
	}
	return lexer.input[position:lexer.position]
}

func (lexer *Lexer) readChar() {
	if lexer.readPostion >= len(lexer.input) {
		lexer.ch = 0
	} else {
		lexer.ch = lexer.input[lexer.readPostion]
	}
	lexer.position = lexer.readPostion
	lexer.readPostion += 1
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
