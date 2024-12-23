package main

import (
	"image/color"
	"time"
)

type WindowData struct {
	Title            string
	Position         Point
	Size, ScreenSize Point
	Padding, Border  float32

	Open, Hovered,
	Closable, Movable, Resizable,
	HoverClose, HoverDragbar, HoverResizeTab bool

	TitleHeight, TitleScreenHeight float32

	BGColor, TitleBGColor, TitleColor, BorderColor,
	SizeTabColor, SizeTabHoverColor, DragbarColor,
	HoverTitleColor, HoverColor, ActiveColor color.RGBA

	Contents []*ItemData
}

type ItemData struct {
	Text      string
	Position  Point
	Size      Point
	FontSize  float32
	LineSpace float32 //Multiplier, 1.0 = no gap between lines
	ItemType  ItemTypeData

	Value float32

	Hovered, Checked,
	Disabled, Invisable bool
	Clicked  time.Time
	FlowType FlowType
	Scroll   Point

	//Style
	Fillet            float32
	Border, BorderPad float32
	Filled, Outlined  bool

	TextColor, Color, HoverColor,
	ClickColor, DisabledColor, CheckedColor color.RGBA

	Action   func()
	Contents []ItemData
}

type RoundRect struct {
	Size, Position Point
	Fillet, Border float32
	Filled         bool
	Color          color.RGBA
}

type Rect struct {
	X0, Y0, X1, Y1 float32
}

type Point struct {
	X, Y float32
}

type FlowType int

const (
	FLOW_HORIZONTAL = iota
	FLOW_VERTICAL
)

type AlignType int

const (
	ALIGN_START = iota
	ALIGN_CENTER
	ALIGN_END
)

type WindowEdge int

const (
	EDGE_NONE = iota
	EDGE_TOP
	EDGE_RIGHT
	EDGE_BOTTOM
	EDGE_LEFT
)

type ItemTypeData int

const (
	ITEM_NONE = iota
	ITEM_TEXT
	ITEM_BUTTON
	ITEM_FLOW
)
