package main

import (
	"image"
	"math"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

func (rect rect) containsPoint(b point) bool {
	return b.X >= rect.X0 && b.Y >= rect.Y0 &&
		b.X <= rect.X1 && b.Y <= rect.Y1
}

func (item *itemData) containsPoint(win *windowData, b point) bool {
	return b.X >= win.getPosition().X+(item.getPosition(win).X*uiScale) &&
		b.X <= win.getPosition().X+(item.getPosition(win).X*uiScale)+(item.Size.X*uiScale) &&
		b.Y >= win.getPosition().Y+(item.getPosition(win).Y*uiScale) &&
		b.Y <= win.getPosition().Y+(item.getPosition(win).Y*uiScale)+(item.Size.Y*uiScale)
}

func (win *windowData) getWinRect() rect {
	winPos := win.getPosition()
	return rect{
		X0: winPos.X,
		Y0: winPos.Y,
		X1: winPos.X + (win.Size.X * uiScale),
		Y1: winPos.Y + (win.Size.Y * uiScale) - (win.TitleHeight * uiScale),
	}
}

func (item *itemData) getItemRect(win *windowData) rect {
	return rect{
		X0: win.getPosition().X + (item.getPosition(win).X * uiScale),
		Y0: win.getPosition().Y + (item.getPosition(win).Y * uiScale),
		X1: win.getPosition().X + (item.getPosition(win).X * uiScale) + (item.Size.X * uiScale),
		Y1: win.getPosition().Y + (item.getPosition(win).Y * uiScale) + (item.Size.Y * uiScale),
	}
}

func (rect rect) getRectangle() image.Rectangle {
	return image.Rectangle{
		Min: image.Point{X: int(math.Ceil(float64(rect.X0))), Y: int(math.Ceil(float64(rect.Y0)))},
		Max: image.Point{X: int(math.Ceil(float64(rect.X1))), Y: int(math.Ceil(float64(rect.Y1)))}}
}

func (win *windowData) getMainRect() rect {
	return rect{
		X0: win.getPosition().X,
		Y0: win.getPosition().Y + (win.TitleHeight * uiScale) + 1,
		X1: win.getPosition().X + win.Size.X,
		Y1: win.getPosition().Y + win.Size.Y - (win.TitleHeight * uiScale),
	}
}

func (win *windowData) getTitleRect() rect {
	if win.TitleHeight <= 0 {
		return rect{}
	}
	return rect{
		X0: win.getPosition().X, Y0: win.getPosition().Y,
		X1: win.getPosition().X + win.Size.X,
		Y1: win.getPosition().Y + (win.TitleHeight * uiScale),
	}
}

func (win *windowData) xRect() rect {
	if win.TitleHeight <= 0 || !win.Closable {
		return rect{}
	}

	var xpad float32 = win.Border
	return rect{
		X0: win.getPosition().X + win.Size.X - (win.TitleHeight * uiScale) + xpad,
		Y0: win.getPosition().Y + xpad,

		X1: win.getPosition().X + win.Size.X - xpad,
		Y1: win.getPosition().Y + (win.TitleHeight * uiScale) - xpad,
	}
}

func (win *windowData) dragbarRect() rect {
	if win.TitleHeight <= 0 && !win.Resizable {
		return rect{}
	}
	textSize := win.titleTextWidth()
	xRect := win.xRect()
	buttonsWidth := xRect.X1 - xRect.X0 + 3

	dpad := (win.TitleHeight * uiScale) / 5
	xStart := textSize.X + float32((win.TitleHeight*uiScale)/1.5)
	xEnd := (win.Size.X - buttonsWidth)
	return rect{
		X0: win.getPosition().X + xStart, Y0: win.getPosition().Y + dpad,
		X1: win.getPosition().X + xEnd, Y1: win.getPosition().Y + (win.TitleHeight * uiScale) - dpad,
	}
}

func (win *windowData) setSize(size point) bool {

	tooSmall := false
	xc, yc := win.itemOverlap(size)
	if !xc {
		win.Size.X = size.X
	}
	if !yc {
		win.Size.Y = size.Y
	}
	if yc && xc {
		tooSmall = true
	}

	if size.X < minWinSizeX {
		size.X = minWinSizeX
		tooSmall = true
	}

	if size.Y < minWinSizeY {
		size.Y = minWinSizeY
		tooSmall = true
	}

	win.BringForward()
	return tooSmall
}

const cornerSize = 14
const tol = 3

func (win *windowData) getWindowEdge(mpos point) windowEdge {

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

func (win *windowData) titleTextWidth() point {
	if win.TitleHeight <= 0 {
		return point{}
	}
	textSize := ((win.TitleHeight * uiScale) / 1.5)
	face := &text.GoTextFace{
		Source: mplusFaceSource,
		Size:   float64(textSize),
	}
	textWidth, textHeight := text.Measure(win.Title, face, 0)
	return point{X: float32(textWidth), Y: float32(textHeight)}
}

func pointAdd(a, b point) point {
	return point{X: a.X + b.X, Y: a.Y + b.Y}
}

func pointMul(a, b point) point {
	return point{X: a.X * b.X, Y: a.Y * b.Y}
}

func pointScaleMul(a point) point {
	return point{X: a.X * uiScale, Y: a.Y * uiScale}
}

func pointScaleDiv(a point) point {
	return point{X: a.X / uiScale, Y: a.Y / uiScale}
}

func pointDiv(a, b point) point {
	return point{X: a.X / b.X, Y: a.Y / b.Y}
}

func pointSub(a, b point) point {
	return point{X: a.X - b.X, Y: a.Y - b.Y}
}

func (win *windowData) getSizeX() float32 {
	return win.Size.X * uiScale
}

func (win *windowData) getSizeY() float32 {
	return win.Size.Y * uiScale
}

func (win *windowData) SetTitleSize(size float32) {
	win.TitleHeight = size / uiScale
}

func SetUIScale(scale float32) {
	uiScale = scale
}
