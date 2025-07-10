//go:build test

package main

import (
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"os"
	"time"
)

var (
	screenWidth  = 1024
	screenHeight = 1024

	signalHandle    chan os.Signal
	mplusFaceSource *text.GoTextFaceSource
	windows         []*windowData
	activeWindow    *windowData
	focusedItem     *itemData
	uiScale         float32 = 1.0
	clickFlash              = time.Millisecond * 100

	whiteImage    interface{}
	whiteSubImage interface{}
)

type Game struct{}

const (
	minWinSizeX = 192
	minWinSizeY = 64
)
