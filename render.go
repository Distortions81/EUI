package main

import (
	"fmt"
	"image/color"
	"math"

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
		win.DrawWinTitle(screen)
		win.DrawBorder(screen)
		win.DrawResizeTab(screen)
		win.DrawContents(screen)
		win.DrawDebug(screen)
	}

	DrawFPS(screen)
}

func (win *WindowData) DrawBG(screen *ebiten.Image) {
	//Window BG Color
	vector.DrawFilledRect(screen,
		win.Position.X, win.Position.Y,
		win.SizeTemp.X, win.SizeTemp.Y-(win.TitleSizeTemp),
		win.ContentsBGColor, false)
}

func (win *WindowData) DrawWinTitle(screen *ebiten.Image) {
	// Window Title
	if win.TitleSize > 0 {

		textSize := (win.TitleSizeTemp / 1.5)
		face := &text.GoTextFace{
			Source: mplusFaceSource,
			Size:   float64(textSize),
		}

		skipTitleText := false
		textWidth, textHeight := text.Measure(win.Title, face, 0)
		if textWidth > float64(win.SizeTemp.X) ||
			textHeight > float64(win.TitleSizeTemp) {
			skipTitleText = true
			//log.Print("Title text too big for title size.")
		}

		//Title text
		if !skipTitleText {
			loo := text.LayoutOptions{
				LineSpacing:    0,
				PrimaryAlign:   text.AlignStart,
				SecondaryAlign: text.AlignCenter,
			}
			tdop := ebiten.DrawImageOptions{}
			tdop.GeoM.Translate(float64(win.Position.X+(win.TitleSizeTemp/4)),
				float64(win.Position.Y+(win.TitleSizeTemp/2)))

			top := &text.DrawOptions{DrawImageOptions: tdop, LayoutOptions: loo}
			top.ColorScale.ScaleWithColor(win.TitleColor)
			text.Draw(screen, win.Title, face, top)
		} else {
			textWidth = 0
		}

		//Close X
		var buttonsWidth float32 = 0
		if win.Closable {
			var xpad float32 = win.TitleSizeTemp / 4.0
			xThick := 3 * UIScale
			if win.HoverX {
				xThick *= (1.5)
				win.HoverX = false
			}
			vector.StrokeLine(screen,
				win.Position.X+win.SizeTemp.X-win.TitleSizeTemp+xpad,
				win.Position.Y+xpad,

				win.Position.X+win.SizeTemp.X-xpad,
				win.Position.Y+win.TitleSizeTemp-xpad,
				xThick, win.TitleColor, true)
			vector.StrokeLine(screen,
				win.Position.X+win.SizeTemp.X-xpad,
				win.Position.Y+xpad,

				win.Position.X+win.SizeTemp.X-win.TitleSizeTemp+xpad,
				win.Position.Y+win.TitleSizeTemp-xpad,
				xThick, win.TitleColor, true)

			buttonsWidth += win.TitleSizeTemp
		}

		//Dragbar
		if win.Movable {
			var xThick float32 = 1
			xColor := win.DragColor
			if win.HoverDragbar {
				xColor = win.DragHoverColor
				win.HoverDragbar = false
			}
			dpad := win.TitleSizeTemp / 5
			for x := textWidth + float64(win.TitleSizeTemp/1.5); x < float64(win.SizeTemp.X-buttonsWidth); x = x + float64(UIScale*5.0) {
				vector.StrokeLine(screen,
					win.Position.X+float32(x), win.Position.Y+dpad,
					win.Position.X+float32(x), win.Position.Y+win.TitleSizeTemp-dpad,
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
		if win.TitleSize > 0 {
			vector.StrokeRect(screen,
				win.Position.X, win.Position.Y,
				win.SizeTemp.X, win.TitleSizeTemp,
				win.Border, FrameColor, false)
		}
		//Window border
		vector.StrokeRect(screen,
			win.Position.X, win.Position.Y,
			win.SizeTemp.X, win.SizeTemp.Y-win.TitleSizeTemp,
			win.Border, FrameColor, false)
	}
}

func (win *WindowData) DrawResizeTab(screen *ebiten.Image) {
	//Resize tab
	if win.Resizable {
		var xThick float32 = 1.0
		xColor := win.SizeColor

		if win.HoverResizeTab {
			xColor = win.SizeHoverColor
			xThick = 2
			win.HoverResizeTab = false
		}
		var Outer, Middle, Inner float32 = 14 * UIScale, 10 * UIScale, 6 * UIScale

		//Outer
		vector.StrokeLine(screen,
			win.Position.X+win.SizeTemp.X-1,
			win.Position.Y+win.SizeTemp.Y-Outer-win.TitleSizeTemp,

			win.Position.X+win.SizeTemp.X-Outer,
			win.Position.Y+win.SizeTemp.Y-win.TitleSizeTemp-1,
			xThick, xColor, true)
		//Middle
		vector.StrokeLine(screen,
			win.Position.X+win.SizeTemp.X-1,
			win.Position.Y+win.SizeTemp.Y-Middle-win.TitleSizeTemp,

			win.Position.X+win.SizeTemp.X-Middle,
			win.Position.Y+win.SizeTemp.Y-win.TitleSizeTemp-1,
			xThick, xColor, true)
		//Inner
		vector.StrokeLine(screen,
			win.Position.X+win.SizeTemp.X-1,
			win.Position.Y+win.SizeTemp.Y-Inner-win.TitleSizeTemp,

			win.Position.X+win.SizeTemp.X-Inner,
			win.Position.Y+win.SizeTemp.Y-win.TitleSizeTemp-1,
			xThick, xColor, true)
	}
}

func (win *WindowData) DrawContents(screen *ebiten.Image) {
	for _, item := range win.Contents {

		if item.ItemType == ITEM_BUTTON {
			itemColor := item.Color
			if item.Clicked {
				itemColor = item.ClickColor
				item.Clicked = false
			} else if item.Hovered {
				itemColor = item.HoverColor
				item.Hovered = false
			}

			vector.DrawFilledRect(screen,
				win.Position.X+item.Position.X, win.Position.Y+item.Position.Y+win.TitleSize,
				item.Size.X*UIScale, item.Size.Y*UIScale, itemColor, false)

			textSize := item.FontSize
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
				float64(win.Position.X+item.Position.X+(item.Size.X/2)),
				float64(win.Position.Y+win.TitleSize+item.Position.Y)+(float64(item.Size.Y/2)))

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

func DrawFPS(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, 0, 0, 58, 16, color.RGBA{R: 0, G: 0, B: 0, A: 192}, false)
	buf := fmt.Sprintf("%4v FPS", int(math.Round(ebiten.ActualFPS())))
	ebitenutil.DebugPrintAt(screen, buf, 0, 0)
}
