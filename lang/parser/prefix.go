package parser

import (
	"strconv"

	"github.com/lindeneg/blue/lang/ast"
	"github.com/lindeneg/blue/lang/token"
)

type prefixFn func(p *P) ast.Expression

var prefixMap = map[token.Type]prefixFn{
	token.IDENTIFIER: parseIdentifier,
	token.STRING:     parseStringLiteral,
	token.INT:        parseNumberLiteral,
	token.FLOAT:      parseNumberLiteral,
	token.TRUE:       parseBooleanLiteral,
	token.FALSE:      parseBooleanLiteral,
}

func parseIdentifier(p *P) ast.Expression {
	return &ast.Identifier{Token: p.cur, Value: p.cur.Literal}
}

func parseNumberLiteral(p *P) ast.Expression {
	n := &ast.Number{Token: p.cur}
	val, err := strconv.ParseFloat(p.cur.Literal, 64)
	if err != nil {
		perr(p, p.cur, "failed to parse %q as a number\n%s", p.cur.Literal, err)
		return nil
	}
	n.Value = val
	return n
}

func parseBooleanLiteral(p *P) ast.Expression {
	b := &ast.Boolean{Token: p.cur}
	b.Value = p.cur.Type == token.TRUE
	return b
}

func parseStringLiteral(p *P) ast.Expression {
	return &ast.String{Token: p.cur, Value: p.cur.Literal}
}
