package states

import "tetrominos/input"

var ctrlCDescription input.KeyDescription

func init() {
	ctrlCDescription = input.KeyDescription{
		Key:         input.CtrlCKey,
		Description: "Exit",
	}
}
