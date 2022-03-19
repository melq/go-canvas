package canvas

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

type Canvas struct {
	W     int
	H     int
	Data  []color.Color
	Color color.Color
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
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer func(f *os.File) {
		err = f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)
	err = png.Encode(f, c.toImage())
	if err != nil {
		log.Fatal(err)
	}
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (c *Canvas) Line(s image.Point, e image.Point) {
	dx := absInt(e.X - s.X)
	dy := absInt(e.Y - s.Y)

	var sx int
	var sy int
	if s.X < e.X {
		sx = 1
	} else {
		sx = -1
	}
	if s.Y < e.Y {
		sy = 1
	} else {
		sy = -1
	}

	err := dx - dy
	x0 := s.X
	y0 := s.Y
	for x0 != e.X || y0 != e.Y {
		c.Data[y0*c.W+x0] = color.Black
		e2 := 2 * err
		if e2 > -dy {
			err = err - dy
			x0 = x0 + sx
		}
		if e2 < dx {
			err = err + dx
			y0 = y0 + sy
		}
	}
}

func (c *Canvas) Rect(p1 image.Point, p2 image.Point) {
	c.Line(p1, image.Point{X: p2.X, Y: p1.Y})
	c.Line(image.Point{X: p2.X, Y: p1.Y}, p2)
	c.Line(p2, image.Point{X: p1.X, Y: p2.Y})
	c.Line(image.Point{X: p1.X, Y: p2.Y}, p1)
}

func (c *Canvas) Triangle(p1 image.Point, p2 image.Point, p3 image.Point) {
	c.Line(p1, p2)
	c.Line(p2, p3)
	c.Line(p3, p1)
}
