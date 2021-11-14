package components

import (
	"tetrominos/view/components/common"
	"tetrominos/view/components/full"
	"tetrominos/view/ui"

	"github.com/gdamore/tcell/v2"
)

type Factory struct {
	canvas               *ui.Canvas
	title                func(*ui.Canvas) common.Label
	score                func(*ui.Canvas) common.Score
	level                func(*ui.Canvas) common.Level
	startControlHints    func(*ui.Canvas) common.ControlHints
	gameControlHints     func(*ui.Canvas) common.ControlHints
	pauseControlHints    func(*ui.Canvas) common.ControlHints
	gameOverControlHints func(*ui.Canvas) common.ControlHints
	earnedScore          func(*ui.Canvas) common.EarnedScore
	pauseMessage         func(*ui.Canvas) common.Label
	gameOverMessage      func(*ui.Canvas) common.Label
}

func NewFullComponentsFactory(screen tcell.Screen) Factory {
	return Factory{
		canvas:               full.NewCanvas(screen),
		title:                full.NewTitle,
		score:                full.NewScore,
		level:                full.NewLevel,
		startControlHints:    full.NewStartControlHints,
		gameControlHints:     full.NewGameControlHints,
		pauseControlHints:    full.NewPauseControlHints,
		gameOverControlHints: full.NewGameOverControlHints,
		earnedScore:          full.NewEarnedScore,
		pauseMessage:         full.NewPauseMessage,
		gameOverMessage:      full.NewGameOverMessage,
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

func (f Factory) StartControlHints() ControlHints {
	return f.startControlHints(f.canvas)
}

func (f Factory) GameControlHints() ControlHints {
	return f.gameControlHints(f.canvas)
}

func (f Factory) PauseControlHints() ControlHints {
	return f.pauseControlHints(f.canvas)
}

func (f Factory) GameOverControlHints() ControlHints {
	return f.gameOverControlHints(f.canvas)
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
