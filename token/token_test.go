package token_test

import (
	"testing"

	"github.com/h4ckedneko/pekolang/token"
)

func TestLookup(t *testing.T) {
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
		tok2 := token.Lookup(ident)
		if tok != tok2 {
			t.Errorf("expected %q but got %q", tok, tok2)
		}
	}
}
