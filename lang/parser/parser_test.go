package parser

import (
	"fmt"
	"testing"

	"github.com/lindeneg/blue/lang/lexer"
)

func TestAssignmentStatements(t *testing.T) {
	tests := []struct {
		input              string
		expectedIdentifier string
		expectedValue      interface{}
	}{
		{"let x = 5;", "x", 5},
		//		{"let y = true;", "y", true},
		//		{"let foobar = y;", "foobar", "y"},
		//		{"const x = 5;", "x", 5},
		//		{"const y = true;", "y", true},
		//		{"const foobar = y;", "foobar", "y"},
	}

	for i, tt := range tests {
		p := New(lexer.FromString(tt.input), fmt.Sprintf("test%d", i))
		program := p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain 1 statements. got=%d",
				len(program.Statements))
		}

		//		stmt := program.Statements[0]
		//		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
		//			return
		//		}
		//
		//		val := stmt.(*ast.LetStatement).Value
		//		if !testLiteralExpression(t, val, tt.expectedValue) {
		//			return
		//		}
	}
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
