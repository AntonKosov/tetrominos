package view

import (
	"tetrominos/input"
	"tetrominos/view/fonts"
	"tetrominos/view/ui"
)

type gameOver struct {
	canvas       *ui.Canvas
	hints        controlHints
	panel        ui.Panel
	gameOverText []string
}

func newGameOver(canvas *ui.Canvas) gameOver {
	gameOverText := fonts.Generate(fonts.Small, " GAME OVER ")
	w := len(gameOverText[0])
	c := gameOver{
		canvas:       canvas,
		hints:        newControlHints(canvas),
		panel:        canvas.CreatePanelInTheCenter(nil, w, len(gameOverText), 2),
		gameOverText: gameOverText,
	}
	return c
}

func (g gameOver) Activate() {
	style := createFontStyle(messageBoxColor, textColor).Bold(true)
	g.panel.OutputAllignedStrings(
		g.gameOverText, ui.HCenterAlligment, ui.TopAlligment, style,
	)
	g.canvas.Draw()
}

func (g gameOver) Deactivate() {
	g.panel.Clear()
	g.hints.clear()
	g.canvas.Draw()
}

func (g gameOver) ShowControlHints(hints []input.KeyDescription) {
	g.hints.output(hints)
	g.canvas.Draw()
}
