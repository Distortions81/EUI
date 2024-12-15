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

		winRect := win.GetRect()
		if winRect.ContainsPoint(mpos) {
			Windows[w].Hovered = true

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
		} else {
			Windows[w].Hovered = false
		}
	}

	return nil
}
