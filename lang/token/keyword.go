package token

// keywords maps a literal string
// containing a keyword to that
// keyword's appropriate token.Type
var keywords = map[string]Type{
	"fn":     FN,
	"let":    LET,
	"const":  CONST,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"elif":   ELIF,
	"else":   ELSE,
	"return": RETURN,
	"for":    FOR,
	"null":   NULL,
}

// Identifier checks if an
// identifier is a keyword. If it is
// it returns that keyword's token.Type.
// If not, it will return token.IDENTIFIER
func Identifier(ident string) Type {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENTIFIER
}
