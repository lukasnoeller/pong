package pong

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
	p.Grid[p.BallCoordinates[1]][p.BallCoordinates[0]] = "X"
}
func (p *Pong) drawInfo() {
	for j := 0; j < p.Border; j++ {
		for i := 0; i < p.Width; i++ {
			p.Grid[j][i] = "*"
			p.Grid[j+(p.Height-p.Border)][i] = "*"
		}
	}
	for j := p.Border; j < p.Height-p.Border; j++ {
		for i := 0; i < p.Border; i++ {
			p.Grid[j][i] = "*"
			p.Grid[j][i+(p.Width-p.Border)] = "*"
		}
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
