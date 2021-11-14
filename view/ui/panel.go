package ui

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
)

type HAlligment string
type VAlligment string

const (
	LeftAlligment    HAlligment = "left"
	HCenterAlligment HAlligment = "hcenter"
	RightAlligment   HAlligment = "right"
	TopAlligment     VAlligment = "top"
	VCenterAlligment VAlligment = "vcenter"
	BottomAlligment  VAlligment = "bottom"
)

type Panel struct {
	canvasX int
	canvasY int
	layer   Layer
	width   int
	height  int

	out func(x, y int, layer Layer, c *Char)
}

func (p Panel) Output(x, y int, c *Char) {
	if x < 0 || y < 0 || x >= p.width || y >= p.height {
		panic("out of bounds")
	}
	p.out(p.canvasX+x, p.canvasY+y, p.layer, c)
}

func (p Panel) OutputStr(x, y int, msg string, s tcell.Style) {
	for i, c := range msg {
		p.Output(x+i, y, &Char{
			R:     c,
			Style: s,
		})
	}
}

func (p Panel) OutputStrings(x, y int, text []string, s tcell.Style) {
	for i, t := range text {
		p.OutputStr(x, y+i, t, s)
	}
}

func (p Panel) OutputHAllignedStr(y int, text string, a HAlligment, s tcell.Style) {
	var x int
	switch a {
	case LeftAlligment:
		x = 0
	case HCenterAlligment:
		x = (p.width - len(text)) / 2
	case RightAlligment:
		x = p.width - len(text)
	default:
		panic(fmt.Sprintf("Unknown alligment: %v", a))
	}
	p.OutputStr(x, y, text, s)
}

func (p Panel) OutputHAllignedStrings(text []string, ha HAlligment, s tcell.Style) {
	for i, t := range text {
		p.OutputHAllignedStr(i, t, ha, s)
	}
}

func (p Panel) OutputAllignedStrings(text []string, ha HAlligment, va VAlligment, s tcell.Style) {
	var y int
	switch va {
	case TopAlligment:
		y = 0
	case VCenterAlligment:
		y = (p.height - len(text)) / 2
	case BottomAlligment:
		y = p.height - len(text)
	default:
		panic(fmt.Sprintf("Unknown allignment: %v", va))
	}

	for i, t := range text {
		p.OutputHAllignedStr(y+i, t, ha, s)
	}
}

func (p Panel) Clear() {
	for y := 0; y < p.height; y++ {
		p.ClearRow(y)
	}
}

func (p Panel) FillAll(s tcell.Style) {
	c := Char{R: ' ', Style: s}
	p.Fill(0, 0, p.width, p.height, &c)
}

func (p Panel) ClearRow(y int) {
	for x := 0; x < p.width; x++ {
		p.Output(x, y, nil)
	}
}

func (p Panel) Fill(x, y, width, height int, c *Char) {
	for i := x; i < x+width; i++ {
		for j := y; j < y+height; j++ {
			p.Output(i, j, c)
		}
	}
}

func (p Panel) Width() int {
	return p.width
}

func (p Panel) Height() int {
	return p.height
}
