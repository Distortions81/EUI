package main

import (
	"time"

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

		//Reduce UI scaling calculations
		win.CalcUIScale()

		if win.Resizable {

			side := win.GetWindowEdge(mposOld)

			//If needed, set cursor
			if !cursorChanged && side != EDGE_NONE {
				c := ebiten.CursorShapeEWResize
				if side == EDGE_TOP || side == EDGE_BOTTOM {
					c = ebiten.CursorShapeNSResize
				} else if side == EDGE_TOP_LEFT || side == EDGE_BOTTOM_RIGHT {
					c = ebiten.CursorShapeNWSEResize
				} else if side == EDGE_TOP_RIGHT || side == EDGE_BOTTOM_LEFT {
					c = ebiten.CursorShapeNESWResize
				}
				if cursorShape != c {
					cursorShape = c
					ebiten.SetCursorShape(cursorShape)
				}
				cursorChanged = true
			}

			//Drag resize edge or corner
			if clickDrag {
				posCh := PointSub(mpos, mposOld)
				sizeCh := Point{X: posCh.X / UIScale, Y: posCh.Y / UIScale}
				if side == EDGE_TOP {
					posCh.X = 0
					sizeCh.X = 0
					if !win.SetSize(PointSub(win.Size, sizeCh)) {
						win.Position = PointAdd(win.Position, posCh)
					}
					continue
				} else if side == EDGE_BOTTOM {
					sizeCh.X = 0
					win.SetSize(PointAdd(win.Size, sizeCh))
					continue
				} else if side == EDGE_LEFT {
					posCh.Y = 0
					sizeCh.Y = 0
					if !win.SetSize(PointSub(win.Size, sizeCh)) {
						win.Position = PointAdd(win.Position, posCh)
					}
					continue
				} else if side == EDGE_RIGHT {
					sizeCh.Y = 0
					win.SetSize(PointAdd(win.Size, sizeCh))
					continue

					//Corner reize
				} else if side == EDGE_TOP_LEFT {
					if !win.SetSize(PointSub(win.Size, sizeCh)) {
						win.Position = PointAdd(win.Position, posCh)
					}
					continue
				} else if side == EDGE_TOP_RIGHT {
					win.Size.X += sizeCh.X
					win.Size.Y -= sizeCh.Y
					win.Position.Y += posCh.Y
					continue
				} else if side == EDGE_BOTTOM_RIGHT {
					win.Size.X += sizeCh.X
					win.Size.Y += sizeCh.Y
					continue
				} else if side == EDGE_BOTTOM_LEFT {
					win.Size.Y += sizeCh.Y
					win.Size.X -= sizeCh.X
					win.Position.X += posCh.X
					continue
				} else if side == EDGE_TOP_LEFT {
					win.Size.Y -= sizeCh.Y
					win.Size.X += sizeCh.X
					win.Position.Y -= posCh.Y
					continue
				}
			}
		}

		//Titlebar items
		if win.TitleHeight > 0 {

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
							win.Position = PointAdd(win.Position, PointSub(mpos, mposOld))
						}
					}
				}
			}
			//Close X
			if win.Closable {
				if win.TitleRect().ContainsPoint(mpos) {
					if win.XRect().ContainsPoint(mpos) {
						win.HoverClose = true
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
					if item.ItemType != ITEM_BUTTON {
						continue
					}
					if item.ContainsPoint(win, mpos) {
						win.Contents[i].Hovered = true
						if click {
							win.Contents[i].Clicked = time.Now()
							if item.Action != nil {
								item.Action()
							}
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
