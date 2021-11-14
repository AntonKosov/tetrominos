package view

import (
	"tetrominos/view/components"
	"tetrominos/view/ui"
)

type pauseComponentsFactory interface {
	Canvas() *ui.Canvas
	PauseControlHints() components.ControlHints
	PauseMessage() components.Label
}

type pause struct {
	canvas  *ui.Canvas
	hints   components.ControlHints
	message components.Label
}

func newPause(factory pauseComponentsFactory) pause {
	return pause{
		canvas:  factory.Canvas(),
		hints:   factory.PauseControlHints(),
		message: factory.PauseMessage(),
	}
}

func (v pause) Activate() {
	v.message.Show()
	v.hints.Show()
	v.canvas.Draw()
}

func (v pause) Deactivate() {
	v.message.Hide()
	v.hints.Hide()
	v.canvas.Draw()
}
