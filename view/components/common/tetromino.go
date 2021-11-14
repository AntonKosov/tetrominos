package common

import (
	"tetrominos/tetrominos"
	"tetrominos/view/ui"
)

type tetromino struct {
	panel ui.Panel
}

func newTetromino(canvas *ui.Canvas, x, y, width, height int, layer ui.Layer,
) tetromino {
	return tetromino{
		panel: canvas.CreatePanel(nil, x, y, width, height, layer),
	}
}

func (t tetromino) Show(x, y int, tetromino tetrominos.Tetromino) {
	prepareTetromino(x, y, tetromino, false, t.panel)
}

func (t tetromino) Clear(x, y int, tetromino tetrominos.Tetromino) {
	prepareTetromino(x, y, tetromino, true, t.panel)
}

func (t tetromino) Hide() {
	t.panel.Clear()
}
