package common

import (
	"tetrominos/view/ui"
)

type Hints struct {
	panel ui.Panel
	title string
	hints []string
}

func NewHints(canvas *ui.Canvas, x, y, width, height int, title string,
	hints []string,
) Hints {
	p := canvas.CreatePanel(
		nil, x, y, width, height, ControlHintsLayer,
	)
	return Hints{
		panel: p,
		title: title,
		hints: hints,
	}
}

func (c Hints) Show() {
	c.panel.OutputStr(0, 0, c.title, ControlHintsCaptionStyle)
	c.panel.OutputStrings(0, 1, c.hints, ControlHintsStyle)
}

func (c Hints) Hide() {
	c.panel.Clear()
}
