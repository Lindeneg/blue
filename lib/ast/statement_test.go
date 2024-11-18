package ast

import (
	"testing"

	"github.com/lindeneg/blue/lib/token"
)

func TestStatements(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.T{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.T{Type: token.IDENTIFIER, Literal: "foo"},
					Value: "foo",
				},
				Value: &Identifier{
					Token: token.T{Type: token.IDENTIFIER, Literal: "bar"},
					Value: "bar",
				},
			},
			&ConstStatement{
				Token: token.T{Type: token.CONST, Literal: "const"},
				Name: &Identifier{
					Token: token.T{Type: token.IDENTIFIER, Literal: "foo"},
					Value: "foo",
				},
				Value: &Identifier{
					Token: token.T{Type: token.IDENTIFIER, Literal: "bar"},
					Value: "bar",
				},
			},
			&BlockStatement{
				Token: token.T{Type: token.LBRACE, Literal: "{"},
				Statements: []Statement{
					&ConstStatement{
						Token: token.T{Type: token.CONST, Literal: "const"},
						Name: &Identifier{
							Token: token.T{Type: token.IDENTIFIER, Literal: "foo"},
							Value: "foo",
						},
						Value: &Identifier{
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
		},
	}
	expected := "let foo = bar;const foo = bar;" +
		"{ const foo = bar; }" +
		"return foo;foo"

	if program.String() != expected {
		t.Errorf("program.String() wrong. got=%q, want=%q", program.String(), expected)
	}
}
