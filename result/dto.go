package result

import (
	"gocv.io/x/gocv"
	"math"
)

type Result struct {
	Content string `json:"content"`
	X       int    `json:"x"`
	Y       int    `json:"y"`
	Width   int    `json:"width"`
	Height  int    `json:"height"`
}

func (r *Result) SetPointVector(vp gocv.PointVector) {
	var minX, minY, maxX, maxY = math.MaxInt, math.MaxInt, 0, 0
	for _, pt := range vp.ToPoints() {
		if pt.X < minX {
			minX = pt.X
		}
		if pt.Y < minY {
			minY = pt.Y
		}
		if pt.X > maxX {
			maxX = pt.X
		}
		if pt.Y > maxY {
			maxY = pt.Y
		}
	}
	r.X = minX
	r.Y = minY
	r.Width = maxX - minX
	r.Height = maxY - minY
}
