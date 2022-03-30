package canvas

import (
	"fmt"
	"image/color"
	"testing"
)

func TestNewCanvas(t *testing.T) {
	fmt.Printf("%#+v", NewCanvas(10, 10))
}

func TestToPng(t *testing.T) {
	c := NewCanvas(800, 400)

	c.SetBackground(color.RGBA{R: 200, G: 200, B: 200, A: 255})
	p1 := Point{X: 0, Y: 0}
	p2 := Point{X: 800, Y: 400}
	c.Line(p1, p2)
	c.ChangeColor(color.RGBA{R: 200, G: 150, B: 150, A: 255})
	p3 := Point{X: 100, Y: 100}
	p4 := Point{X: 700, Y: 300}
	c.Rect(p3, p4)
	p5 := Point{X: 400, Y: 50}
	p6 := Point{X: 600, Y: 350}
	p7 := Point{X: 200, Y: 350}
	c.ChangeColor(color.RGBA{R: 150, G: 150, B: 200, A: 255})
	c.Triangle(p5, p6, p7)
	p8 := Point{X: 400, Y: 100}
	p9 := Point{X: 467, Y: 167}
	p10 := Point{X: 467, Y: 234}
	p11 := Point{X: 400, Y: 300}
	p12 := Point{X: 333, Y: 234}
	p13 := Point{X: 333, Y: 167}
	c.ChangeColor(color.RGBA{R: 125, G: 175, B: 125, A: 255})
	c.Shape(p8, p9, p10, p11, p12, p13)
	c.SetScale(2)
	c.ToPng("./img/sample.png")
}
