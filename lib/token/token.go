package token

import "fmt"

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
	Col  int
	Scope
}

func (t T) String() string {
	return fmt.Sprintf("Type=%d|Literal=%q|Line=%d|Col=%d|Scope=%d",
		t.Type, t.Literal, t.Line, t.Col, t.Scope)
}
