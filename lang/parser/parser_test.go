package parser

import (
	"fmt"
	"testing"

	"github.com/lindeneg/blue/lang/ast"
	"github.com/lindeneg/blue/lang/lexer"
	"github.com/lindeneg/blue/lang/token"
)

func TestAssignmentStatements(t *testing.T) {
	tests := []struct {
		input              string
		expectedIdentifier string
		expectedValue      interface{}
	}{
		{"let x = 5;", "x", 5},
		{"const x = 5;", "x", 5},
		{"let x = 5.57;", "x", 5.57},
		{"let y = true;", "y", true},
		{"const y = false;", "y", false},
		{"foobar = y;", "foobar", "y"},
		{`foobar = "hello"`, "foobar", "hello"},
	}

	for i, tt := range tests {
		program := newProgram(t, tt.input, fmt.Sprintf("assignment-statement-%d", i))
		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain 1 statements. got=%d",
				len(program.Statements))
		}
		stmt := program.Statements[0]
		if !testAssignStatement(t, stmt, tt.expectedIdentifier, tt.expectedValue) {
			t.Fatalf("testAssignStatement failed for test %d: %q", i, stmt)
		}
	}
}

func TestIdentifierExpression(t *testing.T) {
	input := "foobar;"
	program := newProgram(t, input, "identifier.expression")
	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got=%d",
			len(program.Statements))
	}
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T",
			program.Statements[0])
	}
	ident, ok := stmt.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("exp not *ast.Identifier. got=%T", stmt.Expression)
	}
	if ident.Value != "foobar" {
		t.Errorf("ident.Value not %s. got=%s", "foobar", ident.Value)
	}
	if ident.Literal() != "foobar" {
		t.Errorf("ident.TokenLiteral not %s. got=%s", "foobar",
			ident.Literal())
	}
}

func TestStringLiteralExpression(t *testing.T) {
	input := `"hello world";`
	program := newProgram(t, input, "string.literal.expression")
	stmt := program.Statements[0].(*ast.ExpressionStatement)
	literal, ok := stmt.Expression.(*ast.String)
	if !ok {
		t.Fatalf("exp not *ast.String. got=%T", stmt.Expression)
	}
	if literal.Value != "hello world" {
		t.Errorf("literal.Value not %q. got=%q", "hello world", literal.Value)
	}
}

func TestIntegerLiteralExpression(t *testing.T) {
	input := "5;"
	program := newProgram(t, input, "integer.literal.expression")

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got=%d",
			len(program.Statements))
	}
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T",
			program.Statements[0])
	}

	testNumberLiteral(t, stmt.Expression, int64(5))
}

func TestFloatLiteralExpression(t *testing.T) {
	input := "5.62;"
	program := newProgram(t, input, "float.literal.expression")

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got=%d",
			len(program.Statements))
	}
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T",
			program.Statements[0])
	}

	testNumberLiteral(t, stmt.Expression, float64(5.62))
}

func testAssignStatement(t *testing.T, s ast.Statement, identifier string, value any) bool {
	as, ok := s.(*ast.AssignStatement)
	if !ok {
		t.Errorf("unexpected statement want='*ast.AssignStatement', got=%T", s)
		return false
	}
	if s.Literal() != as.Literal() {
		t.Errorf("expected Literal, want=%q, got=%q", as.Literal(), s.Literal())
		return false
	}
	if as.Left.Value != identifier {
		t.Errorf("expected Name.Left, want=%q, got=%s", identifier, as.Left.Value)
		return false
	}
	if as.Left.Literal() != identifier {
		t.Errorf("expected Name.Literal, want=%q, got=%s", identifier, as.Left.Literal())
		return false
	}
	return testLiteralExpression(t, as.Right, value)
}

func testLiteralExpression(
	t *testing.T,
	exp ast.Expression,
	expected interface{},
) bool {
	switch v := expected.(type) {
	case int:
		return testNumberLiteral(t, exp, int64(v))
	case int64:
		return testNumberLiteral(t, exp, v)
	case float32:
		return testNumberLiteral(t, exp, float64(v))
	case float64:
		return testNumberLiteral(t, exp, v)
	case string:
		return testString(t, exp, v)
	case bool:
		return testBooleanLiteral(t, exp, v)
	}
	t.Errorf("type of exp not handled. got=%T", exp)
	return false
}

func testString(t *testing.T, exp ast.Expression, value string) bool {
	switch exp := exp.(type) {
	case *ast.Identifier:
		return testIdentifier(t, value, exp.Value, exp.Literal())
	case *ast.String:
		return testIdentifier(t, value, exp.Value, exp.Literal())
	}
	return false
}

func testIdentifier(t *testing.T, want, gotValue, gotLit string) bool {
	if want != gotValue {
		t.Errorf("unexpected Value want=%q, got=%q", want, gotValue)
		return false
	}
	if want != gotLit {
		t.Errorf("unexpected Literal want=%q, got=%q", want, gotLit)
		return false
	}
	return true
}

func testBooleanLiteral(t *testing.T, exp ast.Expression, value bool) bool {
	b, ok := exp.(*ast.Boolean)
	if !ok {
		t.Errorf("unexpected exp want=*ast.Boolean, got=%T", exp)
		return false
	}
	if b.Value != value {
		t.Errorf("unexpcted b.Value want=%t, got=%t", value, b.Value)
		return false
	}
	if b.Literal() != fmt.Sprintf("%t", value) {
		t.Errorf("unexpected b.Literal want=%t, got=%s", value, b.Literal())
		return false
	}
	return true
}

func testNumberLiteral[T int64 | float64](t *testing.T, num ast.Expression, value T) bool {
	n, ok := num.(*ast.Number)
	if !ok {
		t.Errorf("unexpected num want=*ast.Number, got=%T", num)
		return false
	}
	switch n.Token.Type {
	case token.INT:
		if int64(n.Value) != int64(value) {
			t.Errorf("unexpected num.Value want=%d, got=%d", int64(value), int64(n.Value))
			return false
		}
		if n.Literal() != fmt.Sprintf("%d", int64(value)) {
			t.Errorf("unexpected num.Literal want=%d, got=%s", int64(value), n.Literal())
			return false
		}
	case token.FLOAT:
		if n.Value != float64(value) {
			t.Errorf("unexpected num.Value want=%f, got=%f", float64(value), n.Value)
			return false
		}
		if n.Literal() != fmt.Sprintf("%.2f", float64(value)) {
			t.Errorf("unexpected num.Literal want=%.2f, got=%s", float64(value), n.Literal())
			return false
		}
	default:
		t.Errorf("unexpected token type 'INT' or 'FLOAT', got=%q", n.Token.Type)
	}
	return true
}

func checkParserErrors(t *testing.T, p *P) {
	t.Helper()
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}
	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Error(msg.Msg)
	}
	t.FailNow()
}

func newProgram(t *testing.T, input, name string) *ast.Program {
	t.Helper()
	l := lexer.FromString(input)
	p := New(l, name)
	program := p.ParseProgram()
	checkParserErrors(t, p)
	return program
}
