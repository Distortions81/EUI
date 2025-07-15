package main

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

// ColorWheelImage creates an Ebiten image containing a color wheel of the given size.
// The wheel ranges 0-359 degrees with white on the outside, fully saturated color
// at the midpoint and black at the center.
func ColorWheelImage(size int) *ebiten.Image {
	if size <= 0 {
		return ebiten.NewImage(1, 1)
	}
	img := ebiten.NewImage(size, size)
	r := float64(size) / 2
	mid := r * 0.5
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			dx := float64(x) - r
			dy := float64(y) - r
			dist := math.Hypot(dx, dy)
			if dist > r {
				img.Set(x, y, color.Transparent)
				continue
			}
			ang := math.Atan2(dy, dx) * 180 / math.Pi
			if ang < 0 {
				ang += 360
			}
			var col color.RGBA
			if dist <= mid {
				// from black to full color
				v := dist / mid
				col = hsvaToRGBA(ang, 1, v, 1)
			} else {
				// from full color to white
				t := (dist - mid) / (r - mid)
				col = hsvaToRGBA(ang, 1-t, 1, 1)
			}
			img.Set(x, y, col)
		}
	}
	return img
}
