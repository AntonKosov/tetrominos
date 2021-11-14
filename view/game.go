package view

import (
	"tetrominos/settings"
	"tetrominos/tetrominos"
	"tetrominos/ticker"
	"tetrominos/view/components"
	"tetrominos/view/ui"
	"time"
)

type gameComponentsFactory interface {
	Canvas() *ui.Canvas
	Score() components.Score
	Level() components.Level
	GameHints() components.Hints
	EarnedScore() components.EarnedScore
	NextTetromino() components.NextTetromino
	GameField() components.GameField
}

type game struct {
	canvas        *ui.Canvas
	score         components.Score
	level         components.Level
	earnedScore   components.EarnedScore
	nextTetromino components.NextTetromino
	gameField     components.GameField
	hints         components.Hints
	tickerGroup   *ticker.Group
}

func newGame(factory gameComponentsFactory) *game {
	g := game{
		canvas:        factory.Canvas(),
		score:         factory.Score(),
		level:         factory.Level(),
		earnedScore:   factory.EarnedScore(),
		nextTetromino: factory.NextTetromino(),
		gameField:     factory.GameField(),
		hints:         factory.GameHints(),
	}

	return &g
}

func (g *game) Activate(tickerGroup *ticker.Group) {
	g.tickerGroup = tickerGroup
	g.score.Show()
	g.level.Show()
	g.hints.Show()
	g.gameField.Reset()
	g.canvas.Draw()
}

func (g *game) Deactivate() {
	g.hints.Hide()
	g.canvas.Draw()
}

func (g *game) Resume() {
	g.hints.Show()
	g.canvas.Draw()
}

func (g *game) Draw(c, r int, t tetrominos.Tetromino) {
	g.gameField.Draw(c, r, t)
	g.canvas.Draw()
}

func (g *game) Move(oldC, oldR int, oldT tetrominos.Tetromino, newC, newR int, newT tetrominos.Tetromino) {
	g.gameField.Clear(oldC, oldR, oldT)
	g.gameField.Draw(newC, newR, newT)
	g.canvas.Draw()
}

func (g *game) OutputLevel(level int) {
	g.level.OutputLevel(level)
	g.canvas.Draw()
}

func (g *game) OutputScore(score int) {
	g.score.OutputScore(score)
	g.canvas.Draw()
}

func (g *game) OutputNextTetromino(t tetrominos.Tetromino) {
	g.nextTetromino.Show(t)
	g.canvas.Draw()
}

func (g game) RemoveRows(raws []int, fr []tetrominos.FieldRow, earnedScore int) {
	scoreY := raws[0]
	g.earnedScore.Show(earnedScore, scoreY)
	tickerID, ticker := g.tickerGroup.NewTicker(time.Millisecond * 100)
	defer g.tickerGroup.DeleteTicker(tickerID)
	for c := 0; c < settings.FieldWidth; c++ {
		for _, r := range raws {
			g.gameField.ClearCell(c, r)
		}
		g.canvas.Draw()
		<-ticker
	}
	g.gameField.DrawRows(fr)
	g.earnedScore.Hide()
	g.canvas.Draw()
}
