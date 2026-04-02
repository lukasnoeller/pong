package screen

import (
	"fmt"

	"pong/internal/audio"
	"pong/internal/option"
	"pong/internal/resizer"

	tea "github.com/charmbracelet/bubbletea"
)

type Screen struct {
	Title   string
	Options []option.Option
	Cursor  int
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
		cmd = audio.PlayAudio("persona.mp3")
		switch msg.Type {
		case tea.KeyDown:
			s.Cursor = (s.Cursor + 1) % len(s.Options)
		case tea.KeyUp:
			s.Cursor = (s.Cursor + len(s.Options) - 1) % len(s.Options)
		case tea.KeyEnter:
			if s.Options[s.Cursor].Model != nil {
				if m, ok := s.Options[s.Cursor].Model.(resizer.Resizer); ok {
					m.SetWindowDimensions(s.Width, s.Height)
					return m, tea.Batch(m.Init(), cmd)
				}
				return s.Options[s.Cursor].Model, cmd
			}
			if s.Options[s.Cursor].Name == "Quit" {
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

	for i, o := range s.Options {
		var namestr string

		namestr = o.Name

		if i == s.Cursor {
			str = str + fmt.Sprintf("\n--> %s", namestr)
		} else {
			str = str + fmt.Sprintf("\n    %s", namestr)
		}

	}
	return str
}
