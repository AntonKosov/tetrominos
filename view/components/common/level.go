package common

import (
	"fmt"
	"tetrominos/math"
	"tetrominos/view/fonts"
	"tetrominos/view/ui"
)

type Level struct {
	titleLabel Label
	levelLabel Label
}

func NewLevel(canvas *ui.Canvas, font fonts.Font, x, y int, width int) Level {
	title := "LEVEL"
	alligment := ui.HCenterAlligment
	titleLabel := NewLabel(LabelParams{
		Canvas:     canvas,
		Style:      LevelStyle,
		Layer:      LevelLayer,
		Font:       font,
		Coordinate: &math.Vector{X: x, Y: y},
		Text:       &title,
		Width:      &width,
		HAlligment: &alligment,
	})
	levelLabel := NewLabel(LabelParams{
		Canvas:     canvas,
		Style:      LevelStyle,
		Layer:      LevelLayer,
		Font:       font,
		Width:      &width,
		Coordinate: &math.Vector{X: x, Y: y + titleLabel.Height()},
		HAlligment: &alligment,
	})
	return Level{
		titleLabel: titleLabel,
		levelLabel: levelLabel,
	}
}

func (l *Level) Show() {
	l.titleLabel.Show()
	l.levelLabel.Show()
	l.OutputLevel(0)
}

func (l *Level) OutputLevel(score int) {
	l.levelLabel.SetText(fmt.Sprintf("%v", score))
}
