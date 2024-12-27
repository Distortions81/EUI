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

	drawFPS(screen)
}

func (win *WindowData) drawBG(screen *ebiten.Image) {
	windowArea := screen.SubImage(win.getWinRect().getRectangle()).(*ebiten.Image)
	windowArea.Fill(win.BGColor)
}

func (win *WindowData) drawWinTitle(screen *ebiten.Image) {
	// Window Title
	if win.TitleHeight > 0 {
		screen.Fill(win.TitleBGColor)

		textSize := (win.TitleScreenHeight / 1.5)
		face := &text.GoTextFace{
			Source: mplusFaceSource,
			Size:   float64(textSize),
		}

		skipTitleText := false
		textWidth, textHeight := text.Measure(win.Title, face, 0)
		if textWidth > float64(win.ScreenSize.X) ||
			textHeight > float64(win.TitleScreenHeight) {
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
			tdop.GeoM.Translate(float64(win.Position.X+(win.TitleScreenHeight/4)),
				float64(win.Position.Y+(win.TitleScreenHeight/2)))

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
			var xpad float32 = win.TitleScreenHeight / 4.0
			color := win.TitleColor
			xThick := 3 * uiScale
			if win.HoverClose {
				color = win.HoverTitleColor
				win.HoverClose = false
			}
			vector.StrokeLine(screen,
				win.Position.X+win.ScreenSize.X-win.TitleScreenHeight+xpad,
				win.Position.Y+xpad,

				win.Position.X+win.ScreenSize.X-xpad,
				win.Position.Y+win.TitleScreenHeight-xpad,
				xThick, color, true)
			vector.StrokeLine(screen,
				win.Position.X+win.ScreenSize.X-xpad,
				win.Position.Y+xpad,

				win.Position.X+win.ScreenSize.X-win.TitleScreenHeight+xpad,
				win.Position.Y+win.TitleScreenHeight-xpad,
				xThick, color, true)

			buttonsWidth += win.TitleScreenHeight
		}

		//Dragbar
		if win.Movable {
			var xThick float32 = 1
			xColor := win.DragbarColor
			if win.HoverDragbar {
				xColor = win.HoverTitleColor
				win.HoverDragbar = false
			}
			dpad := win.TitleScreenHeight / 5
			for x := textWidth + float64(win.TitleScreenHeight/1.5); x < float64(win.ScreenSize.X-buttonsWidth); x = x + float64(uiScale*5.0) {
				vector.StrokeLine(screen,
					win.Position.X+float32(x), win.Position.Y+dpad,
					win.Position.X+float32(x), win.Position.Y+win.TitleScreenHeight-dpad,
					xThick, xColor, false)
			}
		}
	}
}

func (win *WindowData) drawBorder(screen *ebiten.Image) {
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
				win.Position.X+1, win.Position.Y+1,
				win.ScreenSize.X, win.TitleScreenHeight,
				win.Border, FrameColor, false)
		}
		//Window border
		vector.StrokeRect(screen,
			win.Position.X+1, win.Position.Y+1,
			win.ScreenSize.X-1, win.ScreenSize.Y-win.TitleScreenHeight-1,
			win.Border, FrameColor, false)
	}
}

func (win *WindowData) drawResizeTab(screen *ebiten.Image) {
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
			win.Position.X+win.ScreenSize.X-1,
			win.Position.Y+win.ScreenSize.Y-Outer-win.TitleScreenHeight,

			win.Position.X+win.ScreenSize.X-Outer,
			win.Position.Y+win.ScreenSize.Y-win.TitleScreenHeight-1,
			xThick, xColor, true)
		//Middle
		vector.StrokeLine(screen,
			win.Position.X+win.ScreenSize.X-1,
			win.Position.Y+win.ScreenSize.Y-Middle-win.TitleScreenHeight,

			win.Position.X+win.ScreenSize.X-Middle,
			win.Position.Y+win.ScreenSize.Y-win.TitleScreenHeight-1,
			xThick, xColor, true)
		//Inner
		vector.StrokeLine(screen,
			win.Position.X+win.ScreenSize.X-1,
			win.Position.Y+win.ScreenSize.Y-Inner-win.TitleScreenHeight,

			win.Position.X+win.ScreenSize.X-Inner,
			win.Position.Y+win.ScreenSize.Y-win.TitleScreenHeight-1,
			xThick, xColor, true)
	}
}

func (win *WindowData) drawContents(screen *ebiten.Image) {
	for _, item := range win.Contents {
		itemImage := screen.SubImage(item.getItemRect(win).getRectangle()).(*ebiten.Image)
		if item.ItemType == ITEM_BUTTON {

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
					win.Position.X+(item.Position.X*uiScale),
					win.Position.Y+(item.Position.Y*uiScale),
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
				float64(win.Position.X+(item.Position.X*uiScale)+((item.Size.X*uiScale)/2)),
				float64(win.Position.Y+(item.Position.Y*uiScale)+((item.Size.Y*uiScale)/2)))

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
				float64(win.Position.X+(item.Position.X*uiScale)),
				float64(win.Position.Y+(item.Position.Y*uiScale)))

			top := &text.DrawOptions{DrawImageOptions: tdop, LayoutOptions: loo}

			top.ColorScale.ScaleWithColor(item.TextColor)
			text.Draw(itemImage, item.Text, face, top)
		}
	}
}

func (win *WindowData) drawDebug(screen *ebiten.Image) {
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
func (win *WindowData) drawRoundRect(screen *ebiten.Image, item *ItemData, BGcolor, BorderColor color.RGBA) {
	drawRoundRect(screen, &RoundRect{
		Size:     pointScaleMul(item.Size),
		Position: pointAdd(win.Position, pointScaleMul(item.Position)),
		Fillet:   item.Fillet * uiScale,
		Color:    BGcolor,
		Filled:   true,
		Border:   item.Border * uiScale,
	})

	offset := Point{X: item.BorderPad * uiScale, Y: item.BorderPad * uiScale}
	drawRoundRect(screen, &RoundRect{
		Size: pointSub(
			pointScaleMul(item.Size),
			offset,
		),
		Position: pointAdd(
			pointAdd(win.Position, pointScaleMul(item.Position)),
			pointDiv(offset, Point{X: 2, Y: 2}),
		),
		Fillet: item.Fillet * uiScale,
		Color:  BorderColor,
		Filled: false,
		Border: item.Border * uiScale,
	})
}

func drawRoundRect(screen *ebiten.Image, item *RoundRect) {
	var (
		path     vector.Path
		vertices []ebiten.Vertex
		indices  []uint16
	)

	//Top left, after corner, clockwise
	path.MoveTo(item.Position.X+item.Fillet, item.Position.Y)
	path.LineTo(item.Position.X+item.Size.X-item.Fillet, item.Position.Y)
	path.QuadTo(
		item.Position.X+item.Size.X,
		item.Position.Y,
		item.Position.X+item.Size.X,
		item.Position.Y+item.Fillet)
	path.LineTo(item.Position.X+item.Size.X, item.Position.Y+item.Size.Y-item.Fillet)
	path.QuadTo(
		item.Position.X+item.Size.X,
		item.Position.Y+item.Size.Y,
		item.Position.X+item.Size.X-item.Fillet,
		item.Position.Y+item.Size.Y)
	path.LineTo(item.Position.X+item.Fillet, item.Position.Y+item.Size.Y)
	path.QuadTo(
		item.Position.X,
		item.Position.Y+item.Size.Y,
		item.Position.X,
		item.Position.Y+item.Size.Y-item.Fillet)
	path.LineTo(item.Position.X, item.Position.Y+item.Fillet)
	path.QuadTo(
		item.Position.X,
		item.Position.Y,
		item.Position.X+item.Fillet,
		item.Position.Y)
	path.Close()

	if item.Filled {
		vertices, indices = path.AppendVerticesAndIndicesForFilling(vertices[:0], indices[:0])
	} else {
		opv := &vector.StrokeOptions{}
		opv.Width = item.Border
		vertices, indices = path.AppendVerticesAndIndicesForStroke(vertices[:0], indices[:0], opv)
	}

	for i := range vertices {
		vertices[i].DstX = (vertices[i].DstX + 0.5)
		vertices[i].DstY = (vertices[i].DstY + 0.5)
		vertices[i].SrcX = 1
		vertices[i].SrcY = 1
		vertices[i].ColorR = float32(item.Color.R) / 255
		vertices[i].ColorG = float32(item.Color.G) / 255
		vertices[i].ColorB = float32(item.Color.B) / 255
		vertices[i].ColorA = float32(item.Color.A) / 255
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
