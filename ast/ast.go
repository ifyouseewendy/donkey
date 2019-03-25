package ast

import (
	"bytes"
	"donkey/token"
)

// Node is an interal presentation of a node on AST. Every node in our AST has
// to implement the Node interface, meaning it has to provide a TokenLiteral()
// method that returns the literal value of the token it’s associated with
// TokenLiteral() will be used only for debugging and testing. The AST we are
// going to construct consists solely of Nodes that are connected to each
// other - it’s a tree after all. Some of these nodes implement the Statemen
// and some the Expression interface.
type Node interface {
	TokenLiteral() string
	String() string
}

// Statement is the meta unit of a program. Every node implements it has to provide a
// statementNode() method returns the literal value of the token it's associated with.
type Statement interface {
	Node
	statementNode() // debug and test use
}

// Expression is a combination of tokens, which resolves to a value.
type Expression interface {
	Node
	expressionNode() // debug and test use
}

// Program is the entrance of AST consists of statements.
type Program struct {
	Statements []Statement
}

// TokenLiteral is a Node implementation for Programm
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

// LetStatement is one of the three types of statements.
// Form: TOKEN NAME = VALUE
type LetStatement struct {
	Token token.Token // token.Token{TYPE: token.LET, LITERAL: "LET"}
	Name  *Identifier // an identifier
	Value Expression  // value binds to the identifier
}

func (ls *LetStatement) statementNode() {}

// TokenLiteral is a Node implementation for LetStatement
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

// ReturnStatement is one of the three types of statements.
// Form: TOKEN ReturnValue
type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}

// TokenLiteral is a Node implementation for ReturnStatement
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}

// ExpressionStatement is a wrapper over Expression, thus we can add it to the
// Statements slice of ast.Program.
type ExpressionStatement struct {
	Token      token.Token // the first token of the expression
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}

// TokenLiteral is a Node implementation for ExpressionStatement
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}

	return ""
}

// Identifier as part of LetStatement, also implements the Expression interface,
// as in other part of Monkey language, it produce values like let x =
// valueProducingIdentifier
type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() {}

// TokenLiteral is a Node implementation for Identifier
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

func (i *Identifier) String() string {
	return i.Value
}
