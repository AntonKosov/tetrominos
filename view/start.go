package view

import (
	"tetrominos/input"
	"tetrominos/view/fonts"
	"tetrominos/view/ui"
)

type start struct {
	canvas *ui.Canvas
	panel  ui.Panel
	hints  controlHints
	name   []string
}

func newStart(canvas *ui.Canvas) start {
	name := fonts.Generate(fonts.Small, " TETROMINOS ")
	h := len(name)
	w := len(name[0])
	p := canvas.CreatePanelInTheCenter(nil, w, h, 2)
	s := start{
		canvas: canvas,
		panel:  p,
		hints:  newControlHints(canvas),
		name:   name,
	}

	return s
}

func (s start) Activate() {
	style := createFontStyle(messageBoxColor, textColor).Bold(true)
	s.panel.OutputAllignedStrings(
		s.name, ui.HCenterAlligment, ui.VCenterAlligment, style,
	)
	s.canvas.Draw()
}

func (s start) Deactivate() {
	s.panel.Clear()
	s.hints.clear()
	s.canvas.Draw()
}

func (s start) ShowControlHints(hints []input.KeyDescription) {
	s.hints.output(hints)
	s.canvas.Draw()
}
