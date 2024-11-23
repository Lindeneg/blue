package parser

import (
	"fmt"

	"github.com/lindeneg/blue/lang/token"
)

// ParseErr describes an error encountered during parsing
type ParseErr struct {
	token.T
	Msg  string
	Line string
}

// Err formats an error with sourceName, line, col and message.
func newParseErr(p *P, t token.T, msg string, args ...any) ParseErr {
	l := t.HighlightErr(p.l.Line(t.Line))
	m := fmt.Sprintf(msg, args...)
	m = fmt.Sprintf("ParseError: %s at\n\t%s:L%d:C%d ------> %s",
		m, p.sourceName, t.Line, t.Col, l)
	return ParseErr{T: t, Msg: m, Line: l}
}

func err(p *P, t token.T, msg string, args ...any) {
	p.errs = append(p.errs, newParseErr(p, t, msg, args...))
}

func expectErr(p *P, want token.T, got token.Type) {
	err(p, want, "unexpected token, got=%q, want=%q", got, want.Type)
}

func parseFnErr(p *P, ty string, t token.T) {
	err(p, t, "no %q function found for token %q", ty, t.Type)
}
