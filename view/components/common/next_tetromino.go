package common

import (
	"tetrominos/tetrominos"
	"tetrominos/view/ui"
)

type NextTetromino struct {
	tetromino tetromino
}

func NewNextTetromino(canvas *ui.Canvas, x, y int) NextTetromino {
	return NextTetromino{
		tetromino: newTetromino(canvas, x, y, 8, 4, NextTetrominoLayer),
	}
}

func (t NextTetromino) Show(tetromino tetrominos.Tetromino) {
	t.tetromino.Hide()
	t.tetromino.Show(0, 0, tetromino)
}

func (t NextTetromino) Hide() {
	t.tetromino.Hide()
}
