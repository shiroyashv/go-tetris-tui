package ui

import (
	"fmt"
	"time"

	"github.com/shiroyashv/go-tetris-tui/internal/config"
	"github.com/shiroyashv/go-tetris-tui/internal/game"

	tea "github.com/charmbracelet/bubbletea"
)

type tickMsg time.Time

type Model struct {
	Game *game.Game
}

func NewModel() Model {
	return Model{
		Game: game.NewGame(),
	}
}

func tickCmd() tea.Cmd {
	return tea.Tick(time.Millisecond*500, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (m Model) Init() tea.Cmd {
	return tickCmd()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "left":
			m.Game.MoveLeft()
		case "right":
			m.Game.MoveRight()
		case "up":
			m.Game.Rotate()
		case "down":
			m.Game.Update()
		}

	case tickMsg:
		m.Game.Update()
		return m, tickCmd()
	}

	return m, nil
}

func (m Model) View() string {
	s := fmt.Sprintf("Score: %d\n\n", m.Game.Score)

	for y := 0; y < config.BoardHeight; y++ {

		s += "<!"

		for x := 0; x < config.BoardWidth; x++ {
			isPiece := false

			pX := x - m.Game.Piece.X
			pY := y - m.Game.Piece.Y

			if pX >= 0 && pX < len(m.Game.Piece.Shape[0]) &&
				pY >= 0 && pY < len(m.Game.Piece.Shape) {

				if m.Game.Piece.Shape[pY][pX] == 1 {
					isPiece = true
				}
			}

			if isPiece {
				s += "[]"
			} else if m.Game.Grid[y][x] == 1 {
				s += "##"
			} else {
				s += " ."
			}
		}

		s += "!>\n"
	}

	s += "<!====================!>\n"
	s += "\nPress 'q' to quit."

	return s
}
