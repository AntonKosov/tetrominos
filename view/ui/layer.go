package ui

type layer struct {
	rep [][]*Char
}

func newLayer(width, height int) *layer {
	rep := make([][]*Char, width)
	for i := 0; i < width; i++ {
		rep[i] = make([]*Char, height)
	}

	l := &layer{
		rep: rep,
	}

	return l
}
