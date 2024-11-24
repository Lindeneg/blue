package parser

import (
	"github.com/lindeneg/blue/lang/ast"
	"github.com/lindeneg/blue/lang/lexer"
	"github.com/lindeneg/blue/lang/token"
)

// P is the parser struct that holds the
// lexer and the current and next token
type P struct {
	l *lexer.L

	cur  token.T
	next token.T

	sourceName string

	errs []ParseErr
}

// New creates a new parser
func New(l *lexer.L, sourceName string) *P {
	p := &P{
		l:          l,
		sourceName: sourceName,
		errs:       make([]ParseErr, 0),
	}
	// initialize cur and next tokens
	p.advance()
	p.advance()
	return p
}

// Errors returns the errors that occured during parsing
func (p *P) Errors() []ParseErr {
	return p.errs
}

// HasErrors returns true if there are any errors
func (p *P) HasErrors() bool {
	return len(p.errs) > 0
}

// ParseProgram parses the program and returns the AST
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

// advance consumes the current token and sets the next token
func (p *P) advance() {
	p.cur = p.next
	p.next = p.l.NextToken()
}

// expect checks if the current token is of type want
// and sets an error if it is not
func (p *P) expect(got token.T, want token.Type) bool {
	if got.Type != want {
		expectErr(p, got, want)
		return false
	}
	return true
}

// expectCur checks if the current token is of type want
func (p *P) expectCur(t token.Type) bool {
	return p.expect(p.cur, t)
}

// expectNext checks if the next token is of type want
func (p *P) expectNext(t token.Type) bool {
	return p.expect(p.next, t)
}

// expectPrefix returns the prefix function for the current token
// or sets an error if the prefix function does not exist
func (p *P) expectPrefix() prefixFn {
	prefix, ok := prefixMap[p.cur.Type]
	if !ok {
		parseFnErr(p, "prefix", p.cur)
		return nil
	}
	return prefix
}

// expectInfix returns the infix function for the next token
// or sets an error if the infix function does not exist
func (p *P) expectInfix() infixFn {
	infix, ok := infixMap[p.next.Type]
	if !ok {
		parseFnErr(p, "infix", p.next)
		return nil
	}
	return infix
}

// parseStatement parses a statement and returns the AST node
func (p *P) parseStatement() ast.Statement {
	switch p.cur.Type {
	case token.LET, token.CONST:
		return p.parseAssignment(&ast.AssignStatement{Token: p.cur}, false)
	case token.IDENTIFIER:
		if p.next.Type == token.ASSIGN {
			return p.parseAssignment(&ast.AssignStatement{Token: p.cur}, true)
		}
	case token.RETURN:
		return nil
		//return p.parseReturnStatement()
	}
	return p.parseExpressionStatement()
}

// parseAssignment parses an assignment statement
func (p *P) parseAssignment(t *ast.AssignStatement, reassign bool) *ast.AssignStatement {
	if !reassign && !p.expectNext(token.IDENTIFIER) {
		return nil
	}
	if !reassign {
		p.advance() // consume 'let' or 'const'
	}
	t.Left = &ast.Identifier{Token: p.cur, Value: p.cur.Literal}
	p.advance() // consume 'identifier'
	if !p.expectCur(token.ASSIGN) {
		return nil
	}
	p.advance() // consume '='
	t.Right = p.parseExpression(LOWEST)
	if p.next.Type == token.SCOLON {
		p.advance() // consume ';'
	}
	return t
}

// parseExpression parses an expression and returns the AST node
func (p *P) parseExpression(pr pred) ast.Expression {
	var (
		prefix prefixFn
		infix  infixFn
	)
	if prefix = p.expectPrefix(); prefix == nil {
		return nil
	}
	leftExp := prefix(p) // consume lhs
	for p.next.Type != token.SCOLON && pr.lt(p.next) {
		if infix = p.expectInfix(); infix == nil {
			return nil
		}
		p.advance()                 // consume infix token
		leftExp = infix(p, leftExp) // consume rhs
	}
	return leftExp
}

func (p *P) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.cur}
	stmt.Expression = p.parseExpression(LOWEST)
	if p.next.Type == token.SCOLON {
		p.advance() // consume ';'
	}
	return stmt
}
