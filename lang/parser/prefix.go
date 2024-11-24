package parser

import (
	"strconv"

	"github.com/lindeneg/blue/lang/ast"
	"github.com/lindeneg/blue/lang/token"
)

type prefixFn func() ast.Expression
type prefixMap map[token.Type]prefixFn

func makePrefixMap(p *P) prefixMap {
	return prefixMap{
		token.IDENTIFIER: p.parseIdentifier,
		token.STRING:     p.parseStringLiteral,
		token.BANG:       p.parsePrefixExpression,
		token.MINUS:      p.parsePrefixExpression,
		token.INT:        p.parseNumberLiteral,
		token.FLOAT:      p.parseNumberLiteral,
		token.TRUE:       p.parseBooleanLiteral,
		token.FALSE:      p.parseBooleanLiteral,

		token.LPAREN:   p.parseGroupedExpression,
		token.LBRACKET: p.parseArrayLiteral,
		//		token.LBRACE:   p.parseHashLiteral,
		//		token.IF:       p.parseIfExpression,
		//		token.DEF:      p.parseFunctionLiteral,
		//		token.FOR:      p.parseForExpression,
	}

}

func (p *P) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.cur, Value: p.cur.Literal}
}

func (p *P) parseNumberLiteral() ast.Expression {
	n := &ast.Number{Token: p.cur}
	val, err := strconv.ParseFloat(p.cur.Literal, 64)
	if err != nil {
		perr(p, p.cur, "failed to parse %q as a number\n%s", p.cur.Literal, err)
		return nil
	}
	n.Value = val
	return n
}

func (p *P) parseBooleanLiteral() ast.Expression {
	b := &ast.Boolean{Token: p.cur}
	b.Value = p.cur.Type == token.TRUE
	return b
}

func (p *P) parseStringLiteral() ast.Expression {
	return &ast.String{Token: p.cur, Value: p.cur.Literal}
}

func (p *P) parsePrefixExpression() ast.Expression {
	expr := &ast.PrefixExpression{Token: p.cur, Operator: p.cur.Literal}
	p.advance() // consume the prefix token
	expr.Right = p.parseExpression(PREFIX)
	return expr
}

func (p *P) parseGroupedExpression() ast.Expression {
	p.advance() // consume '('
	expr := p.parseExpression(LOWEST)
	if !p.expectNext(token.RPAREN) {
		return nil
	}
	p.advance() // consume ')'
	return expr
}

func (p *P) parseArrayLiteral() ast.Expression {
	array := &ast.Array{Token: p.cur}
	array.Elements = p.parseExpressionList(token.RBRACKET)
	return array
}

func (p *P) parseExpressionList(end token.Type) []ast.Expression {
	var list []ast.Expression
	if p.next.Type == end {
		p.advance() // consume ']'
		return list
	}
	p.advance() // consume '['
	list = append(list, p.parseExpression(LOWEST))
	for p.next.Type == token.COMMA {
		p.advance() // consume ','
		p.advance() // consume next element
		list = append(list, p.parseExpression(LOWEST))
	}
	if !p.expectNext(end) {
		return nil
	}
	p.advance() // consume ']'
	return list
}
