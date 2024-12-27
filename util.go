package main

import (
	"image"
	"math"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

func (rect Rect) containsPoint(b Point) bool {
	return b.X >= rect.X0 && b.Y >= rect.Y0 &&
		b.X <= rect.X1 && b.Y <= rect.Y1
}

func (item *ItemData) containsPoint(win *WindowData, b Point) bool {
	return b.X >= win.Position.X+(item.Position.X*uiScale) &&
		b.X <= win.Position.X+(item.Position.X*uiScale)+(item.Size.X*uiScale) &&
		b.Y >= win.Position.Y+(item.Position.Y*uiScale) &&
		b.Y <= win.Position.Y+(item.Position.Y*uiScale)+(item.Size.Y*uiScale)
}

func (win *WindowData) getWinRect() Rect {
	return Rect{
		X0: win.Position.X,
		Y0: win.Position.Y,
		X1: win.Position.X + win.ScreenSize.X,
		Y1: win.Position.Y + win.ScreenSize.Y - win.TitleScreenHeight,
	}
}

func (item *ItemData) getItemRect(win *WindowData) Rect {
	return Rect{
		X0: win.Position.X + (item.Position.X * uiScale),
		Y0: win.Position.Y + (item.Position.Y * uiScale),
		X1: win.Position.X + (item.Position.X * uiScale) + (item.Size.X * uiScale),
		Y1: win.Position.Y + (item.Position.Y * uiScale) + (item.Size.Y * uiScale),
	}
}

func (rect Rect) getRectangle() image.Rectangle {
	return image.Rectangle{
		Min: image.Point{X: int(math.Ceil(float64(rect.X0))), Y: int(math.Ceil(float64(rect.Y0)))},
		Max: image.Point{X: int(math.Ceil(float64(rect.X1))), Y: int(math.Ceil(float64(rect.Y1)))}}
}

func (win *WindowData) getMainRect() Rect {
	return Rect{
		X0: win.Position.X,
		Y0: win.Position.Y + win.TitleScreenHeight + 1,
		X1: win.Position.X + win.ScreenSize.X,
		Y1: win.Position.Y + win.ScreenSize.Y - win.TitleScreenHeight,
	}
}

func (win *WindowData) getTitleRect() Rect {
	if win.TitleHeight <= 0 {
		return Rect{}
	}
	return Rect{
		X0: win.Position.X, Y0: win.Position.Y,
		X1: win.Position.X + win.ScreenSize.X,
		Y1: win.Position.Y + win.TitleScreenHeight,
	}
}

func (win *WindowData) xRect() Rect {
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

func (win *WindowData) dragbarRect() Rect {
	if win.TitleHeight <= 0 && !win.Resizable {
		return Rect{}
	}
	textSize := win.titleTextWidth()
	xRect := win.xRect()
	buttonsWidth := xRect.X1 - xRect.X0 + 3

	dpad := (win.TitleHeight * uiScale) / 5
	xStart := textSize.X + float32((win.TitleHeight*uiScale)/1.5)
	xEnd := (win.ScreenSize.X - buttonsWidth)
	return Rect{
		X0: win.Position.X + xStart, Y0: win.Position.Y + dpad,
		X1: win.Position.X + xEnd, Y1: win.Position.Y + (win.TitleHeight * uiScale) - dpad,
	}
}

func (win *WindowData) setSize(size Point) bool {

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
	win.calcUIScale()

	return tooSmall
}

const cornerSize = 14
const tol = 3

func (win *WindowData) getWindowEdge(mpos Point) WindowEdge {

	if !win.Resizable {
		return EDGE_NONE

	}

	winRect := win.getWinRect()

	//Expand outer window rect
	outRect := winRect
	outRect.X0 -= tol
	outRect.X1 += tol
	outRect.Y0 -= tol
	outRect.Y1 += tol

	//Contract inner window rect
	inRect := winRect
	inRect.X0 += tol
	inRect.X1 -= tol
	inRect.Y0 += tol
	inRect.Y1 -= tol

	//If within outrect, and not within inrect it is window edge
	if outRect.containsPoint(mpos) && !inRect.containsPoint(mpos) {
		if mpos.Y < outRect.Y0+cornerSize {
			if mpos.X < inRect.X0+cornerSize {
				return EDGE_TOP_LEFT
			} else if mpos.X > inRect.X1-cornerSize {
				return EDGE_TOP_RIGHT
			} else {
				return EDGE_TOP
			}
		} else if mpos.Y > inRect.Y1-cornerSize {
			if mpos.X > outRect.X1-cornerSize {
				return EDGE_BOTTOM_RIGHT
			} else if mpos.X < outRect.X0+cornerSize {
				return EDGE_BOTTOM_LEFT
			} else {
				return EDGE_BOTTOM
			}
		} else if mpos.X < inRect.X0 {
			return EDGE_LEFT
		} else if mpos.X < outRect.X1 {
			return EDGE_RIGHT
		}
	}

	return EDGE_NONE
}

func withinRange(a, b float32, tol float32) bool {
	if math.Abs(float64(a-b)) <= float64(tol) {
		return true
	}
	return false
}

func (win *WindowData) titleTextWidth() Point {
	if win.TitleHeight <= 0 {
		return Point{}
	}
	textSize := ((win.TitleHeight * uiScale) / 1.5)
	face := &text.GoTextFace{
		Source: mplusFaceSource,
		Size:   float64(textSize),
	}
	textWidth, textHeight := text.Measure(win.Title, face, 0)
	return Point{X: float32(textWidth), Y: float32(textHeight)}
}

func pointAdd(a, b Point) Point {
	return Point{X: a.X + b.X, Y: a.Y + b.Y}
}

func pointMul(a, b Point) Point {
	return Point{X: a.X * b.X, Y: a.Y * b.Y}
}

func pointScaleMul(a Point) Point {
	return Point{X: a.X * uiScale, Y: a.Y * uiScale}
}

func pointScaleDiv(a Point) Point {
	return Point{X: a.X / uiScale, Y: a.Y / uiScale}
}

func pointDiv(a, b Point) Point {
	return Point{X: a.X / b.X, Y: a.Y / b.Y}
}

func pointSub(a, b Point) Point {
	return Point{X: a.X - b.X, Y: a.Y - b.Y}
}

func (win *WindowData) getSizeX() float32 {
	return win.Size.X * uiScale
}

func (win *WindowData) getSizeY() float32 {
	return win.Size.Y * uiScale
}

// Sets SizeTemp, TitleSizeTemp
func (win *WindowData) calcUIScale() {
	win.ScreenSize = Point{X: win.Size.X * uiScale, Y: win.Size.Y * uiScale}
	win.TitleScreenHeight = win.TitleHeight * uiScale
}

func (win *WindowData) SetTitleSize(size float32) {
	win.TitleHeight = size / uiScale
}

func SetUIScale(scale float32) {
	uiScale = scale

	for _, win := range windows {
		win.calcUIScale()
	}
}
