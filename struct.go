package main

import "image/color"

type WindowData struct {
	Title, Tooltip string
	Position       Point
	Mag, MagTemp   Magnatude

	Open, Hovered, Dumb                                                                                                      bool
	Closable, Movable, Resizable, Scrollable, Maximizable, Minimizeable                                                      bool
	ContentsBGColor, TitleBGColor, TitleColor, BorderColor, SizeColor, SizeHoverColor, DragColor, DragHoverColor, HoverColor color.RGBA
	TitleSize, TitleSizeTemp, Padding, Border                                                                                float32

	HoverX, HoverDragbar, HoverResizeTab bool

	Contents []ItemData
}

type ItemData struct {
	Text     string
	Position Point
	Size     Magnatude
	Rect     Rect

	Value float32

	Hovered, Activated, Checked, Enabled bool
	FlowType                             FlowType
	FlowWrap                             bool
	Padding                              float32
	Scroll                               Point

	Color, HoverColor, ActivatedColor, DisabledColor, CheckedColor color.RGBA

	Contents []ItemData
}

type Point struct {
	X, Y float32
}

type Magnatude Point

type FlowType int

type Rect struct {
	X0, Y0, X1, Y1 float32
}

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
