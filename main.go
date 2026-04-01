package main

import (
	"fmt"
	"log"
	"os"
	"pong/internal/screen"

	tea "github.com/charmbracelet/bubbletea"
	_ "github.com/hajimehoshi/go-mp3"
)

func main() {
	file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.SetOutput(file)
	t := screen.InitializeTitleScreen()
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
