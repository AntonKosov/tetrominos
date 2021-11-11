package view

import (
	"sync"
	"tetrominos/input"

	"github.com/gdamore/tcell/v2"
)

type inputController struct {
	screen             tcell.Screen
	input              chan input.Input
	terminateWaitGroup sync.WaitGroup
	terminateInputCh   chan struct{}
}

func newInputController(screen tcell.Screen) *inputController {
	ic := &inputController{
		screen:           screen,
		input:            make(chan input.Input),
		terminateInputCh: make(chan struct{}),
	}

	ic.runInputHandler()

	return ic
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

func (c *inputController) Close() {
	close(c.terminateInputCh)
	c.terminateWaitGroup.Wait()
}

func (c *inputController) runInputHandler() {
	eventCh := make(chan tcell.Event)
	c.terminateWaitGroup.Add(1)
	go func() {
		defer c.terminateWaitGroup.Done()
		for {
			select {
			case <-c.terminateInputCh:
				return
			case e := <-eventCh:
				if eventKey, ok := e.(*tcell.EventKey); ok {
					if key, ok := keyMapping[eventKey.Key()]; ok {
						c.input <- key
					}
				}
			}
		}
	}()

	c.terminateWaitGroup.Add(1)
	go func() {
		defer c.terminateWaitGroup.Done()
		c.screen.ChannelEvents(eventCh, c.terminateInputCh)
	}()
}
