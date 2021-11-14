package view

import (
	"tetrominos/input"
	"tetrominos/states"
	"tetrominos/view/components"
	"tetrominos/view/components/common"

	"github.com/gdamore/tcell/v2"
)

type compFactory interface {
	startComponentsFactory
	gameComponentsFactory
	pauseComponentsFactory
	gameComponentsFactory
	gameOverComponentsFactory
}

type Terminal struct {
	StartView    states.StartView
	GameView     states.GameView
	PauseView    states.PauseView
	GameOverView states.GameOverView
	Input        <-chan input.Input
	Close        func()
}

func Init(compactMode bool) (*Terminal, error) {
	screen, err := tcell.NewScreen()
	if err != nil {
		return nil, err
	}
	if err := screen.Init(); err != nil {
		return nil, err
	}

	screen.SetStyle(common.BackgroundStyle)

	var compFactory compFactory
	if compactMode {
		// compFactory = components.NewCompactComponentsFactory(screen)
		panic("not implemented")
	} else {
		compFactory = components.NewFullComponentsFactory(screen)
	}

	onResize := func() { compFactory.Canvas().MoveToCenter() }
	ic := newInputController(screen, onResize)
	close := func() {
		ic.Close()
		screen.Fini()
	}

	return &Terminal{
		StartView:    newStart(compFactory),
		GameView:     newGame(compFactory),
		PauseView:    newPause(compFactory),
		GameOverView: newGameOver(compFactory),
		Input:        ic.input,
		Close:        close,
	}, nil
}
