package main

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var (
	mposOld     point
	cursorShape ebiten.CursorShapeType
)

func (g *Game) Update() error {

	mx, my := ebiten.CursorPosition()
	mpos := point{X: float32(mx), Y: float32(my)}

	click := inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0)
	clickTime := inpututil.MouseButtonPressDuration(ebiten.MouseButton0)
	clickDrag := clickTime > 1
	//altClick := inpututil.IsMouseButtonJustPressed(ebiten.MouseButton1)

	posCh := pointScaleDiv(pointSub(mpos, mposOld))
	sizeCh := pointScaleMul(point{X: posCh.X / uiScale, Y: posCh.Y / uiScale})
	c := ebiten.CursorShapeDefault

	//Check all windows
	for i := len(windows) - 1; i >= 0; i-- {
		win := windows[i]
		if !win.Open {
			continue
		}

		part := win.getWindowPart(mposOld, click)

		if part != PART_NONE {

			if c == ebiten.CursorShapeDefault {
				if part == PART_BAR {
					c = ebiten.CursorShapeMove
				} else if part == PART_LEFT || part == PART_RIGHT {
					c = ebiten.CursorShapeEWResize
				} else if part == PART_TOP || part == PART_BOTTOM {
					c = ebiten.CursorShapeNSResize
				} else if part == PART_TOP_LEFT || part == PART_BOTTOM_RIGHT {
					c = ebiten.CursorShapeNWSEResize
				} else if part == PART_TOP_RIGHT || part == PART_BOTTOM_LEFT {
					c = ebiten.CursorShapeNESWResize
				}
			}

			if click {
				if part == PART_CLOSE {
					win.RemoveWindow()
				}
			} else if clickDrag {
				if part == PART_BAR {
					win.Position = pointAdd(win.Position, posCh)
				} else if part == PART_TOP {
					posCh.X = 0
					sizeCh.X = 0
					if !win.setSize(pointSub(win.Size, sizeCh)) {
						win.Position = pointAdd(win.Position, posCh)
					}
				} else if part == PART_BOTTOM {
					sizeCh.X = 0
					win.setSize(pointAdd(win.Size, sizeCh))
				} else if part == PART_LEFT {
					posCh.Y = 0
					sizeCh.Y = 0
					if !win.setSize(pointSub(win.Size, sizeCh)) {
						win.Position = pointAdd(win.Position, posCh)
					}
				} else if part == PART_RIGHT {
					sizeCh.Y = 0
					win.setSize(pointAdd(win.Size, sizeCh))
					//Corner resize
				} else if part == PART_TOP_LEFT {
					if !win.setSize(pointSub(win.Size, sizeCh)) {
						win.Position = pointAdd(win.Position, posCh)
					}
				} else if part == PART_TOP_RIGHT {
					tx := win.Size.X + sizeCh.X
					ty := win.Size.Y - sizeCh.Y
					if !win.setSize(point{X: tx, Y: ty}) {
						win.Position.Y += posCh.Y
					}
				} else if part == PART_BOTTOM_RIGHT {
					tx := win.Size.X + sizeCh.X
					ty := win.Size.Y + sizeCh.Y
					win.setSize(point{X: tx, Y: ty})
				} else if part == PART_BOTTOM_LEFT {
					tx := win.Size.X - sizeCh.X
					ty := win.Size.Y + sizeCh.Y

					if !win.setSize(point{X: tx, Y: ty}) {
						win.Position.X += posCh.X
					}
				}
				break
			}
		}

		//Window items
		win.clickWindowItems(mpos, click)

		//Bring window forward on click, defer
		if win.getWinRect().containsPoint(mpos) {
			if click && activeWindow != win {
				win.BringForward()
			}
			break
		}
	}

	if cursorShape != c {
		ebiten.SetCursorShape(c)
		cursorShape = c
	}
	mposOld = mpos

	return nil
}

func (win *windowData) clickWindowItems(mpos point, click bool) {
	// If the mouse isn't within the window, just return
	if !win.getMainRect().containsPoint(mpos) {
		return
	}
	win.Hovered = true

	for _, item := range win.Contents {
		if item.ItemType == ITEM_FLOW {
			item.clickFlows(mpos, click)
		} else {
			item.clickItem(mpos, click)
		}
	}
}

func (item *itemData) clickFlows(mpos point, click bool) {
	for _, subItem := range item.Contents {
		if subItem.ItemType == ITEM_FLOW {
			subItem.clickFlows(mpos, click)
		} else {
			subItem.clickItem(mpos, click)
		}
	}
}

func (item *itemData) clickItem(mpos point, click bool) {
	// If the mouse isn't within the item, just return
	if !item.DrawRect.containsPoint(mpos) {
		return
	}

	if click {
		item.Clicked = time.Now()
		if item.ItemType == ITEM_CHECKBOX {
			item.Checked = !item.Checked
		}
		if item.Action != nil {
			item.Action()
			return
		}
	} else {
		item.Hovered = true
	}
}
