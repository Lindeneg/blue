package lexer

import (
	"github.com/lindeneg/blue/lang/assert"
	"github.com/lindeneg/blue/lang/token"
)

// L is responsible for lexical analysis.
type L struct {
	// source to be lexed
	source []byte
	// current index
	curIdx int
	// next index
	nextIdx int
	// current char
	char byte
	// current line
	line int
	// current column
	col int
	// scope nesting layer
	scope token.Scope
}

// New creates a new L struct and reads
// the first character of the source
func New(source []byte) *L {
	l := &L{source: source, line: 1}
	l.read()
	return l
}

// FromString calls New with
// source converted to []byte
func FromString(source string) *L {
	return New([]byte(source))
}

// NextToken returns the next token encountered in source
// and advances to the next position as well
func (l *L) NextToken() token.T {
	l.ignoreWhitespace()
	tok := l.slimToken()
	switch l.char {
	case 0:
		assert.A(!l.scope.Global(), "EOF found but scope is '%d'\n", l.scope)
		tok = l.token(token.EOF, "")
		return tok
	case '=':
		if l.peek() == '=' {
			tok = l.tokenRange(token.EQ, 1)
		} else {
			tok = l.token(token.ASSIGN, l.char)
		}
	case '<':
		if l.peek() == '=' {
			tok = l.tokenRange(token.LTOE, 1)
		} else {
			tok = l.token(token.LT, l.char)
		}
	case '>':
		if l.peek() == '=' {
			tok = l.tokenRange(token.GTOE, 1)
		} else {
			tok = l.token(token.GT, l.char)
		}
	case '!':
		if l.peek() == '=' {
			tok = l.tokenRange(token.NEQ, 1)
		} else {
			tok = l.token(token.BANG, l.char)
		}
	case '|':
		if l.peek() == '|' {
			tok = l.tokenRange(token.OR, 1)
		} else {
			tok = l.token(token.UNKNOWN, l.char)
		}
	case '&':
		if l.peek() == '&' {
			tok = l.tokenRange(token.AND, 1)
		} else {
			tok = l.token(token.UNKNOWN, l.char)
		}
	case ';':
		tok = l.token(token.SCOLON, l.char)
	case ',':
		tok = l.token(token.COMMA, l.char)
	case ':':
		tok = l.token(token.COLON, l.char)
	case '+':
		tok = l.token(token.PLUS, l.char)
	case '-':
		tok = l.token(token.MINUS, l.char)
	case '*':
		tok = l.token(token.STAR, l.char)
	case '/':
		if l.peek() == '/' {
			l.ignoreComment()
			return l.NextToken()
		} else {
			tok = l.token(token.FSLASH, l.char)
		}
	case '(':
		tok = l.token(token.LPAREN, l.char)
	case ')':
		tok = l.token(token.RPAREN, l.char)
	case '[':
		tok = l.token(token.LBRACKET, l.char)
	case ']':
		tok = l.token(token.RBRACKET, l.char)
	case '{':
		l.scope++
		tok = l.token(token.LBRACE, l.char)
	case '}':
		tok = l.token(token.RBRACE, l.char)
		l.scope--
	case '"':
		tok.Type = token.STRING
		tok.Literal = string(l.string('"'))
	default:
		return l.handleIdentifier(tok)
	}
	l.read()
	return tok
}

// Line returns the given line as a string
func (l *L) Line(line int) string {
	li := 1
	start := -1
	end := -1
	for i, b := range l.source {
		if li == line && start == -1 {
			start = i
		}
		if b == '\n' {
			if start > -1 {
				end = i
				break
			}
			li += 1
		}
	}
	if start < 0 {
		return ""
	}
	if end < 0 {
		return string(l.source[start:])
	}
	return string(l.source[start:end])
}

// handleIdentifier reads an identifier or a literal
func (l *L) handleIdentifier(tok token.T) token.T {
	if isIdentifierStart(l.char) {
		tok.Literal = string(l.identifier())
		tok.Type = token.Identifier(tok.Literal)
		return tok
	} else if isDigit(l.char) {
		tok.Literal = string(l.digit())
		if l.char == '.' && isDigit(l.peek()) {
			tok.Literal += string(l.char)
			l.read()
			tok.Literal += string(l.digit())
			tok.Type = token.FLOAT
			return tok
		}
		tok.Type = token.INT
		return tok
	}
	tok = l.token(token.UNKNOWN, l.char)
	l.read()
	return tok
}

// ignoreWhitespace ignores spaces, tabs, newlines and carriage return
func (l *L) ignoreWhitespace() {
	for l.char == ' ' || l.char == '\t' || l.char == '\n' || l.char == '\r' {
		l.read()
	}
}

// ignoreWhitespace ignores single-line comments
func (l *L) ignoreComment() {
	for l.char != 0 && l.char != '\n' {
		l.read()
	}
}

// tokenRange return a token with the next r characters
// appended to the literal and advance the position accordingly
func (l *L) tokenRange(tokenType token.Type, r int) token.T {
	tok := l.slimToken()
	tok.Type = tokenType
	tok.Literal = string(l.char)
	for i := 0; i < r; i++ {
		l.read()
		tok.Literal += string(l.char)
	}
	return tok
}

// read reads the next character,
// advances current and next indices
// and updates line and col ints
func (l *L) read() {
	if l.char == '\n' {
		l.line += 1
		l.col = 1
	} else {
		l.col += 1
	}
	if l.nextIdx >= len(l.source) {
		l.char = 0
	} else {
		l.char = l.source[l.nextIdx]
	}
	l.curIdx = l.nextIdx
	l.nextIdx += 1
}

// peek read char without incrementing the next position
func (l *L) peek() byte {
	if l.nextIdx >= len(l.source) {
		return 0
	}
	return l.source[l.nextIdx]
}

// string reads a string until terminator byte is seen
func (l *L) string(terminator byte) []byte {
	position := l.curIdx + 1
	for {
		l.read()
		if l.char == terminator || l.char == 0 {
			break
		}
	}
	return l.source[position:l.curIdx]
}

// read a number from current pos in input string
func (l *L) digit() []byte {
	return l.readWhile(isDigit)
}

// read an identifier from current pos in input string
func (l *L) identifier() []byte {
	return l.readWhile(isIdentifier)
}

// readWhile the callback returns true, returns the bytes read
func (l *L) readWhile(shouldContinue func(byte) bool) []byte {
	pos := l.curIdx
	for shouldContinue(l.char) {
		l.read()
	}
	return l.source[pos:l.curIdx]
}

// token creates a new token.T type with
// line, col and scope set
func (l *L) token(t token.Type, lt any) token.T {
	var lts string
	switch lt := lt.(type) {
	case string:
		lts = lt
	case []byte:
		lts = string(lt)
	case byte:
		lts = string(lt)
	default:
		assert.VerifyNotReached()
	}
	return token.T{Type: t, Literal: lts, Line: l.line, Col: l.col, Scope: l.scope}
}

// slimToken creates a new token.T type with
// line, col and scope set but without type and literal
func (l *L) slimToken() token.T {
	return token.T{Line: l.line, Col: l.col, Scope: l.scope}
}

// isIdentifierStart checks if a byte is a valid
// character for an identifier in the first position
func isIdentifierStart(char byte) bool {
	return char >= 'a' && char <= 'z' ||
		char >= 'A' && char <= 'Z' ||
		char == '_'
}

// isIdentifier checks if a byte is a valid
// character for an identifier in any position
func isIdentifier(char byte) bool {
	return isIdentifierStart(char) || (char >= '0' && char <= '9')
}

// isDigit checks if a byte is a digit
func isDigit(char byte) bool {
	return char >= '0' && char <= '9'
}
