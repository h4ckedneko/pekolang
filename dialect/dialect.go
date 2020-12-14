package dialect

import "github.com/h4ckedneko/pekolang/token"

// Pekolang is the dialect for Pekolang source code.
var Pekolang = map[token.Token]string{
	token.Start:     "AH↓",
	token.TapeFwd:   "HA→",
	token.TapeRwd:   "HA←",
	token.ByteInc:   "HA↑",
	token.ByteDec:   "HA↓",
	token.ByteRead:  "HA↙",
	token.ByteWrite: "HA↘",
	token.LoopBegin: "HA↗",
	token.LoopEnd:   "HA↖",
}

// Brainfuck is the dialect for Brainfuck source code.
var Brainfuck = map[token.Token]string{
	token.TapeFwd:   ">",
	token.TapeRwd:   "<",
	token.ByteInc:   "+",
	token.ByteDec:   "-",
	token.ByteRead:  ",",
	token.ByteWrite: ".",
	token.LoopBegin: "[",
	token.LoopEnd:   "]",
}
