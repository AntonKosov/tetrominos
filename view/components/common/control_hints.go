package common

import (
	"tetrominos/view/ui"
)

type ControlHints struct {
	panel ui.Panel
	hints []string
}

func NewControlHints(canvas *ui.Canvas, x, y, width, height int, hints []string,
) ControlHints {
	p := canvas.CreatePanel(
		nil, x, y, width, height, ControlHintsLayer,
	)
	return ControlHints{
		panel: p,
		hints: hints,
	}
}

func (c ControlHints) Show() {
	c.panel.OutputStr(0, 0, "Controls:", ControlHintsCaptionStyle)
	c.panel.OutputStrings(0, 1, c.hints, ControlHintsStyle)
}

func (c ControlHints) Hide() {
	c.panel.Clear()
}
