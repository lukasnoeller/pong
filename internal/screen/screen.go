package screen

import (
	"fmt"

	"pong/internal/audio"
	"pong/internal/option"
	"pong/internal/pong"
	"pong/internal/resizer"

	tea "github.com/charmbracelet/bubbletea"
)

type Screen struct {
	Title   string
	options []option.Option
	cursor  int
	Width   int
	Height  int
}

func (s Screen) GetWindowDimensions() (int, int) {
	return s.Width, s.Height
}
func (s *Screen) SetWindowDimensions(w int, h int) {
	s.Width = w
	s.Height = h
}
func InitializeTitleScreen() *Screen {
	tp := &Screen{Title: "Pong", cursor: 0, options: []option.Option{{Name: "One Player", Model: &pong.Pong{PaddleTop: pong.PADDLE_TOP, PaddleBottom: pong.PADDLE_BOTTOM, Ball: pong.BALL, BallCoordinates: [2]int{35, 2}, GameStart: true}}, {Name: "Two Players"}, {Name: "Quit"}}}
	tp.cursor = 0
	return tp
}
func (s Screen) Init() tea.Cmd {
	return nil
}
func (s Screen) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		s.Width = msg.Width
		s.Height = msg.Height
	case tea.KeyMsg:
		cmd = audio.PlayAudio()
		switch msg.Type {
		case tea.KeyDown:
			s.cursor = (s.cursor + 1) % len(s.options)
		case tea.KeyUp:
			s.cursor = (s.cursor + len(s.options) - 1) % len(s.options)
		case tea.KeyEnter:
			if s.options[s.cursor].Model != nil {
				if m, ok := s.options[s.cursor].Model.(resizer.Resizer); ok {
					m.SetWindowDimensions(s.Width, s.Height)
					return m, tea.Batch(m.Init(),cmd)
				}
				return s.options[s.cursor].Model, cmd
			}
			if s.options[s.cursor].Name == "Quit" {
				return s, tea.Quit
			}

		case tea.KeyRunes:
			switch string(msg.Runes) {
			case "ctrl+c", "q", "esc":
				return s, tea.Quit
			}

		}
	}

	return s, cmd
}

func (s Screen) View() string {
	var str string
	if s.Title != "" {
		str = s.Title + "\n"
	}

	for i, o := range s.options {
		var namestr string

		namestr = o.Name

		if i == s.cursor {
			str = str + fmt.Sprintf("\n--> %s", namestr)
		} else {
			str = str + fmt.Sprintf("\n    %s", namestr)
		}

	}
	return str
}
