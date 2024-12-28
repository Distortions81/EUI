package main

import (
	"image/color"
	"time"
)

type windowData struct {
	Title           string
	Position        point
	Size            point
	PinTo           pinType
	Padding, Border float32

	Open, Hovered,
	Closable, Movable, Resizable,
	HoverClose, HoverDragbar bool

	TitleHeight float32

	BGColor, TitleBGColor, TitleColor, BorderColor,
	SizeTabColor, DragbarColor,
	HoverTitleColor, HoverColor, ActiveColor color.RGBA

	Contents []*itemData
}

type itemData struct {
	Text      string
	Position  point
	Size      point
	Alignment alignType
	PinTo     pinType
	FontSize  float32
	LineSpace float32 //Multiplier, 1.0 = no gap between lines
	ItemType  itemTypeData

	Value float32

	Hovered, Checked,
	Disabled, Invisable bool
	Clicked  time.Time
	FlowType flowType
	Scroll   point

	//Style
	Fillet            float32
	Border, BorderPad float32
	Filled, Outlined  bool

	TextColor, Color, HoverColor,
	ClickColor, DisabledColor, CheckedColor color.RGBA

	Action   func()
	Contents []itemData
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

type windowEdge int

const (
	EDGE_NONE = iota

	EDGE_TOP
	EDGE_RIGHT
	EDGE_BOTTOM
	EDGE_LEFT

	EDGE_TOP_RIGHT
	EDGE_BOTTOM_RIGHT
	EDGE_BOTTOM_LEFT
	EDGE_TOP_LEFT
)

type itemTypeData int

const (
	ITEM_NONE = iota
	ITEM_TEXT
	ITEM_BUTTON
	ITEM_FLOW
)
