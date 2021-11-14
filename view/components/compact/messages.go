package compact

import (
	"tetrominos/view/components/common"
	"tetrominos/view/fonts"
	"tetrominos/view/ui"
)

func NewTitle(canvas *ui.Canvas) common.Label {
	title := "TETROMINOS"
	return common.NewLabel(common.LabelParams{
		Canvas: canvas,
		Style:  common.TitleStyle,
		Layer:  common.MessageBoxLayer,
		Font:   fonts.Small,
		Text:   &title,
	})
}
