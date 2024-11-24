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

func TestBooleanExpression(t *testing.T) {
	tests := []struct {
		input           string
		expectedBoolean bool
	}{
		{"true;", true},
		{"false;", false},
	}

	for i, tt := range tests {
		program := newProgram(t, tt.input, fmt.Sprintf("boolean-expression-%d", i))
		if len(program.Statements) != 1 {
			t.Fatalf("program has not enough statements. got=%d",
				len(program.Statements))
		}
		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T",
				program.Statements[0])
		}
		boolean, ok := stmt.Expression.(*ast.Boolean)
		if !ok {
			t.Fatalf("exp not *ast.Boolean. got=%T", stmt.Expression)
		}
		if boolean.Value != tt.expectedBoolean {
			t.Errorf("boolean.Value not %t. got=%t", tt.expectedBoolean,
				boolean.Value)
		}
	}
}

func TestParsingPrefixExpressions(t *testing.T) {
	prefixTests := []struct {
		input    string
		operator string
		value    interface{}
	}{
		{"!5;", "!", 5},
		{"-15;", "-", 15},
		{"!foobar;", "!", "foobar"},
		{"-foobar;", "-", "foobar"},
		{"!true;", "!", true},
		{"!false;", "!", false},
	}

	for i, tt := range prefixTests {
		program := newProgram(t, tt.input, fmt.Sprintf("prefix-expression-%d", i))
		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain %d statements. got=%d\n",
				1, len(program.Statements))
		}
		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T",
				program.Statements[0])
		}
		exp, ok := stmt.Expression.(*ast.PrefixExpression)
		if !ok {
			t.Fatalf("stmt is not ast.PrefixExpression. got=%T", stmt.Expression)
		}
		if exp.Operator != tt.operator {
			t.Fatalf("exp.Operator is not '%s'. got=%s",
				tt.operator, exp.Operator)
		}
		if !testLiteralExpression(t, exp.Right, tt.value) {
			return
		}
	}
}

func TestParsingInfixExpressions(t *testing.T) {
	infixTests := []struct {
		input      string
		leftValue  interface{}
		operator   string
		rightValue interface{}
	}{
		{"5 + 5;", 5, "+", 5},
		{"5 - 5;", 5, "-", 5},
		{"5 * 5;", 5, "*", 5},
		{"5 / 5;", 5, "/", 5},
		{"5 > 5;", 5, ">", 5},
		{"5 < 5;", 5, "<", 5},
		{"5 <= 5;", 5, "<=", 5},
		{"5 >= 5;", 5, ">=", 5},
		{"5 == 5;", 5, "==", 5},
		{"5 != 5;", 5, "!=", 5},
		{"foobar + barfoo;", "foobar", "+", "barfoo"},
		{"foobar - barfoo;", "foobar", "-", "barfoo"},
		{"foobar * barfoo;", "foobar", "*", "barfoo"},
		{"foobar / barfoo;", "foobar", "/", "barfoo"},
		{"foobar > barfoo;", "foobar", ">", "barfoo"},
		{"foobar < barfoo;", "foobar", "<", "barfoo"},
		{"foobar == barfoo;", "foobar", "==", "barfoo"},
		{"foobar != barfoo;", "foobar", "!=", "barfoo"},
		{"true == true", true, "==", true},
		{"true != false", true, "!=", false},
		{"false == false", false, "==", false},
	}

	for i, tt := range infixTests {
		program := newProgram(t, tt.input, fmt.Sprintf("infix-expression-%d", i))
		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain %d statements. got=%d\n",
				1, len(program.Statements))
		}
		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T",
				program.Statements[0])
		}
		if !testInfixExpression(t, stmt.Expression, tt.leftValue,
			tt.operator, tt.rightValue) {
			return
		}
	}
}

func TestOperatorPrecedenceParsing(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			"-a * b",
			"((-a) * b)",
		},
		{
			"!-a",
			"(!(-a))",
		},
		{
			"a + b + c",
			"((a + b) + c)",
		},
		{
			"a + b - c",
			"((a + b) - c)",
		},
		{
			"a * b * c",
			"((a * b) * c)",
		},
		{
			"a * [1, 2, 3, 4][b * c] * d",
			"((a * ([1, 2, 3, 4][(b * c)])) * d)",
		},
		{
			"add(a * b[2], b[1], 2 * [1, 2][1])",
			"add((a * (b[2])), (b[1]), (2 * ([1, 2][1])))",
		},
		{
			"a * b / c",
			"((a * b) / c)",
		},
		{
			"a + b / c",
			"(a + (b / c))",
		},
		{
			"a + b * c + d / e - f",
			"(((a + (b * c)) + (d / e)) - f)",
		},
		{
			"3 + 4; -5 * 5",
			"(3 + 4)((-5) * 5)",
		},
		{
			"5 > 4 == 3 < 4",
			"((5 > 4) == (3 < 4))",
		},
		{
			"5 < 4 != 3 > 4",
			"((5 < 4) != (3 > 4))",
		},
		{
			"3 + 4 * 5 == 3 * 1 + 4 * 5",
			"((3 + (4 * 5)) == ((3 * 1) + (4 * 5)))",
		},
		{
			"true",
			"true",
		},
		{
			"false",
			"false",
		},
		{
			"3 > 5 == false",
			"((3 > 5) == false)",
		},
		{
			"3 < 5 == true",
			"((3 < 5) == true)",
		},
		{
			"1 + (2 + 3) + 4",
			"((1 + (2 + 3)) + 4)",
		},
		{
			"(5 + 5) * 2",
			"((5 + 5) * 2)",
		},
		{
			"2 / (5 + 5)",
			"(2 / (5 + 5))",
		},
		{
			"(5 + 5) * 2 * (5 + 5)",
			"(((5 + 5) * 2) * (5 + 5))",
		},
		{
			"-(5 + 5)",
			"(-(5 + 5))",
		},
		{
			"!(true == true)",
			"(!(true == true))",
		},
		{
			"a + add(b * c) + d",
			"((a + add((b * c))) + d)",
		},
		{
			"add(a, b, 1, 2 * 3, 4 + 5, add(6, 7 * 8))",
			"add(a, b, 1, (2 * 3), (4 + 5), add(6, (7 * 8)))",
		},
		{
			"add(a + b + c * d / f + g)",
			"add((((a + b) + ((c * d) / f)) + g))",
		},
	}
	for i, tt := range tests {
		program := newProgram(t, tt.input, fmt.Sprintf("precedence-test-%d", i))
		actual := program.String()
		if actual != tt.expected {
			t.Errorf("expected=%q, got=%q", tt.expected, actual)
		}
	}
}

func TestParsingIndexExpressions(t *testing.T) {
	input := "myArray[1 + 1]"

	program := newProgram(t, input, "index.expression")

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	indexExp, ok := stmt.Expression.(*ast.IndexExpression)
	if !ok {
		t.Fatalf("exp not *ast.IndexExpression. got=%T", stmt.Expression)
	}
	if !testIdentifier(t, indexExp.Left.Literal(), "myArray", "myArray") {
		return
	}
	if !testInfixExpression(t, indexExp.Index, 1, "+", 1) {
		return
	}
}

func TestParsingArrayLiterals(t *testing.T) {
	input := "[1, 2 * 2, 3 + 3]"
	program := newProgram(t, input, "array.literal")
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	array, ok := stmt.Expression.(*ast.Array)
	if !ok {
		t.Fatalf("exp not ast.ArrayLiteral. got=%T", stmt.Expression)
	}
	if len(array.Elements) != 3 {
		t.Fatalf("len(array.Elements) not 3. got=%d", len(array.Elements))
	}

	testNumberLiteral(t, array.Elements[0], int64(1))
	testInfixExpression(t, array.Elements[1], 2, "*", 2)
	testInfixExpression(t, array.Elements[2], 3, "+", 3)
}

func TestParsingCommentLiteral(t *testing.T) {
	input := "// whateverworks hello there"
	program := newProgram(t, input, "comment.literal")
	if len(program.Statements) > 0 {
		t.Fatalf("expected empty program, got %d statements", len(program.Statements))
	}
}

func TestFunctionLiteralParsing(t *testing.T) {
	input := `fn(x, y) { x + y; }`

	program := newProgram(t, input, "function.literal")

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain %d statements. got=%d\n",
			1, len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T",
			program.Statements[0])
	}

	function, ok := stmt.Expression.(*ast.Function)
	if !ok {
		t.Fatalf("stmt.Expression is not ast.FunctionLiteral. got=%T",
			stmt.Expression)
	}

	if len(function.Parameters) != 2 {
		t.Fatalf("function literal parameters wrong. want 2, got=%d\n",
			len(function.Parameters))
	}

	testLiteralExpression(t, function.Parameters[0], "x")
	testLiteralExpression(t, function.Parameters[1], "y")

	if len(function.Body.Statements) != 1 {
		t.Fatalf("function.Body.Statements has not 1 statements. got=%d\n",
			len(function.Body.Statements))
	}
	bodyStmt, ok := function.Body.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("function body stmt is not ast.ExpressionStatement. got=%T",
			function.Body.Statements[0])
	}
	testInfixExpression(t, bodyStmt.Expression, "x", "+", "y")
}

//func TestFunctionParameterParsing(t *testing.T) {
//	tests := []struct {
//		input          string
//		expectedParams []string
//	}{
//		{input: "fn() {};", expectedParams: []string{}},
//		{input: "fn(x) {};", expectedParams: []string{"x"}},
//		{input: "fn(x, y, z) {};", expectedParams: []string{"x", "y", "z"}},
//	}
//
//	for _, tt := range tests {
//		l := lexer.NewLexer(tt.input)
//		p := New(l)
//		program := p.ParseProgram()
//		checkParserErrors(t, p)
//
//		stmt := program.Statements[0].(*ast.ExpressionStatement)
//		function := stmt.Expression.(*ast.FunctionLiteral)
//
//		if len(function.Parameters) != len(tt.expectedParams) {
//			t.Errorf("length parameters wrong. want %d, got=%d\n",
//				len(tt.expectedParams), len(function.Parameters))
//		}
//
//		for i, ident := range tt.expectedParams {
//			testLiteralExpression(t, function.Parameters[i], ident)
//		}
//	}
//}
//
//func TestCallExpressionParsing(t *testing.T) {
//	input := "add(1, 2 * 3, 4 + 5);"
//
//	l := lexer.NewLexer(input)
//	p := New(l)
//	program := p.ParseProgram()
//	checkParserErrors(t, p)
//
//	if len(program.Statements) != 1 {
//		t.Fatalf("program.Statements does not contain %d statements. got=%d\n",
//			1, len(program.Statements))
//	}
//
//	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
//	if !ok {
//		t.Fatalf("stmt is not ast.ExpressionStatement. got=%T",
//			program.Statements[0])
//	}
//
//	exp, ok := stmt.Expression.(*ast.CallExpression)
//	if !ok {
//		t.Fatalf("stmt.Expression is not ast.CallExpression. got=%T",
//			stmt.Expression)
//	}
//
//	if !testIdentifier(t, exp.Function, "add") {
//		return
//	}
//
//	if len(exp.Arguments) != 3 {
//		t.Fatalf("wrong length of arguments. got=%d", len(exp.Arguments))
//	}
//
//	testLiteralExpression(t, exp.Arguments[0], 1)
//	testInfixExpression(t, exp.Arguments[1], 2, "*", 3)
//	testInfixExpression(t, exp.Arguments[2], 4, "+", 5)
//}
//
//func TestCallExpressionParameterParsing(t *testing.T) {
//	tests := []struct {
//		input         string
//		expectedIdent string
//		expectedArgs  []string
//	}{
//		{
//			input:         "add();",
//			expectedIdent: "add",
//			expectedArgs:  []string{},
//		},
//		{
//			input:         "add(1);",
//			expectedIdent: "add",
//			expectedArgs:  []string{"1"},
//		},
//		{
//			input:         "add(1, 2 * 3, 4 + 5);",
//			expectedIdent: "add",
//			expectedArgs:  []string{"1", "(2 * 3)", "(4 + 5)"},
//		},
//	}
//
//	for _, tt := range tests {
//		l := lexer.NewLexer(tt.input)
//		p := New(l)
//		program := p.ParseProgram()
//		checkParserErrors(t, p)
//
//		stmt := program.Statements[0].(*ast.ExpressionStatement)
//		exp, ok := stmt.Expression.(*ast.CallExpression)
//		if !ok {
//			t.Fatalf("stmt.Expression is not ast.CallExpression. got=%T",
//				stmt.Expression)
//		}
//
//		if !testIdentifier(t, exp.Function, tt.expectedIdent) {
//			return
//		}
//
//		if len(exp.Arguments) != len(tt.expectedArgs) {
//			t.Fatalf("wrong number of arguments. want=%d, got=%d",
//				len(tt.expectedArgs), len(exp.Arguments))
//		}
//
//		for i, arg := range tt.expectedArgs {
//			if exp.Arguments[i].String() != arg {
//				t.Errorf("argument %d wrong. want=%q, got=%q", i,
//					arg, exp.Arguments[i].String())
//			}
//		}
//	}
//}

func testInfixExpression(t *testing.T, exp ast.Expression, left interface{},
	operator string, right interface{}) bool {

	opExp, ok := exp.(*ast.InfixExpression)
	if !ok {
		t.Errorf("exp is not ast.InfixExpression. got=%T(%s)", exp, exp)
		return false
	}

	if !testLiteralExpression(t, opExp.Left, left) {
		return false
	}

	if opExp.Operator != operator {
		t.Errorf("exp.Operator is not '%s'. got=%q", operator, opExp.Operator)
		return false
	}

	if !testLiteralExpression(t, opExp.Right, right) {
		return false
	}

	return true
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
