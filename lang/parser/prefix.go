package parser

import (
	"strconv"

	"github.com/lindeneg/blue/lang/ast"
	"github.com/lindeneg/blue/lang/token"
)

type prefixFn func(p *P) ast.Expression

var prefixMap = map[token.Type]prefixFn{
	token.IDENTIFIER: parseIdentifier,
}

func parseIdentifier(p *P) ast.Expression {
	return &ast.Identifier{Token: p.cur, Value: p.cur.Literal}
}

func parseNumberLiteral(p *P) ast.Expression {
	n := &ast.Number{Token: p.cur}
	val, err := strconv.ParseFloat(p.cur.Literal, 64)
	if err != nil {
		//p.parseErr()
		return nil
	}
	n.Value = val
	return n
}
