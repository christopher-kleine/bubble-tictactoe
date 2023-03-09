package tictactoe

import (
	"crypto/rand"
	"math/big"
)

func AI(board []byte, side byte) byte {
	nextPlayer := side%2 + 1

	// Check if the current player can win
	for X := 0; X < 3; X++ {
		for Y := 0; Y < 3; Y++ {
			if board[Y*3+X] == 0 {
				board[Y*3+X] = side
				if id, _ := CheckStatus(board); id == side {
					board[Y*3+X] = 0
					return byte(Y*3 + X)
				}
				board[Y*3+X] = 0
			}
		}
	}

	// Check if the current player can avoid a loss
	for X := 0; X < 3; X++ {
		for Y := 0; Y < 3; Y++ {
			if board[Y*3+X] == 0 {
				board[Y*3+X] = side
				if id, _ := CheckStatus(board); id == nextPlayer {
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
