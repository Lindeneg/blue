package token

import (
	"fmt"
	"strings"
)

type Color string

const (
	ColorRed   Color = "\033[31m"
	ColorReset Color = "\033[0m"
)

// Scope describes layers of nesting.
// scope = 0 => global
// scope >= 1 => local where the number describes layers of nesting
type Scope int

// Global is scope global?
func (s Scope) Global() bool {
	return s == 0
}

// Local is scope local?
func (s Scope) Local() bool {
	return s > 0
}

// T describes a token encountered during lexical analysis.
type T struct {
	// Type of token encountered
	Type
	// Literal value
	Literal string
	// TODO can probably be smarter about this
	// Line the token was encountered at
	Line int
	// Column the token starts at
	Col int
	// scope level (0 global)
	Scope
}

// String returns full context of token
func (t T) String() string {
	return fmt.Sprintf("Type=%d|Literal=%q|Line=%d|Col=%d|Scope=%d",
		t.Type, t.Literal, t.Line, t.Col, t.Scope)
}

// Highlight returns a string that contains m,
// with t being colored by color
func (t T) Highlight(m string, color Color) string {
	if m == "" {
		return m
	}
	errorTokenEscaped := strings.ReplaceAll(t.Literal, "\n", "\\n")
	errorTokenEscaped = strings.ReplaceAll(errorTokenEscaped, "\t", "\\t")
	coloredToken := fmt.Sprintf("%s%s%s", color, errorTokenEscaped, ColorReset)
	return strings.Replace(m, t.Literal, coloredToken, 1)
}

// HighlightErr calls Highlight with ColorRed
func (t T) HighlightErr(m string) string {
	return t.Highlight(m, ColorRed)
}
