package main

import "image/color"

var (
	ColorWhite   = Color(color.RGBA{255, 255, 255, 255})
	ColorBlack   = Color(color.RGBA{0, 0, 0, 255})
	ColorRed     = Color(color.RGBA{203, 67, 53, 255})
	ColorGreen   = Color(color.RGBA{40, 180, 99, 255})
	ColorBlue    = Color(color.RGBA{41, 128, 185, 255})
	ColorYellow  = Color(color.RGBA{244, 208, 63, 255})
	ColorGray    = Color(color.RGBA{128, 128, 128, 255})
	ColorOrange  = Color(color.RGBA{243, 156, 18, 255})
	ColorPink    = Color(color.RGBA{255, 151, 197, 255})
	ColorPurple  = Color(color.RGBA{165, 105, 189, 255})
	ColorSilver  = Color(color.RGBA{209, 209, 209, 255})
	ColorTeal    = Color(color.RGBA{64, 199, 178, 255})
	ColorMaroon  = Color(color.RGBA{199, 54, 103, 255})
	ColorNavy    = Color(color.RGBA{99, 114, 166, 255})
	ColorOlive   = Color(color.RGBA{134, 166, 99, 255})
	ColorLime    = Color(color.RGBA{206, 231, 114, 255})
	ColorFuchsia = Color(color.RGBA{209, 114, 231, 255})
	ColorAqua    = Color(color.RGBA{114, 228, 231, 255})
	ColorBrown   = Color(color.RGBA{176, 116, 78, 255})
	ColorRust    = Color(color.RGBA{210, 113, 52, 255})

	ColorLightRed     = Color(color.RGBA{255, 83, 83, 255})
	ColorLightGreen   = Color(color.RGBA{170, 255, 159, 255})
	ColorLightBlue    = Color(color.RGBA{159, 186, 255, 255})
	ColorLightYellow  = Color(color.RGBA{255, 251, 159, 255})
	ColorLightGray    = Color(color.RGBA{236, 236, 236, 255})
	ColorLightOrange  = Color(color.RGBA{252, 213, 134, 255})
	ColorLightPink    = Color(color.RGBA{254, 163, 182, 255})
	ColorLightPurple  = Color(color.RGBA{254, 163, 245, 255})
	ColorLightSilver  = Color(color.RGBA{228, 228, 228, 255})
	ColorLightTeal    = Color(color.RGBA{152, 221, 210, 255})
	ColorLightMaroon  = Color(color.RGBA{215, 124, 143, 255})
	ColorLightNavy    = Color(color.RGBA{128, 152, 197, 255})
	ColorLightOlive   = Color(color.RGBA{186, 228, 144, 255})
	ColorLightLime    = Color(color.RGBA{219, 243, 153, 255})
	ColorLightFuchsia = Color(color.RGBA{239, 196, 253, 255})
	ColorLightAqua    = Color(color.RGBA{196, 246, 253, 255})

	ColorDarkRed      = Color(color.RGBA{146, 22, 22, 255})
	ColorDarkGreen    = Color(color.RGBA{22, 146, 24, 255})
	ColorDarkBlue     = Color(color.RGBA{22, 98, 146, 255})
	ColorDarkYellow   = Color(color.RGBA{139, 146, 22, 255})
	ColorDarkGray     = Color(color.RGBA{111, 111, 111, 255})
	ColorCharcoal     = Color(color.RGBA{16, 16, 16, 255})
	ColorCharcoalSemi = Color(color.RGBA{16, 16, 16, 128})
	ColorDarkOrange   = Color(color.RGBA{175, 117, 32, 255})
	ColorDarkPink     = Color(color.RGBA{128, 64, 64, 255})
	ColorDarkPurple   = Color(color.RGBA{137, 32, 175, 255})
	ColorDarkSilver   = Color(color.RGBA{162, 162, 162, 255})
	ColorDarkTeal     = Color(color.RGBA{27, 110, 86, 255})
	ColorDarkMaroon   = Color(color.RGBA{110, 27, 55, 255})
	ColorDarkNavy     = Color(color.RGBA{16, 46, 85, 255})
	ColorDarkOlive    = Color(color.RGBA{60, 101, 19, 255})
	ColorDarkLime     = Color(color.RGBA{122, 154, 45, 255})
	ColorDarkFuchsia  = Color(color.RGBA{154, 45, 141, 255})
	ColorDarkAqua     = Color(color.RGBA{45, 154, 154, 255})

	ColorVeryDarkGray = Color(color.RGBA{64, 64, 64, 255})
)

// builtinColorMap maps lowercase color names to their Color values.
var builtinColorMap = map[string]Color{
	"white":   ColorWhite,
	"black":   ColorBlack,
	"red":     ColorRed,
	"green":   ColorGreen,
	"blue":    ColorBlue,
	"yellow":  ColorYellow,
	"gray":    ColorGray,
	"orange":  ColorOrange,
	"pink":    ColorPink,
	"purple":  ColorPurple,
	"silver":  ColorSilver,
	"teal":    ColorTeal,
	"maroon":  ColorMaroon,
	"navy":    ColorNavy,
	"olive":   ColorOlive,
	"lime":    ColorLime,
	"fuchsia": ColorFuchsia,
	"aqua":    ColorAqua,
	"brown":   ColorBrown,
	"rust":    ColorRust,
}
