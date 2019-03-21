// Every node in our AST has to implement the Node interface, meaning it has to provide a TokenLiteral() method that returns the literal value of the token it’s associated with. TokenLiteral() will be used only for debugging and testing. The AST we are going to construct consists solely of Nodes that are connected to each other - it’s a tree after all. Some of these nodes implement the Statement and some the Expression interface.
package ast

import "donkey/token"

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

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// To hold the identifier of the binding, the x in let x = 5;, we have the Identifier struct type, which implements the Expression interface. But the identifier in a let statement doesn’t produce a value, right? So why is it an Expression? It’s to keep things simple. Identifiers in other parts of a Monkey program do produce values, e.g.:  let x = valueProducingIdentifier;. And to keep the number of different node types small, we’ll use Identifier here to represent the name in a variable binding and later reuse it, to represent an identifier as part of or as a complete expression.
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
