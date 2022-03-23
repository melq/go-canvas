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

	c.SetBackground(color.RGBA{R: 200, G: 200, B: 200, A: 255})
	p1 := image.Point{X: 100, Y: 100}
	p2 := image.Point{X: 700, Y: 300}
	c.Line(p1, p2)
	c.ChangeColor(color.RGBA{R: 200, G: 150, B: 150, A: 255})
	c.Rect(p1, p2)
	p3 := image.Point{X: 400, Y: 50}
	p4 := image.Point{X: 600, Y: 350}
	p5 := image.Point{X: 200, Y: 350}
	c.ChangeColor(color.RGBA{R: 150, G: 150, B: 200, A: 255})
	c.Triangle(p3, p4, p5)
	p6 := image.Point{X: 400, Y: 100}
	p7 := image.Point{X: 467, Y: 167}
	p8 := image.Point{X: 467, Y: 234}
	p9 := image.Point{X: 400, Y: 300}
	p10 := image.Point{X: 333, Y: 234}
	p11 := image.Point{X: 333, Y: 167}
	c.ChangeColor(color.RGBA{R: 150, G: 200, B: 150, A: 255})
	c.Shape(p6, p7, p8, p9, p10, p11)
	c.ToPng("./img/sample.png")
}
