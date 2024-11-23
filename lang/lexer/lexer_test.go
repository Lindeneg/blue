package lexer

import (
	"testing"

	"github.com/lindeneg/blue/lang/token"
)

func TestNextToken(t *testing.T) {
	input := `let x = 5;
let y = 10;
const x2 = x + 5;
const x3 = 42.89;

def add(a, b) { return a + b; }

const result = add(x, y);

def foo() {
    if x > y && result < 10 {
        return true;
    } elif x2 > y || result == 15 {
        return true;
    } else {
        return false;
    }
}

const arr = [1, 2];

for let i = range(arr) {
    arr[i];
}

const obj = {"foo": "bar", "baz": 1};

!-/*1;
10 != 9;
10 >= 9;
10 <= 9;
"foobar"
"foo bar"
"foo 0 bar"
null

// this is a single line comment
`
	tests := []struct {
		expectedType    token.Type
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENTIFIER, "x"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SCOLON, ";"},

		{token.LET, "let"},
		{token.IDENTIFIER, "y"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SCOLON, ";"},

		{token.CONST, "const"},
		{token.IDENTIFIER, "x2"},
		{token.ASSIGN, "="},
		{token.IDENTIFIER, "x"},
		{token.PLUS, "+"},
		{token.INT, "5"},
		{token.SCOLON, ";"},

		{token.CONST, "const"},
		{token.IDENTIFIER, "x3"},
		{token.ASSIGN, "="},
		{token.FLOAT, "42.89"},
		{token.SCOLON, ";"},

		{token.DEF, "def"},
		{token.IDENTIFIER, "add"},
		{token.LPAREN, "("},
		{token.IDENTIFIER, "a"},
		{token.COMMA, ","},
		{token.IDENTIFIER, "b"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.IDENTIFIER, "a"},
		{token.PLUS, "+"},
		{token.IDENTIFIER, "b"},
		{token.SCOLON, ";"},
		{token.RBRACE, "}"},

		{token.CONST, "const"},
		{token.IDENTIFIER, "result"},
		{token.ASSIGN, "="},
		{token.IDENTIFIER, "add"},
		{token.LPAREN, "("},
		{token.IDENTIFIER, "x"},
		{token.COMMA, ","},
		{token.IDENTIFIER, "y"},
		{token.RPAREN, ")"},
		{token.SCOLON, ";"},

		{token.DEF, "def"},
		{token.IDENTIFIER, "foo"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IF, "if"},
		{token.IDENTIFIER, "x"},
		{token.GT, ">"},
		{token.IDENTIFIER, "y"},
		{token.AND, "&&"},
		{token.IDENTIFIER, "result"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SCOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELIF, "elif"},
		{token.IDENTIFIER, "x2"},
		{token.GT, ">"},
		{token.IDENTIFIER, "y"},
		{token.OR, "||"},
		{token.IDENTIFIER, "result"},
		{token.EQ, "=="},
		{token.INT, "15"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SCOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SCOLON, ";"},
		{token.RBRACE, "}"},
		{token.RBRACE, "}"},

		{token.CONST, "const"},
		{token.IDENTIFIER, "arr"},
		{token.ASSIGN, "="},
		{token.LBRACKET, "["},
		{token.INT, "1"},
		{token.COMMA, ","},
		{token.INT, "2"},
		{token.RBRACKET, "]"},
		{token.SCOLON, ";"},

		{token.FOR, "for"},
		{token.LET, "let"},
		{token.IDENTIFIER, "i"},
		{token.ASSIGN, "="},
		{token.IDENTIFIER, "range"},
		{token.LPAREN, "("},
		{token.IDENTIFIER, "arr"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENTIFIER, "arr"},
		{token.LBRACKET, "["},
		{token.IDENTIFIER, "i"},
		{token.RBRACKET, "]"},
		{token.SCOLON, ";"},
		{token.RBRACE, "}"},

		{token.CONST, "const"},
		{token.IDENTIFIER, "obj"},
		{token.ASSIGN, "="},
		{token.LBRACE, "{"},
		{token.STRING, "foo"},
		{token.COLON, ":"},
		{token.STRING, "bar"},
		{token.COMMA, ","},
		{token.STRING, "baz"},
		{token.COLON, ":"},
		{token.INT, "1"},
		{token.RBRACE, "}"},
		{token.SCOLON, ";"},

		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.FSLASH, "/"},
		{token.STAR, "*"},
		{token.INT, "1"},
		{token.SCOLON, ";"},
		{token.INT, "10"},
		{token.NEQ, "!="},
		{token.INT, "9"},
		{token.SCOLON, ";"},
		{token.INT, "10"},
		{token.GTOE, ">="},
		{token.INT, "9"},
		{token.SCOLON, ";"},
		{token.INT, "10"},
		{token.LTOE, "<="},
		{token.INT, "9"},
		{token.SCOLON, ";"},
		{token.STRING, "foobar"},
		{token.STRING, "foo bar"},
		{token.STRING, "foo 0 bar"},
		{token.NULL, "null"},
		{token.EOF, ""},
	}
	l := FromString(input)
	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - token.Type wrong. expected=%d, got=%v",
				i, tt.expectedType, tok)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - token.Literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok)
		}
	}
}

func TestLineAndColToken(t *testing.T) {
	input := `// test
let
def
let x2 = x + 5
const x3 = 42.89
`
	tests := []struct {
		expectedLine int
		expectedCol  int
	}{
		{2, 1},
		{3, 1},
		{4, 1},
		{4, 5},
		{4, 8},
		{4, 10},
		{4, 12},
		{4, 14},
		{5, 1},
		{5, 7},
		{5, 10},
		{5, 12},
	}
	l := FromString(input)
	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Line != tt.expectedLine {
			t.Fatalf("tests[%d] - token.Line wrong. expected=%d, got=%v",
				i, tt.expectedLine, tok)
		}
		if tok.Col != tt.expectedCol {
			t.Fatalf("tests[%d] - token.Col wrong. expected=%d, got=%v",
				i, tt.expectedCol, tok)
		}
	}
}

func TestScopedToken(t *testing.T) {
	input := `
let
{
    let
    {
        let
    }
    let
}
let
`
	tests := []struct {
		expectedScope token.Scope
	}{
		{0},
		{1},
		{1},
		{2},
		{2},
		{2},
		{1},
		{1},
		{0},
	}
	l := FromString(input)
	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Scope != tt.expectedScope {
			t.Fatalf("tests[%d] - token.Scope wrong. expected=%d, got=%v",
				i, tt.expectedScope, tok)
		}
	}
}

func TestLexerLine(t *testing.T) {
	input := `let x = 5;
let y = 10;
const x2 = x + 5;
const x3 = 42.89;
`
	tests := []struct {
		line int
		want string
	}{
		{1, "let x = 5;"},
		{3, "const x2 = x + 5;"},
		{4, "const x3 = 42.89;"},
		{-1, ""},
		{0, ""},
	}

	for _, tt := range tests {
		l := FromString(input)
		got := l.Line(tt.line)
		if got != tt.want {
			t.Fatalf("line wrong\nwant=%q\ngot=%q",
				tt.want, got)
		}
	}
}

func TestEmptyLexer(t *testing.T) {
	input := "// comment"
	l := FromString(input)
	tok := l.NextToken()
	if tok.Type != token.EOF {
		t.Fatalf("token.Type wrong. expected=%d, got=%d",
			token.EOF, tok.Type)
	}
	if tok.Literal != "" {
		t.Fatalf("literal wrong. expected=%q, got=%q",
			"", tok.Literal)
	}
}
