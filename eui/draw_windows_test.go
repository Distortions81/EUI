//go:build test

package eui

import (
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
)

// TestDrawWindowsSkipsClosed ensures that Draw renders only open windows.
func TestDrawWindowsSkipsClosed(t *testing.T) {
	winPortal := &windowData{MainPortal: true, open: true, Size: point{X: 50, Y: 50}}
	winVisible := &windowData{open: true, Size: point{X: 50, Y: 50}}
	winHidden := &windowData{open: false, Size: point{X: 50, Y: 50}}

	winPortal.AddWindow(false)
	winVisible.AddWindow(false)
	winHidden.AddWindow(false)

	Draw(ebiten.NewImage(100, 100))

	if winPortal.Render == nil {
		t.Fatalf("expected winPortal to be rendered")
	}
	if winVisible.Render == nil {
		t.Fatalf("expected winVisible to be rendered")
	}
	if winHidden.Render != nil {
		t.Fatalf("expected winHidden not to be rendered")
	}

	windows = nil
	activeWindow = nil
}
