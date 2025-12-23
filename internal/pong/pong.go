package pong

import (
	"bubbletea/internal/resizer"
	"fmt"
	"time"
	"strings"

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
	DisplayInfo bool
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
		if p.BallCoordinates[0] >= p.PaddleCoordinates && p.BallCoordinates[0] <= p.PaddleCoordinates + len(p.PaddleBottom) && p.BallCoordinates[1] > p.Height - 7 || p.BallCoordinates[1] < 0 && p.BallVely <0 {
			p.BallVely = -p.BallVely
			p.BallCoordinates[1]+= p.BallVely
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
			
			case "i", "I":
				p.DisplayInfo = !p.DisplayInfo	
				return p, nil
			}
		}
	}

	return p, nil
}

func (p *Pong) View() string {
	s := p.drawBoard()
	s += p.drawPaddle()
	if p.DisplayInfo {
		
		s += fmt.Sprintf("Width: %v Height: %v\n", p.Width, p.Height) + fmt.Sprintf("PaddleCoordinates: %v\t BallCoordinates: %v BallVely: %v Num lines: %v\n", p.PaddleCoordinates, p.BallCoordinates, p.BallVely, strings.Count(s, "\n"))


}
	return s
	// s := fmt.Sprintf("Width: %v \t Height: %v\n", p.Width, p.Height)
	
}
type GravityTick struct{}

func Tick() tea.Msg {
	time.Sleep(time.Millisecond * 200)
	return GravityTick{}
}
