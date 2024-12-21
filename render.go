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

		//Window BG Color
		vector.DrawFilledRect(screen,
			win.Position.X, win.Position.Y,
			win.GetSizeX(), win.GetSizeY()-(win.GetTitleSize()),
			win.ContentsBGColor, false)

		//Window Title
		if win.TitleSize > 0 {

			textSize := (win.GetTitleSize() / 1.5)
			face := &text.GoTextFace{
				Source: mplusFaceSource,
				Size:   float64(textSize),
			}

			skipTitleText := false
			textWidth, textHeight := text.Measure(win.Title, face, 0)
			if textWidth > float64(win.GetSizeX()) ||
				textHeight > float64(win.GetTitleSize()) {
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
				tdop.GeoM.Translate(float64(win.Position.X+(win.GetTitleSize()/4)),
					float64(win.Position.Y+(win.GetTitleSize()/2)))

				top := &text.DrawOptions{DrawImageOptions: tdop, LayoutOptions: loo}
				top.ColorScale.ScaleWithColor(win.TitleColor)
				text.Draw(screen, win.Title, face, top)
			} else {
				textWidth = 0
			}

			//Close X
			var buttonsWidth float32 = 0
			if win.Closable {
				var xpad float32 = win.GetTitleSize() / 4.0
				xThick := 3 * UIScale
				if win.HoverX {
					xThick *= (1.5)
					win.HoverX = false
				}
				vector.StrokeLine(screen,
					win.Position.X+win.GetSizeX()-win.GetTitleSize()+xpad,
					win.Position.Y+xpad,

					win.Position.X+win.GetSizeX()-xpad,
					win.Position.Y+win.GetTitleSize()-xpad,
					xThick, win.TitleColor, true)
				vector.StrokeLine(screen,
					win.Position.X+win.GetSizeX()-xpad,
					win.Position.Y+xpad,

					win.Position.X+win.GetSizeX()-win.GetTitleSize()+xpad,
					win.Position.Y+win.GetTitleSize()-xpad,
					xThick, win.TitleColor, true)

				buttonsWidth += win.GetTitleSize()
			}

			//Dragbar
			if win.Movable {
				var xThick float32 = 1
				xColor := win.DragColor
				if win.HoverDragbar {
					xColor = win.DragHoverColor
					win.HoverDragbar = false
				}
				dpad := win.GetTitleSize() / 5
				for x := textWidth + float64(win.GetTitleSize()/1.5); x < float64(win.GetSizeX()-buttonsWidth); x = x + float64(UIScale*5.0) {
					vector.StrokeLine(screen,
						win.Position.X+float32(x), win.Position.Y+dpad,
						win.Position.X+float32(x), win.Position.Y+win.GetTitleSize()-dpad,
						xThick, xColor, false)
				}
			}
		}

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
					win.GetSizeX(), win.GetTitleSize(),
					win.Border, FrameColor, false)
			}
			//Window border
			vector.StrokeRect(screen,
				win.Position.X, win.Position.Y,
				win.GetSizeX(), win.GetSizeY()-win.GetTitleSize(),
				win.Border, FrameColor, false)
		}

		//Resize tab
		if win.Resizable {
			var xThick float32 = 1.0
			xColor := win.SizeColor

			if win.HoverResizeTab {
				xColor = win.SizeHoverColor
				xThick = 2
				win.HoverResizeTab = false
			}
			//Outer
			vector.StrokeLine(screen,
				win.Position.X+win.GetSizeX()-1,
				win.Position.Y+win.GetSizeY()-(14*UIScale)-win.GetTitleSize(),

				win.Position.X+win.GetSizeX()-(14*UIScale),
				win.Position.Y+win.GetSizeY()-win.GetTitleSize()-1,
				xThick, xColor, true)
			//Middle
			vector.StrokeLine(screen,
				win.Position.X+win.GetSizeX()-1,
				win.Position.Y+win.GetSizeY()-(10*UIScale)-win.GetTitleSize(),

				win.Position.X+win.GetSizeX()-(10*UIScale),
				win.Position.Y+win.GetSizeY()-win.GetTitleSize()-1,
				xThick, xColor, true)
			//Inner
			vector.StrokeLine(screen,
				win.Position.X+win.GetSizeX()-1,
				win.Position.Y+win.GetSizeY()-(6*UIScale)-win.GetTitleSize(),

				win.Position.X+win.GetSizeX()-(6*UIScale),
				win.Position.Y+win.GetSizeY()-win.GetTitleSize()-1,
				xThick, xColor, true)
		}

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

	vector.DrawFilledRect(screen, 0, 0, 58, 16, color.RGBA{R: 0, G: 0, B: 0, A: 192}, false)
	buf := fmt.Sprintf("%4v FPS", int(math.Round(ebiten.ActualFPS())))
	ebitenutil.DebugPrintAt(screen, buf, 0, 0)
}
