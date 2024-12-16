package main

func (rect Rect) ContainsPoint(b Point) bool {
	return b.X >= rect.X0 && b.Y >= rect.Y0 &&
		b.X <= rect.X1 && b.Y <= rect.Y1
}

func (win WindowData) GetRect() Rect {
	return Rect{
		X0: win.Position.X, Y0: win.Position.Y,
		X1: win.Position.X + (win.Size.X * UIScale),
		Y1: win.Position.Y + (win.Size.Y * UIScale) + (win.TitleSize * UIScale),
	}
}

func (win WindowData) GetTitleRect() Rect {
	return Rect{
		X0: win.Position.X, Y0: win.Position.Y,
		X1: win.Position.X + (win.Size.X * UIScale),
		Y1: win.Position.Y + (win.TitleSize * UIScale),
	}
}
