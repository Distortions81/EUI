package main

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

// ColorWheelImage creates an Ebiten image containing a color wheel of the given size.
// The wheel ranges 0-359 degrees with black at the center and fully saturated
// color on the outer edge.
func ColorWheelImage(size int) *ebiten.Image {
	if size <= 0 {
		return ebiten.NewImage(1, 1)
	}
	img := ebiten.NewImage(size, size)
	r := float64(size) / 2
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
			v := dist / r
			if v < 0 {
				v = 0
			} else if v > 1 {
				v = 1
			}
			col := hsvaToRGBA(ang, 1, v, 1)
			img.Set(x, y, col)
		}
	}
	return img
}
