package main

import (
	"fmt"
	"os"

	"bubbletea/internal/titlescreen"

	tea "github.com/charmbracelet/bubbletea"
	_ "github.com/hajimehoshi/go-mp3"
)

type Button struct {
	Name       string
	Color      string
	CrossedOut bool
}

func (b Button) getName() string {
	return b.Name
}
func (b Button) getCrossedOut() bool {
	return b.CrossedOut
}
func (b *Button) toggleCrossedOut() {
	b.CrossedOut = !b.CrossedOut
}

type Game struct {
	Name       string
	CrossedOut bool
	Model      tea.Model
}

func (g *Game) toggleCrossedOut() {
	g.CrossedOut = !g.CrossedOut
}
func (g Game) getCrossedOut() bool {
	return g.CrossedOut
}
func (g Game) getName() string {
	return g.Name
}

func main() {
	t := titlescreen.InitializeTitleScreen()
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
