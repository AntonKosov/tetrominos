package full

import (
	"tetrominos/view/components/common"
	"tetrominos/view/ui"
)

func NewPauseControlHints(canvas *ui.Canvas) common.ControlHints {
	return createControlHints(canvas, []string{
		"Esc    Resume",
		"Ctrl-C Exit",
	})
}

func NewGameOverControlHints(canvas *ui.Canvas) common.ControlHints {
	return createControlHints(canvas, []string{
		"Esc    Start",
		"Ctrl-C Exit",
	})
}

func NewStartControlHints(canvas *ui.Canvas) common.ControlHints {
	return createControlHints(canvas, []string{
		"Esc    Start",
		"Ctrl-C Exit",
	})
}

func NewGameControlHints(canvas *ui.Canvas) common.ControlHints {
	return createControlHints(canvas, []string{
		"Esc    Pause",
		"Left   Move left",
		"Right  Move right",
		"Up     Rotate left",
		"Down   Rotate right",
		"Enter  Drop",
		"Ctrl-C Exit",
	})
}

func createControlHints(canvas *ui.Canvas, hints []string) common.ControlHints {
	return common.NewControlHints(canvas, controlsHintX, controlsHintY,
		controlsHintWidth, controlsHintHeight, hints)
}
