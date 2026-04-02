package pong

import (
	"strings"
	"time"

	"pong/internal/audio"
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
	BallAcc           int
	PaddleCoordinates [2]int
	State             state
	PaddleWidth       int
	PaddleHeight      int
	PaddleVel         float32
	PaddleAcc         float32
	MaxSpeed          float32
	Friction          float32
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
	p.Border = 5
	p.PaddleHeight = 3
	p.PaddleWidth = int(0.15 * float64(p.Width-2*p.Border))
	p.PaddleCoordinates[1] = p.Height - p.Border - p.PaddleHeight
	p.PaddleAcc = 5.0
	p.Friction = 0.95
	p.MaxSpeed = 8.0
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
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		p.Width = msg.Width
		p.Height = msg.Height
		p.PaddleCoordinates[0] = (p.Width - (2 * p.Border)) / 2
		p.PaddleCoordinates[1] = p.Height - p.Border - 2

		return p, nil
	case GravityTick:
		cmd := Tick
		p.BallCoordinates[1] += p.BallVely
		p.BallCoordinates[0] += p.BallVelx
		if p.PaddleVel > 0 {
			p.PaddleVel -= p.Friction
		}
		if p.PaddleVel < 0 {
			p.PaddleVel += p.Friction
		}
		if p.PaddleVel > p.MaxSpeed {
			p.PaddleVel = p.MaxSpeed
		}
		if p.PaddleVel < -p.MaxSpeed {
			p.PaddleVel = -p.MaxSpeed
		}
		p.PaddleCoordinates[0] += int(p.PaddleVel)
		// if collision with paddle
		if p.BallCoordinates[0] >= p.PaddleCoordinates[0] && p.BallCoordinates[0] <= p.PaddleCoordinates[0]+p.PaddleWidth && p.BallCoordinates[1] > p.PaddleCoordinates[1] {
			p.BallVely = -p.BallVely
			p.BallVelx += int(p.PaddleVel) // Will change later when making ball vel float32
			p.BallCoordinates[1] += p.BallVely
			p.BallCoordinates[0] += p.BallVelx
			cmd = tea.Batch(cmd, audio.PlayAudio("hit.mp3"))
		}
		if p.BallCoordinates[1] < p.Border {
			p.BallCoordinates[1] += p.Border
			p.BallCoordinates[0] += p.BallVelx
			if p.BallVely < 0 {
				p.BallVely = -p.BallVely
			}
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

		return p, cmd
	case tea.KeyMsg:
		//cmd = audio.PlayAudio()
		switch msg.String() {
		case "left", "h":
			if p.PaddleCoordinates[0] > 0 && p.PaddleCoordinates[0] < p.Width-p.Border {

				p.PaddleVel -= p.PaddleAcc
				p.PaddleCoordinates[0] += int(p.PaddleVel)
			}
			return p, nil
		case "right", "l":
			p.PaddleVel += p.PaddleAcc
			p.PaddleCoordinates[0] += int(p.PaddleVel)
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
		p.PaddleVel = 0
	}
	if p.PaddleCoordinates[0]+p.PaddleWidth >= p.Width {
		p.PaddleCoordinates[0] = p.Width - p.PaddleWidth
		p.PaddleVel = 0
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
