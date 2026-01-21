package ui

import "github.com/charmbracelet/lipgloss"

// --- PALETTE (Dracula-inspired) ---
var (
	cBackground = lipgloss.Color("#282A36")
	cCurrentLine= lipgloss.Color("#44475a")
	cForeground = lipgloss.Color("#F8F8F2")
	cComment    = lipgloss.Color("#6272A4")
	cCyan       = lipgloss.Color("#8BE9FD")
	cGreen      = lipgloss.Color("#50FA7B")
	cOrange     = lipgloss.Color("#FFB86C")
	cPink       = lipgloss.Color("#FF79C6")
	cPurple     = lipgloss.Color("#BD93F9")
	cRed        = lipgloss.Color("#FF5555")
	cYellow     = lipgloss.Color("#F1FA8C")
)

var (
	// Block Colors mapped to ID
	blockColors = []lipgloss.Color{
		cCurrentLine, // 0: Empty (used for background dots)
		cCyan,        // 1: I
		cPurple,      // 2: J
		cOrange,      // 3: L
		cYellow,      // 4: O
		cGreen,       // 5: S
		cPink,        // 6: T
		cRed,         // 7: Z
	}
)

// --- STYLES ---

// AppStyle is the main container
var AppStyle = lipgloss.NewStyle().
	Padding(1, 2).
	Background(cBackground).
	Foreground(cForeground)

// HeaderStyle for the "TETRIS" title
var HeaderStyle = lipgloss.NewStyle().
	Foreground(cForeground).
	Background(cComment).
	Padding(0, 1).
	Bold(true).
	MarginBottom(1)

// BoardStyle - The game grid border
var BoardStyle = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder()).
	BorderForeground(cPurple).
	Padding(0, 0).
	MarginRight(2)

// StatsBoxStyle - Container for score/level
var StatsBoxStyle = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder()).
	BorderForeground(cComment).
	Padding(0, 1).
	Width(18) // Fixed width for sidebar

// TextStyle for labels
var LabelStyle = lipgloss.NewStyle().
	Foreground(cComment).
	Bold(true)

// ValueStyle for numbers
var ValueStyle = lipgloss.NewStyle().
	Foreground(cCyan)

// renderBlock decides what to draw for a cell
func renderBlock(colorID int, x, y int) string {
	// If it's a piece (ID > 0)
	if colorID > 0 && colorID < len(blockColors) {
		return lipgloss.NewStyle().
			Background(blockColors[colorID]).
			Width(2).
			Render("  ") // Solid block
	}

	// If it's empty, render a stylish dot grid
	// We use a faint dot to guide the eye without noise
	return lipgloss.NewStyle().
		Foreground(cCurrentLine).
		Background(cBackground).
		Width(2).
		Render(" .")
}