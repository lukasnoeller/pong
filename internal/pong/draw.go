package pong

func (p Pong) drawPaddle(grid [][]string) [][]string {
	for j := p.PaddleCoordinates[1]; j < p.PaddleCoordinates[1]+p.PaddleHeight; j++ {

		for i := p.PaddleCoordinates[0]; i < p.PaddleCoordinates[0]+p.PaddleWidth; i++ {
			if i == p.PaddleCoordinates[0] || i == p.PaddleCoordinates[0]+p.PaddleWidth-1 {
				grid[j][i] = "|"
			} else {
				grid[j][i] = "-"
			}
		}
	}
	return grid
}
func (p Pong) drawBoard() string {
	var s string
	s = "\n" + p.CenterString("P O N G") + "\n"
	for j := range p.Height - 7 {
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
