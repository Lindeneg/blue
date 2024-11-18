package ast

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/lindeneg/blue/lib/token"
)

// Node must be implemented
// by all ast nodes
type Node interface {
	Literal() string
	String() string
}

// Identifier i.e foo, x, len
type Identifier struct {
	Token token.T
	Value string
}

func (i *Identifier) expression()     {}
func (i *Identifier) Literal() string { return i.Token.Literal }
func (i *Identifier) String() string  { return i.Value }

// Number i.e 5, 42.12
type Number struct {
	Token token.T
	Value float64
}

func (i *Number) expression()     {}
func (i *Number) Literal() string { return i.Token.Literal }
func (i *Number) String() string {
	if i.Token.Type == token.INT {
		return fmt.Sprintf("%d", int(i.Value))
	}
	return fmt.Sprintf("%.2f", i.Value)
}

// String i.e "foobar"
type String struct {
	Token token.T
	Value string
}

func (sl *String) expression()     {}
func (sl *String) Literal() string { return sl.Token.Literal }
func (sl *String) String() string {
	return fmt.Sprintf("\"%s\"", sl.Value)
}

// Boolean i.e true, false
type Boolean struct {
	Token token.T
	Value bool
}

func (b *Boolean) expression()     {}
func (b *Boolean) Literal() string { return b.Token.Literal }
func (b *Boolean) String() string {
	return b.Token.Literal
}

// Array i.e [1, 2, 3, 4, 5]
type Array struct {
	Token    token.T
	Elements []Expression
}

func (al *Array) expression()     {}
func (al *Array) Literal() string { return al.Token.Literal }
func (al *Array) String() string {
	var out bytes.Buffer
	elements := []string{}
	for _, el := range al.Elements {
		elements = append(elements, el.String())
	}
	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")
	return out.String()
}

// Function i.e def add(a, b) { return a + b; }
type Function struct {
	Token      token.T
	Name       *Identifier
	Parameters []*Identifier
	Body       *BlockStatement
}

func (fl *Function) expression()     {}
func (fl *Function) Literal() string { return fl.Token.Literal }
func (fl *Function) String() string {
	var out bytes.Buffer
	params := []string{}
	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}
	out.WriteString(fl.Literal())
	out.WriteString(" " + fl.Name.String())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(fl.Body.String())
	return out.String()
}

// Dict i.e {"foo": "bar", "baz": 1}
type Dict struct {
	Token token.T
	Pairs map[Expression]Expression
}

func (d *Dict) expression()     {}
func (d *Dict) Literal() string { return d.Token.Literal }
func (d *Dict) String() string {
	var out bytes.Buffer
	pairs := []string{}
	for key, value := range d.Pairs {
		pairs = append(pairs, key.String()+":"+value.String())
	}
	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")
	return out.String()
}
