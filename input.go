package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (g *Game) Update() error {
	mx, my := ebiten.CursorPosition()
	mpos := Point{X: float32(mx), Y: float32(my)}

	click := inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0)
	//clickHeld := inpututil.MouseButtonPressDuration(ebiten.MouseButton0)

	//altClick := inpututil.IsMouseButtonJustPressed(ebiten.MouseButton1)
	//altClickHeld := inpututil.MouseButtonPressDuration(ebiten.MouseButton1)

	//Check all windows
	for w, win := range Windows {
		if !win.Open {
			continue
		}

		if Windows[w].Hovered = win.GetRect().ContainsPoint(mpos); Windows[w].Hovered {
			if win.GetTitleRect().ContainsPoint(mpos) {
				var xpad float32 = (win.TitleSize * UIScale) / 4.0
				closeRect := Rect{
					X0: win.Position.X + (win.Size.X * UIScale) - (win.TitleSize * UIScale) + xpad,
					Y0: win.Position.Y + xpad,

					X1: win.Position.X + (win.Size.X * UIScale) - xpad,
					Y1: win.Position.Y + (win.TitleSize * UIScale) - xpad,
				}
				if win.HoverX = closeRect.ContainsPoint(mpos); win.HoverX {
					if click {
						win.Open = false
					}
				}
				continue
			}

			//Window contents
			if win.Dumb {
				continue
			}

			for i, item := range win.Contents {
				if item.Rect.ContainsPoint(mpos) {
					if click {
						win.Contents[i].Activated = true
						win.Contents[i].Hovered = false
					} else {
						win.Contents[i].Activated = false
						win.Contents[i].Hovered = true
					}
				} else {
					win.Contents[i].Activated = false
					win.Contents[i].Hovered = false
				}
			}
		}
	}

	return nil
}
