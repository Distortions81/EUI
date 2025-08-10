package eui

import (
	"testing"

	etext "github.com/hajimehoshi/ebiten/v2/text/v2"
)

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

func TestDrawImageOptionsPoolResets(t *testing.T) {
	op := acquireDrawImageOptions()
	op.GeoM.Translate(1, 1)
	op.ColorScale.Scale(0.5, 0.25, 0.75, 0.5)
	releaseDrawImageOptions(op)

	op = acquireDrawImageOptions()
	if tx, ty := op.GeoM.Apply(0, 0); tx != 0 || ty != 0 {
		t.Fatalf("GeoM not reset: translation (%f,%f)", tx, ty)
	}
	if op.ColorScale.R() != 1 || op.ColorScale.G() != 1 || op.ColorScale.B() != 1 || op.ColorScale.A() != 1 {
		t.Fatalf("ColorScale not reset: %v", op.ColorScale)
	}
	releaseDrawImageOptions(op)
}

func TestTextDrawOptionsPoolResets(t *testing.T) {
	op := acquireTextDrawOptions()
	op.DrawImageOptions.GeoM.Translate(1, 1)
	op.DrawImageOptions.ColorScale.Scale(0.5, 0.25, 0.75, 0.5)
	op.LayoutOptions = etext.LayoutOptions{
		LineSpacing:    1,
		PrimaryAlign:   etext.AlignCenter,
		SecondaryAlign: etext.AlignEnd,
	}
	op.ColorScale.Scale(0.5, 0.5, 0.5, 0.5)
	releaseTextDrawOptions(op)

	op = acquireTextDrawOptions()
	if tx, ty := op.GeoM.Apply(0, 0); tx != 0 || ty != 0 {
		t.Fatalf("GeoM not reset: translation (%f,%f)", tx, ty)
	}
	if op.ColorScale.R() != 1 || op.ColorScale.G() != 1 || op.ColorScale.B() != 1 || op.ColorScale.A() != 1 {
		t.Fatalf("ColorScale not reset: %v", op.ColorScale)
	}
	if op.LayoutOptions != (etext.LayoutOptions{}) {
		t.Fatalf("LayoutOptions not reset: %#v", op.LayoutOptions)
	}
	releaseTextDrawOptions(op)
}
