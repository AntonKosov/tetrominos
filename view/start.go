package view

import (
	"tetrominos/view/components"
	"tetrominos/view/ui"
)

type startComponentsFactory interface {
	Canvas() *ui.Canvas
	Title() components.Label
	StartControlHints() components.ControlHints
}

type start struct {
	canvas *ui.Canvas
	title  components.Label
	hints  components.ControlHints
}

func newStart(factory startComponentsFactory) start {
	s := start{
		canvas: factory.Canvas(),
		title:  factory.Title(),
		hints:  factory.StartControlHints(),
	}

	return s
}

func (s start) Activate() {
	s.title.Show()
	s.hints.Show()
	s.canvas.Draw()
}

func (s start) Deactivate() {
	s.title.Hide()
	s.hints.Hide()
	s.canvas.Draw()
}
