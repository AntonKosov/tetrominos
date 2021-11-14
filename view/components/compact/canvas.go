package compact

import (
	"tetrominos/view/components/common"
	"tetrominos/view/ui"

	"github.com/gdamore/tcell/v2"
)

func NewCanvas(screen tcell.Screen) *ui.Canvas {
	canvas := ui.NewCanvas(
		screen, 0, 0, screenWidth, screenHeight, common.Layers, common.BackgroundStyle,
	)
	return canvas
}
