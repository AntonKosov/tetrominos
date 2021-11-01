package view

import (
	"tetrominos/input"
	"tetrominos/view/fonts"
	"tetrominos/view/ui"

	"github.com/gdamore/tcell/v2"
)

type Start struct {
	canvas *ui.Canvas
	panel  ui.Panel
	hints  controlHints
	name   []string
}

func newStart(canvas *ui.Canvas) Start {
	name := fonts.Generate(fonts.Small, " TETROMINOS ")
	h := len(name)
	w := len(name[0])
	p := canvas.CreatePanelInTheCenter(nil, w, h, 2)
	s := Start{
		canvas: canvas,
		panel:  p,
		hints:  newControlHints(canvas),
		name:   name,
	}

	return s
}

func (s Start) Activate() {
	style := tcell.StyleDefault.
		Background(tcell.ColorDarkRed).
		Foreground(tcell.ColorWhite).Bold(true)
	s.panel.OutputAllignedStrings(
		s.name, ui.HCenterAlligment, ui.VCenterAlligment, style,
	)
	s.canvas.Draw()
}

func (s Start) Deactivate() {
	s.panel.Clear()
	s.hints.clear()
	s.canvas.Draw()
}

func (s Start) ShowControlHints(hints []input.KeyDescription) {
	s.hints.output(hints)
	s.canvas.Draw()
}
