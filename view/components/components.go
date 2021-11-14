package components

type Title interface {
	Show()
	Hide()
}

type Score interface {
	Show()
	OutputScore(score int)
}

type Level interface {
	Show()
	OutputLevel(level int)
}

type ControlHints interface {
	Show()
	Hide()
}

type Label interface {
	Show()
	Hide()
	SetText(text string)
}

type EarnedScore interface {
	Show(score int, y int)
	Hide()
}
