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

		//Draw BG Color
		vector.DrawFilledRect(screen,
			win.Position.X, win.Position.Y,
			win.Size.X*UIScale, (win.Size.Y*UIScale)-(win.TitleSize*UIScale),
			win.ContentsBGColor, false)

		//Draw Title
		if win.TitleSize > 0 {

			textSize := ((win.TitleSize * UIScale) / 1.5)
			face := &text.GoTextFace{
				Source: mplusFaceSource,
				Size:   float64(textSize),
			}

			skipTitleText := false
			textWidth, textHeight := text.Measure(win.Title, face, 0)
			if textWidth > float64(win.Size.X*UIScale) ||
				textHeight > float64(win.TitleSize*UIScale) {
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
				tdop.GeoM.Translate(float64(win.Position.X+((win.TitleSize*UIScale)/4)),
					float64(win.Position.Y+((win.TitleSize*UIScale)/2)))

				top := &text.DrawOptions{DrawImageOptions: tdop, LayoutOptions: loo}
				top.ColorScale.ScaleWithColor(win.TitleColor)
				text.Draw(screen, win.Title, face, top)
			} else {
				textWidth = 0
			}

			//Resize tab
			if win.Resizable {
				vector.StrokeLine(screen,
					win.Position.X+(win.Size.X*UIScale)-1,
					win.Position.Y+(win.Size.Y*UIScale)-(14*UIScale)-(win.TitleSize*UIScale),

					win.Position.X+(win.Size.X*UIScale)-(14*UIScale),
					win.Position.Y+(win.Size.Y*UIScale)-(win.TitleSize*UIScale)-1,
					1, win.SizeColor, true)
				vector.StrokeLine(screen,
					win.Position.X+(win.Size.X*UIScale)-1,
					win.Position.Y+(win.Size.Y*UIScale)-(10*UIScale)-(win.TitleSize*UIScale),

					win.Position.X+(win.Size.X*UIScale)-(10*UIScale),
					win.Position.Y+(win.Size.Y*UIScale)-(win.TitleSize*UIScale)-1,
					1, win.SizeColor, true)
				vector.StrokeLine(screen,
					win.Position.X+(win.Size.X*UIScale)-1,
					win.Position.Y+(win.Size.Y*UIScale)-(6*UIScale)-(win.TitleSize*UIScale),

					win.Position.X+(win.Size.X*UIScale)-(6*UIScale),
					win.Position.Y+(win.Size.Y*UIScale)-(win.TitleSize*UIScale)-1,
					1, win.SizeColor, true)
			}

			var buttonsWidth float32 = 0
			if win.Closable {
				var xpad float32 = (win.TitleSize * UIScale) / 4.0
				xThick := 3 * UIScale
				if win.HoverX {
					xThick += 1
					win.HoverX = false
				}
				vector.StrokeLine(screen,
					win.Position.X+(win.Size.X*UIScale)-(win.TitleSize*UIScale)+xpad,
					win.Position.Y+xpad,

					win.Position.X+(win.Size.X*UIScale)-xpad,
					win.Position.Y+(win.TitleSize*UIScale)-xpad,
					xThick, win.TitleColor, true)
				vector.StrokeLine(screen,
					win.Position.X+(win.Size.X*UIScale)-xpad,
					win.Position.Y+xpad,

					win.Position.X+(win.Size.X*UIScale)-(win.TitleSize*UIScale)+xpad,
					win.Position.Y+(win.TitleSize*UIScale)-xpad,
					xThick, win.TitleColor, true)

				buttonsWidth += (win.TitleSize * UIScale)
			}

			//Dragbar
			if win.Movable {
				var xThick float32 = 1
				if win.HoverDragbar {
					xThick += 1
					win.HoverDragbar = false
				}
				dpad := (win.TitleSize * UIScale) / 5
				for x := textWidth + float64((win.TitleSize*UIScale)/1.5); x < float64((win.Size.X*UIScale)-buttonsWidth); x = x + float64(UIScale*5.0) {
					vector.StrokeLine(screen,
						win.Position.X+float32(x), win.Position.Y+dpad,
						win.Position.X+float32(x), win.Position.Y+(win.TitleSize*UIScale)-dpad,
						xThick, win.DragColor, false)
				}
			}
		}

		//Draw frames
		if win.Border > 0 {
			FrameColor := win.BorderColor
			if win.Hovered {
				FrameColor = win.HoverColor
				win.Hovered = false
			}
			if win.TitleSize > 0 {
				vector.StrokeRect(screen,
					win.Position.X, win.Position.Y,
					win.Size.X*UIScale, (win.TitleSize * UIScale),
					win.Border, FrameColor, false)
			}
			//Window border
			vector.StrokeRect(screen,
				win.Position.X, win.Position.Y,
				win.Size.X*UIScale, (win.Size.Y*UIScale)-(win.TitleSize*UIScale),
				win.Border, FrameColor, false)
		}

		if *debugMode {
			grab := win.DragbarRect()
			vector.StrokeRect(screen, grab.X0, grab.Y0, grab.X1-grab.X0, grab.Y1-grab.Y0, 1, color.RGBA{R: 255, A: 255}, false)

			grab = win.XRect()
			vector.StrokeRect(screen, grab.X0, grab.Y0, grab.X1-grab.X0, grab.Y1-grab.Y0, 1, color.RGBA{G: 255, A: 255}, false)

			grab = win.ResizetabRect()
			vector.StrokeRect(screen, grab.X0, grab.Y0, grab.X1-grab.X0, grab.Y1-grab.Y0, 1, color.RGBA{B: 255, A: 255}, false)

			grab = win.GetMainRect()
			vector.StrokeRect(screen, grab.X0, grab.Y0, grab.X1-grab.X0, grab.Y1-grab.Y0, 1, color.RGBA{R: 255, G: 255, A: 255}, false)

			grab = win.GetTitleRect()
			vector.StrokeRect(screen, grab.X0, grab.Y0, grab.X1-grab.X0, grab.Y1-grab.Y0, 1, color.RGBA{B: 255, G: 255, A: 255}, false)
		}
	}

	buf := fmt.Sprintf("%4v FPS", int(math.Round(ebiten.ActualFPS())))
	ebitenutil.DebugPrintAt(screen, buf, defaultWindowWidth-55, defaultWindowHeight-18)
}
