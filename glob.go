package main

import (
	"image"
	"image/color"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var (
	windowWidth  = 512
	windowHeight = 512

	signalHandle    chan os.Signal
	mplusFaceSource *text.GoTextFaceSource
	Windows         []*WindowData
	UIScale         float32 = 1.0

	whiteImage    = ebiten.NewImage(3, 3)
	whiteSubImage = whiteImage.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image)
)

func init() {
	whiteImage.Fill(color.White)
}

const (
	minWinSizeX = 192
	minWinSizeY = 64
)

type Game struct {
}
