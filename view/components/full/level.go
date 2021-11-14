package full

import (
	"tetrominos/view/components/common"
	"tetrominos/view/fonts"
	"tetrominos/view/ui"
)

func NewLevel(canvas *ui.Canvas) common.Level {
	return common.NewLevel(canvas, fonts.Small, screenWidth-sidePanelWidth, 0,
		sidePanelWidth)
}
