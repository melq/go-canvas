package canvas

type Canvas struct {
	H    int
	W    int
	Data []int
}

/*type point struct {
	x int
	y int
}*/

func NewCanvas(h int, w int) *Canvas {
	return &Canvas{
		H:    h,
		W:    w,
		Data: make([]int, h*w),
	}
}

/*func (c *Canvas) Line(s point, e point) {

}*/
