package pong

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type Pong struct {
	width              int
	height             int
	ball_coordinates   [2]int
	paddle_coordinates int
	state              state
}
type state int

const (
	initializing state = iota
	ready
)

func (p Pong) Init() tea.Cmd {
	return nil
}
func (p Pong) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	//var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		p.width = msg.Width
		p.height = msg.Height
		p.state = ready
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
	paddle_top   string = " _____"
	paddle_botom string = "|_____|"
)

func (p Pong) View() string {
	if p.state == initializing {
		return paddle_top + "\n" + paddle_botom
	}

	s := fmt.Sprintf("Width: %v \t Height: %v\n", p.width, p.height)
	return s
}
