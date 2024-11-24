package ast

import (
	"testing"

	"github.com/lindeneg/blue/lang/token"
)

func TestExpressions(t *testing.T) {
	expectedOutput := "(1 + 2.55)" +
		"if ((!x) && (foo < bar)) { return false; }" +
		"elif ((!z) || (z > 10)) { return null; }" +
		`elif done { return "done"; }` +
		"else { return 4; }" +
		"for let i = range(arr);{ (arr[i]) }" +
		`["foo", "bar"]` +
		`|"foo":"bar", "baz":"qux"|`
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
									ReturnValue: &Boolean{
										Token: token.T{Type: token.FALSE, Literal: "false"},
										Value: false,
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
										Token:       token.T{Type: token.RETURN, Literal: "return"},
										ReturnValue: &Null{},
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
										ReturnValue: &String{
											Token: token.T{Type: token.STRING, Literal: "done"},
											Value: "done",
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
					Assignment: &AssignStatement{
						Token: token.T{Type: token.LET, Literal: "let"},
						Left:  &Identifier{Value: "i"},
						Right: &CallExpression{
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
			&ExpressionStatement{
				Expression: &Array{
					Elements: []Expression{
						&String{Value: "foo"},
						&String{Value: "bar"},
					},
				},
			},
			&ExpressionStatement{
				Expression: &Dict{
					Pairs: map[Expression]Expression{
						&String{Value: "foo"}: &String{Value: "bar"},
						&String{Value: "baz"}: &String{Value: "qux"},
					},
				},
			},
		},
	}

	if program.String() != expectedOutput {
		t.Errorf("program.String() wrong.\ngot =%q\nwant=%q", program.String(), expectedOutput)
	}
}
