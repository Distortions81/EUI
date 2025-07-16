package main

import "image/color"

var defaultTheme = &windowData{
	TitleHeight:     24,
	Border:          0,
	Outlined:        false,
	Fillet:          4,
	Padding:         4,
	Margin:          4,
	BorderPad:       0,
	TitleColor:      Color(color.RGBA{R: 255, G: 255, B: 255, A: 255}),
	TitleTextColor:  Color(color.RGBA{R: 255, G: 255, B: 255, A: 255}),
	TitleBGColor:    Color(color.RGBA{R: 64, G: 64, B: 64, A: 255}),
	CloseBGColor:    Color(color.RGBA{R: 48, G: 48, B: 48, A: 255}),
	DragbarSpacing:  5,
	ShowDragbar:     false,
	BorderColor:     Color(color.RGBA{R: 64, G: 64, B: 64, A: 255}),
	SizeTabColor:    Color(color.RGBA{R: 48, G: 48, B: 48, A: 255}),
	DragbarColor:    Color(color.RGBA{R: 64, G: 64, B: 64, A: 255}),
	HoverTitleColor: Color(color.RGBA{R: 0, G: 160, B: 160, A: 255}),
	HoverColor:      Color(color.RGBA{R: 80, G: 80, B: 80, A: 255}),
	BGColor:         Color(color.RGBA{R: 32, G: 32, B: 32, A: 255}),
	ActiveColor:     Color(color.RGBA{R: 0, G: 160, B: 160, A: 255}),

	Movable: true, Closable: true, Resizable: true, Open: true, AutoSize: false,
}

var defaultButton = &itemData{
	Text:      "",
	ItemType:  ITEM_BUTTON,
	Size:      point{X: 128, Y: 64},
	Position:  point{X: 4, Y: 4},
	FontSize:  12,
	LineSpace: 1.2,

	Padding: 0,
	Margin:  4,

	Fillet: 8,
	Filled: true, Outlined: false,
	Border:    0,
	BorderPad: 4,

	TextColor:  Color(color.RGBA{R: 255, G: 255, B: 255, A: 255}),
	Color:      Color(color.RGBA{R: 48, G: 48, B: 48, A: 255}),
	HoverColor: Color(color.RGBA{R: 96, G: 96, B: 96, A: 255}),
	ClickColor: Color(color.RGBA{R: 0, G: 160, B: 160, A: 255}),
}

var defaultText = &itemData{
	Text:      "",
	ItemType:  ITEM_TEXT,
	Size:      point{X: 128, Y: 128},
	Position:  point{X: 4, Y: 4},
	FontSize:  24,
	LineSpace: 1.2,
	Padding:   0,
	Margin:    2,
	TextColor: Color(color.RGBA{R: 255, G: 255, B: 255, A: 255}),
}

var defaultCheckbox = &itemData{
	Text:      "",
	ItemType:  ITEM_CHECKBOX,
	Size:      point{X: 128, Y: 24},
	Position:  point{X: 4, Y: 2},
	AuxSize:   point{X: 16, Y: 16},
	AuxSpace:  4,
	FontSize:  12,
	LineSpace: 1.2,
	Padding:   0,
	Margin:    2,

	Fillet: 8,
	Filled: true, Outlined: false,
	Border:    0,
	BorderPad: 4,

	TextColor:  Color(color.RGBA{R: 255, G: 255, B: 255, A: 255}),
	Color:      Color(color.RGBA{R: 48, G: 48, B: 48, A: 255}),
	HoverColor: Color(color.RGBA{R: 96, G: 96, B: 96, A: 255}),
	ClickColor: Color(color.RGBA{R: 0, G: 160, B: 160, A: 255}),
}

var defaultInput = &itemData{
	ItemType:  ITEM_INPUT,
	Size:      point{X: 128, Y: 24},
	Position:  point{X: 4, Y: 4},
	FontSize:  12,
	LineSpace: 1.2,
	Padding:   0,
	Margin:    2,

	Fillet: 4,
	Filled: true, Outlined: false,
	Border:    0,
	BorderPad: 2,

	TextColor:  Color(color.RGBA{R: 255, G: 255, B: 255, A: 255}),
	Color:      Color(color.RGBA{R: 48, G: 48, B: 48, A: 255}),
	HoverColor: Color(color.RGBA{R: 96, G: 96, B: 96, A: 255}),
	ClickColor: Color(color.RGBA{R: 0, G: 160, B: 160, A: 255}),
}

var defaultRadio = &itemData{
	Text:      "",
	ItemType:  ITEM_RADIO,
	Size:      point{X: 128, Y: 24},
	Position:  point{X: 4, Y: 2},
	AuxSize:   point{X: 16, Y: 16},
	AuxSpace:  4,
	FontSize:  12,
	LineSpace: 1.2,
	Padding:   0,
	Margin:    2,

	Fillet: 8,
	Filled: true, Outlined: false,
	Border:    0,
	BorderPad: 4,

	TextColor:  Color(color.RGBA{R: 255, G: 255, B: 255, A: 255}),
	Color:      Color(color.RGBA{R: 48, G: 48, B: 48, A: 255}),
	HoverColor: Color(color.RGBA{R: 96, G: 96, B: 96, A: 255}),
	ClickColor: Color(color.RGBA{R: 0, G: 160, B: 160, A: 255}),
}

var defaultSlider = &itemData{
	ItemType: ITEM_SLIDER,
	Size:     point{X: 128, Y: 24},
	Position: point{X: 4, Y: 4},
	AuxSize:  point{X: 8, Y: 16},
	AuxSpace: 4,
	FontSize: 12,
	Padding:  0,
	Margin:   4,

	MinValue: 0,
	MaxValue: 100,
	Value:    0,
	IntOnly:  false,

	Fillet: 4,
	Filled: true, Outlined: false,
	Border:    0,
	BorderPad: 2,

	TextColor:  Color(color.RGBA{R: 255, G: 255, B: 255, A: 255}),
	Color:      Color(color.RGBA{R: 48, G: 48, B: 48, A: 255}),
	HoverColor: Color(color.RGBA{R: 96, G: 96, B: 96, A: 255}),
	ClickColor: Color(color.RGBA{R: 0, G: 160, B: 160, A: 255}),
}

var defaultDropdown = &itemData{
	ItemType: ITEM_DROPDOWN,
	Size:     point{X: 128, Y: 24},
	Position: point{X: 4, Y: 4},
	FontSize: 12,
	Padding:  0,
	Margin:   4,

	Fillet: 4,
	Filled: true, Outlined: false,
	Border:    0,
	BorderPad: 2,

	TextColor:  Color(color.RGBA{R: 255, G: 255, B: 255, A: 255}),
	Color:      Color(color.RGBA{R: 48, G: 48, B: 48, A: 255}),
	HoverColor: Color(color.RGBA{R: 96, G: 96, B: 96, A: 255}),
	ClickColor: Color(color.RGBA{R: 0, G: 160, B: 160, A: 255}),
	MaxVisible: 5,
}

var defaultColorWheel = &itemData{
	ItemType: ITEM_COLORWHEEL,
	Size:     point{X: 128, Y: 128},
	Position: point{X: 4, Y: 4},
	Margin:   4,
}

var defaultTab = &itemData{
	ItemType:      ITEM_FLOW,
	FontSize:      12,
	Filled:        true,
	Border:        0,
	BorderPad:     2,
	Fillet:        4,
	ActiveOutline: false,
	TextColor:     Color(color.RGBA{R: 255, G: 255, B: 255, A: 255}),
	Color:         Color(color.RGBA{R: 64, G: 64, B: 64, A: 255}),
	HoverColor:    Color(color.RGBA{R: 96, G: 96, B: 96, A: 255}),
	ClickColor:    Color(color.RGBA{R: 0, G: 160, B: 160, A: 255}),
}

// base copies preserve the initial defaults so that LoadTheme can reset
// to these values before applying theme overrides.
var (
	baseWindow     = *defaultTheme
	baseButton     = *defaultButton
	baseText       = *defaultText
	baseCheckbox   = *defaultCheckbox
	baseRadio      = *defaultRadio
	baseInput      = *defaultInput
	baseSlider     = *defaultSlider
	baseDropdown   = *defaultDropdown
	baseColorWheel = *defaultColorWheel
	baseTab        = *defaultTab
	baseTheme      = &Theme{
		Window:   baseWindow,
		Button:   baseButton,
		Text:     baseText,
		Checkbox: baseCheckbox,
		Radio:    baseRadio,
		Input:    baseInput,
		Slider:   baseSlider,
		Dropdown: baseDropdown,
		Tab:      baseTab,
	}
)
