package eui

const (
	// minWinSizeX and MinWinSizeY enforce a sane minimum window size.
	// Windows should never be smaller than 64x64 pixels.
	minWinSizeX = 64
	minWinSizeY = 64

	defaultTabWidth  = 128
	defaultTabHeight = 24

	// scrollTolerance defines the padding around window edges used to detect
	// resize drags along the sides.
	scrollTolerance = 2
	// cornerTolerance defines the larger area around window corners used to
	// detect diagonal resizing.
	cornerTolerance = 16

	// sliderMaxLabel defines the formatted text used to measure the value
	// field of sliders. Using a constant ensures int and float sliders have
	// identical track lengths regardless of their numeric ranges.
	sliderMaxLabel = "100.00"

	// dropdownOverlayReserve defines the number of option heights reserved
	// at the top and bottom of the screen when positioning dropdown menus
	// to leave room for overlay controls.
	dropdownOverlayReserve = 1
)
