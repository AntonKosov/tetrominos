package view

import (
	"tetrominos/view/components"
	"tetrominos/view/ui"
)

type gameOverComponentsFactory interface {
	Canvas() *ui.Canvas
	GameOverControlHints() components.ControlHints
	GameOverMessage() components.Label
}

type gameOver struct {
	canvas  *ui.Canvas
	message components.Label
	hints   components.ControlHints
}

func newGameOver(factory gameOverComponentsFactory) gameOver {
	return gameOver{
		canvas:  factory.Canvas(),
		message: factory.GameOverMessage(),
		hints:   factory.GameOverControlHints(),
	}
}

func (g gameOver) Activate() {
	g.message.Show()
	g.hints.Show()
	g.canvas.Draw()
}

func (g gameOver) Deactivate() {
	g.message.Hide()
	g.hints.Hide()
	g.canvas.Draw()
}
