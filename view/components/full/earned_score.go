package full

import (
	"tetrominos/settings"
	"tetrominos/view/components/common"
	"tetrominos/view/ui"
)

func NewEarnedScore(canvas *ui.Canvas) common.EarnedScore {
	return common.NewEarnedScore(
		canvas, sidePanelWidth+(settings.FieldWidth+2)*2+1, 0, 2,
		settings.FieldHeight,
	)
}
