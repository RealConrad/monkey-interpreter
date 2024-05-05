package token

type TokenType string

type Token struct {
	Type		TokenType
	Literal		string
}

var keywords = map[string]TokenType {
	"func": FUNCTION,
	"let": LET,
	"const": CONST,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
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
)
