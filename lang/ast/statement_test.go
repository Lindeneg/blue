package ast

import (
	"testing"

	"github.com/lindeneg/blue/lang/token"
)

func TestStatements(t *testing.T) {
	expectedOutput := "let foo = bar;const foo = bar;" +
		"{ const foo = bar; }" +
		"return foo;" +
		"foo" +
		"foobar = true;"
	program := &Program{
		Statements: []Statement{
			&AssignStatement{
				Token: token.T{Type: token.LET, Literal: "let"},
				Left: &Identifier{
					Token: token.T{Type: token.IDENTIFIER, Literal: "foo"},
					Value: "foo",
				},
				Right: &Identifier{
					Token: token.T{Type: token.IDENTIFIER, Literal: "bar"},
					Value: "bar",
				},
			},
			&AssignStatement{
				Token: token.T{Type: token.CONST, Literal: "const"},
				Left: &Identifier{
					Token: token.T{Type: token.IDENTIFIER, Literal: "foo"},
					Value: "foo",
				},
				Right: &Identifier{
					Token: token.T{Type: token.IDENTIFIER, Literal: "bar"},
					Value: "bar",
				},
			},
			&BlockStatement{
				Token: token.T{Type: token.LBRACE, Literal: "{"},
				Statements: []Statement{
					&AssignStatement{
						Token: token.T{Type: token.CONST, Literal: "const"},
						Left: &Identifier{
							Token: token.T{Type: token.IDENTIFIER, Literal: "foo"},
							Value: "foo",
						},
						Right: &Identifier{
							Token: token.T{Type: token.IDENTIFIER, Literal: "bar"},
							Value: "bar",
						},
					},
				},
			},
			&ReturnStatement{
				Token: token.T{Type: token.RETURN, Literal: "return"},
				ReturnValue: &Identifier{
					Token: token.T{Type: token.IDENTIFIER, Literal: "foo"},
					Value: "foo",
				},
			},
			&ExpressionStatement{
				Token: token.T{Type: token.IDENTIFIER, Literal: "foo"},
				Expression: &Identifier{
					Token: token.T{Type: token.IDENTIFIER, Literal: "foo"},
					Value: "foo",
				},
			},
			&AssignStatement{
				Token: token.T{Type: token.IDENTIFIER, Literal: "foobar"},
				Left: &Identifier{
					Token: token.T{Type: token.IDENTIFIER, Literal: "foobar"},
					Value: "foobar",
				},
				Right: &Boolean{
					Token: token.T{Type: token.TRUE, Literal: "true"},
					Value: true,
				},
			},
		},
	}

	if program.String() != expectedOutput {
		t.Errorf("program.String() wrong.\ngot =%q\nwant=%q", program.String(), expectedOutput)
	}
}
