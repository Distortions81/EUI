package main

import (
	"bytes"
	"image/color"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2"
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
		TitleSize:       32,
		Title:           "Test Window Title",
		Tooltip:         "Tooltip stuff here",
		Size:            XYF{X: 300, Y: 300},
		Position:        XYF{X: 32, Y: 32},
		Border:          1,
		ContentsBGColor: color.RGBA{R: 16, G: 16, B: 16, A: 255},
		BorderColor:     color.RGBA{R: 64, G: 64, B: 64, A: 255},
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
		vector.DrawFilledRect(screen, win.Position.X, win.Position.Y, win.Size.X, win.Size.Y-win.TitleSize, win.ContentsBGColor, false)

		//Draw Title
		if win.Title != "" {

			textSize := (win.TitleSize - 10.0)
			face := &text.GoTextFace{
				Source: mplusFaceSource,
				Size:   float64(textSize),
			}

			skipTitleText := false
			textWidth, textHeight := text.Measure(win.Title, face, 0)
			if textWidth > float64(win.Size.X) ||
				textHeight > float64(win.TitleSize) {
				skipTitleText = true
				log.Fatal("Title text too big for title size.")
			}

			textSpacer := (textSize / 5.0)

			if win.Movable {
				for x := int(textWidth + float64(textSpacer*2)); x < int(win.Size.X-textSpacer); x++ {
					if x%6 == 0 {
						vector.StrokeLine(screen,
							win.Position.X+float32(x), win.Position.Y+4,
							win.Position.X+float32(x), win.Position.Y+win.TitleSize-4,
							1, win.BorderColor, false)
					}
				}
			}

			//Title text
			if !skipTitleText {
				loo := text.LayoutOptions{
					LineSpacing:    0,
					PrimaryAlign:   text.AlignStart,
					SecondaryAlign: text.AlignStart,
				}
				tdop := ebiten.DrawImageOptions{}
				tdop.GeoM.Translate(float64(win.Position.X+textSpacer), float64(win.Position.Y))

				top := &text.DrawOptions{DrawImageOptions: tdop, LayoutOptions: loo}
				text.Draw(screen, win.Title, face, top)
			}
		}

		//Draw frames
		if win.Border > 0 {
			if win.TitleSize > 0 {
				vector.StrokeRect(screen, win.Position.X, win.Position.Y, win.Size.X, win.TitleSize, win.Border, win.BorderColor, false)
			}
			//Window border
			vector.StrokeRect(screen, win.Position.X, win.Position.Y, win.Size.X, win.Size.Y-win.TitleSize, win.Border, win.BorderColor, false)
		}
	}
}

var (
	signalHandle    chan os.Signal
	mplusFaceSource *text.GoTextFaceSource
	Windows         []WindowData
)

type Game struct {
}

func (g *Game) Update() error {
	return nil
}

type WindowData struct {
	Title, Tooltip string
	Size, Position XYF

	Open, Closable, Movable, Resizable, Scrollable, Maximizable, Minimizeable bool
	ContentsBGColor, TitleBGColor, TitleColor, BorderColor                    color.RGBA
	TitleSize, Padding, Border                                                float32

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
