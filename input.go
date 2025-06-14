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
					//Corner reize
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
	//If the mouse isn't within the window, just return
	if !win.getMainRect().containsPoint(mpos) {
		return
	}
	win.Hovered = true
	winPos := pointAdd(win.GetPos(), point{X: 0, Y: win.GetTitleSize()})

	for _, item := range win.Contents {
		itemPos := pointAdd(winPos, item.getPosition(win))

		if item.ItemType == ITEM_FLOW {
			item.clickFlows(nil, itemPos, mpos, click)
		} else {
			item.clickItem(nil, itemPos, mpos, click)
		}
	}
}

func (item *itemData) clickFlows(parent *itemData, offset, mpos point, click bool) {

	var flowOffset point

	for _, subItem := range item.Contents {

		if subItem.ItemType == ITEM_FLOW {
			flowPos := pointAdd(offset, item.GetPos())
			flowOff := pointAdd(flowPos, flowOffset)
			itemPos := pointAdd(flowOff, subItem.GetPos())
			subItem.clickFlows(item, itemPos, mpos, click)
		} else {
			flowOff := pointAdd(offset, flowOffset)
			subItem.clickItem(item, flowOff, mpos, click)
		}

		if item.FlowType == FLOW_HORIZONTAL {
			flowOffset = pointAdd(flowOffset, point{X: subItem.GetSize().X, Y: 0})
		} else if item.FlowType == FLOW_VERTICAL {
			flowOffset = pointAdd(flowOffset, point{X: 0, Y: subItem.GetSize().Y})
		}
	}
}

func (item *itemData) clickItem(parent *itemData, offset, mpos point, click bool) {
	if parent == nil {
		parent = item
	}
	maxSize := item.GetSize()
	if item.Size.X > parent.Size.X {
		maxSize.X = parent.GetSize().X
	}
	if item.Size.Y > parent.Size.Y {
		maxSize.Y = parent.GetSize().Y
	}
	itemRect := rect{
		X0: offset.X, X1: offset.X + maxSize.X,
		Y0: offset.Y, Y1: offset.Y + maxSize.Y,
	}

	//If the mouse isn't within the item, just return
	if !itemRect.containsPoint(mpos) {
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
