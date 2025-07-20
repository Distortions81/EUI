package eui

const (
	// minWinSizeX and MinWinSizeY enforce a sane minimum window size.
	// Windows should never be smaller than 64x64 pixels.
	minWinSizeX = 64
	minWinSizeY = 64

	defaultTabWidth  = 128
	defaultTabHeight = 24

	scrollTolerance = 2

	// sliderMaxLabel defines the formatted text used to measure the value
	// field of sliders. Using a constant ensures int and float sliders have
	// identical track lengths regardless of their numeric ranges.
	sliderMaxLabel = "100.00"
)
