package ast

import "github.com/RealConrad/monkey-interpreter/srcs/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

type Identifier struct {
	Token token.Token
	Value string
}

type ReturnStatement struct {
	Token token.Token // the 'return' token
	ReturnValue Expression
}	

type LetStatement struct {
	Token token.Token
	Name *Identifier
	Value Expression
}

/* --------------------------- LET-STATEMENT FUNC --------------------------- */

func (ls *LetStatement) statementNode() {}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

/* ------------------------------- RETURN FUN ------------------------------- */

func (rs *ReturnStatement) statementNode() {}

func (rs *ReturnStatement) TokenLiteral() string { 
	return rs.Token.Literal 
}

/* ----------------------------- IDENTIFIER FUNC ---------------------------- */

func (i *Identifier) expressionNode() {}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

/* ------------------------------ PROGRAM FUNC ------------------------------ */

func (p *Program) TokenLiteral() string {
	if (len(p.Statements) > 0) {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

