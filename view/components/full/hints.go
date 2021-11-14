package full

import (
	"tetrominos/view/components/common"
	"tetrominos/view/ui"
)

func NewPauseHints(canvas *ui.Canvas) common.Hints {
	return createHints(canvas, []string{
		"Esc    Resume",
		"Ctrl-C Exit",
	})
}

func NewGameOverHints(canvas *ui.Canvas) common.Hints {
	return createHints(canvas, []string{
		"Esc    Start",
		"Ctrl-C Exit",
	})
}

func NewStartHints(canvas *ui.Canvas) common.Hints {
	return createHints(canvas, []string{
		"Esc    Start",
		"Ctrl-C Exit",
	})
}

func NewGameHints(canvas *ui.Canvas) common.Hints {
	return createHints(canvas, []string{
		"Esc    Pause",
		"Left   Move left",
		"Right  Move right",
		"Up     Rotate left",
		"Down   Rotate right",
		"Enter  Drop",
		"Ctrl-C Exit",
	})
}

func createHints(canvas *ui.Canvas, hints []string) common.Hints {
	return common.NewHints(canvas, controlsHintX, controlsHintY,
		controlsHintWidth, controlsHintHeight, "Controls:", hints)
}
