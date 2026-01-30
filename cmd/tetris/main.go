package main

import (
	"fmt"
	"os"

	"github.com/shiroyashv/go-tetris-tui/internal/input"
	"github.com/shiroyashv/go-tetris-tui/internal/ui"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	cleanup, err := input.StartBridge()
	if err != nil {
		fmt.Println("Error starting input bridge:", err)
	} else {
		defer cleanup()
	}

	p := tea.NewProgram(ui.NewModel(), tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
