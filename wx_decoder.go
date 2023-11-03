package main

import (
	"github.com/zxfishhack/qrcode/result"
	"gocv.io/x/gocv"
	"gocv.io/x/gocv/contrib"
	"image"
)

var wxDetectors = NewPool[*contrib.WeChatQRCode](func() *contrib.WeChatQRCode {
	return contrib.NewWeChatQRCode("./models/detect.prototxt", "./models/detect.caffemodel", "./models/sr.prototxt", "./models/sr.caffemodel")
})

func WxDecodeQrCodeMulti(input gocv.Mat) (res []*result.Result) {
	detector := wxDetectors.Borrow()
	defer wxDetectors.Put(detector)
	var points []gocv.Mat
	contents := detector.DetectAndDecode(input, &points)

	for idx, c := range contents {
		v := &result.Result{Content: c}
		vp := gocv.NewPointVector()
		p := points[idx]
		for i := 0; i < p.Rows(); i++ {
			vp.Append(image.Pt(int(p.GetFloatAt(i, 0)), int(p.GetFloatAt(i, 1))))
		}
		setPointVector(v, vp)
		vp.Close()
		res = append(res, v)
	}
	return
}
