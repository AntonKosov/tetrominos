package view

import "tetrominos/settings"

const (
	sidePanelWidth = 23
	// double sized characters and the walls
	screenWidth = sidePanelWidth*2 + (settings.FieldWidth+2)*2
	// the 25th line is the floor
	screenHeight       = settings.FieldHeight + 1
	layers             = 3
	controlsHintY      = settings.FieldHeight - 7
	controlsHintHeight = 8
)
