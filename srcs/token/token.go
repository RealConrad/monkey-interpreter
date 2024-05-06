package token

type TokenType string

type Token struct {
	Type		TokenType
	Literal		string
}

var keywords = map[string]TokenType {
	"fn": FUNCTION,
	"let": LET,
	"const": CONST,
	"true": TRUE,
	"false": FALSE,
	"if": IF,
	"else": ELSE,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	tok, ok := keywords[ident]
	if ok {
		return tok
	}
	return IDENT
}

const (
	ILLEGAL 	= "ILLEGAL"
	EOF			= "EOF"

	// identifiers/literals
	IDENT		= "IDENT"
	INT			= "INT"

	// Operators
	ASSIGN		= "="
	PLUS		= "+"
	MINUS		= "-"
	BANG		= "!"
	ASTERISK	= "*"
	SLASH		= "/"
	GT			= ">"
	LT			= "<"

	// Delimiters
	COMMA		= ","
	SEMICOLON	= ";"
	LPAREN		= "("
	RPAREN		= ")"
	LBRACE		= "{"
	RBRACE		= "}"

	// Keywords
	FUNCTION	= "FUNCTION"
	LET			= "LET"
	CONST		= "CONST"
	TRUE		= "TRUE"
	FALSE		= "FALSE"
	IF			= "IF"
	ELSE		= "ELSE"
	RETURN		= "RETURN"

	EQ			= "=="
	NOT_EQ		= "!="
)
