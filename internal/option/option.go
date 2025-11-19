package option

import (
	tea "github.com/charmbracelet/bubbletea"
)

//	type Option interface {
//		GetName() string
//		GetCrossedOut() bool
//		ToggleCrossedOut()
//	}
type Option struct {
	Name  string
	Model tea.Model
	Sound string
}
