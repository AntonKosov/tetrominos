package states

import (
	"tetrominos/tetrominos"
	"tetrominos/ticker"
)

type StartView interface {
	Activate()
	Deactivate()
}

type GameView interface {
	Activate(tickerGroup *ticker.Group)
	Deactivate()
	Resume()
	Draw(c, r int, t tetrominos.Tetromino)
	RemoveRows(rows []int, tr []tetrominos.FieldRow, earnedScore int)
	Move(oldC, oldR int, oldT tetrominos.Tetromino, newC, newR int, newT tetrominos.Tetromino)
	OutputScore(s int)
	OutputLevel(l int)
	OutputNextTetromino(t tetrominos.Tetromino)
}

type PauseView interface {
	Activate()
	Deactivate()
}

type GameOverView interface {
	Activate()
	Deactivate()
}
