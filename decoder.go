package main

import (
	"gocv.io/x/gocv"
	"image"
	"math"
)

func distance(p1 image.Point, p2 image.Point) float64 {
	return math.Sqrt(math.Pow(float64(p2.X-p1.X), 2) + math.Pow(float64(p2.Y-p1.Y), 2))
}

func DecodeQrCode(input gocv.Mat) (content string) {
	detector := gocv.NewQRCodeDetector()
	defer detector.Close()
	corners := gocv.NewMat()
	defer corners.Close()
	res := detector.Detect(input, &corners)
	if !res {
		return
	}
	vp := gocv.NewPointVectorFromMat(corners)
	defer vp.Close()
	rr := gocv.MinAreaRect(vp)
	pointRect := gocv.NewPointVectorFromPoints(rr.Points)
	defer pointRect.Close()
	srcVp := gocv.NewPointVector()
	defer srcVp.Close()
	for _, pDst := range pointRect.ToPoints() {
		var p image.Point
		var dist = float64(input.Rows() + input.Cols())
		for _, pSrc := range vp.ToPoints() {
			d := distance(pSrc, pDst)
			if d < dist {
				p = pSrc
				dist = d
			}
		}
		srcVp.Append(p)
	}
	transform := gocv.GetPerspectiveTransform(srcVp, pointRect)
	defer transform.Close()
	dst := gocv.NewMat()
	defer dst.Close()
	gocv.WarpPerspective(input, &dst, transform, image.Point{X: input.Cols(), Y: input.Rows()})
	dummy := gocv.NewMat()
	defer dummy.Clone()
	return detector.Decode(dst, dummy, &corners)
}
