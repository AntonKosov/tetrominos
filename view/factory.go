package view

import (
	"fmt"
	"sync"
	"tetrominos/input"
	"tetrominos/view/ui"

	"github.com/gdamore/tcell/v2"
)

type Factory struct {
	screen             tcell.Screen
	input              chan input.Input
	canvas             *ui.Canvas
	terminateWaitGroup sync.WaitGroup
	terminateInputCh   chan struct{}
}

func New() (*Factory, error) {
	s, err := tcell.NewScreen()
	if err != nil {
		return nil, err
	}
	if err := s.Init(); err != nil {
		return nil, err
	}
	w, h := s.Size()
	if w < screenWidth || h < screenHeight {
		s.Fini()
		return nil, fmt.Errorf("the screen size must be at least %dx%d (WxH)", screenWidth, screenHeight)
	}

	bgStyle := createFontStyle(backgroundColor, textColor)
	s.SetStyle(bgStyle)

	canvas := ui.NewCanvas(s, (w-screenWidth)/2, (h-screenHeight)/2, screenWidth, screenHeight, layers, bgStyle)
	gs := &Factory{
		screen:           s,
		canvas:           canvas,
		input:            make(chan input.Input),
		terminateInputCh: make(chan struct{}),
	}

	gs.runInputHandler()

	return gs, nil
}

var keyMapping map[tcell.Key]input.Input

func init() {
	keyMapping = map[tcell.Key]input.Input{
		tcell.KeyCtrlC:  input.CtrlCKey,
		tcell.KeyEscape: input.EscKey,
		tcell.KeyRight:  input.RightKey,
		tcell.KeyLeft:   input.LeftKey,
		tcell.KeyUp:     input.UpKey,
		tcell.KeyDown:   input.DownKey,
	}
}

func (s *Factory) Close() {
	close(s.terminateInputCh)
	s.terminateWaitGroup.Wait()
	s.screen.Fini()
}

func (s *Factory) CreateStartView() *Start {
	v := newStart(s.canvas)
	return &v
}

func (s *Factory) CreateGameView() *game {
	v := newGame(s.canvas)
	return &v
}

func (s *Factory) CreatePauseView() *pause {
	v := newPause(s.canvas)
	return &v
}

func (s *Factory) CreateGameOverView() *gameOver {
	v := newGameOver(s.canvas)
	return &v
}

func (s *Factory) Input() <-chan input.Input {
	return s.input
}

func (s *Factory) runInputHandler() {
	eventCh := make(chan tcell.Event)
	s.terminateWaitGroup.Add(1)
	go func() {
		defer s.terminateWaitGroup.Done()
		for {
			select {
			case <-s.terminateInputCh:
				return
			case e := <-eventCh:
				if eventKey, ok := e.(*tcell.EventKey); ok {
					if key, ok := keyMapping[eventKey.Key()]; ok {
						s.input <- key
					}
				}
			}
		}
	}()

	s.terminateWaitGroup.Add(1)
	go func() {
		defer s.terminateWaitGroup.Done()
		s.screen.ChannelEvents(eventCh, s.terminateInputCh)
	}()
}
