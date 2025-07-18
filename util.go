package main

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var (
        strokeLineFn = vector.StrokeLine
        strokeRectFn = vector.StrokeRect
)

func (item *itemData) themeStyle() *itemData {
       if item == nil || item.Theme == nil {
               return nil
       }
       switch item.ItemType {
       case ITEM_BUTTON:
               return &item.Theme.Button
       case ITEM_TEXT:
               return &item.Theme.Text
       case ITEM_CHECKBOX:
               return &item.Theme.Checkbox
       case ITEM_RADIO:
               return &item.Theme.Radio
       case ITEM_INPUT:
               return &item.Theme.Input
       case ITEM_SLIDER:
               return &item.Theme.Slider
       case ITEM_DROPDOWN:
               return &item.Theme.Dropdown
       case ITEM_FLOW:
               if len(item.Tabs) > 0 {
                       return &item.Theme.Tab
               }
       }
       return nil
}

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
	if currentTheme != nil {
		applyThemeToItem(item)
	}
	parent.Contents = append(parent.Contents, item)
}

func (parent *windowData) addItemTo(item *itemData) {
	if currentTheme != nil {
		applyThemeToItem(item)
	}
	parent.Contents = append(parent.Contents, item)
}

func (win *windowData) getMainRect() rect {
	return rect{
		X0: win.getPosition().X,
		Y0: win.getPosition().Y + win.GetTitleSize(),
		X1: win.getPosition().X + win.GetSize().X,
		Y1: win.getPosition().Y + win.GetSize().Y,
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

	// Enforce minimum dimensions and prevent negatives.
	if size.X < MinWinSizeX {
		size.X = MinWinSizeX
		tooSmall = true
	}
	if size.Y < MinWinSizeY {
		size.Y = MinWinSizeY
		tooSmall = true
	}
	if size.X < 0 {
		size.X = MinWinSizeX
		tooSmall = true
	}
	if size.Y < 0 {
		size.Y = MinWinSizeY
		tooSmall = true
	}

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

	win.BringForward()
	win.resizeFlows()

	// Adjust scroll position so content never remains off screen when
	// the window is resized larger than its contents.
	if !win.NoScroll {
		pad := win.Padding + win.BorderPad
		req := win.contentBounds()
		avail := point{
			X: win.GetSize().X - 2*pad,
			Y: win.GetSize().Y - win.GetTitleSize() - 2*pad,
		}
		if req.Y <= avail.Y {
			win.Scroll.Y = 0
		} else {
			max := req.Y - avail.Y
			if win.Scroll.Y > max {
				win.Scroll.Y = max
			}
		}
		if req.X <= avail.X {
			win.Scroll.X = 0
		} else {
			max := req.X - avail.X
			if win.Scroll.X > max {
				win.Scroll.X = max
			}
		}
	}

	return tooSmall
}

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

	if win.Resizable {
		t := Tol * uiScale
		cs := CornerSize * uiScale

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
	}

	if !win.NoScroll {
		pad := win.Padding + win.BorderPad
		req := win.contentBounds()
		avail := point{
			X: win.GetSize().X - 2*pad,
			Y: win.GetSize().Y - win.GetTitleSize() - 2*pad,
		}
		if req.Y > avail.Y {
			barH := avail.Y * avail.Y / req.Y
			maxScroll := req.Y - avail.Y
			pos := float32(0)
			if maxScroll > 0 {
				pos = (win.Scroll.Y / maxScroll) * (avail.Y - barH)
			}
			sbW := currentLayout.BorderPad.Slider * 2
			r := rect{
				X0: win.getPosition().X + win.GetSize().X - win.BorderPad - sbW,
				Y0: win.getPosition().Y + win.GetTitleSize() + win.BorderPad + pos,
				X1: win.getPosition().X + win.GetSize().X - win.BorderPad,
				Y1: win.getPosition().Y + win.GetTitleSize() + win.BorderPad + pos + barH,
			}
			if r.containsPoint(mpos) {
				return PART_SCROLL_V
			}
		}
		if req.X > avail.X {
			barW := avail.X * avail.X / req.X
			maxScroll := req.X - avail.X
			pos := float32(0)
			if maxScroll > 0 {
				pos = (win.Scroll.X / maxScroll) * (avail.X - barW)
			}
			sbW := currentLayout.BorderPad.Slider * 2
			r := rect{
				X0: win.getPosition().X + win.BorderPad + pos,
				Y0: win.getPosition().Y + win.GetSize().Y - win.BorderPad - sbW,
				X1: win.getPosition().X + win.BorderPad + pos + barW,
				Y1: win.getPosition().Y + win.GetSize().Y - win.BorderPad,
			}
			if r.containsPoint(mpos) {
				return PART_SCROLL_H
			}
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
	sz := point{X: item.Size.X * uiScale, Y: item.Size.Y * uiScale}
	if item.Label != "" {
		textSize := (item.FontSize * uiScale) + 2
		sz.Y += textSize + currentLayout.TextPadding
	}
	return sz
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
	first := true
	var b rect

	for _, item := range win.Contents {
		var r rect
		if item.ItemType == ITEM_FLOW {
			cb := item.contentBounds()
			r = rect{
				X0: base.X + item.GetPos().X,
				Y0: base.Y + item.GetPos().Y,
				X1: base.X + item.GetPos().X + cb.X,
				Y1: base.Y + item.GetPos().Y + cb.Y,
			}
		} else {
			r = item.bounds(pointAdd(base, item.GetPos()))
		}
		if first {
			b = r
			first = false
		} else {
			b = unionRect(b, r)
		}
	}

	if first {
		return point{}
	}
	return point{X: b.X1 - base.X, Y: b.Y1 - base.Y}
}

func (win *windowData) updateAutoSize() {
	req := win.contentBounds()
	pad := win.Padding + win.BorderPad

	size := win.GetSize()
	needX := req.X + 2*pad
	if needX > size.X {
		size.X = needX
	}

	// Always include the titlebar height in the calculated size
	size.Y = req.Y + win.GetTitleSize() + 2*pad
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
	first := true
	var b rect
	var flowOffset point

	for _, sub := range list {
		off := pointAdd(base, sub.GetPos())
		if item.ItemType == ITEM_FLOW {
			if item.FlowType == FLOW_HORIZONTAL {
				off = pointAdd(base, point{X: flowOffset.X + sub.GetPos().X, Y: sub.GetPos().Y})
			} else if item.FlowType == FLOW_VERTICAL {
				off = pointAdd(base, point{X: sub.GetPos().X, Y: flowOffset.Y + sub.GetPos().Y})
			} else {
				off = pointAdd(base, pointAdd(flowOffset, sub.GetPos()))
			}
		}

		r := sub.bounds(off)
		if first {
			b = r
			first = false
		} else {
			b = unionRect(b, r)
		}

		if item.ItemType == ITEM_FLOW {
			if item.FlowType == FLOW_HORIZONTAL {
				flowOffset.X += sub.GetSize().X + sub.GetPos().X
			} else if item.FlowType == FLOW_VERTICAL {
				flowOffset.Y += sub.GetSize().Y + sub.GetPos().Y
			}
		}
	}

	if first {
		return point{}
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

	if item.ItemType == ITEM_FLOW {
		req := item.contentBounds()
		size := item.GetSize()
		if req.Y <= size.Y {
			item.Scroll.Y = 0
		} else {
			max := req.Y - size.Y
			if item.Scroll.Y > max {
				item.Scroll.Y = max
			}
		}
		if req.X <= size.X {
			item.Scroll.X = 0
		} else {
			max := req.X - size.X
			if item.Scroll.X > max {
				item.Scroll.X = max
			}
		}
	}
}

func (win *windowData) resizeFlows() {
	for _, item := range win.Contents {
		item.resizeFlow(win.GetSize())
	}
}

func pixelOffset(width float32) float32 {
	if int(math.Round(float64(width)))%2 == 0 {
		return 0
	}
	return 0.5
}

func strokeLine(dst *ebiten.Image, x0, y0, x1, y1, width float32, col color.Color, aa bool) {
	width = float32(math.Round(float64(width)))
	off := pixelOffset(width)
	x0 = float32(math.Round(float64(x0))) + off
	y0 = float32(math.Round(float64(y0))) + off
	x1 = float32(math.Round(float64(x1))) + off
	y1 = float32(math.Round(float64(y1))) + off
	strokeLineFn(dst, x0, y0, x1, y1, width, col, aa)
}

func strokeRect(dst *ebiten.Image, x, y, w, h, width float32, col color.Color, aa bool) {
	width = float32(math.Round(float64(width)))
	off := pixelOffset(width)
	x = float32(math.Round(float64(x))) + off
	y = float32(math.Round(float64(y))) + off
	w = float32(math.Round(float64(w)))
	h = float32(math.Round(float64(h)))
	strokeRectFn(dst, x, y, w, h, width, col, aa)
}

func drawFilledRect(dst *ebiten.Image, x, y, w, h float32, col color.Color, aa bool) {
	x = float32(math.Round(float64(x)))
	y = float32(math.Round(float64(y)))
	w = float32(math.Round(float64(w)))
	h = float32(math.Round(float64(h)))
	vector.DrawFilledRect(dst, x, y, w, h, col, aa)
}
