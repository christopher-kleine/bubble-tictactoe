package tictactoe

type Board [9]byte

func (b Board) IsFull() bool {
	for _, v := range b {
		if v == 0 {
			return false
		}
	}
	return true
}

func (b Board) CanWin(Player byte) (byte, byte, bool) {
	for _, x := range []byte{0, 1, 2} {
		for _, y := range []byte{0, 1, 2} {
			board, _ := b.Set(x, y, Player)
			if id, _ := board.Status(); id == Player {
				return x, y, true
			}
		}
	}

	return 0, 0, false
}

func (b Board) Get(X, Y byte) byte {
	return b[Y*3+X]
}

func (b Board) Set(X, Y, Player byte) (board Board, ok bool) {
	if X > 2 || Y > 2 || b[Y*3+X] != 0 {
		return b, false
	}

	board = b
	board[Y*3+X] = Player

	return board, true
}

func (b Board) Row(X1, Y1, X2, Y2, X3, Y3 byte) (byte, []byte) {
	if b[Y1*3+X1] == b[Y2*3+X2] && b[Y2*3+X2] == b[Y3*3+X3] && b[Y1*3+X1] != 0 {
		return b[Y1*3*X1], []byte{Y1*3 + X1, Y2*3 + X2, Y3*3 + X3}
	}
	return 0, nil
}

func (b Board) Status() (byte, []byte) {
	for _, y := range []byte{0, 1, 2} {
		if id, pos := b.Row(0, y, 1, y, 2, y); id != 0 {
			return id, pos
		}
	}

	for _, x := range []byte{0, 1, 2} {
		if id, pos := b.Row(x, 0, x, 1, x, 2); id != 0 {
			return id, pos
		}
	}

	if id, pos := b.Row(0, 0, 1, 1, 2, 2); id != 0 {
		return id, pos
	}

	if id, pos := b.Row(0, 2, 1, 1, 2, 0); id != 0 {
		return id, pos
	}

	if b.IsFull() {
		return 3, nil
	}

	return 0, nil
}
