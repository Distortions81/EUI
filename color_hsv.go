package main

import (
	"encoding/json"
	"fmt"
	"image/color"
	"math"
)

// hsvaToRGBA converts HSV values (h in degrees [0,360), s and v in [0,1])
// and alpha in [0,1] to color.RGBA.
func hsvaToRGBA(h, s, v, a float64) color.RGBA {
	h = math.Mod(h, 360)
	if h < 0 {
		h += 360
	}
	c := v * s
	x := c * (1 - math.Abs(math.Mod(h/60, 2)-1))
	m := v - c
	var r, g, b float64
	switch {
	case h < 60:
		r, g, b = c, x, 0
	case h < 120:
		r, g, b = x, c, 0
	case h < 180:
		r, g, b = 0, c, x
	case h < 240:
		r, g, b = 0, x, c
	case h < 300:
		r, g, b = x, 0, c
	default:
		r, g, b = c, 0, x
	}
	r += m
	g += m
	b += m
	return color.RGBA{
		R: uint8(clamp(r*255, 0, 255)),
		G: uint8(clamp(g*255, 0, 255)),
		B: uint8(clamp(b*255, 0, 255)),
		A: uint8(clamp(a*255, 0, 255)),
	}
}

func rgbaToHSVA(c color.RGBA) (h, s, v, a float64) {
	r := float64(c.R) / 255
	g := float64(c.G) / 255
	b := float64(c.B) / 255
	max := math.Max(r, math.Max(g, b))
	min := math.Min(r, math.Min(g, b))
	d := max - min
	switch {
	case d == 0:
		h = 0
	case max == r:
		h = math.Mod((g-b)/d, 6) * 60
	case max == g:
		h = ((b-r)/d + 2) * 60
	default:
		h = ((r-g)/d + 4) * 60
	}
	if h < 0 {
		h += 360
	}
	if max == 0 {
		s = 0
	} else {
		s = d / max
	}
	v = max
	a = float64(c.A) / 255
	return
}

func clamp(v, min, max float64) float64 {
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}

// MarshalJSON implements json.Marshaler using HSV representation.
func (c Color) MarshalJSON() ([]byte, error) {
	h, s, v, a := rgbaToHSVA(color.RGBA(c))
	return json.Marshal(struct {
		HSV [4]float64 `json:"HSV"`
	}{[4]float64{h, s, v, a}})
}

// UnmarshalJSON accepts either HSV or RGBA representations.
func (c *Color) UnmarshalJSON(data []byte) error {
	var hstruct struct {
		HSV [4]float64 `json:"HSV"`
	}
	if err := json.Unmarshal(data, &hstruct); err == nil && hstruct.HSV != [4]float64{} {
		*c = Color(hsvaToRGBA(hstruct.HSV[0], hstruct.HSV[1], hstruct.HSV[2], hstruct.HSV[3]))
		return nil
	}
	var rgba struct{ R, G, B, A uint8 }
	if err := json.Unmarshal(data, &rgba); err == nil {
		*c = Color(color.RGBA{rgba.R, rgba.G, rgba.B, rgba.A})
		return nil
	}
	return fmt.Errorf("invalid color format")
}
