type Pong struct {
	width              int
	height             int
	ball_coordinates   [2]int
	paddle_coordinates int
}

func (p Pong) Init() tea.Cmd {
	return nil
}
func (p Pong) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			if t.options[t.cursor].Model != nil {
				return t.options[t.cursor].Model, cmd
			}
			if t.options[t.cursor].Name == "Quit" {
				return t, tea.Quit
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

func (t Screen) View() string {
	var s string
	if t.Title != "" {
		s = t.Title + "\n"
	}

	for i, o := range t.options {
		var namestr string

		namestr = o.Name

		if i == t.cursor {
			s = s + fmt.Sprintf("\n--> %s", namestr)
		} else {
			s = s + fmt.Sprintf("\n    %s", namestr)
		}

	}
	return s
}