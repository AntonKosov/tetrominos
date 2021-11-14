package compact

import "tetrominos/settings"

const (
	scorePanelY    = 5
	sidePanelWidth = 9
	// double sized characters and the walls
	screenWidth = sidePanelWidth + (settings.FieldWidth+2)*2
	// the 25th line is the floor
	screenHeight = settings.FieldHeight + 1
)
