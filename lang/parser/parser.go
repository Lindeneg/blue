package parser

import (
	"github.com/lindeneg/blue/lang/ast"
	"github.com/lindeneg/blue/lang/lexer"
	"github.com/lindeneg/blue/lang/token"
)

type P struct {
	l *lexer.L

	cur  token.T
	next token.T

	sourceName string

	errs []ParseErr
}

func New(l *lexer.L, sourceName string) *P {
	p := &P{
		l:          l,
		sourceName: sourceName,
		errs:       make([]ParseErr, 0),
	}

	p.advance()
	p.advance()

	return p
}

func (p *P) Errors() []ParseErr {
	return p.errs
}

func (p *P) HasErrors() bool {
	return len(p.errs) > 0
}

func (p *P) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.cur.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.advance()
	}
	return program
}

func (p *P) advance() {
	p.cur = p.next
	p.next = p.l.NextToken()
}

func (p *P) expect(got token.T, want token.Type) bool {
	if got.Type != want {
		expectErr(p, got, want)
		return false
	}
	return true
}

func (p *P) expectCurrent(t token.Type) bool {
	return p.expect(p.cur, t)
}

func (p *P) expectPeek(t token.Type) bool {
	return p.expect(p.next, t)
}

func (p *P) expectPrefix() prefixFn {
	prefix, ok := prefixMap[p.cur.Type]
	if !ok {
		parseFnErr(p, "prefix", p.cur)
		return nil
	}
	return prefix
}

func (p *P) expectInfix() infixFn {
	infix, ok := infixMap[p.next.Type]
	if !ok {
		parseFnErr(p, "infix", p.next)
		return nil
	}
	return infix
}

func (p *P) parseStatement() ast.Statement {
	switch p.cur.Type {
	case token.LET:
		return parseAssignment(p, &ast.LetStatement{Token: p.cur})
	case token.CONST:
		return parseAssignment(p, &ast.ConstStatement{Token: p.cur})
	case token.RETURN:
		return nil
		//return p.parseReturnStatement()
	default:
		return nil
		//return p.parseExpressionStatement()
	}
}

func (p *P) parseExpression(pred pred) ast.Expression {
	var (
		prefix prefixFn
		infix  infixFn
	)
	if prefix = p.expectPrefix(); prefix == nil {
		return nil
	}
	leftExp := prefix(p)
	for p.next.Type != token.SCOLON && pred.lt(p.next) {
		if infix = p.expectInfix(); infix == nil {
			return nil
		}
		p.advance()
		leftExp = infix(p, leftExp)
	}
	return leftExp
}

func parseAssignment[T ast.Assignable](p *P, t T) T {
	if !p.expectPeek(token.IDENTIFIER) {
		return t
	}
	p.advance() // consume 'let' or 'const'
	t.SetLeft(&ast.Identifier{Token: p.cur, Value: p.cur.Literal})
	if !p.expectPeek(token.ASSIGN) {
		return t
	}
	p.advance() // consume '='
	t.SetRight(p.parseExpression(LOWEST))
	if p.next.Type == token.SCOLON {
		p.advance() // consume ';'
	}
	return t
}
