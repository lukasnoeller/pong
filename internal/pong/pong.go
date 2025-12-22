package pong

import (
	"bubbletea/internal/resizer"
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type Pong struct {
	Width             int
	Height            int
	BallCoordinates   [2]int
	BallVelx int
	BallVely int
	PaddleCoordinates int
	State             state
	PaddleTop         string
	PaddleBottom      string
	Ball              string
	GameStart         bool
	InitCalls int
}
type state int

const (
	PADDLE_TOP    string = " _____"
	PADDLE_BOTTOM string = "|_____|"
	BALL          string = "•"
)

var _ resizer.Resizer = (*Pong)(nil)

func (p Pong) GetWindowDimensions() (int, int) {
	return p.Width, p.Height
}
func (p *Pong) SetWindowDimensions(w int, h int) {
	p.Width = w
	p.Height = h
}
func (p *Pong) Init() tea.Cmd {
	p.InitCalls++
	p.BallVely = 1
	p.BallVelx = 0
	return Tick
}
func (p *Pong) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	//var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		p.Width = msg.Width
		p.Height = msg.Height
		p.PaddleCoordinates = p.Width / 2
		return p, nil
	case GravityTick:
		p.BallCoordinates[1]+= p.BallVely
		p.BallCoordinates[0]+= p.BallVelx
		if p.BallCoordinates[1] > p.Height {
			p.BallCoordinates[1] = 0
		}
		return p, Tick
	case tea.KeyMsg:
		//cmd = audio.PlayAudio()
		switch msg.Type {
		case tea.KeyLeft:
			p.PaddleCoordinates--
			if p.PaddleCoordinates < 0 {
				p.PaddleCoordinates = 0
			}
			return p, nil
		case tea.KeyRight:
			p.PaddleCoordinates++
			if p.PaddleCoordinates+len(p.PaddleBottom) > p.Width {
				p.PaddleCoordinates = p.Width - len(p.PaddleBottom)
			}
			return p, nil
		case tea.KeyDown: case tea.KeyUp:
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

func (p *Pong) View() string {
	s := p.drawBoard()
	s += p.drawPaddle()
	s += fmt.Sprintf("Width: %v Height: %v\n", p.Width, p.Height) + fmt.Sprintf("PaddleCoordinates: %v\t BallCoordinates: %v\n", p.PaddleCoordinates, p.BallCoordinates)
	s += fmt.Sprintf("Init called %d times\n", p.InitCalls)
	return s
	// s := fmt.Sprintf("Width: %v \t Height: %v\n", p.Width, p.Height)
	
}

func (p Pong) drawPaddle() string {
	var s string
	for _ = range p.PaddleCoordinates {
		s += " "
	}
	s += PADDLE_TOP + "\n" + s + PADDLE_BOTTOM + "\n"
	return s
}
func (p Pong) drawBoard() string {
	var s string
	s = "\n" + p.CenterString("P O N G") + "\n"
	for j := range p.Height - 6 {
		if j == p.BallCoordinates[1] {
			for i := range p.Width {
				if i == p.BallCoordinates[0] {
					s += p.Ball
				} else {
					s += " "
				}
			}
		}
		s += "\n"
	}
	return s
}
func (p Pong) CenterString(str string) string {
	var s string
	for i := range p.Width {
		if i == p.Width/2 {
			s += str
		} else {
			s += " "
		}
	}
	return s
}

type GravityTick struct{}

func Tick() tea.Msg {
	time.Sleep(time.Second * 2)
	return GravityTick{}
}
