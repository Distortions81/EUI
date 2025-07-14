package main

import (
	"fmt"
	"image/color"
	"math"
	"strings"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type dropdownRender struct {
	item   *itemData
	offset point
	clip   rect
}

var pendingDropdowns []dropdownRender

func (g *Game) Draw(screen *ebiten.Image) {

	pendingDropdowns = pendingDropdowns[:0]

	for _, win := range windows {
		if !win.Open {
			continue
		}

		win.Draw(screen)
	}

	for _, dr := range pendingDropdowns {
		drawDropdownOptions(dr.item, dr.offset, dr.clip, screen)
	}

	drawFPS(screen)
}

func (win *windowData) Draw(screen *ebiten.Image) {
	win.drawBG(screen)
	win.drawItems(screen)
	titleArea := screen.SubImage(win.getTitleRect().getRectangle()).(*ebiten.Image)
	win.drawWinTitle(titleArea)
	windowArea := screen.SubImage(win.getWinRect().getRectangle()).(*ebiten.Image)
	//win.drawResizeTab(windowArea)
	win.drawBorder(windowArea)
	win.drawDebug(screen)
}

func (win *windowData) drawBG(screen *ebiten.Image) {
	windowArea := screen.SubImage(win.getWinRect().getRectangle()).(*ebiten.Image)
	windowArea.Fill(win.BGColor)
}

func (win *windowData) drawWinTitle(screen *ebiten.Image) {
	// Window Title
	if win.TitleHeight > 0 {
		screen.Fill(win.TitleBGColor)

		textSize := ((win.GetTitleSize()) / 1.5)
		face := &text.GoTextFace{
			Source: mplusFaceSource,
			Size:   float64(textSize),
		}

		skipTitleText := false
		textWidth, textHeight := text.Measure(win.Title, face, 0)
		if textWidth > float64(win.GetSize().X) ||
			textHeight > float64(win.GetTitleSize()) {
			skipTitleText = true
			//log.Print("Title text too big for title size.")
		}

		//Title text
		if !skipTitleText {
			loo := text.LayoutOptions{
				LineSpacing:    0, //No multi-line titles
				PrimaryAlign:   text.AlignStart,
				SecondaryAlign: text.AlignCenter,
			}
			tdop := ebiten.DrawImageOptions{}
			tdop.GeoM.Translate(float64(win.getPosition().X+((win.GetTitleSize())/4)),
				float64(win.getPosition().Y+((win.GetTitleSize())/2)))

			top := &text.DrawOptions{DrawImageOptions: tdop, LayoutOptions: loo}

			top.ColorScale.ScaleWithColor(win.TitleTextColor)
			buf := strings.ReplaceAll(win.Title, "\n", "") //Remove newline
			buf = strings.ReplaceAll(buf, "\r", "")        //Remove return
			text.Draw(screen, buf, face, top)
		} else {
			textWidth = 0
		}

		//Close X
		var buttonsWidth float32 = 0
		if win.Closable {
			var xpad float32 = (win.GetTitleSize()) / 4.0
			color := win.TitleColor
			// fill background for close area if configured
			if win.CloseBGColor.A > 0 {
				r := win.xRect()
				closeArea := screen.SubImage(r.getRectangle()).(*ebiten.Image)
				closeArea.Fill(win.CloseBGColor)
			}
			xThick := 3 * uiScale
			if win.HoverClose {
				color = win.HoverTitleColor
				win.HoverClose = false
			}
			vector.StrokeLine(screen,
				win.getPosition().X+win.GetSize().X-(win.GetTitleSize())+xpad,
				win.getPosition().Y+xpad,

				win.getPosition().X+win.GetSize().X-xpad,
				win.getPosition().Y+(win.GetTitleSize())-xpad,
				xThick, color, true)
			vector.StrokeLine(screen,
				win.getPosition().X+win.GetSize().X-xpad,
				win.getPosition().Y+xpad,

				win.getPosition().X+win.GetSize().X-(win.GetTitleSize())+xpad,
				win.getPosition().Y+(win.GetTitleSize())-xpad,
				xThick, color, true)

			buttonsWidth += (win.GetTitleSize())
		}

		//Dragbar
		if win.Movable && win.ShowDragbar {
			var xThick float32 = 1
			xColor := win.DragbarColor
			if win.HoverDragbar {
				xColor = win.HoverTitleColor
				win.HoverDragbar = false
			}
			dpad := (win.GetTitleSize()) / 5
			spacing := win.DragbarSpacing
			if spacing <= 0 {
				spacing = 5
			}
			for x := textWidth + float64((win.GetTitleSize())/1.5); x < float64(win.GetSize().X-buttonsWidth); x = x + float64(uiScale*spacing) {
				vector.StrokeLine(screen,
					win.getPosition().X+float32(x), win.getPosition().Y+dpad,
					win.getPosition().X+float32(x), win.getPosition().Y+(win.GetTitleSize())-dpad,
					xThick, xColor, false)
			}
		}
	}
}

func (win *windowData) drawBorder(screen *ebiten.Image) {
	//Draw borders
	if win.Outlined && win.Border > 0 {
		FrameColor := win.BorderColor
		if activeWindow == win {
			FrameColor = win.ActiveColor
		} else if win.Hovered {
			FrameColor = win.HoverColor
			win.Hovered = false
		}
		if win.TitleHeight > 0 {
			vector.StrokeRect(screen,
				win.getPosition().X+1, win.getPosition().Y+1,
				win.GetSize().X-1, (win.GetTitleSize())-1,
				win.Border, FrameColor, false)
		}
		//Window border
		vector.StrokeRect(screen,
			win.getPosition().X+1, win.getPosition().Y+1,
			win.GetSize().X-1, win.GetSize().Y-(win.GetTitleSize())-1,
			win.Border, FrameColor, false)
	}
}

func (win *windowData) drawItems(screen *ebiten.Image) {
	winPos := pointAdd(win.getPosition(), point{X: 0, Y: win.GetTitleSize()})
	clip := win.getMainRect()

	for _, item := range win.Contents {
		itemPos := pointAdd(winPos, item.getPosition(win))

		if item.ItemType == ITEM_FLOW {
			item.drawFlows(nil, itemPos, clip, screen)
		} else {
			item.drawItem(nil, itemPos, clip, screen)
		}
	}
}

func (item *itemData) drawFlows(parent *itemData, offset point, clip rect, screen *ebiten.Image) {
	// Store the drawn rectangle for input handling
	itemRect := rect{
		X0: offset.X,
		Y0: offset.Y,
		X1: offset.X + item.GetSize().X,
		Y1: offset.Y + item.GetSize().Y,
	}
	item.DrawRect = intersectRect(itemRect, clip)

	if item.DrawRect.X1 <= item.DrawRect.X0 || item.DrawRect.Y1 <= item.DrawRect.Y0 {
		return
	}
	subImg := screen.SubImage(item.DrawRect.getRectangle()).(*ebiten.Image)

	var activeContents []*itemData
	drawOffset := pointSub(offset, item.Scroll)

	if len(item.Tabs) > 0 {
		if item.ActiveTab >= len(item.Tabs) {
			item.ActiveTab = 0
		}

		tabHeight := float32(defaultTabHeight) * uiScale
		if th := item.FontSize*uiScale + 4; th > tabHeight {
			tabHeight = th
		}
		x := offset.X
		for i, tab := range item.Tabs {
			face := &text.GoTextFace{Source: mplusFaceSource, Size: float64(tabHeight - 22)}
			tw, _ := text.Measure(tab.Name, face, 0)
			w := float32(tw) + 8
			if w < float32(defaultTabWidth)*uiScale {
				w = float32(defaultTabWidth) * uiScale
			}
			col := item.Color
			if i == item.ActiveTab {
				col = item.ClickColor
			}
			drawTabShape(subImg, point{X: x, Y: offset.Y}, point{X: w, Y: tabHeight}, col)
			loo := text.LayoutOptions{PrimaryAlign: text.AlignCenter, SecondaryAlign: text.AlignCenter}
			dop := ebiten.DrawImageOptions{}
			dop.GeoM.Translate(float64(x+w/2), float64(offset.Y+tabHeight/2))
			text.Draw(subImg, tab.Name, face, &text.DrawOptions{DrawImageOptions: dop, LayoutOptions: loo})
			tab.DrawRect = rect{X0: x, Y0: offset.Y, X1: x + w, Y1: offset.Y + tabHeight}
			x += w
		}
		drawOffset = pointAdd(drawOffset, point{Y: tabHeight})
		vector.StrokeRect(subImg,
			offset.X,
			offset.Y+tabHeight,
			item.GetSize().X,
			item.GetSize().Y-tabHeight,
			1,
			item.Color,
			false)
		activeContents = item.Tabs[item.ActiveTab].Contents
	} else {
		activeContents = item.Contents
	}

	var flowOffset point

	for _, subItem := range activeContents {

		if subItem.ItemType == ITEM_FLOW {
			flowPos := pointAdd(drawOffset, item.GetPos())
			flowOff := pointAdd(flowPos, flowOffset)
			itemPos := pointAdd(flowOff, subItem.GetPos())
			subItem.drawFlows(item, itemPos, item.DrawRect, screen)
		} else {
			flowOff := pointAdd(drawOffset, flowOffset)

			objOff := flowOff
			if parent != nil && parent.ItemType == ITEM_FLOW {
				if parent.FlowType == FLOW_HORIZONTAL {
					objOff = pointAdd(objOff, point{X: subItem.GetPos().X})
				} else if parent.FlowType == FLOW_VERTICAL {
					objOff = pointAdd(objOff, point{Y: subItem.GetPos().Y})
				}
			}

			subItem.drawItem(item, objOff, item.DrawRect, screen)
		}

		if item.ItemType == ITEM_FLOW {
			if item.FlowType == FLOW_HORIZONTAL {
				flowOffset = pointAdd(flowOffset, point{X: subItem.GetSize().X, Y: 0})
				flowOffset = pointAdd(flowOffset, point{X: subItem.GetPos().X})
			} else if item.FlowType == FLOW_VERTICAL {
				flowOffset = pointAdd(flowOffset, point{X: 0, Y: subItem.GetSize().Y})
				flowOffset = pointAdd(flowOffset, point{Y: subItem.GetPos().Y})
			}
		}
	}

	if item.Scrollable {
		req := item.contentBounds()
		size := item.GetSize()
		if item.FlowType == FLOW_VERTICAL && req.Y > size.Y {
			barH := size.Y * size.Y / req.Y
			maxScroll := req.Y - size.Y
			pos := float32(0)
			if maxScroll > 0 {
				pos = (item.Scroll.Y / maxScroll) * (size.Y - barH)
			}
			vector.DrawFilledRect(subImg, item.DrawRect.X1-4, item.DrawRect.Y0+pos, 4, barH, color.RGBA{R: 96, G: 96, B: 96, A: 192}, false)
		} else if item.FlowType == FLOW_HORIZONTAL && req.X > size.X {
			barW := size.X * size.X / req.X
			maxScroll := req.X - size.X
			pos := float32(0)
			if maxScroll > 0 {
				pos = (item.Scroll.X / maxScroll) * (size.X - barW)
			}
			vector.DrawFilledRect(subImg, item.DrawRect.X0+pos, item.DrawRect.Y1-4, barW, 4, color.RGBA{R: 96, G: 96, B: 96, A: 192}, false)
		}
	}
}

func (item *itemData) drawItem(parent *itemData, offset point, clip rect, screen *ebiten.Image) {

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
		X0: offset.X,
		Y0: offset.Y,
		X1: offset.X + maxSize.X,
		Y1: offset.Y + maxSize.Y,
	}
	item.DrawRect = intersectRect(itemRect, clip)
	if item.DrawRect.X1 <= item.DrawRect.X0 || item.DrawRect.Y1 <= item.DrawRect.Y0 {
		return
	}
	subImg := screen.SubImage(item.DrawRect.getRectangle()).(*ebiten.Image)

	if item.ItemType == ITEM_CHECKBOX {

		bThick := float32(1.0)
		itemColor := item.Color
		bColor := item.ClickColor
		if item.Checked {
			itemColor = item.ClickColor
			bColor = item.Color
			bThick = 2.0
		} else if item.Hovered {
			item.Hovered = false
			itemColor = item.HoverColor
		}
		auxSize := pointScaleMul(item.AuxSize)
		drawRoundRect(subImg, &roundRect{
			Size:     auxSize,
			Position: offset, Fillet: item.Fillet, Filled: true, Color: itemColor})
		drawRoundRect(subImg, &roundRect{
			Size:     auxSize,
			Position: offset, Fillet: item.Fillet, Filled: false, Color: bColor, Border: bThick * uiScale})

		if item.Checked {
			xThick := 2 * uiScale
			margin := auxSize.X * 0.25
			vector.StrokeLine(subImg,
				offset.X+margin,
				offset.Y+margin,
				offset.X+auxSize.X-margin,
				offset.Y+auxSize.Y-margin,
				xThick, item.TextColor, true)

			vector.StrokeLine(subImg,
				offset.X+auxSize.X-margin,
				offset.Y+margin,
				offset.X+margin,
				offset.Y+auxSize.Y-margin,
				xThick, item.TextColor, true)
		}

		textSize := (item.FontSize * uiScale) + 2
		face := &text.GoTextFace{
			Source: mplusFaceSource,
			Size:   float64(textSize),
		}
		loo := text.LayoutOptions{
			LineSpacing:    1.2,
			PrimaryAlign:   text.AlignStart,
			SecondaryAlign: text.AlignCenter,
		}
		tdop := ebiten.DrawImageOptions{}
		tdop.GeoM.Translate(
			float64(offset.X+auxSize.X+item.AuxSpace),
			float64(offset.Y+(auxSize.Y/2)),
		)
		top := &text.DrawOptions{DrawImageOptions: tdop, LayoutOptions: loo}
		top.ColorScale.ScaleWithColor(item.TextColor)
		text.Draw(subImg, item.Text, face, top)

	} else if item.ItemType == ITEM_RADIO {

		bThick := float32(1.0)
		itemColor := item.Color
		bColor := item.ClickColor
		if item.Checked {
			itemColor = item.ClickColor
			bColor = item.Color
			bThick = 2.0
		} else if item.Hovered {
			item.Hovered = false
			itemColor = item.HoverColor
		}
		auxSize := pointScaleMul(item.AuxSize)
		drawRoundRect(subImg, &roundRect{
			Size:     auxSize,
			Position: offset,
			Fillet:   auxSize.X / 2,
			Filled:   true,
			Color:    itemColor,
		})
		drawRoundRect(subImg, &roundRect{
			Size:     auxSize,
			Position: offset,
			Fillet:   auxSize.X / 2,
			Filled:   false,
			Color:    bColor,
			Border:   bThick * uiScale,
		})
		if item.Checked {
			inner := auxSize.X / 2.5
			drawRoundRect(subImg, &roundRect{
				Size:     point{X: inner, Y: inner},
				Position: point{X: offset.X + (auxSize.X-inner)/2, Y: offset.Y + (auxSize.Y-inner)/2},
				Fillet:   inner / 2,
				Filled:   true,
				Color:    item.TextColor,
			})
		}

		textSize := (item.FontSize * uiScale) + 2
		face := &text.GoTextFace{
			Source: mplusFaceSource,
			Size:   float64(textSize),
		}
		loo := text.LayoutOptions{
			LineSpacing:    1.2,
			PrimaryAlign:   text.AlignStart,
			SecondaryAlign: text.AlignCenter,
		}
		tdop := ebiten.DrawImageOptions{}
		tdop.GeoM.Translate(
			float64(offset.X+auxSize.X+item.AuxSpace),
			float64(offset.Y+(auxSize.Y/2)),
		)
		top := &text.DrawOptions{DrawImageOptions: tdop, LayoutOptions: loo}
		top.ColorScale.ScaleWithColor(item.TextColor)
		text.Draw(subImg, item.Text, face, top)

	} else if item.ItemType == ITEM_BUTTON {

		if item.Image != nil {
			sop := &ebiten.DrawImageOptions{}
			sop.GeoM.Scale(float64(maxSize.X)/float64(item.Image.Bounds().Dx()),
				float64(maxSize.Y)/float64(item.Image.Bounds().Dy()))
			sop.GeoM.Translate(float64(offset.X), float64(offset.Y))
			subImg.DrawImage(item.Image, sop)
		} else {
			itemColor := item.Color
			if time.Since(item.Clicked) < clickFlash {
				itemColor = item.ClickColor
			} else if item.Hovered {
				item.Hovered = false
				itemColor = item.HoverColor
			}
			drawRoundRect(subImg, &roundRect{
				Size:     maxSize,
				Position: offset, Fillet: item.Fillet, Filled: true, Color: itemColor})
		}

		textSize := (item.FontSize * uiScale) + 2
		face := &text.GoTextFace{
			Source: mplusFaceSource,
			Size:   float64(textSize),
		}
		loo := text.LayoutOptions{
			LineSpacing:    0,
			PrimaryAlign:   text.AlignCenter,
			SecondaryAlign: text.AlignCenter,
		}
		tdop := ebiten.DrawImageOptions{}
		tdop.GeoM.Translate(
			float64(offset.X+((maxSize.X)/2)),
			float64(offset.Y+((maxSize.Y)/2)))
		top := &text.DrawOptions{DrawImageOptions: tdop, LayoutOptions: loo}
		top.ColorScale.ScaleWithColor(item.TextColor)
		text.Draw(subImg, item.Text, face, top)

		//Text
	} else if item.ItemType == ITEM_INPUT {

		itemColor := item.Color
		if item.Focused {
			itemColor = item.ClickColor
		} else if item.Hovered {
			item.Hovered = false
			itemColor = item.HoverColor
		}

		drawRoundRect(subImg, &roundRect{
			Size:     maxSize,
			Position: offset,
			Fillet:   item.Fillet,
			Filled:   true,
			Color:    itemColor,
		})

		textSize := (item.FontSize * uiScale) + 2
		face := &text.GoTextFace{
			Source: mplusFaceSource,
			Size:   float64(textSize),
		}
		loo := text.LayoutOptions{
			LineSpacing:    0,
			PrimaryAlign:   text.AlignStart,
			SecondaryAlign: text.AlignCenter,
		}
		tdop := ebiten.DrawImageOptions{}
		tdop.GeoM.Translate(
			float64(offset.X+item.BorderPad),
			float64(offset.Y+((maxSize.Y)/2)),
		)
		top := &text.DrawOptions{DrawImageOptions: tdop, LayoutOptions: loo}
		top.ColorScale.ScaleWithColor(item.TextColor)
		text.Draw(subImg, item.Text, face, top)

		if item.Focused {
			width, _ := text.Measure(item.Text, face, 0)
			cx := offset.X + item.BorderPad + float32(width)
			vector.StrokeLine(subImg,
				cx, offset.Y+2,
				cx, offset.Y+maxSize.Y-2,
				1, item.TextColor, false)
		}

	} else if item.ItemType == ITEM_SLIDER {

		itemColor := item.Color
		if item.Hovered {
			item.Hovered = false
			itemColor = item.HoverColor
		}

		trackY := offset.Y + maxSize.Y/2
		vector.StrokeLine(subImg,
			offset.X,
			trackY,
			offset.X+maxSize.X-item.AuxSize.X-item.AuxSpace*2,
			trackY,
			2*uiScale, itemColor, true)

		ratio := 0.0
		if item.MaxValue > item.MinValue {
			ratio = float64((item.Value - item.MinValue) / (item.MaxValue - item.MinValue))
		}
		if ratio < 0 {
			ratio = 0
		} else if ratio > 1 {
			ratio = 1
		}
		knobX := offset.X + float32(ratio)*(maxSize.X-item.AuxSize.X-item.AuxSpace*2)
		drawRoundRect(subImg, &roundRect{
			Size:     pointScaleMul(item.AuxSize),
			Position: point{X: knobX, Y: offset.Y + (maxSize.Y-item.AuxSize.Y)/2},
			Fillet:   item.Fillet,
			Filled:   true,
			Color:    item.ClickColor,
		})

		// value text
		valueText := fmt.Sprintf("%.2f", item.Value)
		if item.IntOnly {
			valueText = fmt.Sprintf("%d", int(item.Value))
		}

		textSize := (item.FontSize * uiScale) + 2
		face := &text.GoTextFace{
			Source: mplusFaceSource,
			Size:   float64(textSize),
		}
		loo := text.LayoutOptions{
			LineSpacing:    1.2,
			PrimaryAlign:   text.AlignStart,
			SecondaryAlign: text.AlignCenter,
		}
		tdop := ebiten.DrawImageOptions{}
		tw, _ := text.Measure(valueText, face, 0)
		tdop.GeoM.Translate(
			float64(offset.X+maxSize.X-item.AuxSpace-float32(tw)),
			float64(offset.Y+(maxSize.Y/2)),
		)
		top := &text.DrawOptions{DrawImageOptions: tdop, LayoutOptions: loo}
		top.ColorScale.ScaleWithColor(item.TextColor)
		text.Draw(subImg, valueText, face, top)

	} else if item.ItemType == ITEM_DROPDOWN {

		itemColor := item.Color
		if item.Open {
			itemColor = item.ClickColor
		} else if item.Hovered {
			item.Hovered = false
			itemColor = item.HoverColor
		}

		drawRoundRect(subImg, &roundRect{
			Size:     maxSize,
			Position: offset,
			Fillet:   item.Fillet,
			Filled:   true,
			Color:    itemColor,
		})

		textSize := (item.FontSize * uiScale) + 2
		face := &text.GoTextFace{Source: mplusFaceSource, Size: float64(textSize)}
		loo := text.LayoutOptions{PrimaryAlign: text.AlignStart, SecondaryAlign: text.AlignCenter}
		tdop := ebiten.DrawImageOptions{}
		tdop.GeoM.Translate(float64(offset.X+item.BorderPad), float64(offset.Y+maxSize.Y/2))
		top := &text.DrawOptions{DrawImageOptions: tdop, LayoutOptions: loo}
		top.ColorScale.ScaleWithColor(item.TextColor)
		label := item.Text
		if item.Selected >= 0 && item.Selected < len(item.Options) {
			label = item.Options[item.Selected]
		}
		text.Draw(subImg, label, face, top)

		arrow := maxSize.Y * 0.4
		drawTriangle(subImg,
			point{X: offset.X + maxSize.X - arrow - item.BorderPad,
				Y: offset.Y + (maxSize.Y-arrow)/2},
			arrow,
			item.TextColor)

		if item.Open {
			pendingDropdowns = append(pendingDropdowns, dropdownRender{item: item, offset: offset, clip: clip})
		}

	} else if item.ItemType == ITEM_TEXT {

		textSize := (item.FontSize * uiScale) + 2
		face := &text.GoTextFace{
			Source: mplusFaceSource,
			Size:   float64(textSize),
		}
		loo := text.LayoutOptions{
			LineSpacing:    float64(textSize) * 1.2,
			PrimaryAlign:   text.AlignStart,
			SecondaryAlign: text.AlignStart,
		}
		tdop := ebiten.DrawImageOptions{}
		tdop.GeoM.Translate(
			float64(offset.X),
			float64(offset.Y))

		top := &text.DrawOptions{DrawImageOptions: tdop, LayoutOptions: loo}

		top.ColorScale.ScaleWithColor(item.TextColor)
		text.Draw(subImg, item.Text, face, top)
	}

	if item.Outlined && item.Border > 0 && item.ItemType != ITEM_CHECKBOX && item.ItemType != ITEM_RADIO {
		drawRoundRect(subImg, &roundRect{
			Size:     maxSize,
			Position: offset,
			Fillet:   item.Fillet,
			Filled:   false,
			Color:    item.Color,
			Border:   item.Border * uiScale,
		})
	}

	if *debugMode {
		vector.StrokeRect(subImg,
			item.DrawRect.X0,
			item.DrawRect.Y0,
			item.DrawRect.X1-item.DrawRect.X0,
			item.DrawRect.Y1-item.DrawRect.Y0,
			1, color.RGBA{R: 128}, false)
	}

}

func drawDropdownOptions(item *itemData, offset point, clip rect, screen *ebiten.Image) {
	maxSize := item.GetSize()
	optionH := maxSize.Y
	visible := item.MaxVisible
	if visible <= 0 {
		visible = 5
	}
	startY := offset.Y + maxSize.Y
	first := int(item.Scroll.Y / optionH)
	offY := startY - (item.Scroll.Y - float32(first)*optionH)
	textSize := (item.FontSize * uiScale) + 2
	face := &text.GoTextFace{Source: mplusFaceSource, Size: float64(textSize)}
	loo := text.LayoutOptions{PrimaryAlign: text.AlignStart, SecondaryAlign: text.AlignCenter}
	drawRect := rect{X0: offset.X, Y0: startY, X1: offset.X + maxSize.X, Y1: startY + optionH*float32(visible)}
	visibleRect := intersectRect(drawRect, clip)
	if visibleRect.X1 <= visibleRect.X0 || visibleRect.Y1 <= visibleRect.Y0 {
		return
	}
	subImg := screen.SubImage(visibleRect.getRectangle()).(*ebiten.Image)
	for i := first; i < first+visible && i < len(item.Options); i++ {
		y := offY + float32(i-first)*optionH
		col := item.Color
		if i == item.Selected {
			col = item.ClickColor
		} else if i == item.HoverIndex {
			col = item.HoverColor
		}
		drawRoundRect(subImg, &roundRect{Size: maxSize, Position: point{X: offset.X, Y: y}, Fillet: item.Fillet, Filled: true, Color: col})
		td := ebiten.DrawImageOptions{}
		td.GeoM.Translate(float64(offset.X+item.BorderPad), float64(y+optionH/2))
		text.Draw(subImg, item.Options[i], face, &text.DrawOptions{DrawImageOptions: td, LayoutOptions: loo})
	}
}

func (win *windowData) drawDebug(screen *ebiten.Image) {
	if *debugMode {
		grab := win.getMainRect()
		vector.StrokeRect(screen, grab.X0, grab.Y0, grab.X1-grab.X0, grab.Y1-grab.Y0, 1, color.RGBA{R: 255, G: 255, A: 255}, false)

		grab = win.dragbarRect()
		vector.StrokeRect(screen, grab.X0, grab.Y0, grab.X1-grab.X0, grab.Y1-grab.Y0, 1, color.RGBA{R: 255, A: 255}, false)

		grab = win.xRect()
		vector.StrokeRect(screen, grab.X0, grab.Y0, grab.X1-grab.X0, grab.Y1-grab.Y0, 1, color.RGBA{G: 255, A: 255}, false)

		grab = win.getTitleRect()
		vector.StrokeRect(screen, grab.X0, grab.Y0, grab.X1-grab.X0, grab.Y1-grab.Y0, 1, color.RGBA{B: 255, G: 255, A: 255}, false)
	}
}

func drawRoundRect(screen *ebiten.Image, rrect *roundRect) {
	var (
		path     vector.Path
		vertices []ebiten.Vertex
		indices  []uint16
	)

	//Top left, after corner, clockwise
	path.MoveTo(rrect.Position.X+rrect.Fillet, rrect.Position.Y)
	path.LineTo(rrect.Position.X+rrect.Size.X-rrect.Fillet, rrect.Position.Y)
	path.QuadTo(
		rrect.Position.X+rrect.Size.X,
		rrect.Position.Y,
		rrect.Position.X+rrect.Size.X,
		rrect.Position.Y+rrect.Fillet)
	path.LineTo(rrect.Position.X+rrect.Size.X, rrect.Position.Y+rrect.Size.Y-rrect.Fillet)
	path.QuadTo(
		rrect.Position.X+rrect.Size.X,
		rrect.Position.Y+rrect.Size.Y,
		rrect.Position.X+rrect.Size.X-rrect.Fillet,
		rrect.Position.Y+rrect.Size.Y)
	path.LineTo(rrect.Position.X+rrect.Fillet, rrect.Position.Y+rrect.Size.Y)
	path.QuadTo(
		rrect.Position.X,
		rrect.Position.Y+rrect.Size.Y,
		rrect.Position.X,
		rrect.Position.Y+rrect.Size.Y-rrect.Fillet)
	path.LineTo(rrect.Position.X, rrect.Position.Y+rrect.Fillet)
	path.QuadTo(
		rrect.Position.X,
		rrect.Position.Y,
		rrect.Position.X+rrect.Fillet,
		rrect.Position.Y)
	path.Close()

	if rrect.Filled {
		vertices, indices = path.AppendVerticesAndIndicesForFilling(vertices[:0], indices[:0])
	} else {
		opv := &vector.StrokeOptions{}
		opv.Width = rrect.Border
		vertices, indices = path.AppendVerticesAndIndicesForStroke(vertices[:0], indices[:0], opv)
	}

	for i := range vertices {
		vertices[i].DstX = (vertices[i].DstX + 0.5)
		vertices[i].DstY = (vertices[i].DstY + 0.5)
		vertices[i].SrcX = 1
		vertices[i].SrcY = 1
		vertices[i].ColorR = float32(rrect.Color.R) / 255
		vertices[i].ColorG = float32(rrect.Color.G) / 255
		vertices[i].ColorB = float32(rrect.Color.B) / 255
		vertices[i].ColorA = float32(rrect.Color.A) / 255
	}

	op := &ebiten.DrawTrianglesOptions{}
	op.FillRule = ebiten.FillRuleNonZero
	op.AntiAlias = true
	screen.DrawTriangles(vertices, indices, whiteSubImage, op)
}

func drawTabShape(screen *ebiten.Image, pos point, size point, col color.RGBA) {
	var (
		path     vector.Path
		vertices []ebiten.Vertex
		indices  []uint16
	)

	slope := size.Y / 3
	fillet := size.Y / 5

	path.MoveTo(pos.X, pos.Y+size.Y)
	path.LineTo(pos.X+slope, pos.Y+fillet)
	path.QuadTo(pos.X+slope, pos.Y, pos.X+slope+fillet, pos.Y)
	path.LineTo(pos.X+size.X-slope-fillet, pos.Y)
	path.QuadTo(pos.X+size.X-slope, pos.Y, pos.X+size.X-slope, pos.Y+fillet)
	path.LineTo(pos.X+size.X, pos.Y+size.Y)
	path.Close()

	vertices, indices = path.AppendVerticesAndIndicesForFilling(vertices[:0], indices[:0])
	for i := range vertices {
		vertices[i].DstX = vertices[i].DstX + 0.5
		vertices[i].DstY = vertices[i].DstY + 0.5
		vertices[i].SrcX = 1
		vertices[i].SrcY = 1
		vertices[i].ColorR = float32(col.R) / 255
		vertices[i].ColorG = float32(col.G) / 255
		vertices[i].ColorB = float32(col.B) / 255
		vertices[i].ColorA = float32(col.A) / 255
	}

	op := &ebiten.DrawTrianglesOptions{}
	op.FillRule = ebiten.FillRuleNonZero
	op.AntiAlias = true
	screen.DrawTriangles(vertices, indices, whiteSubImage, op)
}

func drawTriangle(screen *ebiten.Image, pos point, size float32, col color.RGBA) {
	var (
		path     vector.Path
		vertices []ebiten.Vertex
		indices  []uint16
	)

	path.MoveTo(pos.X, pos.Y)
	path.LineTo(pos.X+size, pos.Y)
	path.LineTo(pos.X+size/2, pos.Y+size)
	path.Close()

	vertices, indices = path.AppendVerticesAndIndicesForFilling(vertices[:0], indices[:0])
	for i := range vertices {
		vertices[i].DstX += 0.5
		vertices[i].DstY += 0.5
		vertices[i].SrcX = 1
		vertices[i].SrcY = 1
		vertices[i].ColorR = float32(col.R) / 255
		vertices[i].ColorG = float32(col.G) / 255
		vertices[i].ColorB = float32(col.B) / 255
		vertices[i].ColorA = float32(col.A) / 255
	}

	op := &ebiten.DrawTrianglesOptions{FillRule: ebiten.FillRuleNonZero, AntiAlias: true}
	screen.DrawTriangles(vertices, indices, whiteSubImage, op)
}

func drawFPS(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, 0, 0, 58, 16, color.RGBA{R: 0, G: 0, B: 0, A: 192}, false)
	buf := fmt.Sprintf("%4v FPS", int(math.Round(ebiten.ActualFPS())))
	ebitenutil.DebugPrintAt(screen, buf, 0, 0)
}
