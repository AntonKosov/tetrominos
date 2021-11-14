package common

import (
	"tetrominos/view/colors"

	"github.com/gdamore/tcell/v2"
)

var TitleStyle tcell.Style
var ScoreStyle tcell.Style
var LevelStyle tcell.Style
var BackgroundStyle tcell.Style
var ControlHintsCaptionStyle tcell.Style
var ControlHintsStyle tcell.Style
var EarnedScoreStyle tcell.Style
var PauseMessageStyle tcell.Style
var GameOverMessageStyle tcell.Style

func init() {
	TitleStyle = colors.
		CreateFontStyle(colors.MessageBoxColor, colors.TextColor).
		Bold(true)
	ScoreStyle = colors.CreateFontStyle(
		colors.BackgroundColor, colors.TextColor).Bold(true)
	LevelStyle = ScoreStyle
	BackgroundStyle = colors.CreateFontStyle(colors.BackgroundColor,
		colors.TextColor)
	ControlHintsCaptionStyle = colors.
		CreateFontStyle(colors.BackgroundColor, colors.TextColor).Bold(true)
	ControlHintsStyle = colors.CreateFontStyle(
		colors.BackgroundColor, colors.TextColor)
	EarnedScoreStyle = ScoreStyle
	PauseMessageStyle = colors.
		CreateFontStyle(colors.MessageBoxColor, colors.TextColor).
		Bold(true).Blink(true)
	GameOverMessageStyle = TitleStyle
}
