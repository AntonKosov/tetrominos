package compact

import (
	"tetrominos/view/components/common"
	"tetrominos/view/fonts"
	"tetrominos/view/ui"
)

func NewScore(canvas *ui.Canvas) common.Score {
	return common.NewScore(canvas, fonts.Native, sidePanelX, scorePanelY, sidePanelWidth)
}
