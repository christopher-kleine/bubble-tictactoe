package tictactoe

func CheckStatus(board []byte) (playerID byte, pos [3]byte) {
	for y := 0; y < 3; y++ {
		if id := board[y*3+0]; id != 0 && id == board[y*3+1] && id == board[y*3+2] {
			pos[0] = byte(y*3 + 0)
			pos[1] = byte(y*3 + 1)
			pos[2] = byte(y*3 + 2)
			return id, pos
		}
	}

	for x := 0; x < 3; x++ {
		if id := board[0*3+x]; id != 0 && id == board[1*3+x] && id == board[2*3+x] {
			pos[0] = byte(0*3 + x)
			pos[1] = byte(1*3 + x)
			pos[2] = byte(2*3 + x)
			return id, pos
		}
	}

	if id := board[0*3+0]; id != 0 && id == board[1*3+1] && id == board[2*3+2] {
		pos[0] = byte(0*3 + 0)
		pos[1] = byte(1*3 + 1)
		pos[2] = byte(2*3 + 2)
		return id, pos
	}

	if id := board[0*3+2]; id != 0 && id == board[1*3+1] && id == board[2*3+0] {
		pos[0] = byte(0*3 + 2)
		pos[1] = byte(1*3 + 1)
		pos[2] = byte(2*3 + 2)
		return id, pos
	}

	for index := 0; index <= 8; index++ {
		if board[index] == 0 {
			return 0, pos
		}
	}

	return 3, pos
}
