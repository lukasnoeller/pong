package main

import (
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/hajimehoshi/go-mp3"
	_ "github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto/v2"
)

type TitleScreen struct {
	Title   string
	options []Option
	cursor  int
}
type Option interface {
	getName() string
	getCrossedOut() bool
	toggleCrossedOut()
}
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
func playAudio() tea.Cmd {
	return func() tea.Msg {
		err := run()
		return err
	}
}

func InitializeTitleScreen() TitleScreen {
	tp := TitleScreen{}
	tp.Title = "Pick your game!"
	tp.options = []Option{&Game{Name: "Snake", CrossedOut: false}, &Game{Name: "Pong", CrossedOut: false, Model: TitleScreen{Title: "Pong", cursor: 0, options: []Option{&Button{Name: "Push Me", Color: "red"}}}}, &Game{Name: "Asteroids", CrossedOut: false}}
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
		cmd = playAudio()
		switch msg.Type {

		case tea.KeyDown:
			t.cursor = (t.cursor + 1) % len(t.options)
		case tea.KeyUp:
			t.cursor = (t.cursor + len(t.options) - 1) % len(t.options)
		case tea.KeyEnter:
			t.options[t.cursor].toggleCrossedOut()
			value, ok := t.options[t.cursor].(*Game)
			if t.options[t.cursor].getCrossedOut() && ok {
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
		if o.getCrossedOut() {
			namestr = "--------"
		} else {
			namestr = o.getName()
		}
		if i == t.cursor {
			s = s + fmt.Sprintf("\n--> %s", namestr)
		} else {
			s = s + fmt.Sprintf("\n    %s", namestr)
		}

	}
	return s
}
func run() error {
	f, err := os.Open("persona.mp3")
	if err != nil {
		return err
	}
	defer f.Close()

	d, err := mp3.NewDecoder(f)
	if err != nil {
		return err
	}

	c, ready, err := oto.NewContext(d.SampleRate(), 2, 2)
	if err != nil {
		return err
	}
	<-ready

	p := c.NewPlayer(d)
	defer p.Close()
	p.Play()

	//fmt.Printf("Length: %d[bytes]\n", d.Length())
	for {
		time.Sleep(time.Second)
		if !p.IsPlaying() {
			break
		}
	}

	return nil
}
func main() {
	t := InitializeTitleScreen()
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
