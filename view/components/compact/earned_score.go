package compact

import (
	"tetrominos/settings"
	"tetrominos/view/components/common"
	"tetrominos/view/ui"
)

func NewEarnedScore(canvas *ui.Canvas) common.EarnedScore {
	return common.NewEarnedScore(canvas, earnedScoreX, 0, 2, settings.FieldHeight)
}
