package engine

import (
	"tetrominos/input"
	"tetrominos/states"
)

type Views struct {
	Start    states.StartView
	Game     states.GameView
	Pause    states.PauseView
	GameOver states.GameOverView
}

type eng struct {
	currentState  states.State
	changeStateCh chan states.State
}

func Run(views Views, in <-chan input.Input) {
	e := eng{
		changeStateCh: make(chan states.State, 1),
	}

	stateParams := states.Params{
		StartView:    views.Start,
		GameView:     views.Game,
		PauseView:    views.Pause,
		GameOverView: views.GameOver,

		ChangeState: e.changeStateCh,
	}
	e.changeStateCh <- states.NewStartState(stateParams)

	e.runInputHandler(in)
}

func (e *eng) runInputHandler(in <-chan input.Input) {
	for {
		select {
		case key := <-in:
			switch key {
			case input.CtrlCKey:
				return
			default:
				e.currentState.HandleInput(key)
			}
		case state := <-e.changeStateCh:
			e.changeState(state)
		}
	}
}

func (e *eng) changeState(s states.State) {
	if e.currentState != nil {
		e.currentState.Deactivate()
	}
	e.currentState = s
	e.currentState.Activate()
}
