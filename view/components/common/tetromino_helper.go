package common

import (
	"tetrominos/tetrominos"
	"tetrominos/view/ui"
)

var tetrominoStyle map[tetrominos.TetrominoType]ui.Char

func init() {
	const r = ' '
	tetrominoStyle = map[tetrominos.TetrominoType]ui.Char{
		tetrominos.TetrominoI: {R: r, Style: createBGStyle(tetrominoIColor)},
		tetrominos.TetrominoJ: {R: r, Style: createBGStyle(tetrominoJColor)},
		tetrominos.TetrominoL: {R: r, Style: createBGStyle(tetrominoLColor)},
		tetrominos.TetrominoO: {R: r, Style: createBGStyle(tetrominoOColor)},
		tetrominos.TetrominoS: {R: r, Style: createBGStyle(tetrominoSColor)},
		tetrominos.TetrominoT: {R: r, Style: createBGStyle(tetrominoTColor)},
		tetrominos.TetrominoZ: {R: r, Style: createBGStyle(tetrominoZColor)},
	}
}

func outputTetrominoCell(col, row int, char *ui.Char, p ui.Panel) {
	p.Fill(col*2, row, 2, 1, char)
}

func prepareTetromino(c, r int, t tetrominos.Tetromino, clear bool, p ui.Panel) {
	var char *ui.Char
	if !clear {
		c := tetrominoStyle[t.Type]
		char = &c
	}
	for ty := 0; ty < len(t.Shape); ty++ {
		row := t.Shape[ty]
		for tx := 0; tx < len(row); tx++ {
			if row[tx] {
				outputTetrominoCell(c+tx, r+ty, char, p)
			}
		}
	}
}
