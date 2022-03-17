package canvas

import (
	"fmt"
	"image"
	"image/color"
	"testing"
)

func TestNewCanvas(t *testing.T) {
	fmt.Printf("%#+v", NewCanvas(10, 10))
}

func TestToPng(t *testing.T) {
	c := NewCanvas(1000, 1000)
	for i := range c.Data {
		c.Data[i] = color.RGBA{R: 150, G: 100, B: 100, A: 255}
	}
	c.Line(image.Point{X: 200, Y: 400}, image.Point{X: 800, Y: 600})
	c.ToPng("./img/test.png")
}
