package compact

import (
	"tetrominos/view/components/common"
	"tetrominos/view/ui"
)

func NewNextTetromino(canvas *ui.Canvas) common.NextTetromino {
	return common.NewNextTetromino(canvas, sidePanelX, nextTetrominoY)
}
