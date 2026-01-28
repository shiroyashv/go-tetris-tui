package ui

import (
	"fmt"
	"strings"
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
	return Model{Game: game.NewGame()}
}

func tickCmd(duration time.Duration) tea.Cmd {
	return tea.Tick(duration, func(t time.Time) tea.Msg { return tickMsg(t) })
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
		case "q", "Q", "й", "Й", "ctrl+c":
			return m, tea.Quit
		case "r", "R", "к", "К":
			if m.Game.GameOver {
				m.Game = game.NewGame()
			} else if !m.Game.ConfirmRestart {
				m.Game.Paused = true
				m.Game.ConfirmRestart = true
			}
			return m, nil
		case "y", "Y", "н", "Н":
			if m.Game.ConfirmRestart {
				m.Game = game.NewGame()
			}
			return m, nil
		case "n", "N", "т", "Т":
			if m.Game.ConfirmRestart {
				m.Game.ConfirmRestart = false
				m.Game.Paused = false
			}
			return m, nil
		case "p", "P", "з", "З":
			if !m.Game.GameOver && !m.Game.ConfirmRestart {
				m.Game.Paused = !m.Game.Paused
			}
			return m, nil
		}

		if m.Game.Paused || m.Game.GameOver || m.Game.ConfirmRestart {
			return m, nil
		}
		switch msg.String() {
		case "left":
			m.Game.MoveLeft()
		case "right":
			m.Game.MoveRight()
		case "up":
			m.Game.Rotate()
		case "down":
			m.Game.Update()
		case " ":
			m.Game.HardDrop()
		}
	case tickMsg:
		if !m.Game.Paused && !m.Game.GameOver {
			m.Game.Update()
		}
		return m, tickCmd(m.Game.TickRate)
	}
	return m, nil
}

func renderFullWidth(text string, width int, align lipgloss.Position) string {
	return lipgloss.NewStyle().
		Width(width).
		Background(CBackground).
		Foreground(CForeground).
		Align(align).
		Render(text)
}

func (m Model) View() string {
	if m.WinWidth > 0 && (m.WinWidth < 45 || m.WinHeight < 24) {
		return lipgloss.Place(m.WinWidth, m.WinHeight, lipgloss.Center, lipgloss.Center, "Window too small!",
			lipgloss.WithWhitespaceBackground(CBackground))
	}

	lblStyle := lipgloss.NewStyle().Foreground(CComment).Background(CBackground).Bold(true)
	valStyle := lipgloss.NewStyle().Foreground(CCyan).Background(CBackground)
	dotStyle := lipgloss.NewStyle().Foreground(CCurrentLine).Background(CBackground)

	overlayStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#ffffffff")).Bold(true)

	var overlayLines []string

	centerLine := func(text string, width int) string {
		if len(text) >= width {
			return text
		}
		padding := width - len(text)
		lPad := padding / 2
		rPad := padding - lPad
		return strings.Repeat(" ", lPad) + text + strings.Repeat(" ", rPad)
	}

	if m.Game.GameOver {
		overlayLines = []string{
			centerLine("GAME OVER", config.BoardWidth),
			centerLine("", config.BoardWidth),
			centerLine(fmt.Sprintf("SCORE: %d", m.Game.Score), config.BoardWidth),
			centerLine("", config.BoardWidth),
			centerLine("PRESS 'r'", config.BoardWidth),
			centerLine("TO RESTART", config.BoardWidth),
			centerLine("", config.BoardWidth),
			centerLine("PRESS 'q'", config.BoardWidth),
		}
		} else if m.Game.ConfirmRestart { 
			overlayLines = []string{
				centerLine("RESTART?", config.BoardWidth),
				centerLine("", config.BoardWidth),
				centerLine("ARE YOU", config.BoardWidth),
				centerLine("SURE?", config.BoardWidth),
				centerLine("", config.BoardWidth),
				centerLine("y / n", config.BoardWidth),
		} 
		} else if m.Game.Paused {
		overlayLines = []string{
			centerLine("PAUSED", config.BoardWidth),
			centerLine("", config.BoardWidth),
			centerLine("PRESS 'p'", config.BoardWidth),
		}
	}

	overlayStartY := (config.BoardHeight - len(overlayLines)) / 2

	var boardView string
	for y := 0; y < config.BoardHeight; y++ {
		for x := 0; x < config.BoardWidth; x++ {

			color := 0
			pX := x - m.Game.Piece.X
			pY := y - m.Game.Piece.Y
			if pX >= 0 && pX < len(m.Game.Piece.Shape[0]) && pY >= 0 && pY < len(m.Game.Piece.Shape) {
				if m.Game.Piece.Shape[pY][pX] == 1 {
					color = m.Game.Piece.Color
				}
			}
			if color == 0 {
				color = m.Game.Grid[y][x]
			}

			var bg lipgloss.Color
			if color > 0 && color < len(BlockColors) {
				bg = BlockColors[color]
			} else {
				bg = CBackground
			}

			charToRender := ""
			if len(overlayLines) > 0 {
				lineIndex := y - overlayStartY
				if lineIndex >= 0 && lineIndex < len(overlayLines) {
					line := overlayLines[lineIndex]
					if x < len(line) {
						letter := string(line[x])
						if letter != " " {
							charToRender = letter
						}
					}
				}
			}

			if charToRender != "" {
				boardView += overlayStyle.
					Background(bg).
					Render(fmt.Sprintf(" %s", charToRender))
			} else if color > 0 {

				boardView += lipgloss.NewStyle().Background(bg).Width(2).Render("  ")
			} else {
				boardView += dotStyle.Render(" .")
			}
		}
		if y < config.BoardHeight-1 {
			boardView += "\n"
		}
	}
	boardBox := BoardStyle.Render(boardView)

	const statsW = 22
	var nextView string
	nextShape := m.Game.NextPiece.Shape
	for r := 0; r < len(nextShape); r++ {
		isEmpty := true
		for c := 0; c < len(nextShape[r]); c++ {
			if nextShape[r][c] == 1 {
				isEmpty = false
				break
			}
		}
		if isEmpty {
			continue
		}
		rowStr := ""
		for c := 0; c < len(nextShape[r]); c++ {
			if nextShape[r][c] == 1 {
				rowStr += RenderPreviewBlock(m.Game.NextPiece.Color)
			} else {
				rowStr += RenderPreviewBlock(0)
			}
		}
		if nextView != "" {
			nextView += "\n"
		}
		nextView += lipgloss.NewStyle().Width(statsW).Align(lipgloss.Center).Background(CBackground).Render(rowStr)
	}

	lines := []string{
		lipgloss.NewStyle().Width(statsW).Background(CPurple).Foreground(CBackground).Bold(true).Align(lipgloss.Center).Render("TETRIS GO"),
		renderFullWidth("", statsW, lipgloss.Left),
		renderFullWidth(lblStyle.Render("NEXT"), statsW, lipgloss.Left),
		lipgloss.NewStyle().Height(2).Align(lipgloss.Center).Width(statsW).Background(CBackground).Render(nextView),
		renderFullWidth("", statsW, lipgloss.Left),
		renderFullWidth(lblStyle.Render("SCORE"), statsW, lipgloss.Left),
		renderFullWidth(valStyle.Render(fmt.Sprintf("%d", m.Game.Score)), statsW, lipgloss.Left),
		renderFullWidth("", statsW, lipgloss.Left),
		renderFullWidth(lblStyle.Render("LEVEL"), statsW, lipgloss.Left),
		renderFullWidth(valStyle.Render(fmt.Sprintf("%d", 1+(800-int(m.Game.TickRate.Milliseconds()))/50)), statsW, lipgloss.Left),
		renderFullWidth("", statsW, lipgloss.Left),
		renderFullWidth(lblStyle.Render("CONTROLS"), statsW, lipgloss.Left),
		renderFullWidth("←/→ Move", statsW, lipgloss.Left),
		renderFullWidth("↑   Rotate", statsW, lipgloss.Left),
		renderFullWidth("↓   Down", statsW, lipgloss.Left),
		renderFullWidth("Spc Hard Drop", statsW, lipgloss.Left),
		renderFullWidth("p   Pause", statsW, lipgloss.Left),
		renderFullWidth("r   Restart", statsW, lipgloss.Left),
		renderFullWidth("q   Quit", statsW, lipgloss.Left),
	}
	statsBox := StatsBoxStyle.Copy().Height(config.BoardHeight).Render(lipgloss.JoinVertical(lipgloss.Left, lines...))
	gap := lipgloss.NewStyle().Width(2).Height(config.BoardHeight + 2).Background(CBackground).Render("")

	return lipgloss.Place(m.WinWidth, m.WinHeight, lipgloss.Center, lipgloss.Center,
		lipgloss.JoinHorizontal(lipgloss.Top, boardBox, gap, statsBox),
		lipgloss.WithWhitespaceBackground(CBackground), lipgloss.WithWhitespaceForeground(CForeground),
	)
}
