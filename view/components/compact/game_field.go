package compact

import (
	"tetrominos/view/components/common"
	"tetrominos/view/ui"
)

func NewGameField(canvas *ui.Canvas) common.GameField {
	return common.NewGameField(canvas, 0, 0)
}
