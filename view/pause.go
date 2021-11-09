package view

import (
	"tetrominos/input"
	"tetrominos/view/fonts"
	"tetrominos/view/ui"
)

type pause struct {
	canvas *ui.Canvas
	panel  ui.Panel
	hints  controlHints
	text   []string
}

func newPause(canvas *ui.Canvas) pause {
	pauseText := fonts.Generate(fonts.Small, " PAUSE ")
	h := len(pauseText)
	w := len(pauseText[0])
	p := canvas.CreatePanelInTheCenter(nil, w, h, 2)
	v := pause{
		canvas: canvas,
		panel:  p,
		hints:  newControlHints(canvas),
		text:   pauseText,
	}

	return v
}

func (v pause) Activate() {
	s := createFontStyle(messageBoxColor, textColor).Bold(true).Blink(true)
	v.panel.OutputStrings(0, 0, v.text, s)
}

func (v pause) Deactivate() {
	v.panel.Clear()
	v.hints.clear()
	v.canvas.Draw()
}

func (v pause) ShowControlHints(hints []input.KeyDescription) {
	v.hints.output(hints)
	v.canvas.Draw()
}
