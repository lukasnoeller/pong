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
	Ball              string
	GameStart         bool
	DisplayInfo       bool
	Grid              [][]string
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
	p.Border = 3
	p.PaddleCoordinates[1] = p.Height - p.Border - 2
	p.PaddleHeight = 2
	p.PaddleWidth = 9
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
		if p.BallCoordinates[0] >= p.PaddleCoordinates[0] && p.BallCoordinates[0] <= p.PaddleCoordinates[0]+p.PaddleWidth && p.BallCoordinates[1] > p.PaddleCoordinates[1] || p.BallCoordinates[1] < p.Border && p.BallVely < 0 {
			p.BallVely = -p.BallVely
			p.BallCoordinates[1] += p.BallVely
		}
		if p.BallCoordinates[1] > p.Height-p.Border {
			p.BallCoordinates[1] = 0
		}

		return p, Tick
	case tea.KeyMsg:
		//cmd = audio.PlayAudio()
		switch msg.String() {
		case "left", "h":
			p.PaddleCoordinates[0]--
			p.PaddleVel--
			if p.PaddleCoordinates[0] < 0 {
				p.PaddleCoordinates[0] = 0
			}
			return p, nil
		case "right", "k":
			p.PaddleCoordinates[0]++
			p.PaddleVel++
			if p.PaddleCoordinates[0]+p.PaddleWidth > p.Width {
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
	p.drawBoard()
	var output strings.Builder
	for _, row := range p.Grid {
		output.WriteString(strings.Join(row, "") + "\n")
	}

	return output.String()

}

type GravityTick struct{}

func Tick() tea.Msg {
	time.Sleep(time.Millisecond * 100)
	return GravityTick{}
}
