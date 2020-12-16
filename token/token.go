//go:generate stringer -type=Token

package token

import "fmt"

// A Token represents a lexical token used by Pekolang.
type Token int

// Available tokens.
const (
	EOF Token = iota
	EOL
	Comment

	Begin
	TapeFwd
	TapeRwd
	ByteInc
	ByteDec
	ByteRead
	ByteWrite
	LoopBegin
	LoopEnd
)

var instructions = map[string]Token{
	"AH↓": Begin,
	"HA→": TapeFwd,
	"HA←": TapeRwd,
	"HA↑": ByteInc,
	"HA↓": ByteDec,
	"HA↙": ByteRead,
	"HA↘": ByteWrite,
	"HA↗": LoopBegin,
	"HA↖": LoopEnd,
}

// Lookup returns a token identified by instruction.
func Lookup(ident string) Token {
	if tok, ok := instructions[ident]; ok {
		return tok
	}
	return Comment
}

// A Position represents an arbitrary source position.
type Position struct {
	Line   int
	Column int
}

// String returns the string representation of the position.
func (pos Position) String() string {
	return fmt.Sprintf("line %d: column %d", pos.Line, pos.Column)
}
