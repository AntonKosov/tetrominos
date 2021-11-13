package states

import (
	"tetrominos/input"
	"tetrominos/tetrominos"
	"tetrominos/ticker"
)

type StartView interface {
	Activate()
	Deactivate()
	ShowControlHints(hints []input.KeyDescription)
}

type GameView interface {
	Activate(tickerGroup *ticker.Group)
	Deactivate()
	Draw(c, r int, t tetrominos.Tetromino)
	RemoveRows(rows []int, tr []tetrominos.FieldRow, earnedScore int)
	Move(oldC, oldR int, oldT tetrominos.Tetromino, newC, newR int, newT tetrominos.Tetromino)
	OutputScore(s int)
	OutputLevel(l int)
	OutputNextTetromino(t tetrominos.Tetromino)
	ShowControlHints(hints []input.KeyDescription)
}

type PauseView interface {
	Activate()
	Deactivate()
	ShowControlHints(hints []input.KeyDescription)
}

type GameOverView interface {
	Activate()
	Deactivate()
	ShowControlHints(hints []input.KeyDescription)
}
