package main

import (
	"fmt"
	"os"
	"pong/internal/screen"

	tea "github.com/charmbracelet/bubbletea"
	_ "github.com/hajimehoshi/go-mp3"
)

func main() {
	t := screen.InitializeTitleScreen()
	p := tea.NewProgram(
		t,
		tea.WithAltScreen(),       // use the full size of the terminal in its "alternate screen buffer"
		tea.WithMouseCellMotion(), // turn on mouse support so we can track the mouse wheel
	)

	if _, err := p.Run(); err != nil {
		fmt.Println("could not run program:", err)
		os.Exit(1)
	}
}
