package compact

import "tetrominos/settings"

const (
	gapBetweenPanels = 1
	sidePanelWidth   = 8
	sidePanelX       = (settings.FieldWidth+2)*2 + gapBetweenPanels
	// double sized characters and the walls
	screenWidth = sidePanelWidth + gapBetweenPanels + (settings.FieldWidth+2)*2
	// the 25th line is the floor
	screenHeight   = settings.FieldHeight + 1
	nextTetrominoY = 0
	scorePanelY    = nextTetrominoY + 4
	levelPanelY    = scorePanelY + 3
	earnedScoreX   = settings.FieldWidth + 2
)
