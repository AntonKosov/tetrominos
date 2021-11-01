package states

type Params struct {
	StartView
	GameView
	PauseView
	GameOverView
	ChangeState chan<- State
}
