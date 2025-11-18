package option

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Game struct {
	Name       string
	CrossedOut bool
	Model      tea.Model
}

func (g *Game) ToggleCrossedOut() {
	g.CrossedOut = !g.CrossedOut
}
func (g Game) GetCrossedOut() bool {
	return g.CrossedOut
}
func (g Game) GetName() string {
	return g.Name
}
