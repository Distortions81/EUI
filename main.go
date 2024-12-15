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
	"github.com/hajimehoshi/ebiten/v2/inpututil"
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
		TitleSize:  24,
		Title:      "Test",
		Tooltip:    "Tooltip stuff here",
		Size:       Magnatude{X: 300, Y: 300},
		Position:   Point{X: 32, Y: 32},
		Border:     1,
		TitleColor: color.RGBA{R: 255, G: 255, B: 255, A: 255},

		BorderColor: color.RGBA{R: 64, G: 64, B: 64, A: 255},

		SizeColor:  color.RGBA{R: 48, G: 48, B: 48, A: 255},
		DragColor:  color.RGBA{R: 48, G: 48, B: 48, A: 255},
		HoverColor: color.RGBA{R: 80, G: 80, B: 80, A: 255},

		ContentsBGColor: color.RGBA{R: 16, G: 16, B: 16, A: 255},

		Movable: true, Closable: true, Resizable: true, Open: true,
	}
	Windows = append(Windows, newWindow)

	go startEbiten()

	/*
		go func() {
		UIScale = 0.1
			time.Sleep(time.Second * 5)
			for x := 0; x < 100; x++ {
				time.Sleep(time.Second)
				UIScale = UIScale + 0.1
			}
		}()
	*/

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
		if !win.Open {
			continue
		}

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

			var buttonsWidth float32 = 0
			if win.Closable {
				var xpad float32 = (win.TitleSize * UIScale) / 4.0
				vector.StrokeLine(screen,
					win.Position.X+(win.Size.X*UIScale)-(win.TitleSize*UIScale)+xpad,
					win.Position.Y+xpad,

					win.Position.X+(win.Size.X*UIScale)-xpad,
					win.Position.Y+(win.TitleSize*UIScale)-xpad,
					3*UIScale, win.TitleColor, true)
				vector.StrokeLine(screen,
					win.Position.X+(win.Size.X*UIScale)-xpad,
					win.Position.Y+xpad,

					win.Position.X+(win.Size.X*UIScale)-(win.TitleSize*UIScale)+xpad,
					win.Position.Y+(win.TitleSize*UIScale)-xpad,
					3*UIScale, win.TitleColor, true)

				buttonsWidth += (win.TitleSize * UIScale)
			}

			//Drag bar
			if win.Movable {
				dpad := (win.TitleSize * UIScale) / 5
				for x := textWidth + float64((win.TitleSize*UIScale)/1.5); x < float64((win.Size.X*UIScale)-buttonsWidth); x = x + float64(UIScale*5.0) {
					vector.StrokeLine(screen,
						win.Position.X+float32(x), win.Position.Y+dpad,
						win.Position.X+float32(x), win.Position.Y+(win.TitleSize*UIScale)-dpad,
						1, win.DragColor, false)
				}
			}
		}

		//Draw frames
		if win.Border > 0 {
			FrameColor := win.BorderColor
			if win.Hovered {
				FrameColor = win.HoverColor
			}
			if win.TitleSize > 0 {
				vector.StrokeRect(screen,
					win.Position.X, win.Position.Y,
					win.Size.X*UIScale, (win.TitleSize * UIScale),
					win.Border, FrameColor, false)
			}
			//Window border
			vector.StrokeRect(screen,
				win.Position.X, win.Position.Y,
				win.Size.X*UIScale, (win.Size.Y*UIScale)-(win.TitleSize*UIScale),
				win.Border, FrameColor, false)
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
	mx, my := ebiten.CursorPosition()
	mpos := Point{X: float32(mx), Y: float32(my)}

	click := inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0)
	//clickHeld := inpututil.MouseButtonPressDuration(ebiten.MouseButton0)

	//altClick := inpututil.IsMouseButtonJustPressed(ebiten.MouseButton1)
	//altClickHeld := inpututil.MouseButtonPressDuration(ebiten.MouseButton1)

	//Check all windows
	for w, win := range Windows {
		if !win.Open {
			continue
		}

		winRect := win.GetRect()
		if winRect.ContainsPoint(mpos) {
			Windows[w].Hovered = true

			//Window contents
			if win.Dumb {
				continue
			}
			for i, item := range win.Contents {
				if item.Rect.ContainsPoint(mpos) {
					if click {
						win.Contents[i].Activated = true
						win.Contents[i].Hovered = false
					} else {
						win.Contents[i].Activated = false
						win.Contents[i].Hovered = true
					}
				} else {
					win.Contents[i].Activated = false
					win.Contents[i].Hovered = false
				}
			}
		} else {
			Windows[w].Hovered = false
		}
	}

	return nil
}

func (rect Rect) ContainsPoint(b Point) bool {
	return b.X >= rect.X0 && b.Y >= rect.Y0 &&
		b.X <= rect.X1 && b.Y <= rect.Y1
}

func (win WindowData) GetRect() Rect {
	return Rect{
		X0: win.Position.X, Y0: win.Position.Y,
		X1: win.Position.X + (win.Size.X * UIScale),
		Y1: win.Position.Y + (win.Size.Y * UIScale) + (win.TitleSize * UIScale),
	}
}

type WindowData struct {
	Title, Tooltip string
	Position       Point
	Size           Magnatude

	Open, Hovered, Dumb                                                                      bool
	Closable, Movable, Resizable, Scrollable, Maximizable, Minimizeable                      bool
	ContentsBGColor, TitleBGColor, TitleColor, BorderColor, SizeColor, DragColor, HoverColor color.RGBA
	TitleSize, Padding, Border                                                               float32

	Contents []ItemData
}

type ItemData struct {
	Text     string
	Position Point
	Size     Magnatude
	Rect     Rect

	Value float32

	Hovered, Activated, Checked, Enabled bool
	FlowType                             FlowType
	FlowWrap                             bool
	Padding                              float32
	Scroll                               Point

	Color, HoverColor, ActivatedColor, DisabledColor, CheckedColor color.RGBA

	Contents []ItemData
}

type Point struct {
	X, Y float32
}

type Magnatude Point

type FlowType int

type Rect struct {
	X0, Y0, X1, Y1 float32
}

const (
	FLOW_HORIZONTAL = iota
	FLOW_VERTICAL
)
