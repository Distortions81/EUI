package main

import (
	_ "embed"
	"flag"
	"fmt"
	"image/png"
	"log"
	"os"
	"os/signal"
	"syscall"

	eui "EUI/eui"
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	debugMode    *bool
	themeSel     *eui.WindowData
	signalHandle chan os.Signal
)

//go:embed data/fonts/NotoSans-Regular.ttf
var notoTTF []byte

func main() {

	debugMode = flag.Bool("debug", false, "enable debug visuals")
	flag.Parse()

	signalHandle = make(chan os.Signal, 1)
	signal.Notify(signalHandle, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	// load default themes
	if err := eui.LoadTheme("AccentDark"); err != nil {
		log.Printf("eui.LoadTheme error: %v", err)
	}
	if err := eui.LoadLayout("RoundHybrid"); err != nil {
		log.Printf("eui.LoadLayout error: %v", err)
	}

	// Load default font
	if err := eui.EnsureFontSource(notoTTF); err != nil {
		log.Fatal(err)
	}

	eui.SetUIScale(1.5)

	showcase := makeShowcaseWindow()
	showcase.AddWindow(false)

	flowTest := makeFlowTestWindow()
	flowTest.AddWindow(false)

	themeSel = makeThemeSelector()
	if themeSel != nil {
		themeSel.AddWindow(false)
	}

	// Add a small pinned button to toggle the themes window using an overlay flow
	overlay := &eui.ItemData{
		ItemType: eui.ITEM_FLOW,
		FlowType: eui.FLOW_HORIZONTAL,
		Size:     eui.Point{X: 84, Y: 32},
		Position: eui.Point{X: 4, Y: 4},
		PinTo:    eui.PIN_BOTTOM_RIGHT,
	}
	toggleBtn := eui.NewButton(&eui.ItemData{Text: "Themes", Size: eui.Point{X: 80, Y: 24}, FontSize: 8})
	toggleBtn.Action = func() {
		if themeSel != nil {
			if !themeSel.Open {
				themeSel.Open = true
				themeSel.BringForward()
			} else {
				themeSel.Open = false
				themeSel.ToBack()
			}
		}
	}
	overlay.AddItem(toggleBtn)
	eui.AddOverlayFlow(overlay)

	err := loadIcons()
	if err != nil {
		fmt.Printf("Error: %v\n", err.Error())
	}

	go startEbiten()

	<-signalHandle
}

func loadIcons() error {
	for _, win := range eui.Windows() {
		err := subLoadIcons(win.Contents)
		if err != nil {
			return err
		}
	}
	for _, ov := range eui.Overlays() {
		err := subLoadIcons([]*eui.ItemData{ov})
		if err != nil {
			return err
		}
	}

	return nil
}

func subLoadIcons(parent []*eui.ItemData) error {
	for _, item := range parent {
		if len(item.Tabs) > 0 {
			for _, tab := range item.Tabs {
				subLoadIcons(tab.Contents)
			}
		} else {
			subLoadIcons(item.Contents)
		}

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
	w, h := eui.ScreenSize()
	ebiten.SetWindowSize(w, h)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	ebiten.SetWindowTitle("EUI Prototype")

	if err := ebiten.RunGameWithOptions(newGame(), &ebiten.RunGameOptions{}); err != nil {
		return
	}

	signalHandle <- syscall.SIGINT
}

func newGame() *eui.Game {
	return eui.NewGame()
}
