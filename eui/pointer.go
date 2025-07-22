package eui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// pointerPosition returns the current pointer position.
// If a touch is active, the first touch is used. Otherwise the mouse cursor position is returned.
func pointerPosition() (int, int) {
	ids := ebiten.AppendTouchIDs(nil)
	if len(ids) > 0 {
		return ebiten.TouchPosition(ids[0])
	}
	return ebiten.CursorPosition()
}

// pointerWheel returns the wheel delta when using a mouse.
// For touch input this always returns zero.
func pointerWheel() (float64, float64) {
	if len(ebiten.AppendTouchIDs(nil)) > 0 {
		return 0, 0
	}
	return ebiten.Wheel()
}

// pointerJustPressed reports whether the primary pointer was just pressed.
func pointerJustPressed() bool {
	if len(inpututil.AppendJustPressedTouchIDs(nil)) > 0 {
		return true
	}
	return inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0)
}

// pointerPressed reports whether the primary pointer is currently pressed.
func pointerPressed() bool {
	if len(ebiten.AppendTouchIDs(nil)) > 0 {
		return true
	}
	return ebiten.IsMouseButtonPressed(ebiten.MouseButton0)
}

// pointerPressDuration returns how long the primary pointer has been pressed.
func pointerPressDuration() int {
	ids := ebiten.AppendTouchIDs(nil)
	if len(ids) > 0 {
		return inpututil.TouchPressDuration(ids[0])
	}
	return inpututil.MouseButtonPressDuration(ebiten.MouseButton0)
}
