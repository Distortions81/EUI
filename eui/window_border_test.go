//go:build test

package eui

import (
	"image"
	"image/color"
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func TestWindowBorderScaledRendering(t *testing.T) {
	uiScale = 1.3
	defer func() { uiScale = 1 }()

	win := *defaultTheme
	win.Theme = baseTheme
	win.Margin = 0
	win.Fillet = 0
	win.Border = 2
	win.Outlined = true
	win.Size = normPoint(10, 10)

	w, h := int(win.GetSize().X), int(win.GetSize().Y)
	img := image.NewNRGBA(image.Rect(0, 0, w, h))

	strokeRectFn = func(_ *ebiten.Image, x, y, w, h, width float32, col color.Color, aa bool) {
		c := color.NRGBAModel.Convert(col).(color.NRGBA)
		ix, iy := int(x), int(y)
		iw, ih := int(w), int(h)
		for dx := 0; dx < iw; dx++ {
			img.SetNRGBA(ix+dx, iy, c)
			img.SetNRGBA(ix+dx, iy+ih-1, c)
		}
		for dy := 0; dy < ih; dy++ {
			img.SetNRGBA(ix, iy+dy, c)
			img.SetNRGBA(ix+iw-1, iy+dy, c)
		}
	}
	defer func() { strokeRectFn = vector.StrokeRect }()

	strokeRect(nil, 0, 0, float32(w), float32(h), win.Border, win.Theme.Window.BorderColor, false)

	border := color.NRGBAModel.Convert(win.Theme.Window.BorderColor.ToRGBA()).(color.NRGBA)
	for x := 0; x < w; x++ {
		if c := img.NRGBAAt(x, 0); c != border {
			t.Fatalf("top edge pixel (%d,0) = %#v want %#v", x, c, border)
		}
		if c := img.NRGBAAt(x, h-1); c != border {
			t.Fatalf("bottom edge pixel (%d,%d) = %#v want %#v", x, h-1, c, border)
		}
	}
	for y := 0; y < h; y++ {
		if c := img.NRGBAAt(0, y); c != border {
			t.Fatalf("left edge pixel (0,%d) = %#v want %#v", y, c, border)
		}
		if c := img.NRGBAAt(w-1, y); c != border {
			t.Fatalf("right edge pixel (%d,%d) = %#v want %#v", w-1, y, c, border)
		}
	}
}
