//go:build !test

package main

import (
	"image"
	"image/color"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var (
	screenWidth  = 1024
	screenHeight = 1024

	signalHandle     chan os.Signal
	mplusFaceSource  *text.GoTextFaceSource
	windows          []*windowData
	overlays         []*itemData
	activeWindow     *windowData
	focusedItem      *itemData
	uiScale          float32 = 1.0
	currentTheme     *Theme
	currentThemeName string
	clickFlash       = time.Millisecond * 100

	whiteImage    = ebiten.NewImage(3, 3)
	whiteSubImage = whiteImage.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image)
)

func init() {
	whiteImage.Fill(color.White)
}

// constants moved to const.go

type Game struct {
}
