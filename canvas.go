package canvas

import (
	"errors"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

type Px struct {
	obj   bool
	color color.Color
}

type Canvas struct {
	W     int
	H     int
	Scale int
	Data  []Px
	Color color.Color
}

type Point struct {
	X int
	Y int
}

func NewCanvas(w int, h int) *Canvas {
	c := Canvas{
		W:     w,
		H:     h,
		Scale: 1,
		Data:  make([]Px, w*h),
		Color: color.Black,
	}
	c.SetBackground(color.White)
	return &c
}

func (c *Canvas) toImage() image.Image {
	img := image.NewRGBA(image.Rect(0, 0, c.W*c.Scale, c.H*c.Scale))
	for y := 0; y < c.H*c.Scale; y++ {
		for x := 0; x < c.W*c.Scale; x++ {
			img.Set(x, y, c.Data[y/c.Scale*c.W+x/c.Scale].color)
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

func (c *Canvas) SetScale(scale int) {
	c.Scale = scale
}

func (c *Canvas) SetBackground(color color.Color) {
	for i := range c.Data {
		c.Data[i].color = color
		c.Data[i].obj = true
	}
}

func (c *Canvas) ChangeColor(color color.Color) {
	c.Color = color
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (c *Canvas) setPx(i int, clr color.Color) {
	r1, g1, b1, a1 := clr.RGBA()
	r2, g2, b2, _ := c.Data[i].color.RGBA()
	r := r1*a1/255 + r2*(255-a1)/255
	g := g1*a1/255 + g2*(255-a1)/255
	b := b1*a1/255 + b2*(255-a1)/255
	c.Data[i].color = color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: 255}
	c.Data[i].obj = true
}

func (c *Canvas) Line(s Point, e Point) {
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
		c.setPx(y0*c.W+x0, c.Color)
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

func (c *Canvas) Rect(p1 Point, p2 Point) {
	c.Line(p1, Point{X: p2.X, Y: p1.Y})
	c.Line(Point{X: p2.X, Y: p1.Y}, p2)
	c.Line(p2, Point{X: p1.X, Y: p2.Y})
	c.Line(Point{X: p1.X, Y: p2.Y}, p1)
}

func (c *Canvas) Triangle(p1 Point, p2 Point, p3 Point) {
	c.Line(p1, p2)
	c.Line(p2, p3)
	c.Line(p3, p1)
}

func (c *Canvas) Shape(points ...Point) error {
	if len(points) < 2 {
		return errors.New("not enough points")
	}
	var s Point
	var e Point
	for i := 1; i < len(points); i++ {
		s = points[i-1]
		e = points[i]
		c.Line(s, e)
	}
	c.Line(e, points[0])
	return nil
}
