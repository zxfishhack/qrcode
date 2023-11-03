package main

import (
	"github.com/zxfishhack/qrcode/result"
	"gocv.io/x/gocv"
	"math"
)

func setPointVector(r *result.Result, vp gocv.PointVector) {
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
