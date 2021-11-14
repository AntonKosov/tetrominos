package components

import (
	"tetrominos/view/components/common"
	"tetrominos/view/components/compact"
	"tetrominos/view/components/full"
	"tetrominos/view/ui"

	"github.com/gdamore/tcell/v2"
)

type Factory struct {
	canvas          *ui.Canvas
	title           func(*ui.Canvas) common.Label
	score           func(*ui.Canvas) common.Score
	level           func(*ui.Canvas) common.Level
	startHints      func(*ui.Canvas) common.Hints
	gameHints       func(*ui.Canvas) common.Hints
	pauseHints      func(*ui.Canvas) common.Hints
	gameOverHints   func(*ui.Canvas) common.Hints
	earnedScore     func(*ui.Canvas) common.EarnedScore
	pauseMessage    func(*ui.Canvas) common.Label
	gameOverMessage func(*ui.Canvas) common.Label
	nextTetromino   func(*ui.Canvas) common.NextTetromino
	gameField       func(*ui.Canvas) common.GameField
}

func NewFullComponentsFactory(screen tcell.Screen) Factory {
	return Factory{
		canvas:          full.NewCanvas(screen),
		title:           full.NewTitle,
		score:           full.NewScore,
		level:           full.NewLevel,
		startHints:      full.NewStartHints,
		gameHints:       full.NewGameHints,
		pauseHints:      full.NewPauseHints,
		gameOverHints:   full.NewGameOverHints,
		earnedScore:     full.NewEarnedScore,
		pauseMessage:    full.NewPauseMessage,
		gameOverMessage: full.NewGameOverMessage,
		nextTetromino:   full.NewNextTetromino,
		gameField:       full.NewGameField,
	}
}

func NewCompactComponentsFactory(screen tcell.Screen) Factory {
	return Factory{
		canvas:          compact.NewCanvas(screen),
		title:           compact.NewTitle,
		score:           compact.NewScore,
		level:           compact.NewLevel,
		startHints:      compact.NewStartHints,
		gameHints:       compact.NewGameHints,
		pauseHints:      compact.NewPauseHints,
		gameOverHints:   compact.NewGameOverHints,
		earnedScore:     compact.NewEarnedScore,
		pauseMessage:    compact.NewPauseMessage,
		gameOverMessage: compact.NewGameOverMessage,
		nextTetromino:   compact.NewNextTetromino,
		gameField:       compact.NewGameField,
	}
}

func (f Factory) Canvas() *ui.Canvas {
	return f.canvas
}

func (f Factory) Title() Label {
	l := f.title(f.canvas)
	return &l
}

func (f Factory) Score() Score {
	s := f.score(f.canvas)
	return &s
}

func (f Factory) Level() Level {
	l := f.level(f.canvas)
	return &l
}

func (f Factory) StartHints() Hints {
	return f.startHints(f.canvas)
}

func (f Factory) GameHints() Hints {
	return f.gameHints(f.canvas)
}

func (f Factory) PauseHints() Hints {
	return f.pauseHints(f.canvas)
}

func (f Factory) GameOverHints() Hints {
	return f.gameOverHints(f.canvas)
}

func (f Factory) EarnedScore() EarnedScore {
	return f.earnedScore(f.canvas)
}

func (f Factory) PauseMessage() Label {
	m := f.pauseMessage(f.canvas)
	return &m
}

func (f Factory) GameOverMessage() Label {
	m := f.gameOverMessage(f.canvas)
	return &m
}

func (f Factory) NextTetromino() NextTetromino {
	return f.nextTetromino(f.canvas)
}

func (f Factory) GameField() GameField {
	return f.gameField(f.canvas)
}
