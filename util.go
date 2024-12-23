package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

func (rect Rect) ContainsPoint(b Point) bool {
	return b.X >= rect.X0 && b.Y >= rect.Y0 &&
		b.X <= rect.X1 && b.Y <= rect.Y1
}

func (item *ItemData) ContainsPoint(win *WindowData, b Point) bool {
	return b.X >= win.Position.X+(item.Position.X*UIScale) &&
		b.X <= win.Position.X+(item.Position.X*UIScale)+(item.Size.X*UIScale) &&
		b.Y >= win.Position.Y+(item.Position.Y*UIScale) &&
		b.Y <= win.Position.Y+(item.Position.Y*UIScale)+(item.Size.Y*UIScale)
}

func (win WindowData) GetWinRect() Rect {
	return Rect{
		X0: win.Position.X,
		Y0: win.Position.Y,
		X1: win.Position.X + win.ScreenSize.X,
		Y1: win.Position.Y + win.ScreenSize.Y - win.TitleScreenHeight,
	}
}

func (win WindowData) GetMainRect() Rect {
	return Rect{
		X0: win.Position.X,
		Y0: win.Position.Y + win.TitleScreenHeight + 1,
		X1: win.Position.X + win.ScreenSize.X,
		Y1: win.Position.Y + win.ScreenSize.Y - win.TitleScreenHeight,
	}
}

func (win WindowData) TitleRect() Rect {
	if win.TitleHeight <= 0 {
		return Rect{}
	}
	return Rect{
		X0: win.Position.X, Y0: win.Position.Y,
		X1: win.Position.X + win.ScreenSize.X,
		Y1: win.Position.Y + win.TitleScreenHeight,
	}
}

func (win WindowData) XRect() Rect {
	if win.TitleHeight <= 0 || !win.Closable {
		return Rect{}
	}

	var xpad float32 = win.Border
	return Rect{
		X0: win.Position.X + win.ScreenSize.X - win.TitleScreenHeight + xpad,
		Y0: win.Position.Y + xpad,

		X1: win.Position.X + win.ScreenSize.X - xpad,
		Y1: win.Position.Y + win.TitleScreenHeight - xpad,
	}
}

func (win WindowData) DragbarRect() Rect {
	if win.TitleHeight <= 0 && !win.Resizable {
		return Rect{}
	}
	textSize := win.TitleTextWidth()
	xRect := win.XRect()
	buttonsWidth := xRect.X1 - xRect.X0 + 3

	dpad := (win.TitleHeight * UIScale) / 5
	xStart := textSize.X + float32((win.TitleHeight*UIScale)/1.5)
	xEnd := (win.ScreenSize.X - buttonsWidth)
	return Rect{
		X0: win.Position.X + xStart, Y0: win.Position.Y + dpad,
		X1: win.Position.X + xEnd, Y1: win.Position.Y + (win.TitleHeight * UIScale) - dpad,
	}
}

func (win *WindowData) SetSize(size Point) bool {

	tooSmall := false
	if size.X < minWinSizeX {
		size.X = minWinSizeX
		tooSmall = true
	}

	if size.Y < minWinSizeY {
		size.Y = minWinSizeY
		tooSmall = true
	}
	win.Size = size
	win.CalcUIScale()

	return tooSmall
}

func (win WindowData) ResizeTabRect() Rect {
	if !win.Resizable {
		return Rect{}
	}

	return Rect{
		X1: win.Position.X + win.ScreenSize.X,
		Y0: win.Position.Y + win.ScreenSize.Y - (14 * UIScale) - win.TitleScreenHeight,
		X0: win.Position.X + win.ScreenSize.X - (14 * UIScale),
		Y1: win.Position.Y + win.ScreenSize.Y - win.TitleScreenHeight,
	}
}

func (win WindowData) GetWindowEdge(mpos Point) WindowEdge {

	var BR_Corner = (14 * UIScale)

	if !win.Resizable {
		return EDGE_NONE

	} else if WithinEdgeRange(mpos.X, win.Position.X, tol) &&
		mpos.Y > win.Position.Y && mpos.Y < win.Position.Y+win.ScreenSize.Y-win.TitleScreenHeight {
		return EDGE_LEFT

	} else if WithinEdgeRange(mpos.X, win.Position.X+win.ScreenSize.X, tol) &&
		mpos.Y > win.Position.Y && mpos.Y < win.Position.Y+win.ScreenSize.Y-win.TitleScreenHeight &&
		mpos.Y < win.Position.Y+win.ScreenSize.Y-win.TitleScreenHeight-BR_Corner {
		return EDGE_RIGHT

	} else if WithinEdgeRange(mpos.Y, win.Position.Y, tol) &&
		mpos.X > win.Position.X && mpos.X < win.Position.X+win.ScreenSize.X {
		return EDGE_TOP

	} else if WithinEdgeRange(mpos.Y, win.Position.Y+win.ScreenSize.Y-win.TitleScreenHeight, tol) &&
		mpos.X > win.Position.X && mpos.X < win.Position.X+win.ScreenSize.X-BR_Corner {
		return EDGE_BOTTOM
	}

	return EDGE_NONE
}

const tol = 3

func WithinEdgeRange(a, b float32, tol float32) bool {
	if math.Abs(float64(a-b)) <= float64(tol) {
		return true
	}
	return false
}

func (win WindowData) TitleTextWidth() Point {
	if win.TitleHeight <= 0 {
		return Point{}
	}
	textSize := ((win.TitleHeight * UIScale) / 1.5)
	face := &text.GoTextFace{
		Source: mplusFaceSource,
		Size:   float64(textSize),
	}
	textWidth, textHeight := text.Measure(win.Title, face, 0)
	return Point{X: float32(textWidth), Y: float32(textHeight)}
}

func PointAdd(a, b Point) Point {
	return Point{X: a.X + b.X, Y: a.Y + b.Y}
}

func PointMul(a, b Point) Point {
	return Point{X: a.X * b.X, Y: a.Y * b.Y}
}

func PointScaleMul(a Point) Point {
	return Point{X: a.X * UIScale, Y: a.Y * UIScale}
}

func PointScaleDiv(a Point) Point {
	return Point{X: a.X / UIScale, Y: a.Y / UIScale}
}

func PointDiv(a, b Point) Point {
	return Point{X: a.X / b.X, Y: a.Y / b.Y}
}

func PointSub(a, b Point) Point {
	return Point{X: a.X - b.X, Y: a.Y - b.Y}
}

func (win WindowData) GetSizeX() float32 {
	return win.Size.X * UIScale
}

func (win WindowData) GetSizeY() float32 {
	return win.Size.Y * UIScale
}

// Sets SizeTemp, TitleSizeTemp
func (win *WindowData) CalcUIScale() {
	win.ScreenSize = Point{X: win.Size.X * UIScale, Y: win.Size.Y * UIScale}
	win.TitleScreenHeight = win.TitleHeight * UIScale
}

func (win *WindowData) SetTitleSize(size float32) {
	win.TitleHeight = size / UIScale
}
