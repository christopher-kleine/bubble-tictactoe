package tictactoe

import "math/rand"

type Board [9]byte

func (b Board) Free() []byte {
	list := make([]byte, 0)
	for index, v := range b {
		if v == 0 {
			list = append(list, byte(index))
		}
	}

	return list
}

func (b Board) IsFull() bool {
	return len(b.Free()) == 0
}

func (b Board) CanWin(Player byte) (byte, byte, bool) {
	for _, x := range []byte{0, 1, 2} {
		for _, y := range []byte{0, 1, 2} {
			board, ok := b.Set(x, y, Player)
			if !ok {
				continue
			}

			if _, id := board.Status(); id == Player {
				return x, y, true
			}
		}
	}

	return 0, 0, false
}

func (b Board) WouldLose(Player byte) (byte, byte, bool) {
	opponent := (Player % 2) + 1
	for _, x := range []byte{0, 1, 2} {
		for _, y := range []byte{0, 1, 2} {
			board, ok := b.Set(x, y, opponent)
			if !ok {
				continue
			}

			if _, id := board.Status(); id == opponent {
				return x, y, true
			}
		}
	}

	return 0, 0, false
}

func (b Board) AI(Player byte) (byte, byte, bool) {
	x, y, ok := b.CanWin(Player)
	if ok {
		return x, y, true
	}

	x, y, ok = b.WouldLose(Player)
	if ok {
		return x, y, true
	}

	free := b.Free()
	if len(free) == 0 {
		return 0, 0, false
	}

	index := byte(rand.Intn(len(free)))
	x = free[index] % 3
	y = free[index] / 3

	return x, y, true
}

func (b Board) Get(X, Y byte) byte {
	return b[Y*3+X]
}

func (b Board) Set(X, Y, Player byte) (Board, bool) {
	if X > 2 || Y > 2 || b[Y*3+X] != 0 {
		return b, false
	}

	b[Y*3+X] = Player

	return b, true
}

func (b Board) Row(X1, Y1, X2, Y2, X3, Y3 byte) (byte, []byte) {
	f1 := b.Get(X1, Y1)
	f2 := b.Get(X2, Y2)
	f3 := b.Get(X3, Y3)
	if f1 == f2 && f2 == f3 && f1 != 0 {
		return f1, []byte{Y1*3 + X1, Y2*3 + X2, Y3*3 + X3}
	}
	return 0, nil
}

func (b Board) Status() (Board, byte) {
	for _, y := range []byte{0, 1, 2} {
		if id, pos := b.Row(0, y, 1, y, 2, y); id != 0 {
			b[pos[0]] = id + 2
			b[pos[1]] = id + 2
			b[pos[2]] = id + 2
			return b, id
		}
	}

	for _, x := range []byte{0, 1, 2} {
		if id, pos := b.Row(x, 0, x, 1, x, 2); id != 0 {
			b[pos[0]] = id + 2
			b[pos[1]] = id + 2
			b[pos[2]] = id + 2
			return b, id
		}
	}

	if id, pos := b.Row(0, 0, 1, 1, 2, 2); id != 0 {
		b[pos[0]] = id + 2
		b[pos[1]] = id + 2
		b[pos[2]] = id + 2
		return b, id
	}

	if id, pos := b.Row(2, 0, 1, 1, 0, 2); id != 0 {
		b[pos[0]] = id + 2
		b[pos[1]] = id + 2
		b[pos[2]] = id + 2
		return b, id
	}

	if b.IsFull() {
		return b, 3
	}

	return b, 0
}

func (b Board) Copy() Board {
	return b
}
