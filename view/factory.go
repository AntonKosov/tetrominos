package view

import (
	"fmt"
	"tetrominos/input"
	"tetrominos/states"
	"tetrominos/view/ui"

	"github.com/gdamore/tcell/v2"
)

type Terminal struct {
	StartView    states.StartView
	GameView     states.GameView
	PauseView    states.PauseView
	GameOverView states.GameOverView
	Input        <-chan input.Input
	Close        func()
}

func Init() (*Terminal, error) {
	screen, err := tcell.NewScreen()
	if err != nil {
		return nil, err
	}
	if err := screen.Init(); err != nil {
		return nil, err
	}
	w, h := screen.Size()
	if w < screenWidth || h < screenHeight {
		screen.Fini()
		return nil, fmt.Errorf(
			"the screen size must be at least %dx%d (WxH)",
			screenWidth, screenHeight,
		)
	}

	bgStyle := createFontStyle(backgroundColor, textColor)
	screen.SetStyle(bgStyle)

	originX, originY := canvasOrigin(screen)
	canvas := ui.NewCanvas(
		screen, originX, originY,
		screenWidth, screenHeight, layers, bgStyle,
	)

	onResize := func() { canvas.ChangeOrigin(canvasOrigin(screen)) }
	ic := newInputController(screen, onResize)
	close := func() {
		ic.Close()
		screen.Fini()
	}

	return &Terminal{
		StartView:    newStart(canvas),
		GameView:     newGame(canvas),
		PauseView:    newPause(canvas),
		GameOverView: newGameOver(canvas),
		Input:        ic.input,
		Close:        close,
	}, nil
}

func canvasOrigin(screen tcell.Screen) (int, int) {
	w, h := screen.Size()
	return (w - screenWidth) / 2, (h - screenHeight) / 2
}
