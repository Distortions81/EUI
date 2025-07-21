# EUI API Reference

This document lists all exported functions, types and constants in the
`eui` package. See the README for a short overview of the project.

## Setup

The demo and tests rely on a few system packages in addition to Go. Run the
helper script once to install them (Ubuntu/Debian based distributions):

```sh
./scripts/setup.sh
```

After dependencies are installed you can vet and build the project:

```sh
go vet ./...
go build ./...
```

Format modified Go files before committing:

```sh
gofmt -w <files>
```

package eui // import "EUI/eui"


CONSTANTS

const (
	FLOW_HORIZONTAL = iota
	FLOW_VERTICAL

	FLOW_HORIZONTAL_REV
	FLOW_VERTICAL_REV
)
const (
	ALIGN_NONE = iota
	ALIGN_LEFT
	ALIGN_CENTER
	ALIGN_RIGHT
)
const (
	PIN_TOP_LEFT = iota
	PIN_TOP_CENTER
	PIN_TOP_RIGHT

	PIN_MID_LEFT
	PIN_MID_CENTER
	PIN_MID_RIGHT

	PIN_BOTTOM_LEFT
	PIN_BOTTOM_CENTER
	PIN_BOTTOM_RIGHT
)
const (
	PART_NONE = iota

	PART_BAR
	PART_CLOSE

	PART_TOP
	PART_RIGHT
	PART_BOTTOM
	PART_LEFT

	PART_TOP_RIGHT
	PART_BOTTOM_RIGHT
	PART_BOTTOM_LEFT
	PART_TOP_LEFT

	PART_SCROLL_V
	PART_SCROLL_H
)
const (
	ITEM_NONE = iota
	ITEM_FLOW
	ITEM_TEXT
	ITEM_BUTTON
	ITEM_CHECKBOX
	ITEM_RADIO
	ITEM_INPUT
	ITEM_SLIDER
	ITEM_DROPDOWN
	ITEM_COLORWHEEL
)

VARIABLES

var (
	ColorWhite   = NewColor(255, 255, 255, 255)
	ColorBlack   = NewColor(0, 0, 0, 255)
	ColorRed     = NewColor(203, 67, 53, 255)
	ColorGreen   = NewColor(40, 180, 99, 255)
	ColorBlue    = NewColor(41, 128, 185, 255)
	ColorYellow  = NewColor(244, 208, 63, 255)
	ColorGray    = NewColor(128, 128, 128, 255)
	ColorOrange  = NewColor(243, 156, 18, 255)
	ColorPink    = NewColor(255, 151, 197, 255)
	ColorPurple  = NewColor(165, 105, 189, 255)
	ColorSilver  = NewColor(209, 209, 209, 255)
	ColorTeal    = NewColor(64, 199, 178, 255)
	ColorMaroon  = NewColor(199, 54, 103, 255)
	ColorNavy    = NewColor(99, 114, 166, 255)
	ColorOlive   = NewColor(134, 166, 99, 255)
	ColorLime    = NewColor(206, 231, 114, 255)
	ColorFuchsia = NewColor(209, 114, 231, 255)
	ColorAqua    = NewColor(114, 228, 231, 255)
	ColorBrown   = NewColor(176, 116, 78, 255)
	ColorRust    = NewColor(210, 113, 52, 255)

	ColorLightRed     = NewColor(255, 83, 83, 255)
	ColorLightGreen   = NewColor(170, 255, 159, 255)
	ColorLightBlue    = NewColor(159, 186, 255, 255)
	ColorLightYellow  = NewColor(255, 251, 159, 255)
	ColorLightGray    = NewColor(236, 236, 236, 255)
	ColorLightOrange  = NewColor(252, 213, 134, 255)
	ColorLightPink    = NewColor(254, 163, 182, 255)
	ColorLightPurple  = NewColor(254, 163, 245, 255)
	ColorLightSilver  = NewColor(228, 228, 228, 255)
	ColorLightTeal    = NewColor(152, 221, 210, 255)
	ColorLightMaroon  = NewColor(215, 124, 143, 255)
	ColorLightNavy    = NewColor(128, 152, 197, 255)
	ColorLightOlive   = NewColor(186, 228, 144, 255)
	ColorLightLime    = NewColor(219, 243, 153, 255)
	ColorLightFuchsia = NewColor(239, 196, 253, 255)
	ColorLightAqua    = NewColor(196, 246, 253, 255)

	ColorDarkRed      = NewColor(146, 22, 22, 255)
	ColorDarkGreen    = NewColor(22, 146, 24, 255)
	ColorDarkBlue     = NewColor(22, 98, 146, 255)
	ColorDarkYellow   = NewColor(139, 146, 22, 255)
	ColorDarkGray     = NewColor(111, 111, 111, 255)
	ColorCharcoal     = NewColor(16, 16, 16, 255)
	ColorCharcoalSemi = NewColor(16, 16, 16, 128)
	ColorDarkOrange   = NewColor(175, 117, 32, 255)
	ColorDarkPink     = NewColor(128, 64, 64, 255)
	ColorDarkPurple   = NewColor(137, 32, 175, 255)
	ColorDarkSilver   = NewColor(162, 162, 162, 255)
	ColorDarkTeal     = NewColor(27, 110, 86, 255)
	ColorDarkMaroon   = NewColor(110, 27, 55, 255)
	ColorDarkNavy     = NewColor(16, 46, 85, 255)
	ColorDarkOlive    = NewColor(60, 101, 19, 255)
	ColorDarkLime     = NewColor(122, 154, 45, 255)
	ColorDarkFuchsia  = NewColor(154, 45, 141, 255)
	ColorDarkAqua     = NewColor(45, 154, 154, 255)

	ColorVeryDarkGray = NewColor(64, 64, 64, 255)
)
var (

	// DebugMode enables rendering of debug outlines.
	DebugMode bool

	// DumpMode causes the library to write cached images to disk
	// before exiting when enabled.
	DumpMode bool
)
var (

	// AutoReload enables automatic reloading of theme and layout files
	// when they are modified on disk, only use this for quickly iterating when designing your own themes.
	AutoReload bool
)

FUNCTIONS

func AccentSaturation() float64
    AccentSaturation returns the current accent color saturation value.

func AddOverlayFlow(flow *itemData)
func CurrentLayoutName() string
    CurrentLayoutName returns the active style theme name.

func CurrentThemeName() string
    CurrentThemeName returns the active theme name.

func Draw(screen *ebiten.Image)
    Draw renders the UI to the provided screen image. Call this from your Ebiten
    Draw function.

func DumpCachedImages() error
    DumpCachedImages writes all cached item images and item source images to the
    debug directory. The game must be running so pixels can be read. Any pending
    renders are generated before writing the files.

func EnsureFontSource(ttf []byte) error
    EnsureFontSource initializes the font source from ttf data if needed.

func FontSource() *text.GoTextFaceSource
    FontSource returns the current text face source.

func Layout(outsideWidth, outsideHeight int) (int, int)
    Layout reports the dimensions for the game's screen. Pass Ebiten's outside
    size values to this from your Layout function.

func ListLayouts() ([]string, error)
    ListLayouts returns the available style theme names.

func ListThemes() ([]string, error)
    ListThemes returns the available palette names.

func LoadLayout(name string) error
func LoadTheme(name string) error
    LoadTheme reads a theme JSON file from the themes directory and sets it as
    the current theme without modifying existing windows.

func NewButton(item *itemData) (*itemData, *EventHandler)
    Create a new button from the default theme

func NewCheckbox(item *itemData) (*itemData, *EventHandler)
    Create a new button from the default theme

func NewColorWheel(item *itemData) (*itemData, *EventHandler)
    Create a new color wheel from the default theme

func NewDropdown(item *itemData) (*itemData, *EventHandler)
    Create a new dropdown from the default theme

func NewInput(item *itemData) (*itemData, *EventHandler)
    Create a new input box from the default theme

func NewRadio(item *itemData) (*itemData, *EventHandler)
    Create a new radio button from the default theme

func NewSlider(item *itemData) (*itemData, *EventHandler)
    Create a new slider from the default theme

func NewText(item *itemData) (*itemData, *EventHandler)
    Create a new textbox from the default theme

func SaveTheme(name string) error
    SaveTheme writes the current theme to a JSON file with the given name.

func ScreenSize() (int, int)
    ScreenSize returns the current screen size.

func SetAccentColor(c Color)
    SetAccentColor updates the accent color in the current theme and applies it
    to all windows and widgets.

func SetAccentSaturation(s float64)
    SetAccentSaturation updates the saturation component of the accent color and
    reapplies it to the current theme.

func SetCurrentLayoutName(name string)
    SetCurrentLayoutName updates the active style theme name.

func SetCurrentThemeName(name string)
    SetCurrentThemeName updates the active theme name.

func SetFontSource(src *text.GoTextFaceSource)
    SetFontSource sets the text face source used when rendering text.

func SetScreenSize(w, h int)
    SetScreenSize sets the current screen size used for layout calculations.

func SetUIScale(scale float32)
func Update() error
    Update processes input and updates window state. Programs embedding the UI
    can call this from their Ebiten Update handler.

func NewWindow(win *windowData) *windowData
    Create a new window from the default theme


TYPES

type AlignType = alignType

type Color color.RGBA

func NewColor(r, g, b, a uint8) Color

func (c Color) MarshalJSON() ([]byte, error)
    MarshalJSON implements json.Marshaler using HSV representation.

func (c Color) RGBA() (r, g, b, a uint32)

func (c Color) ToRGBA() color.RGBA

func (c *Color) UnmarshalJSON(data []byte) error
    UnmarshalJSON accepts HSV, RGBA objects or a string. Strings may reference
    a named color from the theme, a hex RGB(A) value like "#RRGGBB" or
    comma-separated HSV components "h,s,v".

type DragType = dragType

type EventHandler struct {
	Events chan UIEvent
	Handle func(UIEvent)
}
    EventHandler holds a channel widgets use to emit events. EventHandler
    provides both channel and callback based event delivery.

func (h *EventHandler) Emit(ev UIEvent)
    Emit delivers the event through the channel and callback if present.

type FlowType = flowType

type Game struct {
}

func NewGame() *Game
    NewGame creates a new Game instance.

func (g *Game) Draw(screen *ebiten.Image)

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int)

func (g *Game) Update() error

type ItemData = itemData

func Overlays() []*ItemData
    Overlays returns the list of active overlays.

func (parent *ItemData) AddItem(child *ItemData)
    AddItem appends a child item to the parent item.

type ItemTypeData = itemTypeData

type LayoutBools struct {
	Window   bool
	Button   bool
	Text     bool
	Checkbox bool
	Radio    bool
	Input    bool
	Slider   bool
	Dropdown bool
	Tab      bool
}

type LayoutNumbers struct {
	Window   float32
	Button   float32
	Text     float32
	Checkbox float32
	Radio    float32
	Input    float32
	Slider   float32
	Dropdown float32
	Tab      float32
}
    LayoutTheme controls spacing and padding used by widgets.

type LayoutTheme struct {
	SliderValueGap   float32
	DropdownArrowPad float32
	TextPadding      float32

	Fillet        LayoutNumbers
	Border        LayoutNumbers
	BorderPad     LayoutNumbers
	Filled        LayoutBools
	Outlined      LayoutBools
	ActiveOutline LayoutBools
}

type PinType = pinType

type Point = point

type Rect = rect

type RoundRect = roundRect

type Theme struct {
	Window   windowData
	Button   itemData
	Text     itemData
	Checkbox itemData
	Radio    itemData
	Input    itemData
	Slider   itemData
	Dropdown itemData
	Tab      itemData
}
    Theme bundles all style information for windows and widgets.

type UIEvent struct {
	Item    *ItemData
	Type    UIEventType
	Value   float32
	Index   int
	Checked bool
	Color   Color
	Text    string
}
    UIEvent describes a user interaction with a widget.

type UIEventType int
    UIEventType defines the kind of event emitted by widgets.

const (
	EventClick UIEventType = iota
	EventSliderChanged
	EventDropdownSelected
	EventCheckboxChanged
	EventRadioSelected
	EventColorChanged
	EventInputChanged
)
type WindowData = windowData

func Windows() []*WindowData
    Windows returns the list of active windows.

func (win *WindowData) AddItem(child *ItemData)
    AddItem appends a child item to the window.

