package full

import (
	"tetrominos/view/components/common"
	"tetrominos/view/ui"
)

func NewNextTetromino(canvas *ui.Canvas) common.NextTetromino {
	return common.NewNextTetromino(canvas,
		screenWidth-sidePanelWidth+(sidePanelWidth-8)/2,
		screenHeight-6,
	)
}
