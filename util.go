package main

import "github.com/hajimehoshi/ebiten/v2/text/v2"

func (rect Rect) ContainsPoint(b Point) bool {
	return b.X >= rect.X0 && b.Y >= rect.Y0 &&
		b.X <= rect.X1 && b.Y <= rect.Y1
}

func (win WindowData) GetWinRect() Rect {
	return Rect{
		X0: win.Position.X,
		Y0: win.Position.Y,
		X1: win.Position.X + (win.Size.X * UIScale),
		Y1: win.Position.Y + (win.Size.Y * UIScale) - (win.TitleSize * UIScale),
	}
}

func (win WindowData) GetMainRect() Rect {
	return Rect{
		X0: win.Position.X,
		Y0: win.Position.Y + (win.TitleSize * UIScale) + 1,
		X1: win.Position.X + (win.Size.X * UIScale),
		Y1: win.Position.Y + (win.Size.Y * UIScale) - (win.TitleSize * UIScale),
	}
}

func (win WindowData) TitleRect() Rect {
	if win.TitleSize <= 0 {
		return Rect{}
	}
	return Rect{
		X0: win.Position.X, Y0: win.Position.Y,
		X1: win.Position.X + (win.Size.X * UIScale),
		Y1: win.Position.Y + (win.TitleSize * UIScale),
	}
}

func (win WindowData) XRect() Rect {
	if win.TitleSize <= 0 || !win.Closable {
		return Rect{}
	}

	var xpad float32 = win.Border
	return Rect{
		X0: win.Position.X + (win.Size.X * UIScale) - (win.TitleSize * UIScale) + xpad,
		Y0: win.Position.Y + xpad,

		X1: win.Position.X + (win.Size.X * UIScale) - xpad,
		Y1: win.Position.Y + (win.TitleSize * UIScale) - xpad,
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
	xEnd := ((win.Size.X * UIScale) - buttonsWidth)
	return Rect{
		X0: win.Position.X + xStart, Y0: win.Position.Y + dpad,
		X1: win.Position.X + xEnd, Y1: win.Position.Y + (win.TitleSize * UIScale) - dpad,
	}
}

func (win WindowData) ResizeTabRect() Rect {
	if win.TitleSize <= 0 && !win.Resizable {
		return Rect{}
	}

	return Rect{
		X1: win.Position.X + (win.Size.X * UIScale) - 1,
		Y0: win.Position.Y + (win.Size.Y * UIScale) - (14 * UIScale) - (win.TitleSize * UIScale),

		X0: win.Position.X + (win.Size.X * UIScale) - (14 * UIScale),
		Y1: win.Position.Y + (win.Size.Y * UIScale) - (win.TitleSize * UIScale) - 1,
	}
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

func MagAdd(a, b Magnatude) Magnatude {
	return Magnatude{X: a.X + b.X, Y: a.Y + b.Y}
}

func MadSubtract(a, b Magnatude) Magnatude {
	return Magnatude{X: a.X - b.X, Y: a.Y - b.Y}
}

func PointToMag(a Point) Magnatude {
	return Magnatude{X: a.X, Y: a.Y}
}
