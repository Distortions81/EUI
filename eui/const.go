package eui

const (
	// minWinSizeX and MinWinSizeY enforce a sane minimum window size.
	// Windows should never be smaller than 100x100 pixels.
	minWinSizeX = 100
	minWinSizeY = 100

	defaultTabWidth  = 128
	defaultTabHeight = 24

	// scrollTolerance defines the padding around window edges used to detect
	// resize drags along the sides, expressed as a fraction of the screen
	// dimension.
	scrollTolerance float32 = 2.0 / 1024.0
	// cornerTolerance defines the larger area around window corners used to
	// detect diagonal resizing, expressed as a fraction of the screen
	// dimension.
	cornerTolerance float32 = 16.0 / 1024.0

	// sliderMaxLabel defines the formatted text used to measure the value
	// field of sliders. Using a constant ensures int and float sliders have
	// identical track lengths regardless of their numeric ranges.
	sliderMaxLabel = "100.00"

	// dropdownOverlayReserve defines the number of option heights reserved
	// at the top and bottom of the screen when positioning dropdown menus
	// to leave room for overlay controls.
	dropdownOverlayReserve = 1
)
