package common

import "github.com/gdamore/tcell/v2"

const (
	backgroundColor = tcell.ColorBlack
	textColor       = tcell.ColorWhite
	messageBoxColor = tcell.ColorDarkRed
	wallColor       = tcell.ColorWhite

	tetrominoIColor = tcell.ColorRed
	tetrominoJColor = tcell.ColorBlue
	tetrominoLColor = tcell.ColorOrange
	tetrominoOColor = tcell.ColorYellow
	tetrominoSColor = tcell.ColorDarkMagenta
	tetrominoTColor = tcell.ColorLightCyan
	tetrominoZColor = tcell.ColorGreen
)

func createBGStyle(bg tcell.Color) tcell.Style {
	return tcell.StyleDefault.Background(bg)
}

func createFontStyle(bg tcell.Color, fg tcell.Color) tcell.Style {
	return createBGStyle(bg).Foreground(fg)
}
