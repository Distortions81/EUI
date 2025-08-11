//go:build test

package eui

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

	signalHandle    chan os.Signal
	mplusFaceSource *text.GoTextFaceSource
	windows         []*windowData
	overlays        []*itemData
	activeWindow    *windowData
	focusedItem     *itemData
	hoveredItem     *itemData
	uiScale         float32 = 1.0
	clickFlash              = time.Millisecond * 100

	// Debug and dump flags used by rendering logic
	DebugMode bool
	DumpMode  bool
	TreeMode  bool

	whiteImage    = ebiten.NewImage(3, 3)
	whiteSubImage = whiteImage.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image)

	currentTheme     *Theme
	currentThemeName string

	notoTTF = defaultTTF

	MinWinSizeX float32 = float32(minWinSizeX) / float32(screenWidth)
	MinWinSizeY float32 = float32(minWinSizeY) / float32(screenHeight)
)

func init() {
	whiteImage.Fill(color.White)
}

func normX(px float32) float32     { return px / float32(screenWidth) }
func normY(py float32) float32     { return py / float32(screenHeight) }
func normPoint(x, y float32) point { return point{X: normX(x), Y: normY(y)} }

type Game struct{}
