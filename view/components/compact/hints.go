package compact

import (
	"tetrominos/view/components/common"
	"tetrominos/view/ui"
)

func NewPauseHints(canvas *ui.Canvas) common.Hints {
	return createHints(canvas, []string{
		"Esc",
		"  Resume",
		"Ctrl-C",
		"  Exit",
	})
}

func NewGameOverHints(canvas *ui.Canvas) common.Hints {
	return createHints(canvas, []string{
		"Esc",
		"  Start",
		"Ctrl-C",
		"  Exit",
	})
}

func NewStartHints(canvas *ui.Canvas) common.Hints {
	return createHints(canvas, []string{
		"Esc",
		"  Start",
		"Ctrl-C",
		"  Exit",
	})
}

func NewGameHints(canvas *ui.Canvas) common.Hints {
	return createHints(canvas, []string{
		"Esc",
		"  Pause",
		"Left",
		"  Move L",
		"Right",
		"  Move R",
		"Up",
		"  Rot L",
		"Down",
		"  Rot R",
		"Enter",
		"  Drop",
		"Ctrl-C",
		"  Exit",
	})
}

func createHints(canvas *ui.Canvas, hints []string) common.Hints {
	h := 1 + len(hints)
	return common.NewHints(canvas, sidePanelX, screenHeight-h,
		sidePanelWidth, h, "Controls", hints)
}
