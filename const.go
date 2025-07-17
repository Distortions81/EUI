package main

const (
	// MinWinSizeX and MinWinSizeY enforce a sane minimum window size.
	// Windows should never be smaller than 64x64 pixels.
	MinWinSizeX = 64
	MinWinSizeY = 64

	DefaultTabWidth  = 128
	DefaultTabHeight = 24

	CornerSize = 14
	Tol        = 2

	// sliderMaxLabel defines the formatted text used to measure the value
	// field of sliders. Using a constant ensures int and float sliders have
	// identical track lengths regardless of their numeric ranges.
	sliderMaxLabel = "100.00"

	// InactiveDim controls the opacity of the black overlay drawn over
	// inactive windows. Values range from 0.0 (no dimming) to 1.0 (fully
	// black)
)

var (
	InactiveDim = 0.20
)
