package eui

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

var sampleOffsets = []float64{0.125, 0.375, 0.625, 0.875}

// colorWheelImage creates an Ebiten image containing a color wheel of the given size.
// The wheel ranges 0-359 degrees with black at the center and fully saturated
// color on the outer edge.
func colorWheelImage(size int) *ebiten.Image {
	if size <= 0 {
		return ebiten.NewImage(1, 1)
	}
	img := ebiten.NewImage(size, size)
	r := float64(size) / 2
	// Use a 4x4 grid of subpixel samples for smoother edges
	maxSamples := len(sampleOffsets) * len(sampleOffsets)
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			var rr, gg, bb, aa float64
			var coverage int
			for _, oy := range sampleOffsets {
				for _, ox := range sampleOffsets {
					dx := float64(x) + ox - r
					dy := float64(y) + oy - r
					dist := math.Hypot(dx, dy)
					if dist > r {
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
					rr += float64(col.R)
					gg += float64(col.G)
					bb += float64(col.B)
					aa += float64(col.A)
					coverage++
				}
			}
			if coverage == 0 {
				img.Set(x, y, color.Transparent)
				continue
			}
			img.Set(x, y, color.RGBA{
				R: uint8(rr / float64(maxSamples)),
				G: uint8(gg / float64(maxSamples)),
				B: uint8(bb / float64(maxSamples)),
				A: uint8(aa / float64(maxSamples)),
			})
		}
	}
	return img
}

// wheelImage returns a cached color wheel image for the item, generating one
// if necessary. The image is recreated when the requested size differs from the
// cached version.
func (it *itemData) wheelImage(size int) *ebiten.Image {
	if size <= 0 {
		size = 1
	}
	if it.wheelImg == nil || it.wheelImg.Bounds().Dx() != size {
		it.wheelImg = colorWheelImage(size)
	}
	return it.wheelImg
}
