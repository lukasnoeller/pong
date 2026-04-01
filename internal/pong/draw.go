package pong

import (
	"errors"
	"fmt"
	"log"
)

func (p *Pong) drawTopHalf(j int) {

	for i := p.PaddleCoordinates[0]; i < p.PaddleCoordinates[0]+p.PaddleWidth; i++ {
		switch i {
		case p.PaddleCoordinates[0]:

			p.Grid[j][i] = "╭"
		case p.PaddleCoordinates[0] + p.PaddleWidth - 1:
			p.Grid[j][i] = "╮"

		default:
			p.Grid[j][i] = "-"
		}
	}
}
func (p *Pong) drawBottomHalf(j int) {

	for i := p.PaddleCoordinates[0]; i < p.PaddleCoordinates[0]+p.PaddleWidth; i++ {
		switch i {
		case p.PaddleCoordinates[0]:

			p.Grid[j][i] = "╰"
		case p.PaddleCoordinates[0] + p.PaddleWidth - 1:
			p.Grid[j][i] = "╯"

		default:
			p.Grid[j][i] = "-"
		}
	}
}
func (p *Pong) drawMiddle(j int) {

	p.Grid[j][p.PaddleCoordinates[0]] = "|"
	p.Grid[j][p.PaddleCoordinates[0]+p.PaddleWidth-1] = "|"
}
func (p *Pong) drawPaddle() {
	p.drawTopHalf(p.PaddleCoordinates[1])
	for j := p.PaddleCoordinates[1] + 1; j < p.PaddleCoordinates[1]+p.PaddleHeight-1; j++ {
		p.drawMiddle(j)
	}
	p.drawBottomHalf(p.PaddleCoordinates[1] + p.PaddleHeight - 1)

}
func (p *Pong) drawBall() {
	p.Grid[p.BallCoordinates[1]-1][p.BallCoordinates[0]-1] = "▄"
	p.Grid[p.BallCoordinates[1]-1][p.BallCoordinates[0]] = "█"
	p.Grid[p.BallCoordinates[1]-1][p.BallCoordinates[0]+1] = "▄"
	p.Grid[p.BallCoordinates[1]][p.BallCoordinates[0]-1] = "█"
	p.Grid[p.BallCoordinates[1]][p.BallCoordinates[0]-2] = "█"
	p.Grid[p.BallCoordinates[1]][p.BallCoordinates[0]] = "█"
	p.Grid[p.BallCoordinates[1]][p.BallCoordinates[0]+1] = "█"
	p.Grid[p.BallCoordinates[1]][p.BallCoordinates[0]+2] = "█"
	p.Grid[p.BallCoordinates[1]+1][p.BallCoordinates[0]-1] = "▀"
	p.Grid[p.BallCoordinates[1]+1][p.BallCoordinates[0]] = "█"
	p.Grid[p.BallCoordinates[1]+1][p.BallCoordinates[0]+1] = "▀"
}
func (p *Pong) Write2Row(r int, info string, value any) error {
	if r >= len(p.Grid) || r < 0 {
		return errors.New("invalid row number provided")
	}
	writeString := fmt.Sprintf("%s: %v", info, value)
	for i := 0; i < len(writeString); i++ {
		if i >= len(p.Grid[r]) {
			return errors.New("exceeded window width")
		}
		p.Grid[r][i] = string(writeString[i])
	}
	return nil
}
func (p *Pong) Write2RowMiddle(r int, info string, value any) error {
	if r >= len(p.Grid) || r < 0 {
		return errors.New("invalid row number provided")
	}
	writeString := fmt.Sprintf("%s: %v", info, value)
	log.Printf("Logging i start: %v\n", p.Width/2-len(writeString)/2)
	for i := 0; i < len(writeString); i++ {
		if i+(p.Width/2-len(writeString)/2) >= len(p.Grid[r]) {
			return errors.New("exceeded window width")
		}
		p.Grid[r][i+p.Width/2-len(writeString)/2] = string(writeString[i])
	}
	return nil
}
func (p *Pong) drawInfo() {
	if err := p.Write2Row(p.Height-p.Border, "PaddleVel", p.PaddleVel); err != nil {
		log.Printf("error during Write2Row: %v\n", err.Error())
	}
	if err := p.Write2Row(p.Height-p.Border+1, "PaddleCoordinates", p.PaddleCoordinates); err != nil {
		log.Printf("error during Write2Row: %v\n", err.Error())
	}
	for j := 0; j < p.Border; j++ {
		for i := 0; i < p.Width; i++ {
			p.Grid[j][i] = " "
			//p.Grid[j+(p.Height-p.Border)][i] = "*"
		}
	}
	for j := p.Border; j < p.Height-p.Border; j++ {
		for i := 0; i < p.Border; i++ {
			p.Grid[j][i] = " "
			p.Grid[j][i+(p.Width-p.Border)] = " "
		}
	}
	if err := p.Write2RowMiddle(p.Border/2, "Score", 3); err != nil {
		log.Printf("error during Write2RowMiddle: %v\n", err.Error())
	}
}

func (p *Pong) drawBoard() {
	p.drawPaddle()
	p.drawBall()
	if p.DisplayInfo {
		p.drawInfo()
	}
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
