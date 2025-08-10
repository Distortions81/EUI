package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Distortions81/EUI/eui"

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

	win := eui.NewWindow()
	win.Open = true
	win.Resizable = true
	win.Closable = true
	win.Title = "Settings"
	win.AutoSize = true
	win.Movable = true

	mainFlow := &eui.ItemData{
		ItemType: eui.ITEM_FLOW,
		FlowType: eui.FLOW_VERTICAL,
	}

	b1, _ := eui.NewCheckbox()
	b1.Text = "Show Item Names"
	b1.Size = eui.Point{X: 150, Y: 24}
	mainFlow.AddItem(b1)
	b2, _ := eui.NewCheckbox()
	b2.Text = "Show Legends"
	b2.Size = eui.Point{X: 150, Y: 24}
	mainFlow.AddItem(b2)
	b3, _ := eui.NewCheckbox()
	b3.Text = "Use Item Numbers"
	b3.Size = eui.Point{X: 150, Y: 24}
	mainFlow.AddItem(b3)

	IconFlow := &eui.ItemData{
		ItemType: eui.ITEM_FLOW,
		FlowType: eui.FLOW_HORIZONTAL,
	}
	t1, _ := eui.NewText()
	t1.Text = "Icon Size:"
	t1.FontSize = 9
	t1.Size = eui.Point{X: 50, Y: 24}
	IconFlow.AddItem(t1)
	t2, _ := eui.NewButton()
	t2.Text = "-"
	t2.Size = eui.Point{X: 16, Y: 16}
	IconFlow.AddItem(t2)
	t3, _ := eui.NewButton()
	t3.Text = "+"
	t3.Size = eui.Point{X: 16, Y: 16}
	IconFlow.AddItem(t3)
	mainFlow.AddItem(IconFlow)

	ScaleFlow := &eui.ItemData{
		ItemType: eui.ITEM_FLOW,
		FlowType: eui.FLOW_HORIZONTAL,
	}
	u1, _ := eui.NewText()
	u1.Text = "UI Scale:"
	u1.FontSize = 9
	u1.Size = eui.Point{X: 50, Y: 24}
	ScaleFlow.AddItem(u1)
	u2, _ := eui.NewButton()
	u2.Text = "-"
	u2.Size = eui.Point{X: 16, Y: 16}
	ScaleFlow.AddItem(u2)
	u3, _ := eui.NewButton()
	u3.Text = "+"
	u3.Size = eui.Point{X: 16, Y: 16}
	ScaleFlow.AddItem(u3)
	mainFlow.AddItem(ScaleFlow)

	c1, _ := eui.NewCheckbox()
	c1.Text = "Textures"
	c1.Size = eui.Point{X: 150, Y: 24}
	mainFlow.AddItem(c1)
	c2, _ := eui.NewCheckbox()
	c2.Text = "VSync"
	c2.Size = eui.Point{X: 150, Y: 24}
	mainFlow.AddItem(c2)
	c3, _ := eui.NewCheckbox()
	c3.Text = "Power Saver"
	c3.Size = eui.Point{X: 150, Y: 24}
	mainFlow.AddItem(c3)
	c4, _ := eui.NewCheckbox()
	c4.Text = "Linear Filtering"
	c4.Size = eui.Point{X: 150, Y: 24}
	mainFlow.AddItem(c4)
	c5, _ := eui.NewCheckbox()
	c5.Text = "HiDPI"
	c5.Size = eui.Point{X: 150, Y: 24}
	mainFlow.AddItem(c5)

	cFPS := fmt.Sprintf("FPS: %2.2f", ebiten.ActualFPS())
	tt1, _ := eui.NewText()
	tt1.Text = cFPS
	tt1.FontSize = 9
	tt1.Size = eui.Point{X: 150, Y: 12}
	mainFlow.AddItem(tt1)

	vers := fmt.Sprintf("Version: %v", "v0.0.9-012345")
	tt2, _ := eui.NewText()
	tt2.Text = vers
	tt2.FontSize = 9
	tt2.Size = eui.Point{X: 150, Y: 12}
	mainFlow.AddItem(tt2)

	win.AddItem(mainFlow)
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

	defer func() {
		signalHandle <- syscall.SIGINT
	}()

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
	//Your layout handling code here
	return eui.Layout(outsideWidth, outsideHeight)
}

func newGame() *demoGame {
	return &demoGame{lastUpdate: time.Now()}
}
