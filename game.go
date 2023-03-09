package tictactoe

import (
	"os"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
)

var (
	templates = []string{
		`
 ┏━━━┳━━━┳━━━┓
 ┃ $1 ┃ $2 ┃ $3 ┃
 ┣━━━╋━━━╋━━━┫
 ┃ $4 ┃ $5 ┃ $6 ┃
 ┣━━━╋━━━╋━━━┫
 ┃ $7 ┃ $8 ┃ $9 ┃
 ┗━━━┻━━━┻━━━┛`,
		`
 ╔═══╦═══╦═══╗
 ║ $1 ║ $2 ║ $3 ║
 ╠═══╬═══╬═══╣
 ║ $4 ║ $5 ║ $6 ║
 ╠═══╬═══╬═══╣
 ║ $7 ║ $8 ║ $9 ║
 ╚═══╩═══╩═══╝`,
		`
    ╻   ╻    
  $1 ┃ $2 ┃ $3 
 ╺━━╋━━━╋━━━╸
  $4 ┃ $5 ┃ $6 
 ╺━━╋━━━╋━━━╸
  $7 ┃ $8 ┃ $9 
    ╹   ╹    `,
		`
    ╥   ╥    
  $1 ║ $2 ║ $3 
 ╞══╬═══╬═══╡
  $4 ║ $5 ║ $6 
 ╞══╬═══╬═══╡
  $7 ║ $8 ║ $9 
    ╨   ╨    `,
	}
	chars = []string{
		" ",
		"O",
		"X",
		"\x1B[32mO\x1B[39m",
		"\x1B[32mX\x1B[39m",
		"\x1B[31mO\x1B[39m",
		"\x1B[31mX\x1B[39m",
	}
)

func New(user string, parent tea.Model) tea.Model {
	model := Model{
		Parent:   parent,
		User:     user,
		Template: 0,
		Board:    [9]byte{0, 0, 0, 0, 0, 0, 0, 0, 0},
		Side:     1,
	}

	return model
}

type Model struct {
	Parent   tea.Model
	User     string
	Board    [9]byte
	Template int
	X, Y     int
	Side     byte
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			if m.Parent == nil {
				return m, tea.Quit
			} else {
				return m.Parent, nil
			}

		case "h":
			return Help{ Parent: m }, nil

		case "t":
			m.Template = (m.Template + 1) % len(templates)

		case "enter":
			if m.Board[m.Y*3+m.X] == 0 {
				m.Board[m.Y*3+m.X] = m.Side
				m.Side = m.Side % 2 + 1
			}

		case "m":
			// Call the A.I.

		case "up":
			if m.Y <= 0 {
				break
			}
			m.Y--

		case "down":
			if m.Y >= 2 {
				break
			}
			m.Y++

		case "left":
			if m.X <= 0 {
				break
			}
			m.X--

		case "right":
			if m.X >= 2 {
				break
			}
			m.X++
		}
	}

	return m, nil
}

func (m Model) View() string {
	board := m.Board[:]
	board[m.Y*3+m.X] = m.Side + 4
	return `
  TicTacToe
═════════════

Press <h> to show the help.
` + os.Expand(templates[m.Template], func(s string) string {
		index, err := strconv.Atoi(s)
		if err != nil {
			return s
		}

		if index >= 1 && index <= 9 {
			return chars[board[index-1]]
		}

		return s
	})
}
