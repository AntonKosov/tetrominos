package common

import (
	"tetrominos/math"
	"tetrominos/view/fonts"
	"tetrominos/view/ui"

	"github.com/gdamore/tcell/v2"
)

type LabelParams struct {
	// Canvas is required
	Canvas *ui.Canvas
	Style  tcell.Style
	Layer  ui.Layer
	Font   fonts.Font
	// if coordinate are no provided, the lable will be in the center
	Coordinate *math.Vector
	// Width or/and Text must be provided
	Width *int
	Text  *string
	// Horizontal alligment
	HAlligment *ui.HAlligment
}

type Label struct {
	isVisible  bool
	text       string
	panel      ui.Panel
	font       fonts.Font
	style      tcell.Style
	hAlligment ui.HAlligment
}

func NewLabel(params LabelParams) Label {
	text := ""
	if params.Text != nil {
		text = *params.Text
	}

	var width int
	if params.Width == nil {
		if len(text) == 0 {
			panic("the size of a label cannot be calculated")
		}
		gt := fonts.Generate(params.Font, text)
		width = len(gt[0])
	} else {
		width = *params.Width
	}

	height := fonts.FontHeight(params.Font)

	var panel ui.Panel
	if params.Coordinate != nil {
		panel = params.Canvas.CreatePanel(nil,
			params.Coordinate.X, params.Coordinate.Y,
			width, height, params.Layer)
	} else {
		panel = params.Canvas.CreatePanelInTheCenter(
			nil, width, height, params.Layer)
	}

	hAlligment := ui.LeftAlligment
	if params.HAlligment != nil {
		hAlligment = *params.HAlligment
	}

	return Label{
		text:       text,
		panel:      panel,
		font:       params.Font,
		style:      params.Style,
		hAlligment: hAlligment,
	}
}

func (l *Label) Width() int {
	return l.panel.Width()
}

func (l *Label) Height() int {
	return l.panel.Height()
}

func (l *Label) SetText(text string) {
	l.text = text
	l.drawText()
}

func (l *Label) Show() {
	l.isVisible = true
	l.drawText()
}

func (l *Label) Hide() {
	l.isVisible = false
	l.panel.Clear()
}

func (l *Label) drawText() {
	if !l.isVisible {
		return
	}
	l.panel.Clear()
	textRows := fonts.Generate(l.font, l.text)
	l.panel.OutputHAllignedStrings(textRows, l.hAlligment, l.style)
}
