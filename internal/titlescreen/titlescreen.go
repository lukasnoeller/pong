package titlescreen

import (
	"fmt"

	"bubbletea/internal/audio"
	"bubbletea/internal/option"

	tea "github.com/charmbracelet/bubbletea"
)

type TitleScreen struct {
	Title   string
	options []option.Option
	cursor  int
}

func InitializeTitleScreen() TitleScreen {
	tp := TitleScreen{}
	tp.Title = "Pick your game!"
	tp.options = []option.Option{&option.Game{Name: "Snake", CrossedOut: false}, &option.Game{Name: "Pong", CrossedOut: false, Model: TitleScreen{Title: "Pong", cursor: 0, options: []option.Option{&option.Button{Name: "Push Me", Color: "red"}}}}, &option.Game{Name: "Asteroids", CrossedOut: false}}
	tp.cursor = 0
	return tp
}
func (t TitleScreen) Init() tea.Cmd {
	return nil
}
func (t TitleScreen) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		cmd = audio.PlayAudio()
		switch msg.Type {

		case tea.KeyDown:
			t.cursor = (t.cursor + 1) % len(t.options)
		case tea.KeyUp:
			t.cursor = (t.cursor + len(t.options) - 1) % len(t.options)
		case tea.KeyEnter:
			t.options[t.cursor].ToggleCrossedOut()
			value, ok := t.options[t.cursor].(*option.Game)
			if t.options[t.cursor].GetCrossedOut() && ok {
				if value.Model != nil {
					return value.Model, cmd
				}

			}
		case tea.KeyRunes:
			switch string(msg.Runes) {
			case "ctrl+c", "q", "esc":
				return t, tea.Quit
			}

		}
	}

	return t, cmd
}

func (t TitleScreen) View() string {
	var s string
	if t.Title != "" {
		s = t.Title + "\n"
	}

	for i, o := range t.options {
		var namestr string
		if o.GetCrossedOut() {
			namestr = "--------"
		} else {
			namestr = o.GetName()
		}
		if i == t.cursor {
			s = s + fmt.Sprintf("\n--> %s", namestr)
		} else {
			s = s + fmt.Sprintf("\n    %s", namestr)
		}

	}
	return s
}
