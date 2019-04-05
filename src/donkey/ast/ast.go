package ast

import (
	"bytes"
	"donkey/token"
	"strings"
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

// IntegerLiteral is a expression of integer literal.
type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode() {}

// TokenLiteral is a Node implementation for IntegerLiteral
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }

func (il *IntegerLiteral) String() string { return il.Token.Literal }

// PrefixExpression is an expression with prefix operator
type PrefixExpression struct {
	Token    token.Token // the prefix token, eg. - !
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) expressionNode() {}

// TokenLiteral is a Node implementation for PrefixExpression
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}

// InfixExpression is an expression with infix operator
type InfixExpression struct {
	Token    token.Token // The operator token, e.g. +
	Left     Expression
	Operator string
	Right    Expression
}

func (ie *InfixExpression) expressionNode() {}

// TokenLiteral is a Node implementation for InfixExpression
func (ie *InfixExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" " + ie.Operator + " ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")

	return out.String()
}

// Boolean as an expression
type Boolean struct {
	Token token.Token
	Value bool
}

func (b *Boolean) expressionNode() {}

// TokenLiteral is a Node implementation for Boolean
func (b *Boolean) TokenLiteral() string { return b.Token.Literal }
func (b *Boolean) String() string       { return b.Token.Literal }

// IfExpression as an expression following the pattern:
// if (<condition>) <consequence> else <alternative>
type IfExpression struct {
	Token       token.Token // The 'if' token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (ie *IfExpression) expressionNode() {}

// TokenLiteral is a Node implementation for IfExpression
func (ie *IfExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IfExpression) String() string {
	var out bytes.Buffer

	out.WriteString("if")
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())

	if ie.Alternative != nil {
		out.WriteString("else ")
		out.WriteString(ie.Alternative.String())
	}

	return out.String()
}

// BlockStatement represents a block of statements: { statements }
type BlockStatement struct {
	Token      token.Token // the { token
	Statements []Statement
}

func (bs *BlockStatement) statementNode() {}

// TokenLiteral is a Node implementation for BlockStatement
func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }
func (bs *BlockStatement) String() string {
	var out bytes.Buffer

	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

// FunctionLiteral represents function as expression following the pattern:
// fn <parameters> <block statement>
type FunctionLiteral struct {
	Token      token.Token // The 'fn' token
	Parameters []*Identifier
	Body       *BlockStatement
}

func (fl *FunctionLiteral) expressionNode() {}

// TokenLiteral is a Node implementation for FunctionLiteral
func (fl *FunctionLiteral) TokenLiteral() string { return fl.Token.Literal }
func (fl *FunctionLiteral) String() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}

	out.WriteString(fl.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(fl.Body.String())

	return out.String()
}

// CallExpression is function calls following the pattern:
// <expression>(<comma separated expressions>)
type CallExpression struct {
	Token     token.Token // The '(' token
	Function  Expression  // Identifier or FunctionLiteral
	Arguments []Expression
}

func (ce *CallExpression) expressionNode() {}

// TokenLiteral is a Node implementation for CallExpression
func (ce *CallExpression) TokenLiteral() string { return ce.Token.Literal }
func (ce *CallExpression) String() string {
	var out bytes.Buffer

	args := []string{}
	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}

	out.WriteString(ce.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")

	return out.String()
}
