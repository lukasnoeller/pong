package resizer

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Resizer interface {
	tea.Model
	GetWindowDimensions() (int, int)
	SetWindowDimensions(int, int)
}
