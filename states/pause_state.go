package states

import "tetrominos/input"

type pauseState struct {
	params Params
	gState *gameState
}

func newPauseState(params Params, gState *gameState) *pauseState {
	s := &pauseState{
		params: params,
		gState: gState,
	}
	return s
}

func (s *pauseState) Activate() {
	s.params.PauseView.Activate()
}

func (s *pauseState) Deactivate() {
	s.params.PauseView.Deactivate()
}

func (s *pauseState) HandleInput(in input.Input) {
	if in != input.EscKey {
		return
	}

	s.params.ChangeState <- s.gState
}
