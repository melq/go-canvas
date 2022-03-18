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
		c.Data[i] = color.RGBA{R: 220, G: 200, B: 200, A: 255}
	}
	p1 := image.Point{X: 100, Y: 100}
	p2 := image.Point{X: 700, Y: 300}
	c.Line(p1, p2)
	c.Rect(p1, p2)
	c.ToPng("./img/sample.png")
}
