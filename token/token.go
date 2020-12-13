package token

import "fmt"

// A Token represents a lexical token used by Pekolang.
type Token int

// Available tokens.
const (
	EOF Token = iota
	EOL
	Start
	TapeFwd
	TapeRwd
	ByteInc
	ByteDec
	ByteRead
	ByteWrite
	LoopBegin
	LoopEnd
)

// A Position represents an arbitrary source position.
type Position struct {
	Line   int
	Column int
}

// String returns the string representation of the position.
func (pos Position) String() string {
	return fmt.Sprintf("line %d: column %d", pos.Line, pos.Column)
}
