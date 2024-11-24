package ast

import (
	"bytes"

	"github.com/lindeneg/blue/lang/token"
)

// Statement node
type Statement interface {
	Node
	statement()
}

// AssignStatement i.e let foo = 5, const foo = 5, foo = 5;
type AssignStatement struct {
	Token token.T
	Left  *Identifier
	Right Expression
}

func (as *AssignStatement) statement()      {}
func (as *AssignStatement) Literal() string { return as.Token.Literal }
func (as *AssignStatement) String() string {
	var out bytes.Buffer
	if as.Token.Type != token.IDENTIFIER {
		out.WriteString(as.Literal() + " ")
	}
	out.WriteString(as.Left.String())
	out.WriteString(" = ")
	if as.Right != nil {
		out.WriteString(as.Right.String())
	}
	out.WriteString(";")
	return out.String()
}

// BlockStatement i.e { ...stuff }
type BlockStatement struct {
	Token      token.T
	Statements []Statement
}

func (bs *BlockStatement) statement()      {}
func (bs *BlockStatement) Literal() string { return bs.Token.Literal }
func (bs *BlockStatement) String() string {
	var out bytes.Buffer
	out.WriteString("{ ")
	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}
	out.WriteString(" }")
	return out.String()
}

// ReturnStatement i.e return 5;
type ReturnStatement struct {
	Token       token.T
	ReturnValue Expression
}

func (rs *ReturnStatement) statement() {}
func (rs *ReturnStatement) Literal() string {
	return rs.Token.Literal
}
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer
	out.WriteString(rs.Literal() + " ")
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
	out.WriteString(";")
	return out.String()
}

// ExpressionStatement statement that is not let or return
type ExpressionStatement struct {
	// first token in the expression
	Token      token.T
	Expression Expression
}

func (es *ExpressionStatement) statement()      {}
func (es *ExpressionStatement) Literal() string { return es.Token.Literal }
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}
