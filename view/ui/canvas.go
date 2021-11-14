package ui

import (
	"github.com/gdamore/tcell/v2"
)

type Layer int

type Canvas struct {
	screen tcell.Screen
	layers []*layer
	bgChar Char

	originX int
	originY int
	width   int
	height  int
}

func NewCanvas(screen tcell.Screen, originX, originY, width, height int,
	layers int, bgStyle tcell.Style,
) *Canvas {
	l := make([]*layer, layers)
	for i := 0; i < layers; i++ {
		l[i] = newLayer(width, height)
	}
	c := &Canvas{
		screen: screen,
		layers: l,
		bgChar: Char{R: ' ', Style: bgStyle},

		originX: originX,
		originY: originY,
		width:   width,
		height:  height,
	}

	return c
}

func (c *Canvas) CreatePanel(parent *Panel, x, y int, width, height int, layer Layer) Panel {
	if parent != nil {
		x += parent.canvasX
		y += parent.canvasY
	}

	if x+width-1 > c.width || y+height-1 > c.height {
		panic("out of bounds")
	}
	if layer < 0 || layer >= Layer(len(c.layers)) {
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

func (c *Canvas) CreatePanelInTheCenter(parent *Panel, width, height int, layer Layer) Panel {
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

// Sync outputs all characters, it's a very expensive operation. It should be
// used in events of resizing the terminal window.
func (c *Canvas) Sync() {
	c.screen.Clear()
	for layerIndex, layer := range c.layers {
		for x := 0; x < len(layer.rep); x++ {
			col := layer.rep[x]
			for y := 0; y < len(col); y++ {
				c.out(x, y, Layer(layerIndex), col[y])
			}
		}
	}
	c.Draw()
}

// ChangeOrigin changes the origin and syncs the screen. It's a very expensive
// operation.
func (c *Canvas) ChangeOrigin(newOriginX, newOriginY int) {
	c.originX = newOriginX
	c.originY = newOriginY
	c.Sync()
}

// MoveToCenter moves the canvas to the terminal center. It's a very expensive
// operation.
func (c *Canvas) MoveToCenter() {
	w, h := c.screen.Size()
	c.ChangeOrigin((w-c.width)/2, (h-c.height)/2)
}

func (c *Canvas) Width() int {
	return c.width
}

func (c *Canvas) Height() int {
	return c.height
}

func (c *Canvas) Screen() tcell.Screen {
	return c.screen
}

func (c *Canvas) out(x, y int, layer Layer, char *Char) {
	l := c.layers[layer]
	l.rep[x][y] = char
	for i := layer + 1; i < Layer(len(c.layers)); i++ {
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
