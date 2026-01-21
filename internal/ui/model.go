package ui

import (
	"fmt"
	"time"

	"github.com/shiroyashv/go-tetris-tui/internal/config"
	"github.com/shiroyashv/go-tetris-tui/internal/game"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type tickMsg time.Time

type Model struct {
	Game      *game.Game
	WinWidth  int
	WinHeight int
}

func NewModel() Model {
	return Model{
		Game: game.NewGame(),
	}
}

func tickCmd(duration time.Duration) tea.Cmd {
	return tea.Tick(duration, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (m Model) Init() tea.Cmd {
	return tickCmd(m.Game.TickRate)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.WinWidth = msg.Width
		m.WinHeight = msg.Height
		return m, nil

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
		return m, tickCmd(m.Game.TickRate)
	}

	return m, nil
}

func (m Model) View() string {
	// Safety check for small windows
	if m.WinWidth > 0 && (m.WinWidth < 40 || m.WinHeight < 24) {
		return lipgloss.Place(m.WinWidth, m.WinHeight, lipgloss.Center, lipgloss.Center, "Window too small!")
	}

	// 1. GAME OVER VIEW
	if m.Game.GameOver {
		title := HeaderStyle.Render("GAME OVER")
		score := fmt.Sprintf("Final Score\n%d", m.Game.Score)
		help := LabelStyle.Render("Press 'q' to quit")

		content := lipgloss.JoinVertical(lipgloss.Center, title, "\n", score, "\n\n", help)
		return AppStyle.Render(lipgloss.Place(m.WinWidth, m.WinHeight, lipgloss.Center, lipgloss.Center, content))
	}

	// 2. RENDER BOARD
	var boardView string
	for y := 0; y < config.BoardHeight; y++ {
		for x := 0; x < config.BoardWidth; x++ {
			color := 0

			// Check active piece
			pX := x - m.Game.Piece.X
			pY := y - m.Game.Piece.Y
			if pX >= 0 && pX < len(m.Game.Piece.Shape[0]) &&
				pY >= 0 && pY < len(m.Game.Piece.Shape) {
				if m.Game.Piece.Shape[pY][pX] == 1 {
					color = m.Game.Piece.Color
				}
			}

			// Check grid
			if color == 0 {
				color = m.Game.Grid[y][x]
			}

			boardView += renderBlock(color, x, y)
		}
		if y < config.BoardHeight-1 {
			boardView += "\n"
		}
	}
	boardBox := BoardStyle.Render(boardView)

	// 3. RENDER SIDEBAR (Stats)
	scoreBlock := lipgloss.JoinVertical(lipgloss.Left,
		LabelStyle.Render("SCORE"),
		ValueStyle.Render(fmt.Sprintf("%d", m.Game.Score)),
	)

	// Calculate level based on TickRate (just for display fun)
	level := 1 + (800-int(m.Game.TickRate.Milliseconds()))/20
	levelBlock := lipgloss.JoinVertical(lipgloss.Left,
		LabelStyle.Render("LEVEL"),
		ValueStyle.Render(fmt.Sprintf("%d", level)),
	)

	controls := LabelStyle.Render("CONTROLS\n\n←/→ Move\n↑   Rotate\n↓   Drop\nq   Quit")

	statsBox := StatsBoxStyle.Render(
		lipgloss.JoinVertical(lipgloss.Left,
			HeaderStyle.Render("TETRIS"),
			"\n",
			scoreBlock,
			"\n",
			levelBlock,
			"\n\n",
			controls,
		),
	)

	// 4. COMBINE
	mainLayout := lipgloss.JoinHorizontal(lipgloss.Top, boardBox, statsBox)

	return AppStyle.Render(
		lipgloss.Place(
			m.WinWidth,
			m.WinHeight,
			lipgloss.Center,
			lipgloss.Center,
			mainLayout,
		),
	)
}