package main

import (
	"flag"
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
		FlowType: eui.FLOW_HORIZONTAL,
	}
	win.AddItem(mainFlow)

	b1, _ := eui.NewButton(&eui.ItemData{Text: "Setting A", Size: eui.Point{X: 100, Y: 24}})
	mainFlow.AddItem(b1)
	b2, _ := eui.NewButton(&eui.ItemData{Text: "Setting B", Size: eui.Point{X: 100, Y: 24}})
	mainFlow.AddItem(b2)
	b3, _ := eui.NewButton(&eui.ItemData{Text: "Setting C", Size: eui.Point{X: 100, Y: 24}})
	mainFlow.AddItem(b3)

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
