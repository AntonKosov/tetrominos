package common

import (
	"fmt"
	"tetrominos/math"
	"tetrominos/view/fonts"
	"tetrominos/view/ui"
)

type Score struct {
	titleLabel Label
	scoreLabel Label
}

func NewScore(canvas *ui.Canvas, font fonts.Font, x, y, width int) Score {
	title := "SCORE"
	alligment := ui.HCenterAlligment
	titleLabel := NewLabel(LabelParams{
		Canvas:     canvas,
		Style:      ScoreStyle,
		Layer:      ScoreLayer,
		Font:       font,
		Coordinate: &math.Vector{X: x, Y: y},
		Width:      &width,
		Text:       &title,
		HAlligment: &alligment,
	})
	scoreLabel := NewLabel(LabelParams{
		Canvas:     canvas,
		Style:      ScoreStyle,
		Layer:      ScoreLayer,
		Font:       font,
		Width:      &width,
		Coordinate: &math.Vector{X: x, Y: y + titleLabel.Height()},
		HAlligment: &alligment,
	})
	return Score{
		titleLabel: titleLabel,
		scoreLabel: scoreLabel,
	}
}

func (s *Score) Show() {
	s.titleLabel.Show()
	s.scoreLabel.Show()
	s.OutputScore(0)
}

func (s *Score) OutputScore(score int) {
	s.scoreLabel.SetText(fmt.Sprintf("%v", score))
}
