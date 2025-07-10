package main

import "image/color"

var defaultTheme = &windowData{
	TitleHeight:     24,
	Border:          1,
	TitleColor:      color.RGBA{R: 255, G: 255, B: 255, A: 255},
	BorderColor:     color.RGBA{R: 64, G: 64, B: 64, A: 255},
	SizeTabColor:    color.RGBA{R: 48, G: 48, B: 48, A: 255},
	DragbarColor:    color.RGBA{R: 64, G: 64, B: 64, A: 255},
	HoverTitleColor: color.RGBA{R: 64, G: 128, B: 128, A: 255},
	HoverColor:      color.RGBA{R: 80, G: 80, B: 80, A: 255},
	BGColor:         color.RGBA{R: 16, G: 16, B: 16, A: 255},
	ActiveColor:     color.RGBA{R: 0, G: 128, B: 128, A: 255},

	Movable: true, Closable: true, Resizable: true, Open: true, AutoSize: false,
}

var defaultButton = &itemData{
	Text:      "Button",
	ItemType:  ITEM_BUTTON,
	Size:      point{X: 128, Y: 64},
	Position:  point{X: 4, Y: 4},
	FontSize:  12,
	LineSpace: 1.2,

	Fillet: 8,
	Filled: true, Outlined: true,
	Border:    1,
	BorderPad: 4,

	TextColor:  color.RGBA{R: 255, G: 255, B: 255, A: 255},
	Color:      color.RGBA{R: 48, G: 48, B: 48, A: 255},
	HoverColor: color.RGBA{R: 48, G: 96, B: 96, A: 255},
	ClickColor: color.RGBA{R: 192, G: 64, B: 64, A: 255},
}

var defaultText = &itemData{
	Text:      "Sample text:\nThe quick brown fox\njumps over the lazy dog.",
	ItemType:  ITEM_TEXT,
	Size:      point{X: 128, Y: 128},
	Position:  point{X: 4, Y: 4},
	FontSize:  24,
	LineSpace: 1.2,
	TextColor: color.RGBA{R: 255, G: 255, B: 255, A: 255},
}

var defaultCheckbox = &itemData{
	Text:      "Option 1",
	ItemType:  ITEM_CHECKBOX,
	Size:      point{X: 128, Y: 32},
	Position:  point{X: 4, Y: 4},
	AuxSize:   point{X: 16, Y: 16},
	AuxSpace:  4,
	FontSize:  12,
	LineSpace: 1.2,

	Fillet: 8,
	Filled: true, Outlined: true,
	Border:    1,
	BorderPad: 4,

	TextColor:  color.RGBA{R: 255, G: 255, B: 255, A: 255},
	Color:      color.RGBA{R: 48, G: 48, B: 48, A: 255},
	HoverColor: color.RGBA{R: 96, G: 96, B: 96, A: 255},
	ClickColor: color.RGBA{R: 0, G: 128, B: 128, A: 255},
}

var defaultInput = &itemData{
	ItemType:  ITEM_INPUT,
	Size:      point{X: 128, Y: 24},
	Position:  point{X: 4, Y: 4},
	FontSize:  12,
	LineSpace: 1.2,

	Fillet: 4,
	Filled: true, Outlined: true,
	Border:    1,
	BorderPad: 2,

	TextColor:  color.RGBA{R: 255, G: 255, B: 255, A: 255},
	Color:      color.RGBA{R: 48, G: 48, B: 48, A: 255},
	HoverColor: color.RGBA{R: 96, G: 96, B: 96, A: 255},
	ClickColor: color.RGBA{R: 0, G: 128, B: 128, A: 255},
}

var defaultSlider = &itemData{
	ItemType:  ITEM_SLIDER,
	Size:      point{X: 128, Y: 16},
	Position:  point{X: 4, Y: 4},
	Fillet:    4,
	Filled:    true,
	Outlined:  true,
	Border:    1,
	BorderPad: 2,

	TextColor:  color.RGBA{255, 255, 255, 255},
	Color:      color.RGBA{64, 64, 64, 255},
	HoverColor: color.RGBA{96, 96, 96, 255},
	ClickColor: color.RGBA{0, 128, 128, 255},
}

var defaultColorSelector = &itemData{
	ItemType:  ITEM_COLORSEL,
	Size:      point{X: 128, Y: 80},
	Position:  point{X: 4, Y: 4},
	Fillet:    4,
	Filled:    true,
	Outlined:  true,
	Border:    1,
	BorderPad: 2,

	TextColor:  color.RGBA{255, 255, 255, 255},
	Color:      color.RGBA{0, 0, 0, 255},
	HoverColor: color.RGBA{96, 96, 96, 255},
	ClickColor: color.RGBA{0, 128, 128, 255},
}
