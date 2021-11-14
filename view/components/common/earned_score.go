package common

import (
	"fmt"
	"tetrominos/view/ui"
)

type EarnedScore struct {
	panel ui.Panel
}

func NewEarnedScore(canvas *ui.Canvas, x, y int, width, height int) EarnedScore {
	return EarnedScore{
		panel: canvas.CreatePanel(nil, x, y, width, height, EarnedScoreLayer),
	}
}

func (es EarnedScore) Show(score int, y int) {
	es.panel.OutputStr(0, y, fmt.Sprintf("+%v", score), EarnedScoreStyle)
}

func (es EarnedScore) Hide() {
	es.panel.Clear()
}
