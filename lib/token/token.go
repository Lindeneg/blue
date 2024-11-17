package token

// T describes a token encoutered during lexical analysis.
type T struct {
	// Type of token encountered
	Type
	// Literal value
	Literal string
	// TODO can probably be smarter about this
	// Line the token was encountered at
	Line int
	// Column the token was encountered at
	Col int
}
