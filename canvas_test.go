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
	p1 := image.Point{X: 100, Y: 100}
	p2 := image.Point{X: 700, Y: 300}
	c.Line(p1, p2)
	fmt.Println(c.Color)
	c.ChangeColor(color.RGBA{R: 200, G: 150, B: 150, A: 255})
	c.Rect(p1, p2)
	fmt.Println(c.Color)
	c.ChangeColor(color.RGBA{R: 150, G: 150, B: 200, A: 255})
	c.Triangle(image.Point{X: 400, Y: 50}, image.Point{X: 600, Y: 350}, image.Point{X: 200, Y: 350})
	fmt.Println(c.Color)
	c.ToPng("./img/sample.png")
}
