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

// LetStatement i.e let foo = 5;
type LetStatement struct {
	Token token.T
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statement()            {}
func (ls *LetStatement) Literal() string       { return ls.Token.Literal }
func (ls *LetStatement) Left() *Identifier     { return ls.Name }
func (ls *LetStatement) Right() Expression     { return ls.Value }
func (ls *LetStatement) SetLeft(i *Identifier) { ls.Name = i }
func (ls *LetStatement) SetRight(e Expression) { ls.Value = e }
func (ls *LetStatement) String() string {
	return assignString(ls)
}

// ConstStatement i.e const foo = 5;
type ConstStatement struct {
	Token token.T
	Name  *Identifier
	Value Expression
}

func (cs *ConstStatement) statement()            {}
func (cs *ConstStatement) Literal() string       { return cs.Token.Literal }
func (cs *ConstStatement) Left() *Identifier     { return cs.Name }
func (cs *ConstStatement) Right() Expression     { return cs.Value }
func (cs *ConstStatement) SetLeft(i *Identifier) { cs.Name = i }
func (cs *ConstStatement) SetRight(e Expression) { cs.Value = e }
func (cs *ConstStatement) String() string {
	return assignString(cs)
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

type Assignable interface {
	Node
	Left() *Identifier
	Right() Expression
	SetLeft(*Identifier)
	SetRight(Expression)
}

func assignString(a Assignable) string {
	var out bytes.Buffer
	out.WriteString(a.Literal() + " ")
	out.WriteString(a.Left().String())
	out.WriteString(" = ")
	if a.Right() != nil {
		out.WriteString(a.Right().String())
	}
	out.WriteString(";")
	return out.String()
}
