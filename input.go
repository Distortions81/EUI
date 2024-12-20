package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const minDrag = 2

var (
	mposOld     Point
	dragStart   Point
	cursorShape ebiten.CursorShapeType
)

func (g *Game) Update() error {
	cursorSet := false
	mx, my := ebiten.CursorPosition()
	mpos := Point{X: float32(mx), Y: float32(my)}
	defer func() { mposOld = mpos }()

	click := inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0)
	clickHeld := inpututil.MouseButtonPressDuration(ebiten.MouseButton0)
	clickDrag := clickHeld > minDrag

	//altClick := inpututil.IsMouseButtonJustPressed(ebiten.MouseButton1)
	//altClickHeld := inpututil.MouseButtonPressDuration(ebiten.MouseButton1)

	//Check all windows
	for _, win := range Windows {
		if !win.Open {
			continue
		}

		//Resize tab
		if win.Resizable {
			if win.ResizeTabRect().ContainsPoint(mposOld) {
				win.HoverResizeTab = true
				if !cursorSet && cursorShape != ebiten.CursorShapeNWSEResize {
					cursorShape = ebiten.CursorShapeNWSEResize
					ebiten.SetCursorShape(cursorShape)
				}
				cursorSet = true
				if clickDrag {
					change := PointToMag(PointSubract(mpos, mposOld))
					win.Size = MagAdd(win.Size, change)
				}
			}
		}

		//Titlebar items
		if win.TitleSize > 0 {

			//Dragbar
			if win.Movable {
				if win.TitleRect().ContainsPoint(mposOld) {
					if win.DragbarRect().ContainsPoint(mposOld) {
						win.HoverDragbar = true

						if !cursorSet && cursorShape != ebiten.CursorShapeMove {
							cursorShape = ebiten.CursorShapeMove
							ebiten.SetCursorShape(cursorShape)
						}
						cursorSet = true

						if clickDrag {
							win.Position = PointAdd(win.Position, PointSubract(mpos, mposOld))
						}
					}
				}
			}
			//Close X
			if win.Closable {
				if win.TitleRect().ContainsPoint(mpos) {
					if win.XRect().ContainsPoint(mpos) {
						win.HoverX = true
						if click {
							win.Open = false
						}
					}
				}
			}
		}

		//Window items
		if win.GetWinRect().ContainsPoint(mpos) {
			win.Hovered = true

			if win.GetMainRect().ContainsPoint(mpos) {
				for i, item := range win.Contents {
					if item.Rect.ContainsPoint(mpos) {
						win.Contents[i].Hovered = true
						if click {
							win.Contents[i].Activated = true
						}
					}
				}
			}
		}
	}

	if !cursorSet && cursorShape != ebiten.CursorShapeDefault {
		if cursorShape != ebiten.CursorShapeDefault {
			cursorShape = ebiten.CursorShapeDefault
			ebiten.SetCursorShape(cursorShape)
		}
	}

	return nil
}
