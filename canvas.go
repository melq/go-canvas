package canvas

import (
	"image"
	"image/color"
)

type Canvas struct {
	W    int
	H    int
	Data []color.Color
}

func NewCanvas(w int, h int) *Canvas {
	return &Canvas{
		W:    w,
		H:    h,
		Data: make([]color.Color, w*h),
	}
}

func (c *Canvas) toImage() image.Image {
	img := image.NewRGBA(image.Rect(0, 0, c.W, c.H))
	for y := 0; y < c.H; y++ {
		for x := 0; x < c.W; x++ {
			img.Set(x, y, c.Data[y*c.W+x])
		}
	}
	return img
}

func (c *Canvas) ToPng(filename string) {

}
