package ast

import (
	"bytes"
	"strings"

	"github.com/lindeneg/blue/lib/token"
)

// Expression node
type Expression interface {
	Node
	expression()
}

// IndexExpression i.e Left[Index]
type IndexExpression struct {
	Token token.T
	Left  Expression
	Index Expression
}

func (ie *IndexExpression) expression()     {}
func (ie *IndexExpression) Literal() string { return ie.Token.Literal }
func (ie *IndexExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString("[")
	out.WriteString(ie.Index.String())
	out.WriteString("])")
	return out.String()
}

// PrefixExpression i.e !true
type PrefixExpression struct {
	Token    token.T
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) expression()     {}
func (pe *PrefixExpression) Literal() string { return pe.Token.Literal }
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")
	return out.String()
}

// InfixExpression i.e 5 + 10
type InfixExpression struct {
	Token    token.T
	Left     Expression
	Operator string
	Right    Expression
}

func (ie *InfixExpression) expression()     {}
func (ie *InfixExpression) Literal() string { return ie.Token.Literal }
func (ie *InfixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" " + ie.Operator + " ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")
	return out.String()
}

type Conditional struct {
	Condition Expression
	Body      *BlockStatement
}

// IfExpression i.e if ... {} elif ... {} else {}
type IfExpression struct {
	Token token.T
	If    Conditional
	Elifs []Conditional
	Else  *BlockStatement
}

func (ie *IfExpression) expressionNode()      {}
func (ie *IfExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IfExpression) String() string {
	var out bytes.Buffer
	out.WriteString("if")
	out.WriteString(ie.If.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.If.Body.String())
	for _, elif := range ie.Elifs {
		out.WriteString("elif ")
		out.WriteString(elif.Condition.String())
		out.WriteString(" ")
		out.WriteString(elif.Body.String())
	}
	if ie.Else != nil {
		out.WriteString("else ")
		out.WriteString(ie.Else.String())
	}
	return out.String()
}

// ForExpression i.e for let i = range(arr) { }
type ForExpression struct {
	Token      token.T
	Assignment *LetStatement
	Body       *BlockStatement
}

func (fe *ForExpression) expressionNode()      {}
func (fe *ForExpression) TokenLiteral() string { return fe.Token.Literal }
func (fe *ForExpression) String() string {
	var out bytes.Buffer
	out.WriteString("for ")
	out.WriteString(fe.Assignment.String())
	out.WriteString(fe.Body.String())
	return out.String()
}

// CallExpression i.e (foo, bar)
type CallExpression struct {
	Token     token.T
	Function  Expression
	Arguments []Expression
}

func (ce *CallExpression) expressionNode()      {}
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
