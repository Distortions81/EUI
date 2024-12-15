package main

import (
	"bytes"
	"fmt"
	"image/color"
	"log"
	"math"
	"os"
	"os/signal"
	"syscall"

	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	defaultWindowWidth  = 1280
	defaultWindowHeight = 720
)

func main() {

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

	newWindow := WindowData{
		TitleSize:       24,
		Title:           "Test",
		Tooltip:         "Tooltip stuff here",
		Size:            XYF{X: 300, Y: 300},
		Position:        XYF{X: 32, Y: 32},
		Border:          1,
		TitleColor:      color.RGBA{R: 255, G: 255, B: 255, A: 255},
		ContentsBGColor: color.RGBA{R: 16, G: 16, B: 16, A: 255},
		BorderColor:     color.RGBA{R: 64, G: 64, B: 64, A: 255},
		SizeColor:       color.RGBA{R: 48, G: 48, B: 48, A: 255},
		Movable:         true, Closable: true, Resizable: true,
	}
	Windows = append(Windows, newWindow)

	go startEbiten()

	<-signalHandle
}

func startEbiten() {

	// Set up ebiten
	ebiten.SetVsyncEnabled(true)
	ebiten.SetTPS(ebiten.SyncWithFPS)

	/* Set up our window */
	ebiten.SetWindowSize(defaultWindowWidth, defaultWindowHeight)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("EUI Prototype")

	if err := ebiten.RunGameWithOptions(newGame(), nil); err != nil {
		return
	}

	signalHandle <- syscall.SIGINT
}

func newGame() *Game {
	return &Game{}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return defaultWindowWidth, defaultWindowHeight
}

func (g *Game) Draw(screen *ebiten.Image) {

	for _, win := range Windows {

		//Draw BG Color
		vector.DrawFilledRect(screen,
			win.Position.X, win.Position.Y,
			win.Size.X*UIScale, (win.Size.Y*UIScale)-(win.TitleSize*UIScale),
			win.ContentsBGColor, false)

		//Draw Title
		if win.TitleSize > 0 {

			textSize := ((win.TitleSize * UIScale) / 1.5)
			face := &text.GoTextFace{
				Source: mplusFaceSource,
				Size:   float64(textSize),
			}

			skipTitleText := false
			textWidth, textHeight := text.Measure(win.Title, face, 0)
			if textWidth > float64(win.Size.X*UIScale) ||
				textHeight > float64(win.TitleSize*UIScale) {
				skipTitleText = true
				//log.Print("Title text too big for title size.")
			}

			//Title text
			if !skipTitleText {
				loo := text.LayoutOptions{
					LineSpacing:    0,
					PrimaryAlign:   text.AlignStart,
					SecondaryAlign: text.AlignCenter,
				}
				tdop := ebiten.DrawImageOptions{}
				tdop.GeoM.Translate(float64(win.Position.X+((win.TitleSize*UIScale)/4)),
					float64(win.Position.Y+((win.TitleSize*UIScale)/2)))

				top := &text.DrawOptions{DrawImageOptions: tdop, LayoutOptions: loo}
				top.ColorScale.ScaleWithColor(win.TitleColor)
				text.Draw(screen, win.Title, face, top)
			} else {
				textWidth = 0
			}

			//Drag bar
			if win.Movable {
				for x := textWidth + float64((win.TitleSize*UIScale)/1.5); x < float64(win.Size.X*UIScale); x = x + float64(UIScale*5.0) {
					vector.StrokeLine(screen,
						win.Position.X+float32(x), win.Position.Y+4,
						win.Position.X+float32(x), win.Position.Y+(win.TitleSize*UIScale)-4,
						1, win.DragColor, false)
				}
			}
		}

		//Resize bar
		if win.Resizable {
			vector.StrokeLine(screen,
				win.Position.X+(win.Size.X*UIScale)-1,
				win.Position.Y+(win.Size.Y*UIScale)-(14*UIScale)-(win.TitleSize*UIScale),

				win.Position.X+(win.Size.X*UIScale)-(14*UIScale),
				win.Position.Y+(win.Size.Y*UIScale)-(win.TitleSize*UIScale)-1,
				1, win.SizeColor, true)
			vector.StrokeLine(screen,
				win.Position.X+(win.Size.X*UIScale)-1,
				win.Position.Y+(win.Size.Y*UIScale)-(10*UIScale)-(win.TitleSize*UIScale),

				win.Position.X+(win.Size.X*UIScale)-(10*UIScale),
				win.Position.Y+(win.Size.Y*UIScale)-(win.TitleSize*UIScale)-1,
				1, win.SizeColor, true)
			vector.StrokeLine(screen,
				win.Position.X+(win.Size.X*UIScale)-1,
				win.Position.Y+(win.Size.Y*UIScale)-(6*UIScale)-(win.TitleSize*UIScale),

				win.Position.X+(win.Size.X*UIScale)-(6*UIScale),
				win.Position.Y+(win.Size.Y*UIScale)-(win.TitleSize*UIScale)-1,
				1, win.SizeColor, true)
		}

		if win.Closable {
			var xpad float32 = win.TitleSize / 4
			vector.StrokeLine(screen,
				win.Position.X+(win.Size.X*UIScale)-(win.TitleSize)+xpad,
				win.Position.Y+xpad,

				win.Position.X+(win.Size.X*UIScale)-xpad,
				win.Position.Y+win.TitleSize-xpad,
				3, win.TitleColor, true)
			vector.StrokeLine(screen,
				win.Position.X+(win.Size.X*UIScale)-xpad,
				win.Position.Y+xpad,

				win.Position.X+(win.Size.X*UIScale)-(win.TitleSize)+xpad,
				win.Position.Y+win.TitleSize-xpad,
				3, win.TitleColor, true)
		}

		//Draw frames
		if win.Border > 0 {
			if win.TitleSize > 0 {
				vector.StrokeRect(screen,
					win.Position.X, win.Position.Y,
					win.Size.X*UIScale, (win.TitleSize * UIScale),
					win.Border, win.BorderColor, false)
			}
			//Window border
			vector.StrokeRect(screen,
				win.Position.X, win.Position.Y,
				win.Size.X*UIScale, (win.Size.Y*UIScale)-(win.TitleSize*UIScale),
				win.Border, win.BorderColor, false)
		}
	}

	buf := fmt.Sprintf("%4v FPS", int(math.Round(ebiten.ActualFPS())))
	ebitenutil.DebugPrintAt(screen, buf, defaultWindowWidth-55, defaultWindowHeight-18)
}

var (
	signalHandle    chan os.Signal
	mplusFaceSource *text.GoTextFaceSource
	Windows         []WindowData
	UIScale         float32 = 1.0
)

type Game struct {
}

func (g *Game) Update() error {
	return nil
}

type WindowData struct {
	Title, Tooltip string
	Size, Position XYF

	Open, Closable, Movable, Resizable, Scrollable, Maximizable, Minimizeable    bool
	ContentsBGColor, TitleBGColor, TitleColor, BorderColor, SizeColor, DragColor color.RGBA
	TitleSize, Padding, Border                                                   float32

	Contents []ItemData
}

type ItemData struct {
	Text                                 string
	Position, Size, Value                XYF
	Hovered, Activated, Checked, Enabled bool
	FlowType                             FlowType
	FlowWrap                             bool
	Padding                              float32
	Scroll                               XYF

	Color, HoverColor, ActivatedColor, DisabledColor, CheckedColor color.RGBA

	Contents []ItemData
}

type XYF struct {
	X, Y float32
}

type FlowType int

const (
	FLOW_HORIZONTAL = iota
	FLOW_VERTICAL
)
