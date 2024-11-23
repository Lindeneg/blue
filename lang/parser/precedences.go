package parser

import "github.com/lindeneg/blue/lang/token"

// pred is a type that describes token precedence
type pred int

const (
	_ pred = iota
	LOWEST
	EQUALS      // ==
	LESSGREATER // > or <
	SUM         // +
	PRODUCT     // *
	PREFIX      // -X or !X
	CALL        // myFunction(X)
	INDEX       // array[index]
)

// lt checks if pred p has less predcedence than token t
func (p pred) lt(t token.T) bool {
	return p < predMap.find(t)
}

// PredMap maps token.Type to its respective precedecnce
type PredMap map[token.Type]pred

var predMap = PredMap{
	token.EQ:       EQUALS,
	token.NEQ:      EQUALS,
	token.LT:       LESSGREATER,
	token.GT:       LESSGREATER,
	token.LTOE:     LESSGREATER,
	token.GTOE:     LESSGREATER,
	token.PLUS:     SUM,
	token.MINUS:    SUM,
	token.FSLASH:   PRODUCT,
	token.STAR:     PRODUCT,
	token.LPAREN:   CALL,
	token.LBRACKET: INDEX,
}

// find searches PredMap for precedence of token t
func (p PredMap) find(t token.T) pred {
	if p, ok := predMap[t.Type]; ok {
		return p
	}
	return LOWEST
}
