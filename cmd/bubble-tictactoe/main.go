package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	tictactoe "github.com/christopher-kleine/bubble-tictactoe"
)

func main() {
	model := tictactoe.New("Demo", nil)
	p := tea.NewProgram(model)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Whoops! An error occurred: %v\n", err)
		os.Exit(1)
	}
}

