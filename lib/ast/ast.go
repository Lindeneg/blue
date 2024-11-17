package ast

import "github.com/lindeneg/blue/lib/token"

type Node interface {
	Literal() string
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Identifier struct {
	Token token.T
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string       { return i.Value }

type IntegerLiteral struct {
	Token token.T
	Value int64
}

func (i *IntegerLiteral) expressionNode()      {}
func (i *IntegerLiteral) TokenLiteral() string { return i.Token.Literal }
func (i *IntegerLiteral) String() string       { return i.Token.Literal }

type FloatLiteral struct {
	Token token.T
	Value float64
}

func (i *FloatLiteral) expressionNode()      {}
func (i *FloatLiteral) TokenLiteral() string { return i.Token.Literal }
func (i *FloatLiteral) String() string       { return i.Token.Literal }
