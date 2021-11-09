package view

import (
	"strings"
	"tetrominos/input"
	"tetrominos/view/ui"
)

type controlHints struct {
	panel ui.Panel
}

func newControlHints(canvas *ui.Canvas) controlHints {
	p := canvas.CreatePanel(nil, 0, controlsHintY, sidePanelWidth, controlsHintHeight, 0)
	return controlHints{
		panel: p,
	}
}

func (c controlHints) output(hints []input.KeyDescription) {
	captionStyle := createFontStyle(backgroundColor, textColor).Bold(true)
	c.panel.OutputStr(0, 0, "Controls:", captionStyle)
	hintStyle := createFontStyle(backgroundColor, textColor)
	longestKeyName := 0
	for _, kd := range hints {
		l := len(kd.Key)
		if longestKeyName < l {
			longestKeyName = l
		}
	}
	for i, h := range hints {
		sb := strings.Builder{}
		sb.WriteString(string(h.Key))
		for sb.Len() < longestKeyName {
			sb.WriteRune(' ')
		}
		sb.WriteRune(' ')
		sb.WriteString(h.Description)
		c.panel.OutputStr(2, i+1, sb.String(), hintStyle)
	}
}

func (c controlHints) clear() {
	c.panel.Clear()
}
