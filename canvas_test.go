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
	c := NewCanvas(800, 400)
	for i := range c.Data {
		c.Data[i] = color.RGBA{R: 200, G: 200, B: 200, A: 255}
	}
	c.Line(image.Point{X: 100, Y: 100}, image.Point{X: 700, Y: 300})
	c.ToPng("./img/sample.png")
}
