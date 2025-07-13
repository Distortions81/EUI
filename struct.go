package main

import (
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type windowData struct {
	Title           string
	Position        point
	Size            point
	PinTo           pinType
	Padding, Border float32
	Outlined        bool

	Open, Hovered, Flow,
	Closable, Movable, Resizable,
	HoverClose, HoverDragbar,
	AutoSize bool

	TitleHeight float32

	BGColor, TitleBGColor, TitleColor, BorderColor,
	SizeTabColor, DragbarColor,
	HoverTitleColor, HoverColor, ActiveColor color.RGBA

	Contents []*itemData
}

type itemData struct {
	Parent *itemData
	// Name is used when the item is part of a tabbed flow
	Name      string
	Text      string
	Position  point
	Size      point
	Alignment alignType
	PinTo     pinType
	FontSize  float32
	LineSpace float32 //Multiplier, 1.0 = no gap between lines
	ItemType  itemTypeData

	Value      float32
	MinValue   float32
	MaxValue   float32
	IntOnly    bool
	RadioGroup string

	Hovered, Checked, Focused,
	Disabled, Invisible bool
	Clicked  time.Time
	FlowType flowType
	Scroll   point

	Fixed, Scrollable bool

	ImageName string
	Image     *ebiten.Image

	//Style
	Fillet            float32
	Border, BorderPad float32
	Filled, Outlined  bool
	AuxSize           point
	AuxSpace          float32

	TextColor, Color, HoverColor,
	ClickColor, DisabledColor, CheckedColor color.RGBA

	Action   func()
	Contents []*itemData

	// Tabs allows a flow to contain multiple tabbed flows. Only the
	// flow referenced by ActiveTab will be drawn and receive input.
	Tabs      []*itemData
	ActiveTab int

	// DrawRect stores the last drawn rectangle of the item in screen
	// coordinates so input handling can use the exact same area that was
	// rendered.
	DrawRect rect
}

type roundRect struct {
	Size, Position point
	Fillet, Border float32
	Filled         bool
	Color          color.RGBA
}

type rect struct {
	X0, Y0, X1, Y1 float32
}

type point struct {
	X, Y float32
}

type flowType int

const (
	FLOW_HORIZONTAL = iota
	FLOW_VERTICAL

	FLOW_HORIZONTAL_REV
	FLOW_VERTICAL_REV
)

type alignType int

const (
	ALIGN_NONE = iota
	ALIGN_LEFT
	ALIGN_CENTER
	ALIGN_RIGHT
)

type pinType int

const (
	PIN_TOP_LEFT = iota
	PIN_TOP_CENTER
	PIN_TOP_RIGHT

	PIN_MID_LEFT
	PIN_MID_CENTER
	PIN_MID_RIGHT

	PIN_BOTTOM_LEFT
	PIN_BOTTOM_CENTER
	PIN_BOTTOM_RIGHT
)

type dragType int

const (
	PART_NONE = iota

	PART_BAR
	PART_CLOSE

	PART_TOP
	PART_RIGHT
	PART_BOTTOM
	PART_LEFT

	PART_TOP_RIGHT
	PART_BOTTOM_RIGHT
	PART_BOTTOM_LEFT
	PART_TOP_LEFT
)

type itemTypeData int

const (
	ITEM_NONE = iota
	ITEM_FLOW
	ITEM_TEXT
	ITEM_BUTTON
	ITEM_CHECKBOX
	ITEM_RADIO
	ITEM_INPUT
	ITEM_SLIDER
)
