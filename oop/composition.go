package oop

import "fmt"

type Coordinates struct {
	X float64
	Y float64
}

func (c Coordinates) LogCoordinates() {
	fmt.Printf("x-%.2f, y-%.2f\n", c.X, c.Y)
}

type ColorPicker struct {
	Coordinates
	ColorRGBA string
}

func NewColorPicker(x, y float64) *ColorPicker {
	return &ColorPicker{
		Coordinates: Coordinates{X: x, Y: y},
		ColorRGBA:   "(0,0,255,0.5)",
	}

}

func (p ColorPicker) PickColor() {
	fmt.Printf("color is %s, coordinates are\n", p.ColorRGBA)
	p.LogCoordinates()
	p.Coordinates.LogCoordinates() // the same
}

var ColorPickerInstance = NewColorPicker(2.5424, 24.2314)
