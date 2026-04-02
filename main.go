package main

import (
	"fmt"
	"log"
	"os"
	"pong/internal/option"
	"pong/internal/pong"
	"pong/internal/screen"

	tea "github.com/charmbracelet/bubbletea"
	_ "github.com/hajimehoshi/go-mp3"
)

func initializeTitleScreen() *screen.Screen {
	tp := &screen.Screen{Title: "Pong", Cursor: 0, Options: []option.Option{{Name: "One Player", Model: &pong.Pong{BallCoordinates: [2]int{35, 2}, GameStart: true}}, {Name: "Two Players"}, {Name: "Quit"}}}
	tp.Cursor = 0
	return tp
}
func main() {
	file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.SetOutput(file)
	t := initializeTitleScreen()
	p := tea.NewProgram(
		t,
		tea.WithAltScreen(),       // use the full size of the terminal in its "alternate screen buffer"
		tea.WithMouseCellMotion(), // turn on mouse support so we can track the mouse wheel
	)

	if _, err := p.Run(); err != nil {
		fmt.Println("could not run program:", err)
		os.Exit(1)
	}
}
