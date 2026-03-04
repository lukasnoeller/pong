package pong

func (p Pong) drawTopHalf(grid [][]string, j int) [][]string {

	for i := p.PaddleCoordinates[0]; i < p.PaddleCoordinates[0]+p.PaddleWidth; i++ {
		switch i {
		case p.PaddleCoordinates[0]:

			grid[j][i] = "╭"
		case p.PaddleCoordinates[0] + p.PaddleWidth - 1:
			grid[j][i] = "╮"

		default:
			grid[j][i] = "-"
		}
	}
	return grid
}
func (p Pong) drawBottomHalf(grid [][]string, j int) [][]string {

	for i := p.PaddleCoordinates[0]; i < p.PaddleCoordinates[0]+p.PaddleWidth; i++ {
		switch i {
		case p.PaddleCoordinates[0]:

			grid[j][i] = "╰"
		case p.PaddleCoordinates[0] + p.PaddleWidth - 1:
			grid[j][i] = "╯"

		default:
			grid[j][i] = "-"
		}
	}
	return grid
}
func (p Pong) drawMiddle(grid [][]string, j int) [][]string {

	grid[j][p.PaddleCoordinates[0]] = "|"
	grid[j][p.PaddleCoordinates[0]+p.PaddleWidth-1] = "|"
	return grid
}
func (p Pong) drawPaddle(grid [][]string) [][]string {
	grid = p.drawTopHalf(grid, p.PaddleCoordinates[1])
	for j := p.PaddleCoordinates[1] + 1; j < p.PaddleCoordinates[1]+p.PaddleHeight-1; j++ {
		p.drawMiddle(grid, j)
	}
	grid = p.drawBottomHalf(grid, p.PaddleCoordinates[1]+p.PaddleHeight-1)

	return grid
}
func (p Pong) drawBall(grid [][]string) [][]string {
	grid[p.BallCoordinates[1]][p.BallCoordinates[0]] = "X"
	return grid
}
func (p Pong) drawInfo(grid [][]string) [][]string {
	for j := 0; j < p.Border; j++ {
		for i := 0; i < p.Width; i++ {
			grid[j][i] = "*"
			grid[j+(p.Height-p.Border)][i] = "*"
		}
	}
	for j := p.Border; j < p.Height-p.Border; j++ {
		for i := 0; i < p.Border; i++ {
			grid[j][i] = "*"
			grid[j][i+(p.Width-p.Border)] = "*"
		}
	}
	return grid
}

func (p Pong) drawBoard(grid [][]string) [][]string {
	grid = p.drawPaddle(grid)
	grid = p.drawBall(grid)
	if p.DisplayInfo {
		grid = p.drawInfo(grid)
	}
	return grid
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
