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

	for _, win := range Windows {
		if !win.Open {
			continue
		}

		//Reduce UI scaling calculations
		win.CalcUIScale()
		win.DrawBG(screen)
		win.DrawContents(screen)
		win.DrawWinTitle(screen)
		win.DrawBorder(screen)
		win.DrawResizeTab(screen)
		win.DrawDebug(screen)
	}

	DrawFPS(screen)
}

func (win *WindowData) DrawBG(screen *ebiten.Image) {
	//Window BG Color
	vector.DrawFilledRect(screen,
		win.Position.X, win.Position.Y,
		win.ScreenSize.X, win.ScreenSize.Y-(win.TitleScreenHeight),
		win.BGColor, false)
}

func (win *WindowData) DrawWinTitle(screen *ebiten.Image) {
	// Window Title
	if win.TitleHeight > 0 {

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
			xThick := 3 * UIScale
			if win.HoverClose {
				xThick *= (1.5)
				win.HoverClose = false
			}
			vector.StrokeLine(screen,
				win.Position.X+win.ScreenSize.X-win.TitleScreenHeight+xpad,
				win.Position.Y+xpad,

				win.Position.X+win.ScreenSize.X-xpad,
				win.Position.Y+win.TitleScreenHeight-xpad,
				xThick, win.TitleColor, true)
			vector.StrokeLine(screen,
				win.Position.X+win.ScreenSize.X-xpad,
				win.Position.Y+xpad,

				win.Position.X+win.ScreenSize.X-win.TitleScreenHeight+xpad,
				win.Position.Y+win.TitleScreenHeight-xpad,
				xThick, win.TitleColor, true)

			buttonsWidth += win.TitleScreenHeight
		}

		//Dragbar
		if win.Movable {
			var xThick float32 = 1
			xColor := win.DragbarColor
			if win.HoverDragbar {
				xColor = win.DragbarHoverColor
				win.HoverDragbar = false
			}
			dpad := win.TitleScreenHeight / 5
			for x := textWidth + float64(win.TitleScreenHeight/1.5); x < float64(win.ScreenSize.X-buttonsWidth); x = x + float64(UIScale*5.0) {
				vector.StrokeLine(screen,
					win.Position.X+float32(x), win.Position.Y+dpad,
					win.Position.X+float32(x), win.Position.Y+win.TitleScreenHeight-dpad,
					xThick, xColor, false)
			}
		}
	}
}

func (win *WindowData) DrawBorder(screen *ebiten.Image) {
	//Draw borders
	if win.Border > 0 {
		FrameColor := win.BorderColor
		if win.Hovered {
			FrameColor = win.HoverColor
			win.Hovered = false
		}
		if win.TitleHeight > 0 {
			vector.StrokeRect(screen,
				win.Position.X, win.Position.Y,
				win.ScreenSize.X, win.TitleScreenHeight,
				win.Border, FrameColor, false)
		}
		//Window border
		vector.StrokeRect(screen,
			win.Position.X, win.Position.Y,
			win.ScreenSize.X, win.ScreenSize.Y-win.TitleScreenHeight,
			win.Border, FrameColor, false)
	}
}

func (win *WindowData) DrawResizeTab(screen *ebiten.Image) {
	//Resize tab
	if win.Resizable {
		var xThick float32 = 1.0
		xColor := win.SizeTabColor

		if win.HoverResizeTab {
			xColor = win.SizeTabHoverColor
			xThick = 2
			win.HoverResizeTab = false
		}
		var Outer, Middle, Inner float32 = 14 * UIScale, 10 * UIScale, 6 * UIScale

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

func (win *WindowData) DrawContents(screen *ebiten.Image) {
	for _, item := range win.Contents {

		if item.Position.X > win.Size.X ||
			item.Position.Y > win.Size.Y-win.TitleHeight {
			continue
		} else if item.Position.X+item.Size.X > win.Size.X ||
			item.Position.Y+item.Size.Y > win.Size.Y-win.TitleHeight {
			continue
		}

		if item.ItemType == ITEM_BUTTON {

			BGColor := item.Color
			BorderColor := item.HoverColor
			if time.Since(item.Clicked) < FlashTime {
				BGColor = item.ClickColor
				BorderColor = item.Color
			} else if item.Hovered {
				BGColor = item.HoverColor
				BorderColor = item.Color
				item.Hovered = false
			}

			if item.Fillet < 1 {
				vector.DrawFilledRect(screen,
					win.Position.X+(item.Position.X*UIScale),
					win.Position.Y+(item.Position.Y*UIScale),
					item.Size.X*UIScale, item.Size.Y*UIScale, BGColor, false)
			} else {
				win.DrawRoundRect(screen, item, BGColor, BorderColor)
			}

			textSize := item.FontSize * UIScale
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
				float64(win.Position.X+(item.Position.X*UIScale)+((item.Size.X*UIScale)/2)),
				float64(win.Position.Y+(item.Position.Y*UIScale)+((item.Size.Y*UIScale)/2)))

			top := &text.DrawOptions{DrawImageOptions: tdop, LayoutOptions: loo}

			top.ColorScale.ScaleWithColor(item.TextColor)
			text.Draw(screen, item.Text, face, top)

			//Text
		} else if item.ItemType == ITEM_TEXT {

			textSize := item.FontSize * UIScale
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
				float64(win.Position.X+(item.Position.X*UIScale)),
				float64(win.Position.Y+(item.Position.Y*UIScale)))

			top := &text.DrawOptions{DrawImageOptions: tdop, LayoutOptions: loo}

			top.ColorScale.ScaleWithColor(item.TextColor)
			text.Draw(screen, item.Text, face, top)
		}
	}
}

func (win *WindowData) DrawDebug(screen *ebiten.Image) {
	if *debugMode {
		grab := win.GetMainRect()
		vector.StrokeRect(screen, grab.X0, grab.Y0, grab.X1-grab.X0, grab.Y1-grab.Y0, 1, color.RGBA{R: 255, G: 255, A: 255}, false)

		grab = win.DragbarRect()
		vector.StrokeRect(screen, grab.X0, grab.Y0, grab.X1-grab.X0, grab.Y1-grab.Y0, 1, color.RGBA{R: 255, A: 255}, false)

		grab = win.XRect()
		vector.StrokeRect(screen, grab.X0, grab.Y0, grab.X1-grab.X0, grab.Y1-grab.Y0, 1, color.RGBA{G: 255, A: 255}, false)

		grab = win.ResizeTabRect()
		vector.StrokeRect(screen, grab.X0, grab.Y0, grab.X1-grab.X0, grab.Y1-grab.Y0, 1, color.RGBA{B: 255, A: 255}, false)

		grab = win.TitleRect()
		vector.StrokeRect(screen, grab.X0, grab.Y0, grab.X1-grab.X0, grab.Y1-grab.Y0, 1, color.RGBA{B: 255, G: 255, A: 255}, false)
	}
}

// Break this up, make a generic draw function as well
func (win *WindowData) DrawRoundRect(screen *ebiten.Image, item *ItemData, BGcolor, BorderColor color.RGBA) {
	DrawRoundRect(screen, &RoundRect{
		Size:     PointScaleMul(item.Size),
		Position: PointAdd(win.Position, PointScaleMul(item.Position)),
		Fillet:   item.Fillet * UIScale,
		Color:    BGcolor,
		Filled:   true,
		Border:   item.Border * UIScale,
	})

	offset := Point{X: item.BorderPad * UIScale, Y: item.BorderPad * UIScale}
	DrawRoundRect(screen, &RoundRect{
		Size: PointSub(
			PointScaleMul(item.Size),
			offset,
		),
		Position: PointAdd(
			PointAdd(win.Position, PointScaleMul(item.Position)),
			PointDiv(offset, Point{X: 2, Y: 2}),
		),
		Fillet: item.Fillet * UIScale,
		Color:  BorderColor,
		Filled: false,
		Border: item.Border * UIScale,
	})
}

func DrawRoundRect(screen *ebiten.Image, item *RoundRect) {
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

func DrawFPS(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, 0, 0, 58, 16, color.RGBA{R: 0, G: 0, B: 0, A: 192}, false)
	buf := fmt.Sprintf("%4v FPS", int(math.Round(ebiten.ActualFPS())))
	ebitenutil.DebugPrintAt(screen, buf, 0, 0)
}
