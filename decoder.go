package main

import (
	"github.com/zxfishhack/qrcode/result"
	"gocv.io/x/gocv"
	"image"
	"math"
)

var detectors = NewPool[*gocv.QRCodeDetector](func() *gocv.QRCodeDetector {
	v := gocv.NewQRCodeDetector()
	return &v
})

func distance(p1 image.Point, p2 image.Point) float64 {
	return math.Sqrt(math.Pow(float64(p2.X-p1.X), 2) + math.Pow(float64(p2.Y-p1.Y), 2))
}

func decodeQrCode(detector *gocv.QRCodeDetector, input gocv.Mat, vp gocv.PointVector) (res *result.Result) {
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
	rect := gocv.NewMatWithSize(pointRect.Size(), 1, gocv.MatTypeCV32SC2)
	defer rect.Clone()
	for idx, p := range pointRect.ToPoints() {
		rect.SetIntAt(idx, 0, int32(p.X))
		rect.SetIntAt(idx, 1, int32(p.Y))
	}
	straightQrcode := gocv.NewMat()
	defer straightQrcode.Close()
	content := detector.Decode(dst, rect, &straightQrcode)

	res = &result.Result{
		Content: content,
	}
	res.SetPointVector(vp)
	return
}

// DecodeQrCode
// Result: First QrCode in image
func DecodeQrCode(input gocv.Mat) (res *result.Result) {
	detector := detectors.Borrow()
	defer detectors.Put(detector)
	corners := gocv.NewMat()
	defer corners.Close()
	found := detector.Detect(input, &corners)
	if !found {
		return
	}
	vp := gocv.NewPointVectorFromMat(corners)
	defer vp.Close()
	return decodeQrCode(detector, input, vp)
}

// DecodeQrCodeMulti
// Result: csv format
// content,Left,Top,Width,Height
func DecodeQrCodeMulti(input gocv.Mat) (res []*result.Result) {
	detector := detectors.Borrow()
	defer detectors.Put(detector)
	corners := gocv.NewMat()
	defer corners.Close()
	found := detector.DetectMulti(input, &corners)
	if !found {
		return
	}
	for j := 0; j < corners.Rows(); j++ {
		vp := gocv.NewPointVector()
		for i := 0; i < corners.Cols(); i++ {
			p := corners.GetVecfAt(j, i)
			vp.Append(image.Pt(int(p[0]), int(p[1])))
		}
		v := decodeQrCode(detector, input, vp)
		vp.Close()
		if v != nil {
			res = append(res, v)
		}
	}
	return
}
