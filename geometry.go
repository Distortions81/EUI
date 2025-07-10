package main

import (
	"image"
	"math"
)

// containsPoint checks whether the given point lies within the rectangle.
func (r rect) containsPoint(p point) bool {
	return p.X >= r.X0 && p.Y >= r.Y0 && p.X <= r.X1 && p.Y <= r.Y1
}

// containsPoint determines if the point is inside the item's rectangle on the window.
func (it *itemData) containsPoint(win *windowData, p point) bool {
	return p.X >= win.getPosition().X+it.getPosition(win).X &&
		p.X <= win.getPosition().X+it.getPosition(win).X+it.GetSize().X &&
		p.Y >= win.getPosition().Y+it.getPosition(win).Y &&
		p.Y <= win.getPosition().Y+it.getPosition(win).Y+it.GetSize().Y
}

// getRectangle converts a rect to the standard image.Rectangle type.
func (r rect) getRectangle() image.Rectangle {
	return image.Rectangle{
		Min: image.Point{X: int(math.Ceil(float64(r.X0))), Y: int(math.Ceil(float64(r.Y0)))},
		Max: image.Point{X: int(math.Ceil(float64(r.X1))), Y: int(math.Ceil(float64(r.Y1)))},
	}
}

// withinRange returns true if a and b are within the provided tolerance.
func withinRange(a, b float32, tol float32) bool {
	return math.Abs(float64(a-b)) <= float64(tol)
}

func pointAdd(a, b point) point { return point{X: a.X + b.X, Y: a.Y + b.Y} }
func pointSub(a, b point) point { return point{X: a.X - b.X, Y: a.Y - b.Y} }
func pointMul(a, b point) point { return point{X: a.X * b.X, Y: a.Y * b.Y} }
func pointDiv(a, b point) point { return point{X: a.X / b.X, Y: a.Y / b.Y} }

func pointScaleMul(a point) point { return point{X: a.X * uiScale, Y: a.Y * uiScale} }
func pointScaleDiv(a point) point { return point{X: a.X / uiScale, Y: a.Y / uiScale} }

// unionRect expands a to encompass b and returns the result.
func unionRect(a, b rect) rect {
	if b.X0 < a.X0 {
		a.X0 = b.X0
	}
	if b.Y0 < a.Y0 {
		a.Y0 = b.Y0
	}
	if b.X1 > a.X1 {
		a.X1 = b.X1
	}
	if b.Y1 > a.Y1 {
		a.Y1 = b.Y1
	}
	return a
}
