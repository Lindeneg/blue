package parser

import (
	"github.com/lindeneg/blue/lang/ast"
	"github.com/lindeneg/blue/lang/token"
)

type infixFn func(ast.Expression) ast.Expression
type infixMap map[token.Type]infixFn

func makeInfixMap(p *P) infixMap {
	return infixMap{
		token.PLUS:     p.parseInfixExpression,
		token.MINUS:    p.parseInfixExpression,
		token.STAR:     p.parseInfixExpression,
		token.FSLASH:   p.parseInfixExpression,
		token.EQ:       p.parseInfixExpression,
		token.NEQ:      p.parseInfixExpression,
		token.LT:       p.parseInfixExpression,
		token.LTOE:     p.parseInfixExpression,
		token.GT:       p.parseInfixExpression,
		token.GTOE:     p.parseInfixExpression,
		token.LPAREN:   p.parseCallExpression,
		token.LBRACKET: p.parseIndexExpression,
	}
}

func (p *P) parseInfixExpression(left ast.Expression) ast.Expression {
	expression := &ast.InfixExpression{
		Token:    p.cur,
		Operator: p.cur.Literal,
		Left:     left,
	}
	precedence := predMap.find(p.cur)
	p.advance() // consume the infix token
	expression.Right = p.parseExpression(precedence)

	return expression
}

func (p *P) parseIndexExpression(left ast.Expression) ast.Expression {
	expression := &ast.IndexExpression{
		Token: p.cur,
		Left:  left,
	}
	p.advance() // consume '['
	expression.Index = p.parseExpression(LOWEST)
	if !p.expectNext(token.RBRACKET) {
		return nil
	}
	p.advance() // consume ']'
	return expression
}

func (p *P) parseCallExpression(left ast.Expression) ast.Expression {
	expression := &ast.CallExpression{
		Token:    p.cur,
		Function: left,
	}
	expression.Arguments = p.parseExpressionList(token.RPAREN)
	return expression
}
