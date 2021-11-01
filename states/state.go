package states

import (
	"tetrominos/input"
)

type State interface {
	Activate()
	Deactivate()
	HandleInput(input.Input)
}
