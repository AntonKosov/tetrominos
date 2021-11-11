package states

import (
	"tetrominos/input"
)

type gameOverState struct {
	params Params
	score  int
}

func newGameOverState(params Params, score int) State {
	gs := &gameOverState{
		params: params,
		score:  score,
	}

	return gs
}

func (s *gameOverState) Activate() {
	s.params.GameOverView.Activate()
	s.params.GameOverView.ShowControlHints([]input.KeyDescription{
		{
			Key:         input.EscKey,
			Description: "Start",
		},
		ctrlCDescription,
	})
}

func (s *gameOverState) Deactivate() {
	s.params.GameOverView.Deactivate()
}

func (s *gameOverState) HandleInput(in input.Input) {
	if in == input.EscKey {
		s.params.ChangeState <- newGameState(s.params)
	}
}
