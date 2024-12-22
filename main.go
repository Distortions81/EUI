package main

import (
	"bytes"
	"flag"
	"image/color"
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

	//UIScale = 1

	//Done button
	newButton := DefaultButton
	newButton.Text = "Generate"
	newButton.Position = Point{
		X: 300 - 128 - 16,
		Y: 300 - 24 - 64 - 16}
	newButton.Size = Point{X: 128, Y: 64}

	//Scaleup button
	newScaleup := DefaultButton
	newScaleup.Text = "Scale Up"
	newScaleup.Position = Point{
		X: 16,
		Y: 24 + 128}
	newScaleup.Size = Point{X: 128, Y: 24}
	newScaleup.FontSize = 18

	//Scaledown button
	newScaledown := DefaultButton
	newScaledown.Text = "Scale Down"
	newScaledown.Position = Point{
		X: 16,
		Y: 24 + 128 + 32}
	newScaledown.Size = Point{X: 128, Y: 24}
	newScaledown.FontSize = 18

	//Text
	newText := ItemData{}
	newText.ItemType = ITEM_TEXT
	newText.Text = "Click 'generate' to\ngenerate a new code."
	newText.FontSize = 24
	newText.Position = Point{
		X: 16,
		Y: 24 + 16}
	newText.Size = Point{X: 128, Y: 128}
	newText.TextColor = color.RGBA{R: 255, G: 255, B: 255, A: 255}

	newWindow := NewWindow(
		&WindowData{
			TitleSize: 24,
			Title:     "Test Window",
			Size:      Point{X: 300, Y: 300},
			Position:  Point{X: 32, Y: 32},
			Contents: []*ItemData{
				&newButton, &newText, &newScaleup, &newScaledown},
		})
	newWindow.AddWindow()

	newButton.Action = func() {
		newButton.Text = "Okay"
		newText.Text = "Secret code: 1234"
		newButton.Action = func() {
			newWindow.Open = false
		}
	}

	newScaleup.Action = func() {
		if UIScale < 8 {
			UIScale += 0.1
		}
	}

	newScaledown.Action = func() {
		if UIScale > 0.2 {
			UIScale -= 0.1
		}
	}

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
