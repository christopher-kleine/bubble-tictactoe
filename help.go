package tictactoe

import tea "github.com/charmbracelet/bubbletea"

type Help struct {
	Parent tea.Model
}

func (h Help) Init() tea.Cmd {
	return nil
}

func (h Help) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tea.KeyMsg:
		return h.Parent, nil
	}

	return h, nil
}

func (h Help) View() string {
	return `
   TicTacToe - Help
 ════════════════════

 Rules:
 Two players place their piece ("X" or "O") on the board.
 Goal is to get 3 pieces in a row. A "row" is vertical, horizontal or diagonal.
 A player can't place a piece into a cell that's already occupied.
 Also: Once a piece was placed into a cell, it can't be moved anymore.
 Can no player make a valid move, but no player got 3 pieces in a row,
 the game ends in a draw.

 Keys:
 <Arrow Keys>  = Move your cursor across the board
 <Enter>       = Place your piece into the current cell if possible
 <t>           = Change the board layout
 <CTRL+C>, <q> = Quit the game
 <h>           = Show this help
 <m>           = Let the A.I. make a move

 Press any key to return to the game`
}

