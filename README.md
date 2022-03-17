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
	for i := range c.Data {                                             // 背景の塗りつぶし
		c.Data[i] = color.RGBA{R: 200, G: 200, B: 200, A: 255}
	}
	c.Line(image.Point{X: 100, Y: 100}, image.Point{X: 700, Y: 300})    // 線分の描画
	c.ToPng("./img/sample.png")
}

```
![](img/sample.png)
