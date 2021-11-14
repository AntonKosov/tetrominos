package common

import "tetrominos/view/ui"

const (
	Layers = 3

	ControlHintsLayer  ui.Layer = 0
	ScoreLayer         ui.Layer = 0
	LevelLayer         ui.Layer = 0
	NextTetrominoLayer ui.Layer = 0
	WallsLayer         ui.Layer = 0

	FallingTetrominoLayer ui.Layer = 1
	EarnedScoreLayer      ui.Layer = 1

	MessageBoxLayer ui.Layer = 2
)
