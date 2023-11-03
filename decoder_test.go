package main

import (
	"gocv.io/x/gocv"
	"testing"
)

func BenchmarkDecodeQrCode(b *testing.B) {
	m := gocv.IMRead("test.png", gocv.IMReadColor)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		WxDecodeQrCodeMulti(m)
	}
	b.ReportAllocs()
}

func TestDecodeQrCode(t *testing.T) {
	m := gocv.IMRead("multi.png", gocv.IMReadColor)
	res := DecodeQrCode(m)
	t.Logf("%#v", res)
}

func TestDecodeQrCodeMulti(t *testing.T) {
	m := gocv.IMRead("multi.png", gocv.IMReadColor)
	res := DecodeQrCodeMulti(m)
	for _, r := range res {
		t.Logf("%#v", r)
	}
}

func TestWxDecodeQrCodeMulti(t *testing.T) {
	m := gocv.IMRead("multi.png", gocv.IMReadColor)
	res := WxDecodeQrCodeMulti(m)
	for _, r := range res {
		t.Logf("%#v", r)
	}
}

func TestWxDecodeQrCode2(t *testing.T) {
	m := gocv.IMRead("test.png", gocv.IMReadColor)
	res := WxDecodeQrCodeMulti(m)
	for _, r := range res {
		t.Logf("%#v", r)
	}
}
