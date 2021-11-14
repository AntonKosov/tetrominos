package view

import (
	"tetrominos/view/components"
	"tetrominos/view/ui"
)

type startComponentsFactory interface {
	Canvas() *ui.Canvas
	Title() components.Label
	StartHints() components.Hints
	GameField() components.GameField
}

type start struct {
	canvas    *ui.Canvas
	title     components.Label
	hints     components.Hints
	gameField components.GameField
}

func newStart(factory startComponentsFactory) start {
	s := start{
		canvas:    factory.Canvas(),
		title:     factory.Title(),
		hints:     factory.StartHints(),
		gameField: factory.GameField(),
	}

	return s
}

func (s start) Activate() {
	s.title.Show()
	s.hints.Show()
	s.gameField.Reset()
	s.canvas.Draw()
}

func (s start) Deactivate() {
	s.title.Hide()
	s.hints.Hide()
	s.gameField.Hide()
	s.canvas.Draw()
}
