package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const minDrag = 10

var mposOld Point

func (g *Game) Update() error {
	mx, my := ebiten.CursorPosition()
	mpos := Point{X: float32(mx), Y: float32(my)}

	click := inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0)
	clickHeld := inpututil.MouseButtonPressDuration(ebiten.MouseButton0)
	clickDrag := clickHeld > minDrag

	//altClick := inpututil.IsMouseButtonJustPressed(ebiten.MouseButton1)
	//altClickHeld := inpututil.MouseButtonPressDuration(ebiten.MouseButton1)

	//Check all windows
	for w, win := range Windows {
		if !win.Open {
			continue
		}

		if win.GetRect().ContainsPoint(mpos) {
			Windows[w].Hovered = true

			if win.TitleSize > 0 && win.GetTitleRect().ContainsPoint(mpos) {
				if win.Closable {
					if win.XRect().ContainsPoint(mpos) {
						win.HoverX = true
						if click {
							win.Open = false
						}
						continue
					}
				}
				if win.Resizable {
					if win.DragbarRect().ContainsPoint(mpos) {
						win.HoverDragbar = true
						if clickDrag {
							win.Position = PointAdd(win.Position, PointSubract(mpos, mposOld))
						}
						continue
					}
				}
			}

			//Window contents
			if win.Dumb {
				continue
			}

			for i, item := range win.Contents {
				if item.Rect.ContainsPoint(mpos) {
					win.Contents[i].Hovered = true
					if click {
						win.Contents[i].Activated = true
					}
				}
			}
		}
	}

	mposOld = mpos

	return nil
}
