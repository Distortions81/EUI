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
	cursorChanged := false

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

		if win.Resizable {
			//Resize Tab
			if win.ResizeTabRect().ContainsPoint(mposOld) {
				win.HoverResizeTab = true
				if !cursorChanged {
					if cursorShape != ebiten.CursorShapeNWSEResize {
						cursorShape = ebiten.CursorShapeNWSEResize
						ebiten.SetCursorShape(cursorShape)
					}
					cursorChanged = true
				}
				if clickDrag {
					change := PointToMag(PointSubract(mpos, mposOld))
					win.Size = MagAdd(win.Size, change)
				}
			} else {
				//Resize Edge
				side := win.GetWindowEdge(mposOld)

				//If needed, set cursoer
				if !cursorChanged && side != SIDE_NONE {
					c := ebiten.CursorShapeEWResize
					if side == SIDE_TOP || side == SIDE_BOTTOM {
						c = ebiten.CursorShapeNSResize
					}
					if cursorShape != c {
						cursorShape = c
						ebiten.SetCursorShape(cursorShape)
					}
					cursorChanged = true
				}

				//Drag rezie edge
				if clickDrag {
					if side == SIDE_TOP {
						change := PointSubract(mpos, mposOld)
						change.X = 0
						win.Position = PointAdd(win.Position, change)
						win.Size = Magnatude(PointSubract(MagToPoint(win.Size), change))
					} else if side == SIDE_BOTTOM {
						change := PointSubract(mpos, mposOld)
						change.X = 0
						win.Size = Magnatude(PointAdd(MagToPoint(win.Size), change))
					} else if side == SIDE_LEFT {
						change := PointSubract(mpos, mposOld)
						change.Y = 0
						win.Position = PointAdd(win.Position, change)
						win.Size = Magnatude(PointSubract(MagToPoint(win.Size), change))
					} else if side == SIDE_RIGHT {
						change := PointSubract(mpos, mposOld)
						change.Y = 0
						win.Size = Magnatude(PointAdd(MagToPoint(win.Size), change))
					}
				}
			}
		}

		//Titlebar items
		if win.TitleSize > 0 {

			//Dragbar
			if !cursorChanged && win.Movable {
				if win.TitleRect().ContainsPoint(mposOld) {
					if win.DragbarRect().ContainsPoint(mposOld) {
						win.HoverDragbar = true

						if !cursorChanged {
							if cursorShape != ebiten.CursorShapeMove {
								cursorShape = ebiten.CursorShapeMove
								ebiten.SetCursorShape(cursorShape)
							}
							cursorChanged = true
						}

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

	if !cursorChanged && cursorShape != ebiten.CursorShapeDefault {
		if cursorShape != ebiten.CursorShapeDefault {
			cursorShape = ebiten.CursorShapeDefault
			ebiten.SetCursorShape(cursorShape)
		}
	}

	return nil
}
