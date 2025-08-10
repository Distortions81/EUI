package eui

import "testing"

func TestSetTitleSizeZeroBeforeOpen(t *testing.T) {
	uiScale = 1
	screenWidth, screenHeight = 800, 600
	windows = nil

	win := NewWindow()
	win.AutoSize = true
	item := &itemData{ItemType: ITEM_BUTTON, Size: point{X: 100, Y: 50}}
	win.addItemTo(item)

	win.SetTitleSize(0)
	win.Open()

	if got := win.GetTitleSize(); got != 0 {
		t.Fatalf("title size got %v want %v", got, 0)
	}

	pad := (win.Padding + win.BorderPad) * uiScale
	wantHeight := item.GetSize().Y + 2*pad
	if got := win.GetSize().Y; got != wantHeight {
		t.Fatalf("window height got %v want %v", got, wantHeight)
	}

	windows = nil
}
