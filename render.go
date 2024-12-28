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

func (g *Game) Draw(screen *ebiten.Image) {

	for _, win := range windows {
		if !win.Open {
			continue
		}

		win.Draw(screen)
	}

	drawFPS(screen)
}

func (win *windowData) Draw(screen *ebiten.Image) {
	mainArea := screen.SubImage(win.getWinRect().getRectangle()).(*ebiten.Image)
	titleArea := screen.SubImage(win.getTitleRect().getRectangle()).(*ebiten.Image)
	windowArea := screen.SubImage(win.getWinRect().getRectangle()).(*ebiten.Image)

	win.drawBG(mainArea)
	win.drawContents(mainArea)
	win.drawWinTitle(titleArea)
	win.drawResizeTab(mainArea)
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

		textSize := ((win.TitleHeight * uiScale) / 1.5)
		face := &text.GoTextFace{
			Source: mplusFaceSource,
			Size:   float64(textSize),
		}

		skipTitleText := false
		textWidth, textHeight := text.Measure(win.Title, face, 0)
		if textWidth > float64(win.Size.X) ||
			textHeight > float64(win.TitleHeight*uiScale) {
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
			tdop.GeoM.Translate(float64(win.getPosition().X+((win.TitleHeight*uiScale)/4)),
				float64(win.getPosition().Y+((win.TitleHeight*uiScale)/2)))

			top := &text.DrawOptions{DrawImageOptions: tdop, LayoutOptions: loo}

			top.ColorScale.ScaleWithColor(win.TitleColor)
			buf := strings.ReplaceAll(win.Title, "\n", "") //Remove newline
			buf = strings.ReplaceAll(win.Title, "\r", "")  //Remove return
			text.Draw(screen, buf, face, top)
		} else {
			textWidth = 0
		}

		//Close X
		var buttonsWidth float32 = 0
		if win.Closable {
			var xpad float32 = (win.TitleHeight * uiScale) / 4.0
			color := win.TitleColor
			xThick := 3 * uiScale
			if win.HoverClose {
				color = win.HoverTitleColor
				win.HoverClose = false
			}
			vector.StrokeLine(screen,
				win.getPosition().X+win.Size.X-(win.TitleHeight*uiScale)+xpad,
				win.getPosition().Y+xpad,

				win.getPosition().X+win.Size.X-xpad,
				win.getPosition().Y+(win.TitleHeight*uiScale)-xpad,
				xThick, color, true)
			vector.StrokeLine(screen,
				win.getPosition().X+win.Size.X-xpad,
				win.getPosition().Y+xpad,

				win.getPosition().X+win.Size.X-(win.TitleHeight*uiScale)+xpad,
				win.getPosition().Y+(win.TitleHeight*uiScale)-xpad,
				xThick, color, true)

			buttonsWidth += (win.TitleHeight * uiScale)
		}

		//Dragbar
		if win.Movable {
			var xThick float32 = 1
			xColor := win.DragbarColor
			if win.HoverDragbar {
				xColor = win.HoverTitleColor
				win.HoverDragbar = false
			}
			dpad := (win.TitleHeight * uiScale) / 5
			for x := textWidth + float64((win.TitleHeight*uiScale)/1.5); x < float64(win.Size.X-buttonsWidth); x = x + float64(uiScale*5.0) {
				vector.StrokeLine(screen,
					win.getPosition().X+float32(x), win.getPosition().Y+dpad,
					win.getPosition().X+float32(x), win.getPosition().Y+(win.TitleHeight*uiScale)-dpad,
					xThick, xColor, false)
			}
		}
	}
}

func (win *windowData) drawBorder(screen *ebiten.Image) {
	//Draw borders
	if win.Border > 0 {
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
				win.Size.X-1, (win.TitleHeight*uiScale)-1,
				win.Border, FrameColor, false)
		}
		//Window border
		vector.StrokeRect(screen,
			win.getPosition().X+1, win.getPosition().Y+1,
			win.Size.X-1, win.Size.Y-(win.TitleHeight*uiScale)-1,
			win.Border, FrameColor, false)
	}
}

func (win *windowData) drawResizeTab(screen *ebiten.Image) {
	//Resize tab
	if win.Resizable {
		var xThick float32 = 1.0
		xColor := win.SizeTabColor

		if activeWindow == win {
			xColor = win.ActiveColor
		}
		var Outer, Middle, Inner float32 = 14 * uiScale, 10 * uiScale, 6 * uiScale

		//Outer
		vector.StrokeLine(screen,
			win.getPosition().X+win.Size.X-1,
			win.getPosition().Y+win.Size.Y-Outer-(win.TitleHeight*uiScale),

			win.getPosition().X+win.Size.X-Outer,
			win.getPosition().Y+win.Size.Y-(win.TitleHeight*uiScale)-1,
			xThick, xColor, true)
		//Middle
		vector.StrokeLine(screen,
			win.getPosition().X+win.Size.X-1,
			win.getPosition().Y+win.Size.Y-Middle-(win.TitleHeight*uiScale),

			win.getPosition().X+win.Size.X-Middle,
			win.getPosition().Y+win.Size.Y-(win.TitleHeight*uiScale)-1,
			xThick, xColor, true)
		//Inner
		vector.StrokeLine(screen,
			win.getPosition().X+win.Size.X-1,
			win.getPosition().Y+win.Size.Y-Inner-(win.TitleHeight*uiScale),

			win.getPosition().X+win.Size.X-Inner,
			win.getPosition().Y+win.Size.Y-(win.TitleHeight*uiScale)-1,
			xThick, xColor, true)
	}
}

func (win *windowData) drawContents(screen *ebiten.Image) {
	for _, item := range win.Contents {
		itemImage := screen.SubImage(item.getItemRect(win).getRectangle()).(*ebiten.Image)
		if item.ItemType == ITEM_FLOW {
			newWin := windowData{Size: win.Size, Position: pointSub(win.Position, point{X: 0, Y: win.TitleHeight}), Contents: item.Contents}
			newWin.drawContents(itemImage)

		} else if item.ItemType == ITEM_BUTTON {
			BGColor := item.Color
			BorderColor := item.HoverColor
			if time.Since(item.Clicked) < clickFlash {
				BGColor = item.ClickColor
				BorderColor = item.Color
			} else if item.Hovered {
				BGColor = item.HoverColor
				BorderColor = item.Color
				item.Hovered = false
			}

			if item.Fillet < 1 {
				vector.DrawFilledRect(itemImage,
					win.getPosition().X+(item.getPosition(win).X*uiScale),
					win.getPosition().Y+(item.getPosition(win).Y*uiScale),
					item.Size.X*uiScale, item.Size.Y*uiScale, BGColor, false)
			} else {
				win.drawRoundRect(itemImage, item, BGColor, BorderColor)
			}

			textSize := item.FontSize * uiScale
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
				float64(win.getPosition().X+(item.getPosition(win).X*uiScale)+((item.Size.X*uiScale)/2)),
				float64(win.getPosition().Y+(item.getPosition(win).Y*uiScale)+((item.Size.Y*uiScale)/2)))

			top := &text.DrawOptions{DrawImageOptions: tdop, LayoutOptions: loo}

			top.ColorScale.ScaleWithColor(item.TextColor)
			text.Draw(itemImage, item.Text, face, top)

			//Text
		} else if item.ItemType == ITEM_TEXT {

			textSize := item.FontSize * uiScale
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
				float64(win.getPosition().X+(item.getPosition(win).X*uiScale)),
				float64(win.getPosition().Y+(win.TitleHeight*uiScale)+(item.getPosition(win).Y*uiScale)))

			top := &text.DrawOptions{DrawImageOptions: tdop, LayoutOptions: loo}

			top.ColorScale.ScaleWithColor(item.TextColor)
			text.Draw(itemImage, item.Text, face, top)
		}
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

// Break this up, make a generic draw function as well
func (win *windowData) drawRoundRect(screen *ebiten.Image, item *itemData, BGcolor, BorderColor color.RGBA) {
	drawRoundRect(screen, &roundRect{
		Size:     pointScaleMul(item.Size),
		Position: pointAdd(win.getPosition(), pointScaleMul(item.getPosition(win))),
		Fillet:   item.Fillet * uiScale,
		Color:    BGcolor,
		Filled:   true,
		Border:   item.Border * uiScale,
	})

	offset := point{X: item.BorderPad * uiScale, Y: item.BorderPad * uiScale}
	drawRoundRect(screen, &roundRect{
		Size: pointSub(
			pointScaleMul(item.Size),
			offset,
		),
		Position: pointAdd(
			pointAdd(win.getPosition(), pointScaleMul(item.getPosition(win))),
			pointDiv(offset, point{X: 2, Y: 2}),
		),
		Fillet: item.Fillet * uiScale,
		Color:  BorderColor,
		Filled: false,
		Border: item.Border * uiScale,
	})
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

func drawFPS(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, 0, 0, 58, 16, color.RGBA{R: 0, G: 0, B: 0, A: 192}, false)
	buf := fmt.Sprintf("%4v FPS", int(math.Round(ebiten.ActualFPS())))
	ebitenutil.DebugPrintAt(screen, buf, 0, 0)
}
