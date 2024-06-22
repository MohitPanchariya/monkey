package ast

import "github.com/MohitPanchariya/monkey/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode() // dummy method to throw error when a Statement is used instead of an Expression
}

type Expression interface {
	Node
	expressionNode() // dummy method to throw error when an Expression is used instead of a Statement
}

// Program node represents the root node of the AST.
// Each monkey program is a series of statements
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type LetStatement struct {
	Token token.Token // Every node in the tree keeps track of the token its associated with
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

type Identifier struct {
	Token token.Token // token.IDENT
	Value string      // Identifiers can be value producing as well (e.g. let x = valueProducingIdentifier)
}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Identifier) expressionNode() {}

type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}

func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}
