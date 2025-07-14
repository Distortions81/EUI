package main

import (
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

func (win *windowData) getWinRect() rect {
	winPos := win.getPosition()
	return rect{
		X0: winPos.X,
		Y0: winPos.Y,
		X1: winPos.X + win.GetSize().X,
		Y1: winPos.Y + win.GetSize().Y,
	}
}

func (item *itemData) getItemRect(win *windowData) rect {
	return rect{
		X0: win.getPosition().X + (item.getPosition(win).X),
		Y0: win.getPosition().Y + (item.getPosition(win).Y),
		X1: win.getPosition().X + (item.getPosition(win).X) + (item.GetSize().X),
		Y1: win.getPosition().Y + (item.getPosition(win).Y) + (item.GetSize().Y),
	}
}

func (parent *itemData) addItemTo(item *itemData) {
	item.Parent = parent
	parent.Contents = append(parent.Contents, item)
}

func (parent *windowData) addItemTo(item *itemData) {
	parent.Contents = append(parent.Contents, item)
}

func (win *windowData) getMainRect() rect {
	pad := win.Padding + win.BorderPad
	return rect{
		X0: win.getPosition().X + pad,
		Y0: win.getPosition().Y + win.GetTitleSize() + pad + 1,
		X1: win.getPosition().X + win.GetSize().X - pad,
		Y1: win.getPosition().Y + win.GetSize().Y - win.GetTitleSize() - pad,
	}
}

func (win *windowData) getTitleRect() rect {
	if win.TitleHeight <= 0 {
		return rect{}
	}
	return rect{
		X0: win.getPosition().X, Y0: win.getPosition().Y,
		X1: win.getPosition().X + win.GetSize().X,
		Y1: win.getPosition().Y + (win.GetTitleSize()),
	}
}

func (win *windowData) xRect() rect {
	if win.TitleHeight <= 0 || !win.Closable {
		return rect{}
	}

	var xpad float32 = win.Border
	return rect{
		X0: win.getPosition().X + win.GetSize().X - (win.GetTitleSize()) + xpad,
		Y0: win.getPosition().Y + xpad,

		X1: win.getPosition().X + win.GetSize().X - xpad,
		Y1: win.getPosition().Y + (win.GetTitleSize()) - xpad,
	}
}

func (win *windowData) dragbarRect() rect {
	if win.TitleHeight <= 0 && !win.Resizable {
		return rect{}
	}
	textSize := win.titleTextWidth()
	xRect := win.xRect()
	buttonsWidth := xRect.X1 - xRect.X0 + 3

	dpad := (win.GetTitleSize()) / 5
	xStart := textSize.X + float32((win.GetTitleSize())/1.5)
	xEnd := (win.GetSize().X - buttonsWidth)
	return rect{
		X0: win.getPosition().X + xStart, Y0: win.getPosition().Y + dpad,
		X1: win.getPosition().X + xEnd, Y1: win.getPosition().Y + (win.GetTitleSize()) - dpad,
	}
}

func (win *windowData) setSize(size point) bool {

	tooSmall := false
	xc, yc := win.itemOverlap(size)
	if !xc {
		win.Size.X = size.X
	}
	if !yc {
		win.Size.Y = size.Y
	}
	if yc && xc {
		tooSmall = true
	}

	if size.X < minWinSizeX {
		size.X = minWinSizeX
		tooSmall = true
	}

	if size.Y < minWinSizeY {
		size.Y = minWinSizeY
		tooSmall = true
	}

	win.BringForward()
	win.resizeFlows()
	return tooSmall
}

const cornerSize = 14
const tol = 2

func (win *windowData) getWindowPart(mpos point, click bool) dragType {

	//Titlebar items
	if win.TitleHeight > 0 {
		if win.getTitleRect().containsPoint(mpos) {
			//Close X
			if win.Closable {
				if win.xRect().containsPoint(mpos) {
					win.HoverClose = true
					return PART_CLOSE
				}
			}
			//Drag bar
			if win.Movable {
				if win.dragbarRect().containsPoint(mpos) {
					win.HoverDragbar = true
					return PART_BAR
				}
			}
		}
	}

	if !win.Resizable {
		return PART_NONE
	}

	t := tol * uiScale
	cs := cornerSize * uiScale

	winRect := win.getWinRect()

	//Expand outer window rect
	outRect := winRect
	outRect.X0 -= t
	outRect.X1 += t
	outRect.Y0 -= t
	outRect.Y1 += t

	//Contract inner window rect
	inRect := winRect
	inRect.X0 += t
	inRect.X1 -= t
	inRect.Y0 += t
	inRect.Y1 -= t

	//If within outrect, and not within inrect it is window DRAG
	if outRect.containsPoint(mpos) && !inRect.containsPoint(mpos) {
		if mpos.Y < outRect.Y0+cs {
			if mpos.X < inRect.X0+cs {
				return PART_TOP_LEFT
			} else if mpos.X > inRect.X1-cs {
				return PART_TOP_RIGHT
			} else {
				return PART_TOP
			}
		} else if mpos.Y > inRect.Y1-cs {
			if mpos.X > outRect.X1-cs {
				return PART_BOTTOM_RIGHT
			} else if mpos.X < outRect.X0+cs {
				return PART_BOTTOM_LEFT
			} else {
				return PART_BOTTOM
			}
		} else if mpos.X < inRect.X0 {
			return PART_LEFT
		} else if mpos.X < outRect.X1 {
			return PART_RIGHT
		}
	}

	return PART_NONE
}

func (win *windowData) titleTextWidth() point {
	if win.TitleHeight <= 0 {
		return point{}
	}
	textSize := ((win.GetTitleSize()) / 1.5)
	face := &text.GoTextFace{
		Source: mplusFaceSource,
		Size:   float64(textSize),
	}
	textWidth, textHeight := text.Measure(win.Title, face, 0)
	return point{X: float32(textWidth), Y: float32(textHeight)}
}

func (win *windowData) SetTitleSize(size float32) {
	win.TitleHeight = size / uiScale
}

func SetUIScale(scale float32) {
	uiScale = scale
}

func (win *windowData) GetTitleSize() float32 {
	return win.TitleHeight * uiScale
}

func (win *windowData) GetSize() point {
	return point{X: win.Size.X * uiScale, Y: win.Size.Y * uiScale}
}

func (win *windowData) GetPos() point {
	return point{X: win.Position.X * uiScale, Y: win.Position.Y * uiScale}
}

func (item *itemData) GetSize() point {
	return point{X: item.Size.X * uiScale, Y: item.Size.Y * uiScale}
}

func (item *itemData) GetPos() point {
	return point{X: item.Position.X * uiScale, Y: item.Position.Y * uiScale}
}

func (item *itemData) bounds(offset point) rect {
	r := rect{
		X0: offset.X,
		Y0: offset.Y,
		X1: offset.X + item.GetSize().X,
		Y1: offset.Y + item.GetSize().Y,
	}
	if item.ItemType == ITEM_FLOW {
		var flowOffset point
		var subItems []*itemData
		if len(item.Tabs) > 0 {
			if item.ActiveTab >= len(item.Tabs) {
				item.ActiveTab = 0
			}
			subItems = item.Tabs[item.ActiveTab].Contents
		} else {
			subItems = item.Contents
		}
		for _, sub := range subItems {
			var off point
			if item.FlowType == FLOW_HORIZONTAL {
				off = pointAdd(offset, point{X: flowOffset.X + sub.GetPos().X, Y: sub.GetPos().Y})
			} else if item.FlowType == FLOW_VERTICAL {
				off = pointAdd(offset, point{X: sub.GetPos().X, Y: flowOffset.Y + sub.GetPos().Y})
			} else {
				off = pointAdd(offset, pointAdd(flowOffset, sub.GetPos()))
			}
			sr := sub.bounds(off)
			r = unionRect(r, sr)
			if item.FlowType == FLOW_HORIZONTAL {
				flowOffset.X += sub.GetSize().X + sub.GetPos().X
			} else if item.FlowType == FLOW_VERTICAL {
				flowOffset.Y += sub.GetSize().Y + sub.GetPos().Y
			}
		}
	} else {
		for _, sub := range item.Contents {
			off := pointAdd(offset, sub.GetPos())
			r = unionRect(r, sub.bounds(off))
		}
	}
	return r
}

func (win *windowData) contentBounds() point {
	if len(win.Contents) == 0 {
		return point{}
	}
	base := point{X: 0, Y: win.GetTitleSize()}
	b := win.Contents[0].bounds(pointAdd(base, win.Contents[0].GetPos()))
	for _, item := range win.Contents[1:] {
		r := item.bounds(pointAdd(base, item.GetPos()))
		b = unionRect(b, r)
	}
	return point{X: b.X1 - base.X, Y: b.Y1 - base.Y}
}

func (win *windowData) updateAutoSize() {
	req := win.contentBounds()
	size := win.GetSize()
	if req.X > size.X {
		size.X = req.X
	}
	if req.Y+win.GetTitleSize() > size.Y {
		size.Y = req.Y + win.GetTitleSize()
	}
	if size.X > float32(screenWidth) {
		size.X = float32(screenWidth)
	}
	if size.Y > float32(screenHeight) {
		size.Y = float32(screenHeight)
	}
	win.Size = point{X: size.X / uiScale, Y: size.Y / uiScale}
	win.resizeFlows()
}

func (item *itemData) contentBounds() point {
	list := item.Contents
	if len(item.Tabs) > 0 {
		if item.ActiveTab >= len(item.Tabs) {
			item.ActiveTab = 0
		}
		list = item.Tabs[item.ActiveTab].Contents
	}
	if len(list) == 0 {
		return point{}
	}
	base := point{}
	b := list[0].bounds(pointAdd(base, list[0].GetPos()))
	for _, sub := range list[1:] {
		r := sub.bounds(pointAdd(base, sub.GetPos()))
		b = unionRect(b, r)
	}
	return point{X: b.X1 - base.X, Y: b.Y1 - base.Y}
}

func (item *itemData) resizeFlow(parentSize point) {
	available := parentSize

	if item.ItemType == ITEM_FLOW {
		size := available
		if item.Fixed {
			size = item.GetSize()
		}
		if !item.Scrollable {
			req := item.contentBounds()
			if req.X > size.X {
				size.X = req.X
			}
			if req.Y > size.Y {
				size.Y = req.Y
			}
		}
		item.Size = point{X: size.X / uiScale, Y: size.Y / uiScale}
		available = item.GetSize()
	} else {
		available = item.GetSize()
	}

	var list []*itemData
	if len(item.Tabs) > 0 {
		if item.ActiveTab >= len(item.Tabs) {
			item.ActiveTab = 0
		}
		list = item.Tabs[item.ActiveTab].Contents
	} else {
		list = item.Contents
	}
	for _, sub := range list {
		sub.resizeFlow(available)
	}
}

func (win *windowData) resizeFlows() {
	for _, item := range win.Contents {
		item.resizeFlow(win.GetSize())
	}
}
