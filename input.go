package main

import (
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var (
	mposOld     point
	cursorShape ebiten.CursorShapeType

	dragPart   dragType
	dragWin    *windowData
	activeItem *itemData
)

func (g *Game) Update() error {

	mx, my := ebiten.CursorPosition()
	mpos := point{X: float32(mx), Y: float32(my)}

	click := inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0)
	if click {
		if !dropdownOpenContainsAnywhere(mpos) {
			closeAllDropdowns()
		}
		if focusedItem != nil {
			focusedItem.Focused = false
		}
		focusedItem = nil
	}
	clickTime := inpututil.MouseButtonPressDuration(ebiten.MouseButton0)
	clickDrag := clickTime > 1

	if !ebiten.IsMouseButtonPressed(ebiten.MouseButton0) {
		dragPart = PART_NONE
		dragWin = nil
		activeItem = nil
	}
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

		var part dragType
		if dragPart != PART_NONE && dragWin == win {
			part = dragPart
		} else {
			part = win.getWindowPart(mpos, click)
		}

		if part != PART_NONE {

			if dragPart == PART_NONE && c == ebiten.CursorShapeDefault {
				switch part {
				case PART_BAR:
					c = ebiten.CursorShapeMove
				case PART_LEFT, PART_RIGHT:
					c = ebiten.CursorShapeEWResize
				case PART_TOP, PART_BOTTOM:
					c = ebiten.CursorShapeNSResize
				case PART_TOP_LEFT, PART_BOTTOM_RIGHT:
					c = ebiten.CursorShapeNWSEResize
				case PART_TOP_RIGHT, PART_BOTTOM_LEFT:
					c = ebiten.CursorShapeNESWResize
				case PART_SCROLL_V, PART_SCROLL_H:
					c = ebiten.CursorShapePointer
				}
			}

			if click && dragPart == PART_NONE {
				if part == PART_CLOSE {
					win.RemoveWindow()
					continue
				}
				dragPart = part
				dragWin = win
			} else if clickDrag && dragPart != PART_NONE && dragWin == win {
				switch dragPart {
				case PART_BAR:
					win.Position = pointAdd(win.Position, posCh)
				case PART_TOP:
					posCh.X = 0
					sizeCh.X = 0
					if !win.setSize(pointSub(win.Size, sizeCh)) {
						win.Position = pointAdd(win.Position, posCh)
					}
				case PART_BOTTOM:
					sizeCh.X = 0
					win.setSize(pointAdd(win.Size, sizeCh))
				case PART_LEFT:
					posCh.Y = 0
					sizeCh.Y = 0
					if !win.setSize(pointSub(win.Size, sizeCh)) {
						win.Position = pointAdd(win.Position, posCh)
					}
				case PART_RIGHT:
					sizeCh.Y = 0
					win.setSize(pointAdd(win.Size, sizeCh))
				case PART_TOP_LEFT:
					if !win.setSize(pointSub(win.Size, sizeCh)) {
						win.Position = pointAdd(win.Position, posCh)
					}
				case PART_TOP_RIGHT:
					tx := win.Size.X + sizeCh.X
					ty := win.Size.Y - sizeCh.Y
					if !win.setSize(point{X: tx, Y: ty}) {
						win.Position.Y += posCh.Y
					}
				case PART_BOTTOM_RIGHT:
					tx := win.Size.X + sizeCh.X
					ty := win.Size.Y + sizeCh.Y
					win.setSize(point{X: tx, Y: ty})
				case PART_BOTTOM_LEFT:
					tx := win.Size.X - sizeCh.X
					ty := win.Size.Y + sizeCh.Y
					if !win.setSize(point{X: tx, Y: ty}) {
						win.Position.X += posCh.X
					}
				case PART_SCROLL_V:
					dragWindowScroll(win, mpos, true)
				case PART_SCROLL_H:
					dragWindowScroll(win, mpos, false)
				}
				break
			}
		}

		//Window items
		win.clickWindowItems(mpos, click)

		// Bring window forward on click if the cursor is over it or an
		// expanded dropdown. Break so windows behind don't receive the
		// event.
		if win.getWinRect().containsPoint(mpos) || dropdownOpenContains(win.Contents, mpos) {
			if click && activeWindow != win {
				win.BringForward()
			}
			break
		}
	}

	for i := len(overlays) - 1; i >= 0; i-- {
		if clickOverlay(overlays[i], mpos, click) {
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
			if win.getMainRect().containsPoint(mpos) || dropdownOpenContains(win.Contents, mpos) {
				if scrollDropdown(win.Contents, mpos, wheelDelta) {
					break
				}
				if scrollFlow(win.Contents, mpos, wheelDelta) {
					break
				}
				if scrollWindow(win, wheelDelta) {
					break
				}
			}
		}
		for i := len(overlays) - 1; i >= 0; i-- {
			ov := overlays[i]
			if ov.DrawRect.containsPoint(mpos) || dropdownOpenContains([]*itemData{ov}, mpos) {
				if scrollDropdown([]*itemData{ov}, mpos, wheelDelta) {
					break
				}
				if scrollFlow([]*itemData{ov}, mpos, wheelDelta) {
					break
				}
			}
		}
	}

	// Refresh flow layouts so scroll bars update when content size changes
	for _, win := range windows {
		if win.Open {
			win.resizeFlows()
		}
	}
	for _, ov := range overlays {
		ov.resizeFlow(ov.GetSize())
	}

	return nil
}

func (win *windowData) clickWindowItems(mpos point, click bool) {
	// If the mouse isn't within the window or any open dropdown, just return
	if !win.getMainRect().containsPoint(mpos) && !dropdownOpenContains(win.Contents, mpos) {
		return
	}
	if clickOpenDropdown(win.Contents, mpos, click) {
		return
	}
	win.Hovered = true

	for _, item := range win.Contents {
		handled := false
		if item.ItemType == ITEM_FLOW {
			handled = item.clickFlows(mpos, click)
		} else {
			handled = item.clickItem(mpos, click)
		}
		if handled {
			return
		}
	}
}

func clickOverlay(root *itemData, mpos point, click bool) bool {
	if root == nil {
		return false
	}
	if !root.DrawRect.containsPoint(mpos) && !dropdownOpenContains([]*itemData{root}, mpos) {
		return false
	}
	if clickOpenDropdown([]*itemData{root}, mpos, click) {
		return true
	}
	if root.ItemType == ITEM_FLOW {
		return root.clickFlows(mpos, click)
	}
	return root.clickItem(mpos, click)
}

func (item *itemData) clickFlows(mpos point, click bool) bool {
	if len(item.Tabs) > 0 {
		if item.ActiveTab >= len(item.Tabs) {
			item.ActiveTab = 0
		}
		for i, tab := range item.Tabs {
			tab.Hovered = false
			if tab.DrawRect.containsPoint(mpos) {
				tab.Hovered = true
				if click {
					activeItem = tab
					item.ActiveTab = i
				}
				return true
			}
		}
		for _, subItem := range item.Tabs[item.ActiveTab].Contents {
			if subItem.ItemType == ITEM_FLOW {
				if subItem.clickFlows(mpos, click) {
					return true
				}
			} else {
				if subItem.clickItem(mpos, click) {
					return true
				}
			}
		}
	} else {
		for _, subItem := range item.Contents {
			if subItem.ItemType == ITEM_FLOW {
				if subItem.clickFlows(mpos, click) {
					return true
				}
			} else {
				if subItem.clickItem(mpos, click) {
					return true
				}
			}
		}
	}
	return item.DrawRect.containsPoint(mpos)
}

func (item *itemData) clickItem(mpos point, click bool) bool {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) && activeItem != nil && activeItem != item {
		return false
	}
	// For dropdowns check the expanded option area as well
	if !item.DrawRect.containsPoint(mpos) {
		if !(item.ItemType == ITEM_DROPDOWN && item.Open && func() bool {
			optionH := item.GetSize().Y
			visible := item.MaxVisible
			if visible <= 0 {
				visible = 5
			}
			startY := item.DrawRect.Y1
			openHeight := optionH * float32(visible)
			r := rect{X0: item.DrawRect.X0, Y0: startY, X1: item.DrawRect.X1, Y1: startY + openHeight}
			return r.containsPoint(mpos)
		}()) {
			return false
		}
	}

	if click {
		activeItem = item
		item.Clicked = time.Now()
		if item.ItemType == ITEM_COLORWHEEL {
			if col, ok := item.colorAt(mpos); ok {
				item.SelectedColor = col
				if item.OnColorChange != nil {
					item.OnColorChange(col)
				} else {
					SetAccentColor(col)
				}
			}
		}
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
		} else if item.ItemType == ITEM_DROPDOWN {
			if item.Open {
				optionH := item.GetSize().Y
				visible := item.MaxVisible
				if visible <= 0 {
					visible = 5
				}
				startY := item.DrawRect.Y1
				openHeight := optionH * float32(visible)
				r := rect{X0: item.DrawRect.X0, Y0: startY, X1: item.DrawRect.X1, Y1: startY + openHeight}
				if r.containsPoint(mpos) {
					idx := int((mpos.Y - startY + item.Scroll.Y) / optionH)
					if idx >= 0 && idx < len(item.Options) {
						item.Selected = idx
						item.Open = false
						if item.OnSelect != nil {
							item.OnSelect(idx)
						}
					}
				} else {
					item.Open = false
				}
			} else {
				item.Open = true
			}
		}
		if item.Action != nil {
			item.Action()
			return true
		}
	} else {
		item.Hovered = true
		if item.ItemType == ITEM_COLORWHEEL && ebiten.IsMouseButtonPressed(ebiten.MouseButton0) {
			if col, ok := item.colorAt(mpos); ok {
				item.SelectedColor = col
				if item.OnColorChange != nil {
					item.OnColorChange(col)
				} else {
					SetAccentColor(col)
				}
			}
		} else if item.ItemType == ITEM_DROPDOWN && item.Open {
			optionH := item.GetSize().Y
			visible := item.MaxVisible
			if visible <= 0 {
				visible = 5
			}
			startY := item.DrawRect.Y1
			openHeight := optionH * float32(visible)
			r := rect{X0: item.DrawRect.X0, Y0: startY, X1: item.DrawRect.X1, Y1: startY + openHeight}
			if r.containsPoint(mpos) {
				idx := int((mpos.Y - startY + item.Scroll.Y) / optionH)
				if idx >= 0 && idx < len(item.Options) {
					if idx != item.HoverIndex {
						item.HoverIndex = idx
						if item.OnHover != nil {
							item.OnHover(idx)
						}
					}
				}
			} else {
				if item.HoverIndex != -1 {
					item.HoverIndex = -1
					if item.OnHover != nil {
						item.OnHover(item.Selected)
					}
				}
			}
		}
		if item.ItemType == ITEM_SLIDER && ebiten.IsMouseButtonPressed(ebiten.MouseButton0) {
			item.setSliderValue(mpos)
			if item.Action != nil {
				item.Action()
			}
		}
	}
	return true
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
	// Determine the width of the slider track accounting for the
	// displayed value text to the right of the knob.
	// Measure against a consistent label width so sliders with
	// different ranges have identical track lengths.
	maxLabel := sliderMaxLabel
	textSize := (item.FontSize * uiScale) + 2
	face := &text.GoTextFace{Source: mplusFaceSource, Size: float64(textSize)}
	maxW, _ := text.Measure(maxLabel, face, 0)

	knobW := item.AuxSize.X * uiScale
	width := item.DrawRect.X1 - item.DrawRect.X0 - knobW - currentLayout.SliderValueGap - float32(maxW)
	if width <= 0 {
		return
	}
	start := item.DrawRect.X0 + knobW/2
	val := (mpos.X - start)
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

func (item *itemData) colorAt(mpos point) (Color, bool) {
	size := point{X: item.Size.X * uiScale, Y: item.Size.Y * uiScale}
	offsetY := float32(0)
	if item.Label != "" {
		offsetY = (item.FontSize*uiScale + 2) + currentLayout.TextPadding
	}
	cx := item.DrawRect.X0 + size.X/2
	cy := item.DrawRect.Y0 + offsetY + size.Y/2
	dx := float64(mpos.X - cx)
	dy := float64(mpos.Y - cy)
	r := float64(size.X) / 2
	dist := math.Hypot(dx, dy)
	if dist > r {
		return Color{}, false
	}
	ang := math.Atan2(dy, dx) * 180 / math.Pi
	if ang < 0 {
		ang += 360
	}
	v := dist / r
	if v < 0 {
		v = 0
	} else if v > 1 {
		v = 1
	}
	col := hsvaToRGBA(ang, 1, v, 1)
	return Color(col), true
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
				} else {
					if req.Y <= size.Y {
						it.Scroll.Y = 0
					}
					if req.X <= size.X {
						it.Scroll.X = 0
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

func scrollDropdown(items []*itemData, mpos point, delta point) bool {
	for _, it := range items {
		if it.ItemType == ITEM_DROPDOWN && it.Open {
			optionH := it.GetSize().Y
			visible := it.MaxVisible
			if visible <= 0 {
				visible = 5
			}
			startY := it.DrawRect.Y1
			openH := optionH * float32(visible)
			r := rect{X0: it.DrawRect.X0, Y0: startY, X1: it.DrawRect.X1, Y1: startY + openH}
			if r.containsPoint(mpos) {
				maxScroll := optionH*float32(len(it.Options)) - openH
				it.Scroll.Y -= delta.Y * optionH
				if it.Scroll.Y < 0 {
					it.Scroll.Y = 0
				}
				if it.Scroll.Y > maxScroll {
					it.Scroll.Y = maxScroll
				}
				return true
			}
		}
		if len(it.Tabs) > 0 {
			if it.ActiveTab >= len(it.Tabs) {
				it.ActiveTab = 0
			}
			if scrollDropdown(it.Tabs[it.ActiveTab].Contents, mpos, delta) {
				return true
			}
		}
		if scrollDropdown(it.Contents, mpos, delta) {
			return true
		}
	}
	return false
}

func scrollWindow(win *windowData, delta point) bool {
	if win.NoScroll {
		return false
	}
	pad := win.Padding + win.BorderPad
	req := win.contentBounds()
	avail := point{
		X: win.GetSize().X - 2*pad,
		Y: win.GetSize().Y - win.GetTitleSize() - 2*pad,
	}
	handled := false
	if req.Y > avail.Y {
		win.Scroll.Y -= delta.Y * 16
		if win.Scroll.Y < 0 {
			win.Scroll.Y = 0
		}
		max := req.Y - avail.Y
		if win.Scroll.Y > max {
			win.Scroll.Y = max
		}
		handled = true
	} else {
		win.Scroll.Y = 0
	}
	if req.X > avail.X {
		win.Scroll.X -= delta.X * 16
		if win.Scroll.X < 0 {
			win.Scroll.X = 0
		}
		max := req.X - avail.X
		if win.Scroll.X > max {
			win.Scroll.X = max
		}
		handled = true
	} else {
		win.Scroll.X = 0
	}
	return handled
}

func dragWindowScroll(win *windowData, mpos point, vert bool) {
	if win.NoScroll {
		return
	}
	pad := win.Padding + win.BorderPad
	req := win.contentBounds()
	avail := point{
		X: win.GetSize().X - 2*pad,
		Y: win.GetSize().Y - win.GetTitleSize() - 2*pad,
	}
	if vert && req.Y > avail.Y {
		barH := avail.Y * avail.Y / req.Y
		maxScroll := req.Y - avail.Y
		track := win.getPosition().Y + win.GetTitleSize() + win.BorderPad
		pos := mpos.Y - (track + barH/2)
		if pos < 0 {
			pos = 0
		}
		if pos > avail.Y-barH {
			pos = avail.Y - barH
		}
		if avail.Y-barH > 0 {
			win.Scroll.Y = (pos / (avail.Y - barH)) * maxScroll
		} else {
			win.Scroll.Y = 0
		}
	} else if vert {
		win.Scroll.Y = 0
	}
	if !vert && req.X > avail.X {
		barW := avail.X * avail.X / req.X
		maxScroll := req.X - avail.X
		track := win.getPosition().X + win.BorderPad
		pos := mpos.X - (track + barW/2)
		if pos < 0 {
			pos = 0
		}
		if pos > avail.X-barW {
			pos = avail.X - barW
		}
		if avail.X-barW > 0 {
			win.Scroll.X = (pos / (avail.X - barW)) * maxScroll
		} else {
			win.Scroll.X = 0
		}
	} else if !vert {
		win.Scroll.X = 0
	}
}
func dropdownOpenContains(items []*itemData, mpos point) bool {
	for _, it := range items {
		if it.ItemType == ITEM_DROPDOWN && it.Open {
			optionH := it.GetSize().Y
			visible := it.MaxVisible
			if visible <= 0 {
				visible = 5
			}
			startY := it.DrawRect.Y1
			openH := optionH * float32(visible)
			r := rect{X0: it.DrawRect.X0, Y0: startY, X1: it.DrawRect.X1, Y1: startY + openH}
			if r.containsPoint(mpos) {
				return true
			}
		}
		if len(it.Tabs) > 0 {
			if it.ActiveTab >= len(it.Tabs) {
				it.ActiveTab = 0
			}
			if dropdownOpenContains(it.Tabs[it.ActiveTab].Contents, mpos) {
				return true
			}
		}
		if dropdownOpenContains(it.Contents, mpos) {
			return true
		}
	}
	return false
}

func clickOpenDropdown(items []*itemData, mpos point, click bool) bool {
	for _, it := range items {
		if it.ItemType == ITEM_DROPDOWN && it.Open {
			optionH := it.GetSize().Y
			visible := it.MaxVisible
			if visible <= 0 {
				visible = 5
			}
			startY := it.DrawRect.Y1
			openH := optionH * float32(visible)
			r := rect{X0: it.DrawRect.X0, Y0: startY, X1: it.DrawRect.X1, Y1: startY + openH}
			if r.containsPoint(mpos) {
				it.clickItem(mpos, click)
				return true
			}
		}
		if len(it.Tabs) > 0 {
			if it.ActiveTab >= len(it.Tabs) {
				it.ActiveTab = 0
			}
			if clickOpenDropdown(it.Tabs[it.ActiveTab].Contents, mpos, click) {
				return true
			}
		}
		if clickOpenDropdown(it.Contents, mpos, click) {
			return true
		}
	}
	return false
}

func dropdownOpenContainsAnywhere(mpos point) bool {
	for _, win := range windows {
		if win.Open && dropdownOpenContains(win.Contents, mpos) {
			return true
		}
	}
	for _, ov := range overlays {
		if dropdownOpenContains([]*itemData{ov}, mpos) {
			return true
		}
	}
	return false
}

func closeDropdowns(items []*itemData) {
	for _, it := range items {
		if it.ItemType == ITEM_DROPDOWN {
			it.Open = false
		}
		for _, tab := range it.Tabs {
			closeDropdowns(tab.Contents)
		}
		closeDropdowns(it.Contents)
	}
}

func closeAllDropdowns() {
	for _, win := range windows {
		if win.Open {
			closeDropdowns(win.Contents)
		}
	}
	for _, ov := range overlays {
		closeDropdowns([]*itemData{ov})
	}
}
