package pong


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

