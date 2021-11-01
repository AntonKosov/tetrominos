package ui

import (
	"github.com/gdamore/tcell/v2"
)

type Canvas struct {
	screen tcell.Screen
	layers []*layer
	bgChar Char

	originX int
	originY int
	width   int
	height  int
}

func NewCanvas(screen tcell.Screen, x, y, width, height int, layers int, bgStyle tcell.Style) *Canvas {
	l := make([]*layer, layers)
	for i := 0; i < layers; i++ {
		l[i] = newLayer(width, height)
	}
	c := &Canvas{
		screen: screen,
		layers: l,
		bgChar: Char{
			R:     ' ',
			Style: bgStyle,
		},

		originX: x,
		originY: y,
		width:   width,
		height:  height,
	}

	return c
}

func (c *Canvas) CreatePanel(parent *Panel, x, y int, width, height int, layer int) Panel {
	if parent != nil {
		x += parent.canvasX
		y += parent.canvasY
	}

	if x+width-1 > c.width || y+height-1 > c.height {
		panic("out of bounds")
	}
	if layer < 0 || layer >= len(c.layers) {
		panic("wrong layer index")
	}

	return Panel{
		canvasX: x,
		canvasY: y,
		layer:   layer,
		width:   width,
		height:  height,
		out:     c.out,
	}
}

func (c *Canvas) CreatePanelInTheCenter(parent *Panel, width, height int, layer int) Panel {
	return c.CreatePanel(
		parent,
		(c.width-width)/2,
		(c.height-height)/2,
		width,
		height,
		layer,
	)
}
func (c *Canvas) Draw() {
	c.screen.Show()
}

func (c *Canvas) out(x, y int, layer int, char *Char) {
	l := c.layers[layer]
	l.rep[x][y] = char
	for i := layer + 1; i < len(c.layers); i++ {
		ul := c.layers[i]
		if ul.rep[x][y] != nil {
			return
		}
	}
	if char != nil {
		c.outChar(x, y, *char)
		return
	}
	for i := layer - 1; i >= 0; i-- {
		ul := c.layers[i]
		ch := ul.rep[x][y]
		if ch != nil {
			c.outChar(x, y, *ch)
			return
		}
	}
	c.outChar(x, y, c.bgChar)
}

func (c *Canvas) outChar(x, y int, char Char) {
	c.screen.SetContent(c.originX+x, c.originY+y, char.R, nil, char.Style)
}
