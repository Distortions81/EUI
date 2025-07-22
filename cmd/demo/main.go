package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	eui "EUI/eui"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	debugMode    *bool
	dumpMode     *bool
	themeSel     *eui.WindowData
	signalHandle chan os.Signal
	currentScale float32
)

func main() {

	debugMode = flag.Bool("debug", false, "enable debug visuals")
	dumpMode = flag.Bool("dump", false, "dump cached images and exit")
	flag.Parse()
	eui.DebugMode = *debugMode
	eui.DumpMode = *dumpMode

	signalHandle = make(chan os.Signal, 1)
	signal.Notify(signalHandle, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	// Load a custom theme or style from files if desired
	// if err := eui.LoadTheme("AccentDark"); err != nil {
	//         log.Printf("eui.LoadTheme error: %v", err)
	// }
	// if err := eui.LoadStyle("RoundHybrid"); err != nil {
	//         log.Printf("eui.LoadStyle error: %v", err)
	// }

	// Use your own font
	// if err := eui.EnsureFontSource(customTTF); err != nil {
	//         log.Fatal(err)
	// }

	currentScale = 1.5
	eui.SetUIScale(currentScale)

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
		Size:     eui.Point{X: 80, Y: 24},
		Position: eui.Point{X: 4, Y: 4},
		PinTo:    eui.PIN_BOTTOM_RIGHT,
	}
	toggleBtn, toggleEvents := eui.NewButton(&eui.ItemData{Text: "Themes", Size: eui.Point{X: 80, Y: 24}, FontSize: 8})
	toggleEvents.Handle = func(ev eui.UIEvent) {
		if ev.Type != eui.EventClick || themeSel == nil {
			return
		}

		wins := eui.Windows()
		if themeSel.Open {
			if len(wins) == 0 || wins[len(wins)-1] != themeSel {
				themeSel.BringForward()
				return
			}
			themeSel.Open = false
			themeSel.ToBack()
			return
		}

		themeSel.MarkOpen()
	}
	overlay.AddItem(toggleBtn)
	eui.AddOverlayFlow(overlay)

	statusOverlay := &eui.ItemData{
		ItemType: eui.ITEM_FLOW,
		FlowType: eui.FLOW_HORIZONTAL,
		Size:     eui.Point{X: 320, Y: 24},
		Position: eui.Point{X: 110, Y: 0},
		PinTo:    eui.PIN_BOTTOM_LEFT,
	}
	statusText, _ = eui.NewText(&eui.ItemData{Size: eui.Point{X: 300, Y: 24}, FontSize: 8})
	statusOverlay.AddItem(statusText)
	eui.AddOverlayFlow(statusOverlay)

	scaleOverlay := &eui.ItemData{
		ItemType: eui.ITEM_FLOW,
		FlowType: eui.FLOW_HORIZONTAL,
		PinTo:    eui.PIN_BOTTOM_LEFT,
	}

	textItem, _ := eui.NewText(&eui.ItemData{FontSize: 8, Size: eui.Point{X: 80, Y: 24}})
	textItem.Text = fmt.Sprintf("Scale: %2.2f", currentScale)

	minusBtn, minusEvents := eui.NewButton(&eui.ItemData{Text: "-", Size: eui.Point{X: 24, Y: 24}, FontSize: 8})
	minusEvents.Handle = func(ev eui.UIEvent) {
		if ev.Type == eui.EventClick {
			currentScale = eui.UIScale()
			currentScale -= 0.25
			if currentScale < 0.25 {
				currentScale = 0.25
			}
			eui.SetUIScale(currentScale)
			textItem.Text = fmt.Sprintf("Scale: %2.2f", currentScale)
		}
	}
	plusBtn, plusEvents := eui.NewButton(&eui.ItemData{Text: "+", Size: eui.Point{X: 24, Y: 24}, FontSize: 8})
	plusEvents.Handle = func(ev eui.UIEvent) {
		if ev.Type == eui.EventClick {
			currentScale = eui.UIScale()
			currentScale += 0.25
			if currentScale > 4.0 {
				currentScale = 4
			}
			eui.SetUIScale(currentScale)
			textItem.Text = fmt.Sprintf("Scale: %2.2f", currentScale)
		}
	}
	scaleOverlay.AddItem(minusBtn)
	scaleOverlay.AddItem(plusBtn)
	scaleOverlay.AddItem(textItem)

	eui.AddOverlayFlow(scaleOverlay)

	go startEbiten()

	<-signalHandle
}

func startEbiten() {

	// Set up ebiten
	ebiten.SetVsyncEnabled(true)
	ebiten.SetTPS(ebiten.SyncWithFPS)

	/* Set up our window */
	w, h := eui.ScreenSize()
	scale := ebiten.Monitor().DeviceScaleFactor()
	if scale <= 0 {
		scale = 1
	}
	ebiten.SetWindowSize(int(float64(w)/scale), int(float64(h)/scale))
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	ebiten.SetWindowTitle("EUI Prototype")

	if err := ebiten.RunGameWithOptions(newGame(), &ebiten.RunGameOptions{}); err != nil {
		return
	}

	signalHandle <- syscall.SIGINT
}

type demoGame struct {
	lastUpdate time.Time
}

func (g *demoGame) Update() error {
	return eui.Update()
}
func (g *demoGame) Draw(screen *ebiten.Image) {
	//Your draw code here
	eui.Draw(screen)
}
func (g *demoGame) Layout(outsideWidth, outsideHeight int) (int, int) {
	//Your layout handling code here
	return eui.Layout(outsideWidth, outsideHeight)
}

func newGame() *demoGame {
	return &demoGame{lastUpdate: time.Now()}
}
