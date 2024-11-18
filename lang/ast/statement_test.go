package ast

import (
	"testing"

	"github.com/lindeneg/blue/lang/token"
)

func TestStatements(t *testing.T) {
	expectedOutput := "let foo = bar;const foo = bar;" +
		"{ const foo = bar; }" +
		"return foo;foo"
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

	if program.String() != expectedOutput {
		t.Errorf("program.String() wrong.\ngot =%q\nwant=%q", program.String(), expectedOutput)
	}
}
