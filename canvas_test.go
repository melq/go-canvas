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
	c := NewCanvas(100, 100)
	for i := range c.Data {
		c.Data[i] = color.RGBA{R: 100, G: 100, B: 100, A: 255}
	}
	c.ToPng("./img/test.png")
}
