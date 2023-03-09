package tictactoe

import (
	"crypto/rand"
	"math/big"
)

func AI(board Board, player byte) (x byte, y byte) {
	nextPlayer := player%2 + 1

	// Check if the current player can win
	for _, X := range []byte{0, 1, 2} {
		for _, Y := range []byte{0, 1, 2} {
			if board.Get(X, Y) == 0 {
				b, _ := board.Set(X, Y, player)
				if id, _ := b.Status(); id == player {
					return X, Y
				}
			}
		}
	}

	// Check if the current player can avoid a loss
	for X := 0; X < 3; X++ {
		for Y := 0; Y < 3; Y++ {
			if board[Y*3+X] == 0 {
				board[Y*3+X] = player
				if id, _ := board.Status(); id == nextPlayer {
					board[Y*3+X] = 0
					return byte(Y*3 + X)
				}
				board[Y*3+X] = 0
			}
		}
	}

	// Pick a random field
	fields := make([]byte, 0)
	for X := 0; X < 3; X++ {
		for Y := 0; Y < 3; Y++ {
			if board[Y*3+X] == 0 {
				fields = append(fields, byte(Y*3+X))
			}
		}
	}
	l := int64(len(fields))
	index, _ := rand.Int(rand.Reader, big.NewInt(l))
	return byte(index.Int64())
}
