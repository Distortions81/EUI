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
	treeMode     *bool
	signalHandle chan os.Signal
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

	signalHandle = make(chan os.Signal, 1)
	signal.Notify(signalHandle, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	currentScale = 1.5
	eui.SetUIScale(currentScale)

	win := eui.NewWindow(&eui.WindowData{
		Open: true, Resizable: true,
		Closable: true, Title: "Settings",
		AutoSize: true, Movable: true,
	})

	mainFlow := &eui.ItemData{
		ItemType: eui.ITEM_FLOW,
		Size:     win.Size,
		FlowType: eui.FLOW_VERTICAL,
	}
	win.AddItem(mainFlow)

	b1, _ := eui.NewCheckbox(&eui.ItemData{Text: "Show Item Names", Size: eui.Point{X: 150, Y: 24}})
	mainFlow.AddItem(b1)
	b2, _ := eui.NewCheckbox(&eui.ItemData{Text: "Show Legends", Size: eui.Point{X: 150, Y: 24}})
	mainFlow.AddItem(b2)
	b3, _ := eui.NewCheckbox(&eui.ItemData{Text: "Use Item Numbers", Size: eui.Point{X: 150, Y: 24}})
	mainFlow.AddItem(b3)

	IconFlow := &eui.ItemData{
		ItemType:   eui.ITEM_FLOW,
		Size:       eui.Point{X: 150, Y: 24},
		FlowType:   eui.FLOW_HORIZONTAL,
		Fixed:      true,
		Scrollable: false,
	}
	mainFlow.AddItem(IconFlow)

	t1, _ := eui.NewText(&eui.ItemData{Text: "Icon Size:", FontSize: 9, Size: eui.Point{X: 50, Y: 24}})
	IconFlow.AddItem(t1)
	t2, _ := eui.NewButton(&eui.ItemData{Text: "-", Size: eui.Point{X: 16, Y: 16}})
	IconFlow.AddItem(t2)
	t3, _ := eui.NewButton(&eui.ItemData{Text: "+", Size: eui.Point{X: 16, Y: 16}})
	IconFlow.AddItem(t3)

	ScaleFlow := &eui.ItemData{
		ItemType:   eui.ITEM_FLOW,
		Size:       eui.Point{X: 150, Y: 24},
		FlowType:   eui.FLOW_HORIZONTAL,
		Fixed:      true,
		Scrollable: false,
	}
	mainFlow.AddItem(ScaleFlow)

	u1, _ := eui.NewText(&eui.ItemData{Text: "UI Scale:", FontSize: 9, Size: eui.Point{X: 50, Y: 24}})
	ScaleFlow.AddItem(u1)
	u2, _ := eui.NewButton(&eui.ItemData{Text: "-", Size: eui.Point{X: 16, Y: 16}})
	ScaleFlow.AddItem(u2)
	u3, _ := eui.NewButton(&eui.ItemData{Text: "+", Size: eui.Point{X: 16, Y: 16}})
	ScaleFlow.AddItem(u3)

	c1, _ := eui.NewCheckbox(&eui.ItemData{Text: "Textures", Size: eui.Point{X: 150, Y: 24}})
	mainFlow.AddItem(c1)
	c2, _ := eui.NewCheckbox(&eui.ItemData{Text: "VSync", Size: eui.Point{X: 150, Y: 24}})
	mainFlow.AddItem(c2)
	c3, _ := eui.NewCheckbox(&eui.ItemData{Text: "Power Saver", Size: eui.Point{X: 150, Y: 24}})
	mainFlow.AddItem(c3)
	c4, _ := eui.NewCheckbox(&eui.ItemData{Text: "Linear Filtering", Size: eui.Point{X: 150, Y: 24}})
	mainFlow.AddItem(c4)
	c5, _ := eui.NewCheckbox(&eui.ItemData{Text: "HiDPI", Size: eui.Point{X: 150, Y: 24}})
	mainFlow.AddItem(c5)

	cFPS := fmt.Sprintf("FPS: %2.2f", ebiten.ActualFPS())
	tt1, _ := eui.NewText(&eui.ItemData{Text: cFPS, FontSize: 9, Size: eui.Point{X: 150, Y: 12}})
	mainFlow.AddItem(tt1)

	vers := fmt.Sprintf("Version: %v", "v0.0.9-012345")
	tt2, _ := eui.NewText(&eui.ItemData{Text: vers, FontSize: 9, Size: eui.Point{X: 150, Y: 12}})
	mainFlow.AddItem(tt2)

	win.AddWindow(false)
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
