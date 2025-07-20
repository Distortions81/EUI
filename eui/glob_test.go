//go:build test

package eui

import (
	"os"
	"time"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var (
	screenWidth  = 1024
	screenHeight = 1024

	signalHandle    chan os.Signal
	mplusFaceSource *text.GoTextFaceSource
	windows         []*windowData
	overlays        []*itemData
	activeWindow    *windowData
	focusedItem     *itemData
	uiScale         float32 = 1.0
	clickFlash              = time.Millisecond * 100

	whiteImage    interface{}
	whiteSubImage interface{}

	currentTheme     *Theme
	currentThemeName string
)

type Game struct{}
