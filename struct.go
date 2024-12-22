package main

import "image/color"

type WindowData struct {
	Title, Tooltip string
	Position       Point
	Size, SizeTemp Point

	Open, Hovered, Dumb                                                                                                      bool
	Closable, Movable, Resizable, Scrollable, Maximizable, Minimizeable                                                      bool
	ContentsBGColor, TitleBGColor, TitleColor, BorderColor, SizeColor, SizeHoverColor, DragColor, DragHoverColor, HoverColor color.RGBA
	TitleSize, TitleSizeTemp, Padding, Border                                                                                float32

	HoverX, HoverDragbar, HoverResizeTab bool

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

	Hovered, Clicked, Checked, Enabled bool
	FlowType                           FlowType
	FlowWrap                           bool
	Padding                            float32
	Scroll                             Point

	TextColor, Color, HoverColor, ClickColor, DisabledColor, CheckedColor color.RGBA

	Action   func()
	Contents []ItemData
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

type WindowSide int

const (
	SIDE_NONE = iota
	SIDE_TOP
	SIDE_RIGHT
	SIDE_BOTTOM
	SIDE_LEFT
)

type ItemTypeData int

const (
	ITEM_NONE = iota
	ITEM_TEXT
	ITEM_BUTTON
	ITEM_FLOW
)
