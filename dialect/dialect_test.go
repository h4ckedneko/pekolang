package dialect_test

import (
	"testing"

	"github.com/h4ckedneko/pekolang/dialect"
	"github.com/h4ckedneko/pekolang/token"
)

func TestLookupPekolang(t *testing.T) {
	tests := map[string]token.Token{
		"AH↓":        token.Begin,
		"HA→":        token.TapeFwd,
		"HA←":        token.TapeRwd,
		"HA↑":        token.ByteInc,
		"HA↓":        token.ByteDec,
		"HA↙":        token.ByteRead,
		"HA↘":        token.ByteWrite,
		"HA↗":        token.LoopBegin,
		"HA↖":        token.LoopEnd,
		"PE↗KO↘":     token.Comment,
		"Oreeeee!":   token.Comment,
		"Hey Moona!": token.Comment,
	}
	for ident, tok := range tests {
		tok2 := dialect.Lookup(dialect.Pekolang, ident)
		if tok != tok2 {
			t.Errorf("expected %q but got %q", tok, tok2)
		}
	}
}

func TestLookupBrainfuck(t *testing.T) {
	tests := map[string]token.Token{
		">": token.TapeFwd,
		"<": token.TapeRwd,
		"+": token.ByteInc,
		"-": token.ByteDec,
		",": token.ByteRead,
		".": token.ByteWrite,
		"[": token.LoopBegin,
		"]": token.LoopEnd,
		";": token.Comment,
		"a": token.Comment,
		"A": token.Comment,
	}
	for ident, tok := range tests {
		tok2 := dialect.Lookup(dialect.Brainfuck, ident)
		if tok != tok2 {
			t.Errorf("expected %q but got %q", tok, tok2)
		}
	}
}
