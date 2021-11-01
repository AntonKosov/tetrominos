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
	factory, err := view.New()
	if err != nil {
		log.Fatalln(fmt.Sprintf("Cannot initialize terminal: %v\n", err.Error()))
	}
	defer factory.Close()

	rand.Seed(time.Now().UnixMicro())

	views := engine.Views{
		Start:    factory.CreateStartView(),
		Game:     factory.CreateGameView(),
		Pause:    factory.CreatePauseView(),
		GameOver: factory.CreateGameOverView(),
	}
	engine.Run(views, factory.Input())
}
