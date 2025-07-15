package main

import "testing"

func TestWithinRange(t *testing.T) {
	if !withinRange(1, 1.1, 0.2) {
		t.Errorf("expected true")
	}
	if withinRange(1, 1.3, 0.2) {
		t.Errorf("expected false")
	}
}

func TestPointOperations(t *testing.T) {
	a := point{X: 1, Y: 2}
	b := point{X: 3, Y: 4}
	if r := pointAdd(a, b); r.X != 4 || r.Y != 6 {
		t.Errorf("pointAdd result %+v", r)
	}
	if r := pointSub(b, a); r.X != 2 || r.Y != 2 {
		t.Errorf("pointSub result %+v", r)
	}
	if r := pointMul(a, b); r.X != 3 || r.Y != 8 {
		t.Errorf("pointMul result %+v", r)
	}
	if r := pointDiv(b, a); r.X != 3 || r.Y != 2 {
		t.Errorf("pointDiv result %+v", r)
	}
	uiScale = 2
	if r := pointScaleMul(a); r.X != 2 || r.Y != 4 {
		t.Errorf("pointScaleMul result %+v", r)
	}
	if r := pointScaleDiv(point{X: 4, Y: 6}); r.X != 2 || r.Y != 3 {
		t.Errorf("pointScaleDiv result %+v", r)
	}
	uiScale = 1
}

func TestUnionRect(t *testing.T) {
	a := rect{X0: 0, Y0: 0, X1: 10, Y1: 10}
	b := rect{X0: 5, Y0: 5, X1: 15, Y1: 20}
	exp := rect{X0: 0, Y0: 0, X1: 15, Y1: 20}
	if r := unionRect(a, b); r != exp {
		t.Errorf("unionRect got %+v want %+v", r, exp)
	}
}

func TestMergeData(t *testing.T) {
	orig := &windowData{Title: "orig", Size: point{X: 10, Y: 10}, TitleHeight: 5}
	upd := &windowData{Title: "new", Size: point{X: 20, Y: 30}}
	res := mergeData(orig, upd).(*windowData)
	if res.Title != "new" {
		t.Errorf("Title=%v", res.Title)
	}
	if res.Size != (point{X: 20, Y: 30}) {
		t.Errorf("Size=%v", res.Size)
	}
	if res.TitleHeight != 5 {
		t.Errorf("TitleHeight=%v", res.TitleHeight)
	}
}

func TestPinPositions(t *testing.T) {
	screenWidth = 800
	screenHeight = 600
	win := &windowData{Position: point{X: 10, Y: 10}, Size: point{X: 100, Y: 80}, TitleHeight: 10}
	var pin pinType = PIN_TOP_RIGHT
	pos := pin.getWinPosition(win)
	exp := point{X: 800 - win.GetSize().X - win.GetPos().X, Y: win.GetPos().Y}
	if pos != exp {
		t.Errorf("top right got %+v want %+v", pos, exp)
	}
	item := &itemData{Position: point{X: 0, Y: 0}, Size: point{X: 20, Y: 20}, PinTo: PIN_BOTTOM_CENTER}
	res := item.PinTo.getItemPosition(win, item)
	expItem := point{X: win.GetSize().X/2 - item.GetSize().X/2 + item.GetPos().X,
		Y: win.GetSize().Y - win.GetTitleSize() - item.GetSize().Y - item.GetPos().Y}
	if res != expItem {
		t.Errorf("item position got %+v want %+v", res, expItem)
	}
}

func TestItemOverlap(t *testing.T) {
	win := &windowData{Size: point{X: 100, Y: 100}, Position: point{X: 0, Y: 0}}
	a := &itemData{Position: point{X: 0, Y: 0}, Size: point{X: 60, Y: 60}}
	b := &itemData{Position: point{X: 50, Y: 50}, Size: point{X: 60, Y: 60}}
	win.Contents = []*itemData{a, b}
	xc, yc := win.itemOverlap(win.Size)
	if !xc || !yc {
		t.Errorf("expected overlap got %v %v", xc, yc)
	}
	b.Position = point{X: 70, Y: 70}
	xc, yc = win.itemOverlap(win.Size)
	if xc || yc {
		t.Errorf("expected no overlap got %v %v", xc, yc)
	}
}

func TestSetSliderValue(t *testing.T) {
	item := &itemData{MinValue: 0, MaxValue: 10, AuxSize: point{X: 8}, AuxSpace: 4}
	item.DrawRect = rect{X0: 0, Y0: 0, X1: 100, Y1: 20}
	item.setSliderValue(point{X: 42})
	if item.Value < 4.9 || item.Value > 5.1 {
		t.Errorf("slider value got %v", item.Value)
	}
	item.IntOnly = true
	item.setSliderValue(point{X: 42})
	if item.Value != 5 {
		t.Errorf("int slider got %v", item.Value)
	}
}

func TestMarkOpen(t *testing.T) {
	win1 := &windowData{Title: "win1", Open: true}
	win2 := &windowData{Title: "win2", Open: false}
	windows = []*windowData{win2, win1}
	activeWindow = win1
	win2.MarkOpen()
	if !win2.Open {
		t.Errorf("expected window to be open")
	}
	if activeWindow != win2 {
		t.Errorf("expected active window to be win2")
	}
	if len(windows) != 2 || windows[1] != win2 {
		t.Errorf("window order incorrect: %v", windows)
	}
}
