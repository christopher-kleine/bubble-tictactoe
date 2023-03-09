package tictactoe_test

import (
	"testing"

	tictactoe "github.com/christopher-kleine/bubble-tictactoe"
)

func TestSet(t *testing.T) {
	b := tictactoe.Board{0, 0, 0, 0, 0, 0, 0, 0, 0}
	if l := len(b); l != 9 {
		t.Errorf("Should be 9, got %d", l)
	}

	b1 := b.Set(0, 0, 1)
	if b[0] == b1[0] {
		t.Errorf("b should not be modified")
	}

	t.Logf("%v / %v", b, b1)
}
