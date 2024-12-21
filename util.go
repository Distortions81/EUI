package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

func (rect Rect) ContainsPoint(b Point) bool {
	return b.X >= rect.X0 && b.Y >= rect.Y0 &&
		b.X <= rect.X1 && b.Y <= rect.Y1
}

func (win WindowData) GetWinRect() Rect {
	return Rect{
		X0: win.Position.X,
		Y0: win.Position.Y,
		X1: win.Position.X + win.MagTemp.X,
		Y1: win.Position.Y + win.MagTemp.Y - win.TitleSizeTemp,
	}
}

func (win WindowData) GetMainRect() Rect {
	return Rect{
		X0: win.Position.X,
		Y0: win.Position.Y + win.TitleSizeTemp + 1,
		X1: win.Position.X + win.MagTemp.X,
		Y1: win.Position.Y + win.MagTemp.Y - win.TitleSizeTemp,
	}
}

func (win WindowData) TitleRect() Rect {
	if win.TitleSize <= 0 {
		return Rect{}
	}
	return Rect{
		X0: win.Position.X, Y0: win.Position.Y,
		X1: win.Position.X + win.MagTemp.X,
		Y1: win.Position.Y + win.TitleSizeTemp,
	}
}

func (win WindowData) XRect() Rect {
	if win.TitleSize <= 0 || !win.Closable {
		return Rect{}
	}

	var xpad float32 = win.Border
	return Rect{
		X0: win.Position.X + win.MagTemp.X - win.TitleSizeTemp + xpad,
		Y0: win.Position.Y + xpad,

		X1: win.Position.X + win.MagTemp.X - xpad,
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
	xEnd := (win.MagTemp.X - buttonsWidth)
	return Rect{
		X0: win.Position.X + xStart, Y0: win.Position.Y + dpad,
		X1: win.Position.X + xEnd, Y1: win.Position.Y + (win.TitleSize * UIScale) - dpad,
	}
}

func (win WindowData) ResizeTabRect() Rect {
	if !win.Resizable {
		return Rect{}
	}

	return Rect{
		X1: win.Position.X + win.MagTemp.X,
		Y0: win.Position.Y + win.MagTemp.Y - (14 * UIScale) - win.TitleSizeTemp,
		X0: win.Position.X + win.MagTemp.X - (14 * UIScale),
		Y1: win.Position.Y + win.MagTemp.Y - win.TitleSizeTemp,
	}
}

func (win WindowData) GetWindowEdge(mpos Point) WindowSide {

	if !win.Resizable {
		return SIDE_NONE
	}
	if WithinEdgeRange(mpos.X, win.Position.X, tol) &&
		mpos.Y > win.Position.Y && mpos.Y < win.Position.Y+win.MagTemp.Y-win.TitleSizeTemp {
		return SIDE_LEFT
	}
	if WithinEdgeRange(mpos.X, win.Position.X+win.MagTemp.X, tol) &&
		mpos.Y > win.Position.Y && mpos.Y < win.Position.Y+win.MagTemp.Y-win.TitleSizeTemp &&
		mpos.Y < win.Position.Y+win.MagTemp.Y-win.TitleSizeTemp-14 {
		return SIDE_RIGHT
	}
	if WithinEdgeRange(mpos.Y, win.Position.Y, tol) &&
		mpos.X > win.Position.X && mpos.X < win.Position.X+win.MagTemp.X {

		return SIDE_TOP
	}
	if WithinEdgeRange(mpos.Y, win.Position.Y+win.MagTemp.Y-win.TitleSizeTemp, tol) &&
		mpos.X > win.Position.X && mpos.X < win.Position.X+win.MagTemp.Y &&
		mpos.X < win.Position.X+win.MagTemp.X-14 {
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

func (win WindowData) TitleTextWidth() Magnatude {
	if win.TitleSize <= 0 {
		return Magnatude{}
	}
	textSize := ((win.TitleSize * UIScale) / 1.5)
	face := &text.GoTextFace{
		Source: mplusFaceSource,
		Size:   float64(textSize),
	}
	textWidth, textHeight := text.Measure(win.Title, face, 0)
	return Magnatude{X: float32(textWidth), Y: float32(textHeight)}
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

func MagAdd(a, b Magnatude) Magnatude {
	return Magnatude{X: a.X + b.X, Y: a.Y + b.Y}
}

func MagSubtract(a, b Magnatude) Magnatude {
	return Magnatude{X: a.X - b.X, Y: a.Y - b.Y}
}

func MagScale(a Magnatude) Magnatude {
	return Magnatude{X: a.X / UIScale, Y: a.Y / UIScale}
}

func (win WindowData) GetSizeX() float32 {
	return win.Mag.X * UIScale
}

func (win WindowData) GetSizeY() float32 {
	return win.Mag.Y * UIScale
}

// Sets MagTemp, TitleSizeTemp
func (win *WindowData) PreCalcSize() {
	win.MagTemp = Magnatude{X: win.Mag.X * UIScale, Y: win.Mag.Y * UIScale}
	win.TitleSizeTemp = win.TitleSize * UIScale
}

func (win *WindowData) SetSize(size Magnatude) {
	win.Mag = Magnatude{X: size.X / UIScale, Y: size.Y / UIScale}
	win.PreCalcSize()
}

func (win *WindowData) SetTitleSize(size float32) {
	win.TitleSize = size / UIScale
}
