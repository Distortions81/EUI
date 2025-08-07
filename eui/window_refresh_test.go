//go:build test

package eui

import (
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
)

func TestWindowRefreshRerenders(t *testing.T) {
	DebugMode = true
	defer func() { DebugMode = false }()

	textItem := *defaultText
	textItem.Text = "before"
	textItem.Theme = baseTheme

	win := *defaultTheme
	win.Theme = baseTheme
	win.Contents = []*itemData{&textItem}
	win.Open = true

	windows = []*windowData{&win}
	screen := ebiten.NewImage(200, 200)

	win.Dirty = true
	Draw(screen)
	rc0 := textItem.RenderCount

	textItem.Text = "after"
	win.Refresh()
	Draw(screen)

	if textItem.RenderCount <= rc0 {
		t.Fatalf("expected render count to increase after Refresh")
	}
}
