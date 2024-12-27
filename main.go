package main

import (
	"bytes"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var debugMode *bool

func main() {

	debugMode = flag.Bool("debug", false, "enable debug visuals")
	flag.Parse()

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

	SetUIScale(1)

	newWindow := makeTestWindow()
	newWindow.AddWindow()

	newWindow = makeTestWindow()
	newWindow.Position.X += 64
	newWindow.Position.Y += 64
	newWindow.AddWindow()

	go startEbiten()

	<-signalHandle
}

func startEbiten() {

	// Set up ebiten
	ebiten.SetVsyncEnabled(true)
	ebiten.SetTPS(ebiten.SyncWithFPS)

	/* Set up our window */
	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	ebiten.SetWindowTitle("EUI Prototype")

	if err := ebiten.RunGameWithOptions(newGame(), &ebiten.RunGameOptions{}); err != nil {
		return
	}

	signalHandle <- syscall.SIGINT
}

func newGame() *Game {
	return &Game{}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	windowWidth, windowHeight = outsideWidth, outsideHeight
	return outsideWidth, outsideHeight
}
