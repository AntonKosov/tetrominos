package states

import (
	"tetrominos/input"
)

type startState struct {
	params Params
}

func NewStartState(params Params) State {
	ss := &startState{
		params: params,
	}

	return ss
}

func (s *startState) Activate() {
	s.params.StartView.Activate()
}

func (s *startState) Deactivate() {
	s.params.StartView.Deactivate()
}

func (s *startState) HandleInput(in input.Input) {
	if in != input.EscKey {
		return
	}

	s.params.ChangeState <- newGameState(s.params)
}
