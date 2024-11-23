package token

import (
	"fmt"
	"strings"
)

type Color string

const (
	COLOR_RED   Color = "\033[31m"
	COLOR_RESET Color = "\033[0m"
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

// T describes a token encoutered during lexical analysis.
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

func (t T) String() string {
	return fmt.Sprintf("Type=%d|Literal=%q|Line=%d|Col=%d|Scope=%d",
		t.Type, t.Literal, t.Line, t.Col, t.Scope)
}

func (t T) Highlight(m string, color Color) string {
	if m == "" {
		return m
	}
	errorTokenEscaped := strings.ReplaceAll(t.Literal, "\n", "\\n")
	errorTokenEscaped = strings.ReplaceAll(errorTokenEscaped, "\t", "\\t")
	coloredToken := fmt.Sprintf("%s%s%s", color, errorTokenEscaped, COLOR_RESET)
	return strings.Replace(m, t.Literal, coloredToken, 1)
}

func (t T) HighlightErr(m string) string {
	return t.Highlight(m, COLOR_RED)
}
