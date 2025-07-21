package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	eui "EUI/eui"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	debugMode    *bool
	dumpMode     *bool
	themeSel     *eui.WindowData
	signalHandle chan os.Signal
)

func main() {

	debugMode = flag.Bool("debug", false, "enable debug visuals")
	dumpMode = flag.Bool("dump", false, "dump cached images and exit")
	flag.Parse()
	eui.DebugMode = *debugMode
	eui.DumpMode = *dumpMode

	signalHandle = make(chan os.Signal, 1)
	signal.Notify(signalHandle, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	// Load a custom theme or layout from files if desired
	// if err := eui.LoadTheme("AccentDark"); err != nil {
	//         log.Printf("eui.LoadTheme error: %v", err)
	// }
	// if err := eui.LoadLayout("RoundHybrid"); err != nil {
	//         log.Printf("eui.LoadLayout error: %v", err)
	// }

	// Use your own font
	// if err := eui.EnsureFontSource(customTTF); err != nil {
	//         log.Fatal(err)
	// }

	eui.SetUIScale(1.5)

	showcase := makeShowcaseWindow()
	showcase.AddWindow(false)

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

	go startEbiten()

	<-signalHandle
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
