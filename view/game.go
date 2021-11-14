package view

import (
	"tetrominos/settings"
	"tetrominos/tetrominos"
	"tetrominos/ticker"
	"tetrominos/view/colors"
	"tetrominos/view/components"
	"tetrominos/view/components/common"
	"tetrominos/view/ui"
	"time"
)

type gameComponentsFactory interface {
	Canvas() *ui.Canvas
	Score() components.Score
	Level() components.Level
	GameControlHints() components.ControlHints
	EarnedScore() components.EarnedScore
}

type game struct {
	canvas                *ui.Canvas
	score                 components.Score
	level                 components.Level
	earnedScore           components.EarnedScore
	fallingTetrominoPanel ui.Panel
	nextTetrominoPanel    ui.Panel
	hints                 components.ControlHints
	tickerGroup           *ticker.Group
}

func newGame(factory gameComponentsFactory) *game {
	drawContainer(factory.Canvas())

	g := game{
		canvas:      factory.Canvas(),
		score:       factory.Score(),
		level:       factory.Level(),
		earnedScore: factory.EarnedScore(),
		fallingTetrominoPanel: factory.Canvas().CreatePanel(
			nil, sidePanelWidth+2, 0,
			settings.FieldWidth*2, settings.FieldHeight, common.FallingTetrominoLayer,
		),
		nextTetrominoPanel: factory.Canvas().CreatePanel(
			nil, screenWidth-sidePanelWidth+(sidePanelWidth-8)/2,
			screenHeight-6, 8, 4, common.NextTetrominoLayer,
		),
		hints: factory.GameControlHints(),
	}

	return &g
}

var tetrominoStyle map[tetrominos.TetrominoType]ui.Char

func init() {
	const r = ' '
	tetrominoStyle = map[tetrominos.TetrominoType]ui.Char{
		tetrominos.TetrominoI: {
			R:     r,
			Style: colors.CreateBGStyle(colors.TetrominoIColor),
		},
		tetrominos.TetrominoJ: {
			R:     r,
			Style: colors.CreateBGStyle(colors.TetrominoJColor),
		},
		tetrominos.TetrominoL: {
			R:     r,
			Style: colors.CreateBGStyle(colors.TetrominoLColor),
		},
		tetrominos.TetrominoO: {
			R:     r,
			Style: colors.CreateBGStyle(colors.TetrominoOColor),
		},
		tetrominos.TetrominoS: {
			R:     r,
			Style: colors.CreateBGStyle(colors.TetrominoSColor),
		},
		tetrominos.TetrominoT: {
			R:     r,
			Style: colors.CreateBGStyle(colors.TetrominoTColor),
		},
		tetrominos.TetrominoZ: {
			R:     r,
			Style: colors.CreateBGStyle(colors.TetrominoZColor),
		},
	}
}

func (g *game) Activate(tickerGroup *ticker.Group) {
	g.tickerGroup = tickerGroup
	g.fallingTetrominoPanel.Clear()
	g.score.Show()
	g.level.Show()
	g.hints.Show()
	g.canvas.Draw()
}

func (g *game) Deactivate() {
	g.hints.Hide()
	g.canvas.Draw()
}

func (g *game) Draw(c, r int, t tetrominos.Tetromino) {
	prepareTetromino(c, r, t, false, g.fallingTetrominoPanel)
	g.canvas.Draw()
}

func (g *game) Move(oldC, oldR int, oldT tetrominos.Tetromino, newC, newR int, newT tetrominos.Tetromino) {
	prepareTetromino(oldC, oldR, oldT, true, g.fallingTetrominoPanel)
	prepareTetromino(newC, newR, newT, false, g.fallingTetrominoPanel)
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
	g.nextTetrominoPanel.Clear()
	prepareTetromino(0, 0, t, false, g.nextTetrominoPanel)
	g.canvas.Draw()
}

func (g game) RemoveRows(raws []int, fr []tetrominos.FieldRow, earnedScore int) {
	scoreY := raws[0]
	g.earnedScore.Show(earnedScore, scoreY)
	tickerID, ticker := g.tickerGroup.NewTicker(time.Millisecond * 100)
	defer g.tickerGroup.DeleteTicker(tickerID)
	for c := 0; c < settings.FieldWidth; c++ {
		for _, r := range raws {
			outputCell(c, r, nil, g.fallingTetrominoPanel)
		}
		g.canvas.Draw()
		<-ticker
	}
	for i, tr := range fr {
		g.drawRaw(i, tr)
	}
	g.earnedScore.Hide()
	g.canvas.Draw()
}

func (g *game) drawRaw(r int, fr tetrominos.FieldRow) {
	for c := 0; c < settings.FieldWidth; c++ {
		var char *ui.Char
		if fr[c] != nil {
			c := tetrominoStyle[*fr[c]]
			char = &c
		}

		outputCell(c, r, char, g.fallingTetrominoPanel)
	}
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
				outputCell(c+tx, r+ty, char, p)
			}
		}
	}
}

func outputCell(col, row int, char *ui.Char, p ui.Panel) {
	p.Fill(col*2, row, 2, 1, char)
}

func drawContainer(canvas *ui.Canvas) {
	const x = sidePanelWidth
	const width = (settings.FieldWidth + 2) * 2
	const height = settings.FieldHeight + 1
	walls := canvas.CreatePanel(nil, x, 0, width, height, common.WallsLayer)
	wallChar := &ui.Char{
		R:     ' ',
		Style: colors.CreateBGStyle(colors.WallColor),
	}
	walls.Fill(0, 0, 2, height, wallChar)
	walls.Fill(width-2, 0, 2, height, wallChar)
	walls.Fill(2, settings.FieldHeight, settings.FieldWidth*2, 1, wallChar)
	canvas.Draw()
}
