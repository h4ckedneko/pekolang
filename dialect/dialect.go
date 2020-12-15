package dialect

import "github.com/h4ckedneko/pekolang/token"

// Pekolang is the dialect for Pekolang source code.
var Pekolang = map[string]token.Token{
	"AH↓": token.Begin,
	"HA→": token.TapeFwd,
	"HA←": token.TapeRwd,
	"HA↑": token.ByteInc,
	"HA↓": token.ByteDec,
	"HA↙": token.ByteRead,
	"HA↘": token.ByteWrite,
	"HA↗": token.LoopBegin,
	"HA↖": token.LoopEnd,
}

// Brainfuck is the dialect for Brainfuck source code.
var Brainfuck = map[string]token.Token{
	">": token.TapeFwd,
	"<": token.TapeRwd,
	"+": token.ByteInc,
	"-": token.ByteDec,
	",": token.ByteRead,
	".": token.ByteWrite,
	"[": token.LoopBegin,
	"]": token.LoopEnd,
}

// Lookup returns a token mapped to a keyword identification from a dialect.
func Lookup(dialect map[string]token.Token, ident string) token.Token {
	if tok, ok := dialect[ident]; ok {
		return tok
	}
	return token.Comment
}
