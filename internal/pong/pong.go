package pong

import (
	"strings"
	"time"

	"pong/internal/resizer"

	tea "github.com/charmbracelet/bubbletea"
)

type Pong struct {
	Width             int
	Height            int
	Border            int
	BallCoordinates   [2]int
	BallVelx          int
	BallVely          int
	PaddleCoordinates [2]int
	State             state
	PaddleWidth       int
	PaddleHeight      int
	PaddleVel         int
	PaddleAcc         int
	Friction          int
	GameStart         bool
	DisplayInfo       bool
	Grid              [][]string
}
type state int

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
	p.Border = 3
	p.PaddleCoordinates[1] = p.Height - p.Border - 2
	p.PaddleHeight = 2
	p.PaddleWidth = int(0.1 * float64(p.Width-2*p.Border))
	p.PaddleAcc = 5
	p.Friction = 1
	grid := make([][]string, p.Height)
	for j, _ := range grid {
		row := make([]string, p.Width)
		for i, _ := range row {

			row[i] = " "
		}
		grid[j] = row
	}
	p.Grid = grid
	return Tick
}
func (p *Pong) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	//var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		p.Width = msg.Width
		p.Height = msg.Height
		p.PaddleCoordinates[0] = (p.Width - (2 * p.Border)) / 2
		p.PaddleCoordinates[1] = p.Height - p.Border - 2

		return p, nil
	case GravityTick:
		p.BallCoordinates[1] += p.BallVely
		p.BallCoordinates[0] += p.BallVelx
		if p.PaddleVel > 0 {
			p.PaddleVel -= p.Friction
		}
		if p.PaddleVel < 0 {
			p.PaddleVel += p.Friction
		}
		p.PaddleCoordinates[0] += p.PaddleVel
		if p.PaddleCoordinates[0] < 0 {
			p.PaddleCoordinates[0] = 0
		}
		if p.PaddleCoordinates[0]+p.PaddleWidth >= p.Width {
			p.PaddleCoordinates[0] = p.Width - p.PaddleWidth
		}
		if p.BallCoordinates[0] >= p.PaddleCoordinates[0] && p.BallCoordinates[0] <= p.PaddleCoordinates[0]+p.PaddleWidth && p.BallCoordinates[1] > p.PaddleCoordinates[1] || p.BallCoordinates[1] < p.Border && p.BallVely < 0 {
			p.BallVely = -p.BallVely
			p.BallVelx = p.PaddleVel
			p.BallCoordinates[1] += p.BallVely
			p.BallCoordinates[0] += p.BallVelx
		}
		if p.BallCoordinates[1] > p.Height-p.Border {
			p.BallCoordinates[1] = p.Border
			p.BallCoordinates[0] = (p.Width - 2*p.Border) / 2
			p.BallVely = 1
			p.BallVelx = 0
		}
		if p.BallCoordinates[0] > p.Width-p.Border {
			p.BallCoordinates[0] = p.Width - p.Border
			p.BallVely = -p.BallVely
			p.BallVelx = -p.BallVelx
		}
		if p.BallCoordinates[0] < p.Border {
			p.BallCoordinates[0] = p.Border
			p.BallVely = -p.BallVely
			p.BallVelx = -p.BallVelx

		}

		return p, Tick
	case tea.KeyMsg:
		//cmd = audio.PlayAudio()
		switch msg.String() {
		case "left", "h":
			p.PaddleVel -= p.PaddleAcc
			p.PaddleCoordinates[0] += p.PaddleVel
			if p.PaddleCoordinates[0] < 0 {
				p.PaddleCoordinates[0] = 0
			}
			return p, nil
		case "right", "l":
			p.PaddleVel += p.PaddleAcc
			p.PaddleCoordinates[0] += p.PaddleVel
			if p.PaddleCoordinates[0]+p.PaddleWidth >= p.Width {
				p.PaddleCoordinates[0] = p.Width - p.PaddleWidth
			}
			return p, nil

		case "ctrl+c", "q", "esc":
			return p, tea.Quit

		case "i", "I":
			p.DisplayInfo = !p.DisplayInfo
			return p, nil

		}
	}

	return p, nil
}

func (p *Pong) View() string {
	grid := make([][]string, p.Height)
	for j, _ := range grid {
		row := make([]string, p.Width)
		for i, _ := range row {

			row[i] = " "
		}
		grid[j] = row
	}
	p.Grid = grid

	if p.PaddleCoordinates[0] < 0 {
		p.PaddleCoordinates[0] = 0
	}
	if p.PaddleCoordinates[0]+p.PaddleWidth >= p.Width {
		p.PaddleCoordinates[0] = p.Width - p.PaddleWidth
	}
	p.drawBoard()
	var output strings.Builder
	for _, row := range p.Grid {
		output.WriteString(strings.Join(row, "") + "\n")
	}

	return output.String()

}

type GravityTick struct{}

func Tick() tea.Msg {
	time.Sleep(time.Millisecond * 90)
	return GravityTick{}
}
