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
	return b.X >= win.Position.X+item.Position.X &&
		b.X <= win.Position.X+item.Position.X+item.Size.X &&
		b.Y >= win.Position.Y+item.Position.Y &&
		b.Y <= win.Position.Y+item.Position.Y+item.Size.Y
}

func (win WindowData) GetWinRect() Rect {
	return Rect{
		X0: win.Position.X,
		Y0: win.Position.Y,
		X1: win.Position.X + win.SizeTemp.X,
		Y1: win.Position.Y + win.SizeTemp.Y - win.TitleSizeTemp,
	}
}

func (win WindowData) GetMainRect() Rect {
	return Rect{
		X0: win.Position.X,
		Y0: win.Position.Y + win.TitleSizeTemp + 1,
		X1: win.Position.X + win.SizeTemp.X,
		Y1: win.Position.Y + win.SizeTemp.Y - win.TitleSizeTemp,
	}
}

func (win WindowData) TitleRect() Rect {
	if win.TitleSize <= 0 {
		return Rect{}
	}
	return Rect{
		X0: win.Position.X, Y0: win.Position.Y,
		X1: win.Position.X + win.SizeTemp.X,
		Y1: win.Position.Y + win.TitleSizeTemp,
	}
}

func (win WindowData) XRect() Rect {
	if win.TitleSize <= 0 || !win.Closable {
		return Rect{}
	}

	var xpad float32 = win.Border
	return Rect{
		X0: win.Position.X + win.SizeTemp.X - win.TitleSizeTemp + xpad,
		Y0: win.Position.Y + xpad,

		X1: win.Position.X + win.SizeTemp.X - xpad,
		Y1: win.Position.Y + win.TitleSizeTemp - xpad,
	}
}

func (win WindowData) DragbarRect() Rect {
	if win.TitleSize <= 0 && !win.Resizable {
		return Rect{}
	}
	textSize := win.TitleTextWidth()
	xRect := win.XRect()
	buttonsWidth := xRect.X1 - xRect.X0 + 3

	dpad := (win.TitleSize * UIScale) / 5
	xStart := textSize.X + float32((win.TitleSize*UIScale)/1.5)
	xEnd := (win.SizeTemp.X - buttonsWidth)
	return Rect{
		X0: win.Position.X + xStart, Y0: win.Position.Y + dpad,
		X1: win.Position.X + xEnd, Y1: win.Position.Y + (win.TitleSize * UIScale) - dpad,
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
		X1: win.Position.X + win.SizeTemp.X,
		Y0: win.Position.Y + win.SizeTemp.Y - (14 * UIScale) - win.TitleSizeTemp,
		X0: win.Position.X + win.SizeTemp.X - (14 * UIScale),
		Y1: win.Position.Y + win.SizeTemp.Y - win.TitleSizeTemp,
	}
}

func (win WindowData) GetWindowEdge(mpos Point) WindowSide {

	var BR_Corner = (14 * UIScale)

	if !win.Resizable {
		return SIDE_NONE

	} else if WithinEdgeRange(mpos.X, win.Position.X, tol) &&
		mpos.Y > win.Position.Y && mpos.Y < win.Position.Y+win.SizeTemp.Y-win.TitleSizeTemp {
		return SIDE_LEFT

	} else if WithinEdgeRange(mpos.X, win.Position.X+win.SizeTemp.X, tol) &&
		mpos.Y > win.Position.Y && mpos.Y < win.Position.Y+win.SizeTemp.Y-win.TitleSizeTemp &&
		mpos.Y < win.Position.Y+win.SizeTemp.Y-win.TitleSizeTemp-BR_Corner {
		return SIDE_RIGHT

	} else if WithinEdgeRange(mpos.Y, win.Position.Y, tol) &&
		mpos.X > win.Position.X && mpos.X < win.Position.X+win.SizeTemp.X {
		return SIDE_TOP

	} else if WithinEdgeRange(mpos.Y, win.Position.Y+win.SizeTemp.Y-win.TitleSizeTemp, tol) &&
		mpos.X > win.Position.X && mpos.X < win.Position.X+win.SizeTemp.X-BR_Corner {
		return SIDE_BOTTOM
	}

	return SIDE_NONE
}

const tol = 3

func WithinEdgeRange(a, b float32, tol float32) bool {
	if math.Abs(float64(a-b)) <= float64(tol) {
		return true
	}
	return false
}

func (win WindowData) TitleTextWidth() Point {
	if win.TitleSize <= 0 {
		return Point{}
	}
	textSize := ((win.TitleSize * UIScale) / 1.5)
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

func PointSubract(a, b Point) Point {
	return Point{X: a.X - b.X, Y: a.Y - b.Y}
}

func PointScale(a Point) Point {
	return Point{X: a.X / UIScale, Y: a.Y / UIScale}
}

func FloatScale(a float32) float32 {
	return a / UIScale
}

func (win WindowData) GetSizeX() float32 {
	return win.Size.X * UIScale
}

func (win WindowData) GetSizeY() float32 {
	return win.Size.Y * UIScale
}

// Sets SizeTemp, TitleSizeTemp
func (win *WindowData) CalcUIScale() {
	win.SizeTemp = Point{X: win.Size.X * UIScale, Y: win.Size.Y * UIScale}
	win.TitleSizeTemp = win.TitleSize * UIScale
}

func (win *WindowData) SetTitleSize(size float32) {
	win.TitleSize = size / UIScale
}
