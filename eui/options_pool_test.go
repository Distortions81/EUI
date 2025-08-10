package eui

import (
	"image"
	"reflect"
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
	text "github.com/hajimehoshi/ebiten/v2/text/v2"
)

func mutateDrawImageOptions(op *ebiten.DrawImageOptions) {
	op.GeoM.Translate(1, 1)
	op.ColorScale.Scale(1, 2, 3, 4)
	op.Filter = ebiten.FilterLinear
	op.CompositeMode = ebiten.CompositeModeSourceOver

	v := reflect.ValueOf(op).Elem()
	if f := v.FieldByName("Address"); f.IsValid() && f.CanSet() {
		f.Set(reflect.ValueOf(ebiten.AddressClampToZero))
	}
	if f := v.FieldByName("SourceRect"); f.IsValid() && f.CanSet() {
		r := image.Rect(1, 2, 3, 4)
		f.Set(reflect.ValueOf(&r))
	}
}

func TestAcquireDrawImageOptions(t *testing.T) {
	op := acquireDrawImageOptions()
	mutateDrawImageOptions(op)
	releaseDrawImageOptions(op)

	op = acquireDrawImageOptions()
	var expected ebiten.DrawImageOptions
	if !reflect.DeepEqual(*op, expected) {
		t.Errorf("acquired options not reset: %+v", op)
	}
}

func TestAcquireTextDrawOptions(t *testing.T) {
	op := acquireTextDrawOptions()
	mutateDrawImageOptions(&op.DrawImageOptions)
	op.LayoutOptions.LineSpacing = 1
	op.LayoutOptions.PrimaryAlign = text.AlignCenter
	releaseTextDrawOptions(op)

	op = acquireTextDrawOptions()
	var expected text.DrawOptions
	if !reflect.DeepEqual(*op, expected) {
		t.Errorf("acquired text options not reset: %+v", op)
	}
}

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
