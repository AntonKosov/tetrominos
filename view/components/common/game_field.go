package common

import (
	"tetrominos/settings"
	"tetrominos/tetrominos"
	"tetrominos/view/ui"
)

type GameField struct {
	walls ui.Panel
	panel ui.Panel
}

func NewGameField(canvas *ui.Canvas, x, y int) GameField {
	return GameField{
		walls: canvas.CreatePanel(nil, x, y, (settings.FieldWidth+2)*2,
			settings.FieldHeight+1, WallsLayer),
		panel: canvas.CreatePanel(nil, x+2, y,
			settings.FieldWidth*2, settings.FieldHeight, GameFieldLayer),
	}
}

func (f GameField) Draw(x, y int, tetromino tetrominos.Tetromino) {
	prepareTetromino(x, y, tetromino, false, f.panel)
}

func (f GameField) Clear(x, y int, tetromino tetrominos.Tetromino) {
	prepareTetromino(x, y, tetromino, true, f.panel)
}

func (f GameField) ClearCell(x, y int) {
	outputTetrominoCell(x, y, nil, f.panel)
}

func (f GameField) DrawRows(rows []tetrominos.FieldRow) {
	for i, tr := range rows {
		f.drawRaw(i, tr)
	}
}

func (f GameField) Hide() {
	f.panel.Clear()
	f.walls.Clear()
}

func (f GameField) Reset() {
	wallChar := &ui.Char{R: ' ', Style: WallStyle}
	f.walls.Fill(0, 0, 2, f.walls.Height(), wallChar)
	f.walls.Fill(f.walls.Width()-2, 0, 2, f.panel.Height(), wallChar)
	f.walls.Fill(0, settings.FieldHeight, (settings.FieldWidth+2)*2, 1, wallChar)
	f.panel.Clear()
}

func (f GameField) drawRaw(r int, fr tetrominos.FieldRow) {
	for c := 0; c < settings.FieldWidth; c++ {
		var char *ui.Char
		if fr[c] != nil {
			c := tetrominoStyle[*fr[c]]
			char = &c
		}

		outputTetrominoCell(c, r, char, f.panel)
	}
}
