package input

type Input string

const (
	CtrlCKey = "Ctrl-C"
	EscKey   = "Esc"
	LeftKey  = "Left"
	RightKey = "Right"
	UpKey    = "Up"
	DownKey  = "Down"
	Enter    = "Enter"
)

type KeyDescription struct {
	Key         Input
	Description string
}
