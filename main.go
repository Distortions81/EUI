package main

import (
	"bytes"
	"flag"
	"fmt"
	"image/png"
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

	SetUIScale(2)
	newWindow := makeTestWindow()
	newWindow.AddWindow(false)

	newWindow2 := makeTestWindow()
	newWindow2.Position.X += 192
	newWindow2.Position.Y += 192
	newWindow2.AddWindow(false)

	err := loadIcons()
	if err != nil {
		fmt.Printf("Error: %v\n", err.Error())
	}

	go startEbiten()

	<-signalHandle
}

func loadIcons() error {
	for _, win := range windows {
		err := subLoadIcons(win.Contents)
		if err != nil {
			return err
		}
	}

	return nil
}

func subLoadIcons(parent []*itemData) error {
	for _, item := range parent {
		subLoadIcons(item.Contents)

		if item.ImageName != "" {
			image, err := loadImage(item.ImageName)
			if err != nil {
				return err
			}
			item.Image = image
		}
	}

	return nil
}

func loadImage(name string) (*ebiten.Image, error) {
	fileData, err := os.OpenFile("data/"+name+".png", os.O_RDONLY, 0755)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Loaded %v.png\n", name)

	png, err := png.Decode(fileData)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Decoded %v.png\n", name)

	image := ebiten.NewImageFromImage(png)
	fmt.Printf("Image %v.png added to ebiten\n", name)

	return image, nil
}

func startEbiten() {

	// Set up ebiten
	ebiten.SetVsyncEnabled(true)
	ebiten.SetTPS(ebiten.SyncWithFPS)

	/* Set up our window */
	ebiten.SetWindowSize(screenWidth, screenHeight)
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
	screenWidth, screenHeight = outsideWidth, outsideHeight
	return outsideWidth, outsideHeight
}
