package ui

import "github.com/charmbracelet/lipgloss"

var (
	CBackground  = lipgloss.Color("#282A36")
	CCurrentLine = lipgloss.Color("#44475a")
	CForeground  = lipgloss.Color("#F8F8F2")
	CComment     = lipgloss.Color("#6272A4")
	CCyan        = lipgloss.Color("#8BE9FD")
	CPurple      = lipgloss.Color("#BD93F9")
	CRed         = lipgloss.Color("#FF5555")
)

var BlockColors = []lipgloss.Color{
	CCurrentLine,
	CCyan, CPurple, lipgloss.Color("#FFB86C"), lipgloss.Color("#F1FA8C"),
	lipgloss.Color("#50FA7B"), lipgloss.Color("#FF79C6"), CRed,
}

var BoardStyle = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder()).
	BorderForeground(CPurple).
	BorderBackground(CBackground).
	Background(CBackground).
	Padding(0, 0)

var StatsBoxStyle = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder()).
	BorderForeground(CComment).
	BorderBackground(CBackground).
	Background(CBackground)

func RenderBlock(colorID int, x, y int) string {
	var color lipgloss.Color
	if colorID > 0 && colorID < len(BlockColors) {
		color = BlockColors[colorID]
	} else {

		return lipgloss.NewStyle().
			Width(2).
			Background(CBackground).
			Foreground(CCurrentLine).
			Render(" .")
	}

	return lipgloss.NewStyle().
		Width(2).
		Background(color).
		Render("  ")
}

func RenderPreviewBlock(colorID int) string {
	if colorID > 0 && colorID < len(BlockColors) {
		return lipgloss.NewStyle().Width(2).Background(BlockColors[colorID]).Render("  ")
	}
	return lipgloss.NewStyle().Width(2).Background(CBackground).Render("  ")
}
