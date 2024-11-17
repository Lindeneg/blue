package lexer

import (
	"testing"

	"github.com/lindeneg/blue/lib/token"
)

func TestNextToken(t *testing.T) {
	input := `let x = 5;
let y = 10;
const x2 = x + 5;

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
"foobar";
"foo bar";
"foo 0 bar"

// this is a single line comment
/* this is a 
multiline 
block-comment */
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

		{token.FUNCTION, "def"},
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

		{token.FUNCTION, "def"},
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
		{token.TRUE, "false"},
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
	}
	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestEmptyLexer(t *testing.T) {
	input := "// comment"
	l := New(input)
	tok := l.NextToken()
	if tok.Type != token.EOF {
		t.Fatalf("tokentype wrong. expected=%q, got=%q",
			token.EOF, tok.Type)
	}
	if tok.Literal != "" {
		t.Fatalf("literal wrong. expected=%q, got=%q",
			"", tok.Literal)
	}
}
