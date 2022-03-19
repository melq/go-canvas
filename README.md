# go-canvas

図を描画するのに利用できるモジュールを製作中

## Usage
```go
package main
import (
    "image"
    "image/color"
	
    "github.com/melq/go-canvas"
)

func main() {
	c := NewCanvas(800, 400)

	c.SetBackground(color.RGBA{R: 200, G: 200, B: 200, A: 255})
	p1 := image.Point{X: 100, Y: 100}
	p2 := image.Point{X: 700, Y: 300}
	p3 := image.Point{X: 400, Y: 50}
	p4 := image.Point{X: 600, Y: 350}
	p5 := image.Point{X: 200, Y: 350}
	c.Line(p1, p2)
	c.ChangeColor(color.RGBA{R: 200, G: 150, B: 150, A: 255})
	c.Rect(p1, p2)
	c.ChangeColor(color.RGBA{R: 150, G: 150, B: 200, A: 255})
	c.Triangle(p3, p4, p5)
	c.ToPng("./img/sample.png")
}

```
![](img/sample.png)
