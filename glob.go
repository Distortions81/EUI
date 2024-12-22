package main

import (
	"os"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var (
	windowWidth  = 1280
	windowHeight = 720

	signalHandle    chan os.Signal
	mplusFaceSource *text.GoTextFaceSource
	Windows         []*WindowData
	UIScale         float32 = 1.0
)

const (
	minWinSizeX = 192
	minWinSizeY = 64
)

type Game struct {
}
