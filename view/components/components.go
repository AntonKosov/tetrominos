package components

import "tetrominos/tetrominos"

type Title interface {
	Show()
	Hide()
}

type Score interface {
	Show()
	OutputScore(score int)
}

type Level interface {
	Show()
	OutputLevel(level int)
}

type Hints interface {
	Show()
	Hide()
}

type Label interface {
	Show()
	Hide()
	SetText(text string)
}

type EarnedScore interface {
	Show(score int, y int)
	Hide()
}

type NextTetromino interface {
	Show(tetromino tetrominos.Tetromino)
	Hide()
}

type GameField interface {
	Draw(x, y int, tetromino tetrominos.Tetromino)
	DrawRows(rows []tetrominos.FieldRow)
	Clear(x, y int, tetromino tetrominos.Tetromino)
	ClearCell(x, y int)
	Reset()
	Hide()
}
