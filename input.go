package main

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const minDrag = 0

var (
	mposOld     point
	dragStart   point
	cursorShape ebiten.CursorShapeType
)

func (g *Game) Update() error {

	cursorChanged := false

	mx, my := ebiten.CursorPosition()
	mpos := point{X: float32(mx), Y: float32(my)}
	defer func() { mposOld = mpos }()

	click := inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0)
	clickHeld := inpututil.MouseButtonPressDuration(ebiten.MouseButton0)
	clickDrag := clickHeld > minDrag

	//altClick := inpututil.IsMouseButtonJustPressed(ebiten.MouseButton1)
	//altClickHeld := inpututil.MouseButtonPressDuration(ebiten.MouseButton1)

	//Check all windows
	for i := len(windows) - 1; i >= 0; i-- {
		win := windows[i]
		if !win.Open {
			continue
		}

		defer func() {
			if !cursorChanged && cursorShape != ebiten.CursorShapeDefault {
				cursorShape = ebiten.CursorShapeDefault
				ebiten.SetCursorShape(cursorShape)
			}
		}()

		if win.Resizable {
			side := win.getWindowEdge(mposOld)

			//If needed, set cursor
			if !cursorChanged && side != EDGE_NONE {
				c := ebiten.CursorShapeDefault
				if side == EDGE_LEFT || side == EDGE_RIGHT {
					c = ebiten.CursorShapeEWResize
				}
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
					cursorChanged = true
				}

				//Drag resize edge or corner
				if clickDrag {
					posCh := pointSub(mpos, mposOld)
					sizeCh := pointScaleMul(point{X: posCh.X / uiScale, Y: posCh.Y / uiScale})
					if side == EDGE_TOP {
						posCh.X = 0
						sizeCh.X = 0
						if !win.setSize(pointSub(win.Size, sizeCh)) {
							win.Position = pointAdd(win.Position, posCh)
						}
						break
					} else if side == EDGE_BOTTOM {
						sizeCh.X = 0
						win.setSize(pointAdd(win.Size, sizeCh))
						break
					} else if side == EDGE_LEFT {
						posCh.Y = 0
						sizeCh.Y = 0
						if !win.setSize(pointSub(win.Size, sizeCh)) {
							win.Position = pointAdd(win.Position, posCh)
						}
						break
					} else if side == EDGE_RIGHT {
						sizeCh.Y = 0
						win.setSize(pointAdd(win.Size, sizeCh))
						break

						//Corner reize
					} else if side == EDGE_TOP_LEFT {
						if !win.setSize(pointSub(win.Size, sizeCh)) {
							win.Position = pointAdd(win.Position, posCh)
						}
						break
					} else if side == EDGE_TOP_RIGHT {
						tx := win.Size.X + sizeCh.X
						ty := win.Size.Y - sizeCh.Y
						if !win.setSize(point{X: tx, Y: ty}) {
							win.Position.Y += posCh.Y
						}
						break
					} else if side == EDGE_BOTTOM_RIGHT {
						tx := win.Size.X + sizeCh.X
						ty := win.Size.Y + sizeCh.Y
						win.setSize(point{X: tx, Y: ty})
						break
					} else if side == EDGE_BOTTOM_LEFT {
						tx := win.Size.Y + sizeCh.Y
						ty := win.Size.X - sizeCh.X
						if !win.setSize(point{X: tx, Y: ty}) {
							win.Position.X += posCh.X
						}
						break
					} else if side == EDGE_TOP_LEFT {
						tx := win.Size.Y - sizeCh.Y
						ty := win.Size.X + sizeCh.X
						if !win.setSize(point{X: tx, Y: ty}) {
							win.Position.Y -= posCh.Y
						}
						break
					}
				}
			}
		}

		//Titlebar items
		if win.TitleHeight > 0 {

			//Dragbar
			if !cursorChanged && win.Movable {
				if win.getTitleRect().containsPoint(mposOld) {
					if win.dragbarRect().containsPoint(mposOld) {
						win.HoverDragbar = true

						if !cursorChanged {
							if cursorShape != ebiten.CursorShapeMove {
								cursorShape = ebiten.CursorShapeMove
								ebiten.SetCursorShape(cursorShape)
							}
							cursorChanged = true
						}

						if clickDrag {
							win.Position = pointAdd(win.Position, pointSub(mpos, mposOld))
							break
						}
					}
				}
			}
			//Close X
			if win.Closable {
				if win.getTitleRect().containsPoint(mpos) {
					if win.xRect().containsPoint(mpos) {
						win.HoverClose = true
						if click {
							win.RemoveWindow()
							break
						}
					}
				}
			}
		}

		//Window items
		win.clickWindowItems(mpos, click)

		//Bring window forward on click
		if win.getWinRect().containsPoint(mpos) {
			if click {
				win.BringForward()
			}
			return nil
		}
	}

	return nil
}

func (win *windowData) clickWindowItems(mpos point, click bool) {
	if win.getWinRect().containsPoint(mpos) {
		win.Hovered = true

		if win.getMainRect().containsPoint(mpos) {
			for i, item := range win.Contents {
				if item.ItemType == ITEM_FLOW {
					newWin := windowData{Size: win.Size, Position: pointSub(win.Position, point{X: 0, Y: win.TitleHeight}), Contents: item.Contents}
					newWin.clickWindowItems(mpos, click)
					continue
				}
				if item.ItemType != ITEM_BUTTON {
					continue
				}
				if item.containsPoint(win, mpos) {
					win.Contents[i].Hovered = true
					if click {
						win.Contents[i].Clicked = time.Now()
						if item.Action != nil {
							item.Action()
							break
						}
					}
				}
			}
		}
	}
}
