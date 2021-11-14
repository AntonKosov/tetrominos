package common

import (
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
var WallStyle tcell.Style

func init() {
	TitleStyle = createFontStyle(messageBoxColor, textColor).Bold(true)
	ScoreStyle = createFontStyle(backgroundColor, textColor).Bold(true)
	LevelStyle = ScoreStyle
	BackgroundStyle = createFontStyle(backgroundColor, textColor)
	ControlHintsCaptionStyle = createFontStyle(backgroundColor, textColor).Bold(true)
	ControlHintsStyle = createFontStyle(backgroundColor, textColor)
	EarnedScoreStyle = ScoreStyle
	PauseMessageStyle = createFontStyle(messageBoxColor, textColor).Bold(true).Blink(true)
	GameOverMessageStyle = TitleStyle
	WallStyle = createBGStyle(wallColor)
}
