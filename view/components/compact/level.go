package compact

import (
	"tetrominos/view/components/common"
	"tetrominos/view/fonts"
	"tetrominos/view/ui"
)

func NewLevel(canvas *ui.Canvas) common.Level {
	return common.NewLevel(canvas, fonts.Native, sidePanelX, levelPanelY, sidePanelWidth)
}
