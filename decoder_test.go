package main

import (
	"gocv.io/x/gocv"
	"testing"
)

func BenchmarkDecodeQrCode(b *testing.B) {
	m := gocv.IMRead("test.png", gocv.IMReadColor)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DecodeQrCode(m)
	}
	b.ReportAllocs()
}
