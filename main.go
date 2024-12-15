package main

import (
	"bytes"
	"image/color"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

func main() {

	signalHandle = make(chan os.Signal, 1)
	signal.Notify(signalHandle, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	//Load default font
	if mplusFaceSource == nil {
		s, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
		if err != nil {
			log.Fatal(err)
		}
		mplusFaceSource = s
	}

	newWindow := WindowData{
		TitleSize:  24,
		Title:      "Test",
		Tooltip:    "Tooltip stuff here",
		Size:       Magnatude{X: 300, Y: 300},
		Position:   Point{X: 32, Y: 32},
		Border:     1,
		TitleColor: color.RGBA{R: 255, G: 255, B: 255, A: 255},

		BorderColor: color.RGBA{R: 64, G: 64, B: 64, A: 255},

		SizeColor:  color.RGBA{R: 48, G: 48, B: 48, A: 255},
		DragColor:  color.RGBA{R: 48, G: 48, B: 48, A: 255},
		HoverColor: color.RGBA{R: 80, G: 80, B: 80, A: 255},

		ContentsBGColor: color.RGBA{R: 16, G: 16, B: 16, A: 255},

		Movable: true, Closable: true, Resizable: true, Open: true,
	}
	Windows = append(Windows, newWindow)

	go startEbiten()

	/*
		go func() {
		UIScale = 0.1
			time.Sleep(time.Second * 5)
			for x := 0; x < 100; x++ {
				time.Sleep(time.Second)
				UIScale = UIScale + 0.1
			}
		}()
	*/

	<-signalHandle
}

func startEbiten() {

	// Set up ebiten
	ebiten.SetVsyncEnabled(true)
	ebiten.SetTPS(ebiten.SyncWithFPS)

	/* Set up our window */
	ebiten.SetWindowSize(defaultWindowWidth, defaultWindowHeight)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("EUI Prototype")

	if err := ebiten.RunGameWithOptions(newGame(), nil); err != nil {
		return
	}

	signalHandle <- syscall.SIGINT
}

func newGame() *Game {
	return &Game{}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return defaultWindowWidth, defaultWindowHeight
}
