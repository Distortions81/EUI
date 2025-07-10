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
	if click {
		if focusedItem != nil {
			focusedItem.Focused = false
		}
		focusedItem = nil
	}
	clickTime := inpututil.MouseButtonPressDuration(ebiten.MouseButton0)
	clickDrag := clickTime > 1
	//altClick := inpututil.IsMouseButtonJustPressed(ebiten.MouseButton1)

	wx, wy := ebiten.Wheel()
	wheelDelta := point{X: float32(wx), Y: float32(wy)}

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
				} else if win.Resizable {
					if part == PART_LEFT || part == PART_RIGHT {
						c = ebiten.CursorShapeEWResize
					} else if part == PART_TOP || part == PART_BOTTOM {
						c = ebiten.CursorShapeNSResize
					} else if part == PART_TOP_LEFT || part == PART_BOTTOM_RIGHT {
						c = ebiten.CursorShapeNWSEResize
					} else if part == PART_TOP_RIGHT || part == PART_BOTTOM_LEFT {
						c = ebiten.CursorShapeNESWResize
					}
				}
			}

			if click {
				if part == PART_CLOSE {
					win.RemoveWindow()
				}
			} else if clickDrag {
				if part == PART_BAR {
					win.Position = pointAdd(win.Position, posCh)
				} else if win.Resizable && part == PART_TOP {
					posCh.X = 0
					sizeCh.X = 0
					if !win.setSize(pointSub(win.Size, sizeCh)) {
						win.Position = pointAdd(win.Position, posCh)
					}
				} else if win.Resizable && part == PART_BOTTOM {
					sizeCh.X = 0
					win.setSize(pointAdd(win.Size, sizeCh))
				} else if win.Resizable && part == PART_LEFT {
					posCh.Y = 0
					sizeCh.Y = 0
					if !win.setSize(pointSub(win.Size, sizeCh)) {
						win.Position = pointAdd(win.Position, posCh)
					}
				} else if win.Resizable && part == PART_RIGHT {
					sizeCh.Y = 0
					win.setSize(pointAdd(win.Size, sizeCh))
					//Corner resize
				} else if win.Resizable && part == PART_TOP_LEFT {
					if !win.setSize(pointSub(win.Size, sizeCh)) {
						win.Position = pointAdd(win.Position, posCh)
					}
				} else if win.Resizable && part == PART_TOP_RIGHT {
					tx := win.Size.X + sizeCh.X
					ty := win.Size.Y - sizeCh.Y
					if !win.setSize(point{X: tx, Y: ty}) {
						win.Position.Y += posCh.Y
					}
				} else if win.Resizable && part == PART_BOTTOM_RIGHT {
					tx := win.Size.X + sizeCh.X
					ty := win.Size.Y + sizeCh.Y
					win.setSize(point{X: tx, Y: ty})
				} else if win.Resizable && part == PART_BOTTOM_LEFT {
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

	if focusedItem != nil {
		for _, r := range ebiten.AppendInputChars(nil) {
			if r >= 32 && r != 127 && r != '\r' && r != '\n' {
				focusedItem.Text += string(r)
			}
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyBackspace) {
			runes := []rune(focusedItem.Text)
			if len(runes) > 0 {
				focusedItem.Text = string(runes[:len(runes)-1])
			}
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
			focusedItem.Focused = false
			focusedItem = nil
		}
	}

	mposOld = mpos

	if wheelDelta.X != 0 || wheelDelta.Y != 0 {
		for i := len(windows) - 1; i >= 0; i-- {
			win := windows[i]
			if !win.Open {
				continue
			}
			if win.getMainRect().containsPoint(mpos) {
				if scrollFlow(win.Contents, mpos, wheelDelta) {
					break
				}
			}
		}
	}

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
	if len(item.Tabs) > 0 {
		if item.ActiveTab >= len(item.Tabs) {
			item.ActiveTab = 0
		}
		for i, tab := range item.Tabs {
			if tab.DrawRect.containsPoint(mpos) {
				if click {
					item.ActiveTab = i
				}
				return
			}
		}
		for _, subItem := range item.Tabs[item.ActiveTab].Contents {
			if subItem.ItemType == ITEM_FLOW {
				subItem.clickFlows(mpos, click)
			} else {
				subItem.clickItem(mpos, click)
			}
		}
	} else {
		for _, subItem := range item.Contents {
			if subItem.ItemType == ITEM_FLOW {
				subItem.clickFlows(mpos, click)
			} else {
				subItem.clickItem(mpos, click)
			}
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
		} else if item.ItemType == ITEM_RADIO {
			item.Checked = true
			// uncheck others in group
			if item.RadioGroup != "" {
				uncheckRadioGroup(item.Parent, item.RadioGroup, item)
			}
		} else if item.ItemType == ITEM_INPUT {
			focusedItem = item
			item.Focused = true
		}
		if item.Action != nil {
			item.Action()
			return
		}
	} else {
		item.Hovered = true
		if item.ItemType == ITEM_SLIDER && ebiten.IsMouseButtonPressed(ebiten.MouseButton0) {
			item.setSliderValue(mpos)
			if item.Action != nil {
				item.Action()
			}
		}
	}
}

func uncheckRadioGroup(parent *itemData, group string, except *itemData) {
	if parent == nil {
		for _, win := range windows {
			subUncheckRadio(win.Contents, group, except)
		}
	} else {
		subUncheckRadio(parent.Contents, group, except)
	}
}

func subUncheckRadio(list []*itemData, group string, except *itemData) {
	for _, it := range list {
		if it.ItemType == ITEM_RADIO && it.RadioGroup == group && it != except {
			it.Checked = false
		}
		if len(it.Tabs) > 0 {
			for _, tab := range it.Tabs {
				subUncheckRadio(tab.Contents, group, except)
			}
		}
		subUncheckRadio(it.Contents, group, except)
	}
}

func (item *itemData) setSliderValue(mpos point) {
	width := item.DrawRect.X1 - item.DrawRect.X0 - item.AuxSize.X - item.AuxSpace*2
	if width <= 0 {
		return
	}
	val := (mpos.X - item.DrawRect.X0)
	if val < 0 {
		val = 0
	}
	if val > width {
		val = width
	}
	ratio := val / width
	item.Value = item.MinValue + ratio*(item.MaxValue-item.MinValue)
	if item.IntOnly {
		item.Value = float32(int(item.Value + 0.5))
	}
}

func scrollFlow(items []*itemData, mpos point, delta point) bool {
	for _, it := range items {
		if it.ItemType == ITEM_FLOW {
			if it.DrawRect.containsPoint(mpos) {
				req := it.contentBounds()
				size := it.GetSize()
				if it.Scrollable {
					if it.FlowType == FLOW_VERTICAL && req.Y > size.Y {
						it.Scroll.Y -= delta.Y * 16
						if it.Scroll.Y < 0 {
							it.Scroll.Y = 0
						}
						max := req.Y - size.Y
						if it.Scroll.Y > max {
							it.Scroll.Y = max
						}
						return true
					} else if it.FlowType == FLOW_HORIZONTAL && req.X > size.X {
						it.Scroll.X -= delta.X * 16
						if it.Scroll.X < 0 {
							it.Scroll.X = 0
						}
						max := req.X - size.X
						if it.Scroll.X > max {
							it.Scroll.X = max
						}
						return true
					}
				}
			}
			var sub []*itemData
			if len(it.Tabs) > 0 {
				if it.ActiveTab >= len(it.Tabs) {
					it.ActiveTab = 0
				}
				sub = it.Tabs[it.ActiveTab].Contents
			} else {
				sub = it.Contents
			}
			if scrollFlow(sub, mpos, delta) {
				return true
			}
		}
	}
	return false
}
