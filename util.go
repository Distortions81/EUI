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
	return b.X >= win.getPosition().X+(item.getPosition(win).X) &&
		b.X <= win.getPosition().X+(item.getPosition(win).X)+(item.GetSize().X) &&
		b.Y >= win.getPosition().Y+(item.getPosition(win).Y) &&
		b.Y <= win.getPosition().Y+(item.getPosition(win).Y)+(item.GetPos().Y)
}

func (win *windowData) getWinRect() rect {
	winPos := win.getPosition()
	return rect{
		X0: winPos.X,
		Y0: winPos.Y,
		X1: winPos.X + (win.GetSize().X),
		Y1: winPos.Y + (win.GetSize().Y) - (win.GetTitleSize()),
	}
}

func (item *itemData) getItemRect(win *windowData) rect {
	return rect{
		X0: win.getPosition().X + (item.getPosition(win).X),
		Y0: win.getPosition().Y + (item.getPosition(win).Y),
		X1: win.getPosition().X + (item.getPosition(win).X) + (item.GetSize().X),
		Y1: win.getPosition().Y + (item.getPosition(win).Y) + (item.GetSize().Y),
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
		Y0: win.getPosition().Y + (win.GetTitleSize()) + 1,
		X1: win.getPosition().X + win.GetSize().X,
		Y1: win.getPosition().Y + win.GetSize().Y - (win.GetTitleSize()),
	}
}

func (win *windowData) getTitleRect() rect {
	if win.TitleHeight <= 0 {
		return rect{}
	}
	return rect{
		X0: win.getPosition().X, Y0: win.getPosition().Y,
		X1: win.getPosition().X + win.GetSize().X,
		Y1: win.getPosition().Y + (win.GetTitleSize()),
	}
}

func (win *windowData) xRect() rect {
	if win.TitleHeight <= 0 || !win.Closable {
		return rect{}
	}

	var xpad float32 = win.Border
	return rect{
		X0: win.getPosition().X + win.GetSize().X - (win.GetTitleSize()) + xpad,
		Y0: win.getPosition().Y + xpad,

		X1: win.getPosition().X + win.GetSize().X - xpad,
		Y1: win.getPosition().Y + (win.GetTitleSize()) - xpad,
	}
}

func (win *windowData) dragbarRect() rect {
	if win.TitleHeight <= 0 && !win.Resizable {
		return rect{}
	}
	textSize := win.titleTextWidth()
	xRect := win.xRect()
	buttonsWidth := xRect.X1 - xRect.X0 + 3

	dpad := (win.GetTitleSize()) / 5
	xStart := textSize.X + float32((win.GetTitleSize())/1.5)
	xEnd := (win.GetSize().X - buttonsWidth)
	return rect{
		X0: win.getPosition().X + xStart, Y0: win.getPosition().Y + dpad,
		X1: win.getPosition().X + xEnd, Y1: win.getPosition().Y + (win.GetTitleSize()) - dpad,
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
	textSize := ((win.GetTitleSize()) / 1.5)
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

func (win *windowData) SetTitleSize(size float32) {
	win.TitleHeight = size / uiScale
}

func SetUIScale(scale float32) {
	uiScale = scale
}

func (win *windowData) GetTitleSize() float32 {
	return win.TitleHeight * uiScale
}

func (win *windowData) GetSize() point {
	return point{X: win.Size.X * uiScale, Y: win.Size.Y * uiScale}
}

func (win *windowData) GetPos() point {
	return point{X: win.Position.X * uiScale, Y: win.Position.Y * uiScale}
}

func (item *itemData) GetSize() point {
	return point{X: item.Size.X * uiScale, Y: item.Size.Y * uiScale}
}

func (item *itemData) GetPos() point {
	return point{X: item.Position.X * uiScale, Y: item.Position.Y * uiScale}
}
