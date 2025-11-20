package pong

import (
	"bubbletea/internal/resizer"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type Pong struct {
	Width              int
	Height             int
	Ball_coordinates   [2]int
	Paddle_coordinates int
	State              state
}
type state int

const (
	initializing state = iota
	ready
)

var _ resizer.Resizer = (*Pong)(nil)

func (p Pong) GetWindowDimensions() (int, int) {
	return p.Width, p.Height
}
func (p *Pong) SetWindowDimensions(w int, h int) {
	p.Width = w
	p.Height = h
}
func (p Pong) Init() tea.Cmd {
	return nil
}
func (p Pong) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	//var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		p.Width = msg.Width
		p.Height = msg.Height
		p.State = ready
		return p, nil
	case tea.KeyMsg:
		//cmd = audio.PlayAudio()
		switch msg.Type {
		case tea.KeyDown:

		case tea.KeyUp:

		case tea.KeyEnter:

		case tea.KeyRunes:
			switch string(msg.Runes) {
			case "ctrl+c", "q", "esc":
				return p, tea.Quit
			}

		}
	}

	return p, nil
}

const (
	paddle_top    string = " _____"
	paddle_bottom string = "|_____|"
)

func (p Pong) View() string {
	if p.State == initializing {
		return paddle_top + "\n" + paddle_bottom + "\n" + fmt.Sprintf("Width: %v Height: %v\n", p.Width, p.Height)
	}

	s := fmt.Sprintf("Width: %v \t Height: %v\n", p.Width, p.Height)
	return s
}
