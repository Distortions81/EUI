package eui

import "testing"

func BenchmarkAcquireDrawImageOptions(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		op := acquireDrawImageOptions()
		op.GeoM.Translate(1, 1)
		releaseDrawImageOptions(op)
	}
}

func BenchmarkAcquireTextDrawOptions(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		op := acquireTextDrawOptions()
		op.DrawImageOptions.GeoM.Translate(1, 1)
		releaseTextDrawOptions(op)
	}
}
