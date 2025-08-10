package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/Distortions81/EUI/eui"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	debugMode    *bool
	dumpMode     *bool
	treeMode     *bool
	currentScale float32
)

func main() {

	debugMode = flag.Bool("debug", false, "enable debug visuals")
	dumpMode = flag.Bool("dump", false, "dump cached images and exit")
	treeMode = flag.Bool("tree", false, "dump window tree and exit")
	flag.Parse()
	eui.DebugMode = *debugMode
	eui.DumpMode = *dumpMode
	eui.TreeMode = *treeMode

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

	makeShowcaseWindow()
	makeThemeSelector()

	statusOverlay := &eui.ItemData{
		ItemType: eui.ITEM_FLOW,
		FlowType: eui.FLOW_HORIZONTAL,
		Size:     eui.Point{X: 320, Y: 500},
		Position: eui.Point{X: 110, Y: 0},
		PinTo:    eui.PIN_BOTTOM_LEFT,
	}
	statusText, _ = eui.NewText()
	statusText.Size = eui.Point{X: 300, Y: 24}
	statusText.FontSize = 8
	statusOverlay.AddItem(statusText)
	eui.AddOverlayFlow(statusOverlay)

	scaleOverlay := &eui.ItemData{
		ItemType: eui.ITEM_FLOW,
		FlowType: eui.FLOW_HORIZONTAL,
		PinTo:    eui.PIN_BOTTOM_LEFT,
	}

	textItem, _ := eui.NewText()
	textItem.FontSize = 8
	textItem.Size = eui.Point{X: 80, Y: 24}
	textItem.Text = fmt.Sprintf("Scale: %2.2f", currentScale)

	minusBtn, minusEvents := eui.NewButton()
	minusBtn.Text = "-"
	minusBtn.Size = eui.Point{X: 24, Y: 24}
	minusBtn.FontSize = 8
	minusEvents.Handle = func(ev eui.UIEvent) {
		if ev.Type == eui.EventClick {
			currentScale = eui.UIScale()
			currentScale -= 0.25
			eui.SetUIScale(currentScale)
			currentScale = eui.UIScale()
			textItem.Text = fmt.Sprintf("Scale: %2.2f", currentScale)
		}
	}
	plusBtn, plusEvents := eui.NewButton()
	plusBtn.Text = "+"
	plusBtn.Size = eui.Point{X: 24, Y: 24}
	plusBtn.FontSize = 8
	plusEvents.Handle = func(ev eui.UIEvent) {
		if ev.Type == eui.EventClick {
			currentScale = eui.UIScale()
			currentScale += 0.25
			eui.SetUIScale(currentScale)
			currentScale = eui.UIScale()
			textItem.Text = fmt.Sprintf("Scale: %2.2f", currentScale)
		}
	}
	scaleOverlay.AddItem(minusBtn)
	scaleOverlay.AddItem(plusBtn)
	scaleOverlay.AddItem(textItem)

	eui.AddOverlayFlow(scaleOverlay)

	startEbiten()
}

func startEbiten() {

	// Set up ebiten
	ebiten.SetVsyncEnabled(true)
	ebiten.SetTPS(ebiten.SyncWithFPS)

	/* Set up our window */
	ebiten.SetWindowSize(1000, 1000)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	ebiten.SetWindowTitle("EUI Prototype")

	if err := ebiten.RunGameWithOptions(newGame(), &ebiten.RunGameOptions{}); err != nil {
		log.Printf("ebiten.RunGameWithOptions error: %v", err)
	}
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
	eui.Layout(outsideWidth, outsideHeight)
	return outsideWidth, outsideHeight
}

func newGame() *demoGame {
	return &demoGame{lastUpdate: time.Now()}
}
