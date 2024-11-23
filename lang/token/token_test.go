package token

import "testing"

func TestHighlight(t *testing.T) {
	input := "let x != 5;"
	tests := []struct {
		color    string
		expected string
	}{
		{"foo", "let x != foo5\x1b[0m;"},
		{string(COLOR_RED), "let x != \x1b[31m5\x1b[0m;"},
	}
	for _, tt := range tests {
		tok := T{
			Type:    INT,
			Literal: "5",
		}
		got := tok.Highlight(input, Color(tt.color))
		if got != tt.expected {
			t.Fatalf("highlight failed\nwant=%q\ngot=%q", tt.expected, got)
		}
	}
}

func TestHighlightErr(t *testing.T) {
	input := "let x != 5;"
	want := "let x != \x1b[31m5\x1b[0m;"
	got := T{
		Type:    INT,
		Literal: "5",
	}.HighlightErr(input)
	if got != want {
		t.Fatalf("highlight failed\nwant=%q\ngot=%q", want, got)
	}
}
