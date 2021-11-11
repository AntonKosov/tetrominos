package main

import (
	"fmt"
	"log"
	"math/rand"
	"tetrominos/engine"
	"tetrominos/view"
	"time"
)

func main() {
	terminal, err := view.Init()
	if err != nil {
		log.Fatalln(fmt.Sprintf("Cannot initialize terminal: %v\n", err.Error()))
	}
	defer terminal.Close()

	rand.Seed(time.Now().UnixMicro())

	views := engine.Views{
		Start:    terminal.StartView,
		Game:     terminal.GameView,
		Pause:    terminal.PauseView,
		GameOver: terminal.GameOverView,
	}
	engine.Run(views, terminal.Input)
}
