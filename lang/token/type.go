package token

// Type is all tokens currently understood
type Type int

const (
	UNKNOWN    Type = iota // unknown token
	EOF                    // end of file
	IDENTIFIER             // x, y, foobar etc..
	INT                    // 1, 2 etc..
	FLOAT                  // 1.5, 2.3 etc..
	STRING                 // "hello"
	ASSIGN                 // =
	PLUS                   // +
	MINUS                  // -
	BANG                   // !
	STAR                   // *
	FSLASH                 // /
	COLON                  // :
	SCOLON                 // ;
	COMMA                  // ,
	LT                     // <
	GT                     // >
	EQ                     // ==
	NEQ                    // !=
	LTOE                   // <=
	GTOE                   // >=
	AND                    // &&
	OR                     // ||
	LPAREN                 // (
	RPAREN                 // )
	LBRACE                 // {
	RBRACE                 // }
	LBRACKET               // [
	RBRACKET               // ]
	DEF                    // def keyword
	LET                    // let keyword
	CONST                  // const keyword
	TRUE                   // true keyword
	FALSE                  // false keyword
	IF                     // if keyword
	ELIF                   // elif keyword
	ELSE                   // else keyword
	RETURN                 // return keyword
	FOR                    // for keyword
	NULL                   // null keyword
)
