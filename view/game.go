package view

import (
	"fmt"
	"tetrominos/input"
	"tetrominos/settings"
	"tetrominos/tetrominos"
	"tetrominos/view/fonts"
	"tetrominos/view/ui"

	"github.com/gdamore/tcell/v2"
)

type game struct {
	canvas             *ui.Canvas
	tetrominoPanel     ui.Panel
	scoreTextPanel     ui.Panel
	scorePanel         ui.Panel
	levelTextPanel     ui.Panel
	levelPanel         ui.Panel
	nextTetrominoPanel ui.Panel
	hints              controlHints
	scoreText          []string
}

func newGame(canvas *ui.Canvas) game {
	drawContainer(canvas)

	scoreText := fonts.Generate(fonts.Small, "SCORE")

	g := game{
		canvas: canvas,
		tetrominoPanel: canvas.CreatePanel(
			nil, sidePanelWidth+2, 0, settings.FieldWidth*2, settings.FieldHeight, 1,
		),
		scoreTextPanel: canvas.CreatePanel(
			nil, 0, 0, sidePanelWidth, len(scoreText), 0,
		),
		scorePanel: canvas.CreatePanel(
			nil, 0, len(scoreText), sidePanelWidth, len(scoreText), 0,
		),
		levelTextPanel: canvas.CreatePanel(
			nil, screenWidth-sidePanelWidth, 0, sidePanelWidth, len(scoreText), 0,
		),
		levelPanel: canvas.CreatePanel(
			nil, screenWidth-sidePanelWidth, len(scoreText), sidePanelWidth, len(scoreText), 0,
		),
		nextTetrominoPanel: canvas.CreatePanel(
			nil, screenWidth-sidePanelWidth+(sidePanelWidth-8)/2,
			screenHeight-6, 8, 4, 0,
		),
		hints:     newControlHints(canvas),
		scoreText: scoreText,
	}

	return g
}

var scoreStyle tcell.Style
var tetrimonoStyle map[tetrominos.TetrominoType]ui.Char

func init() {
	scoreStyle = tcell.StyleDefault.
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorWhite).Bold(true)
	const r = ' '
	tetrimonoStyle = map[tetrominos.TetrominoType]ui.Char{
		tetrominos.TetrominoI: {
			R:     r,
			Style: tcell.StyleDefault.Background(tcell.ColorRed),
		},
		tetrominos.TetrominoJ: {
			R:     r,
			Style: tcell.StyleDefault.Background(tcell.ColorBlue),
		},
		tetrominos.TetrominoL: {
			R:     r,
			Style: tcell.StyleDefault.Background(tcell.ColorOrange),
		},
		tetrominos.TetrominoO: {
			R:     r,
			Style: tcell.StyleDefault.Background(tcell.ColorYellow),
		},
		tetrominos.TetrominoS: {
			R:     r,
			Style: tcell.StyleDefault.Background(tcell.ColorDarkMagenta),
		},
		tetrominos.TetrominoT: {
			R:     r,
			Style: tcell.StyleDefault.Background(tcell.ColorLightCyan),
		},
		tetrominos.TetrominoZ: {
			R:     r,
			Style: tcell.StyleDefault.Background(tcell.ColorGreen),
		},
	}
}

func (g *game) Activate() {
	g.tetrominoPanel.Clear()
	g.scoreTextPanel.OutputAllignedStrings(
		g.scoreText, ui.HCenterAlligment, ui.VCenterAlligment, scoreStyle,
	)
	g.levelTextPanel.OutputAllignedStrings(
		fonts.Generate(fonts.Small, "LEVEL"),
		ui.HCenterAlligment, ui.VCenterAlligment, scoreStyle,
	)
	g.OutputScore(0)
	g.OutputLevel(0)
	g.canvas.Draw()
}

func (g *game) Deactivate() {
	g.hints.clear()
	g.canvas.Draw()
}

func (g *game) Draw(c, r int, t tetrominos.Tetromino) {
	prepareTetromino(c, r, t, false, g.tetrominoPanel)
	g.canvas.Draw()
}

func (g *game) Move(oldC, oldR int, oldT tetrominos.Tetromino, newC, newR int, newT tetrominos.Tetromino) {
	prepareTetromino(oldC, oldR, oldT, true, g.tetrominoPanel)
	prepareTetromino(newC, newR, newT, false, g.tetrominoPanel)
	g.canvas.Draw()
}

func (g *game) OutputLevel(l int) {
	levelText := fonts.Generate(fonts.Small, fmt.Sprint(l))
	g.levelPanel.Clear()
	g.levelPanel.OutputAllignedStrings(
		levelText, ui.HCenterAlligment, ui.VCenterAlligment, scoreStyle,
	)
	g.canvas.Draw()
}

func (g *game) OutputScore(s int) {
	scoreText := fonts.Generate(fonts.Small, fmt.Sprint(s))
	g.scorePanel.Clear()
	g.scorePanel.OutputAllignedStrings(
		scoreText, ui.HCenterAlligment, ui.VCenterAlligment, scoreStyle,
	)
	g.canvas.Draw()
}

func (g *game) OutputNextTetromino(t tetrominos.Tetromino) {
	g.nextTetrominoPanel.Clear()
	prepareTetromino(0, 0, t, false, g.nextTetrominoPanel)
	g.canvas.Draw()
}

func (g *game) DrawRaws(startRaw int, fr []tetrominos.FieldRow) {
	for i, tr := range fr {
		g.drawRaw(startRaw+i, tr)
	}
	g.canvas.Draw()
}

func (g *game) ShowControlHints(hints []input.KeyDescription) {
	g.hints.output(hints)
	g.canvas.Draw()
}

func (g *game) drawRaw(r int, fr tetrominos.FieldRow) {
	for c := 0; c < settings.FieldWidth; c++ {
		var char *ui.Char
		if fr[c] != nil {
			c := tetrimonoStyle[*fr[c]]
			char = &c
		}

		outputCell(c, r, char, g.tetrominoPanel)
	}
}

func prepareTetromino(c, r int, t tetrominos.Tetromino, clear bool, p ui.Panel) {
	var char *ui.Char
	if !clear {
		c := tetrimonoStyle[t.Type]
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
	container := canvas.CreatePanel(nil, x, 0, width, height, 0)
	wallChar := &ui.Char{
		R:     ' ',
		Style: tcell.StyleDefault.Background(tcell.ColorWhite),
	}
	container.Fill(0, 0, 2, height, wallChar)
	container.Fill(width-2, 0, 2, height, wallChar)
	container.Fill(2, settings.FieldHeight, settings.FieldWidth*2, 1, wallChar)
	canvas.Draw()
}
