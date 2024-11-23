package parser

import (
	"github.com/lindeneg/blue/lang/ast"
	"github.com/lindeneg/blue/lang/token"
)

type infixFn func(*P, ast.Expression) ast.Expression

var infixMap = map[token.Type]infixFn{}
