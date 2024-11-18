package ast

import (
	"testing"

	"github.com/lindeneg/blue/lib/token"
)

func TestExpressions(t *testing.T) {
	expectedOutput := "(1 + 2.55)" +
		"if ((!x) && (foo < bar)) { return 1; }" +
		"elif ((!z) || (z > 10)) { return 2; }" +
		"elif done { return 3; }" +
		"else { return 4; }" +
		"for let i = range(arr);{ (arr[i]) }"
	program := &Program{
		Statements: []Statement{
			&ExpressionStatement{
				Expression: &InfixExpression{
					Token:    token.T{Type: token.PLUS, Literal: "+"},
					Left:     &Number{Token: token.T{Type: token.INT}, Value: 1},
					Operator: "+",
					Right:    &Number{Token: token.T{Type: token.FLOAT}, Value: 2.55},
				},
			},
			&ExpressionStatement{
				Expression: &IfExpression{
					Token: token.T{Type: token.IF, Literal: "if"},
					If: Conditional{
						Condition: &InfixExpression{
							Token: token.T{Type: token.AND, Literal: "&&"},
							Left: &PrefixExpression{
								Token:    token.T{Type: token.BANG, Literal: "!"},
								Operator: "!",
								Right:    &Identifier{Value: "x"},
							},
							Operator: "&&",
							Right: &InfixExpression{
								Token:    token.T{Type: token.LT, Literal: "<"},
								Left:     &Identifier{Value: "foo"},
								Operator: "<",
								Right:    &Identifier{Value: "bar"},
							},
						},
						Body: &BlockStatement{
							Statements: []Statement{
								&ReturnStatement{
									Token: token.T{Type: token.RETURN, Literal: "return"},
									ReturnValue: &Number{
										Token: token.T{Type: token.INT, Literal: "1"},
										Value: 1,
									},
								},
							},
						},
					},
					Elifs: []Conditional{
						{
							Condition: &InfixExpression{
								Token: token.T{Type: token.OR, Literal: "||"},
								Left: &PrefixExpression{
									Token:    token.T{Type: token.BANG, Literal: "!"},
									Operator: "!",
									Right:    &Identifier{Value: "z"},
								},
								Operator: "||",
								Right: &InfixExpression{
									Token:    token.T{Type: token.GT, Literal: ">"},
									Left:     &Identifier{Value: "z"},
									Operator: ">",
									Right:    &Identifier{Value: "10"},
								},
							},
							Body: &BlockStatement{
								Statements: []Statement{
									&ReturnStatement{
										Token: token.T{Type: token.RETURN, Literal: "return"},
										ReturnValue: &Number{
											Token: token.T{Type: token.INT, Literal: "2"},
											Value: 2,
										},
									},
								},
							},
						},
						{
							Condition: &Identifier{Value: "done"},
							Body: &BlockStatement{
								Statements: []Statement{
									&ReturnStatement{
										Token: token.T{Type: token.RETURN, Literal: "return"},
										ReturnValue: &Number{
											Token: token.T{Type: token.INT, Literal: "3"},
											Value: 3,
										},
									},
								},
							},
						},
					},
					Else: &BlockStatement{
						Statements: []Statement{
							&ReturnStatement{
								Token: token.T{Type: token.RETURN, Literal: "return"},
								ReturnValue: &Number{
									Token: token.T{Type: token.INT, Literal: "4"},
									Value: 4,
								},
							},
						},
					},
				},
			},
			&ExpressionStatement{
				Expression: &ForExpression{
					Token: token.T{Type: token.FOR, Literal: "for"},
					Assignment: &LetStatement{
						Token: token.T{Type: token.LET, Literal: "let"},
						Name:  &Identifier{Value: "i"},
						Value: &CallExpression{
							Function: &Function{
								Name: &Identifier{Value: "range"},
							},
							Arguments: []Expression{
								&Identifier{Value: "arr"},
							},
						},
					},
					Body: &BlockStatement{
						Statements: []Statement{
							&ExpressionStatement{
								Expression: &IndexExpression{
									Left:  &Identifier{Value: "arr"},
									Index: &Identifier{Value: "i"},
								},
							},
						},
					},
				},
			},
		},
	}

	if program.String() != expectedOutput {
		t.Errorf("program.String() wrong.\ngot =%q\nwant=%q", program.String(), expectedOutput)
	}
}
