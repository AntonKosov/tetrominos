package colors

import "github.com/gdamore/tcell/v2"

const (
	BackgroundColor = tcell.ColorBlack
	TextColor       = tcell.ColorWhite
	MessageBoxColor = tcell.ColorDarkRed
	WallColor       = tcell.ColorWhite

	TetrominoIColor = tcell.ColorRed
	TetrominoJColor = tcell.ColorBlue
	TetrominoLColor = tcell.ColorOrange
	TetrominoOColor = tcell.ColorYellow
	TetrominoSColor = tcell.ColorDarkMagenta
	TetrominoTColor = tcell.ColorLightCyan
	TetrominoZColor = tcell.ColorGreen
)

func CreateBGStyle(bg tcell.Color) tcell.Style {
	return tcell.StyleDefault.Background(bg)
}

func CreateFontStyle(bg tcell.Color, fg tcell.Color) tcell.Style {
	return CreateBGStyle(bg).Foreground(fg)
}
