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
```

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

	// DebugMode enables rendering of debug outlines.
	DebugMode bool

        // DumpMode causes the library to write cached images to disk
        // before exiting when enabled.
        DumpMode bool

        // TreeMode dumps the window hierarchy to debug/tree.json
        // before exiting when enabled.
        TreeMode bool
)
var (

        // AutoReload enables automatic reloading of theme and layout files
        // when they are modified on disk, only use this for quickly iterating when designing your own themes.
        AutoReload bool

       // AutoHiDPI enables automatic scaling when the device scale factor changes.
       // The active UI scale is adjusted so the interface keeps the same size on screen.
       // This defaults to true and can be disabled if necessary. Applications
       // that track their own scaling variables should call `UIScale()` after
       // `Layout` to read the updated value.
       AutoHiDPI bool
)

FUNCTIONS

func AccentSaturation() float64
    AccentSaturation returns the current accent color saturation value.

func AddOverlayFlow(flow *itemData)
func CurrentStyleName() string
    CurrentStyleName returns the active style theme name.

func CurrentThemeName() string
    CurrentThemeName returns the active theme name.

func Draw(screen *ebiten.Image)
    Draw renders the UI to the provided screen image. Call this from your Ebiten
    Draw function.

func DumpCachedImages() error
    DumpCachedImages writes all cached item images and item source images to the
    debug directory. The game must be running so pixels can be read. Any pending
    renders are generated before writing the files.

func DumpTree() error
    DumpTree writes the window and overlay hierarchy to debug/tree.json.

func EnsureFontSource(ttf []byte) error
    EnsureFontSource initializes the font source from ttf data if needed.

func FontSource() *text.GoTextFaceSource
    FontSource returns the current text face source.

func Layout(outsideWidth, outsideHeight int) (int, int)
    Layout reports the screen dimensions and scales the resolution using the
    device scale factor. Pass Ebiten's outside size values to this from your
    Layout function. Disable this behavior by setting `AutoHiDPI` to false.

func ListStyles() ([]string, error)
    ListStyles returns the available style theme names.

func ListThemes() ([]string, error)
    ListThemes returns the available palette names.

func LoadStyle(name string) error
func LoadTheme(name string) error
    LoadTheme reads a theme JSON file from the themes directory and sets it as
    the current theme without modifying existing windows.

func NewButton(item *itemData) (*itemData, *EventHandler)
    Create a new button from the default theme. Unspecified fields
    inherit their values from the theme.

func NewCheckbox(item *itemData) (*itemData, *EventHandler)
    Create a new button from the default theme. Unspecified fields
    inherit their values from the theme.

func NewColorWheel(item *itemData) (*itemData, *EventHandler)
    Create a new color wheel from the default theme. Unspecified fields
    inherit their values from the theme.

func NewDropdown(item *itemData) (*itemData, *EventHandler)
    Create a new dropdown from the default theme. Unspecified fields
    inherit their values from the theme.

func NewInput(item *itemData) (*itemData, *EventHandler)
    Create a new input box from the default theme. Unspecified fields
    inherit their values from the theme.

func NewRadio(item *itemData) (*itemData, *EventHandler)
    Create a new radio button from the default theme. Unspecified fields
    inherit their values from the theme.

func NewSlider(item *itemData) (*itemData, *EventHandler)
    Create a new slider from the default theme. Unspecified fields
    inherit their values from the theme.

func NewText(item *itemData) (*itemData, *EventHandler)
    Create a new textbox from the default theme. Unspecified fields
    inherit their values from the theme.

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

func SetCurrentStyleName(name string)
    SetCurrentStyleName updates the active style theme name.

func SetCurrentThemeName(name string)
    SetCurrentThemeName updates the active theme name.

func SetFontSource(src *text.GoTextFaceSource)
    SetFontSource sets the text face source used when rendering text.

func SetScreenSize(w, h int)
    SetScreenSize sets the current screen size used for layout calculations
    and clamps existing windows to the new bounds.

func SetUIScale(scale float32)
    SetUIScale updates layout metrics for the given scale and resizes
    windows created with AutoSize. The value is clamped to the range
    0.5–2.5.
func UIScale() float32
    UIScale returns the current UI scale factor. When `AutoHiDPI` is enabled
    the value may change after `Layout` applies the device scale factor.
func SyncHiDPIScale()
    SyncHiDPIScale adjusts the UI scale automatically when the device scale
    factor changes.
func Update() error
    Update processes input and updates window state. Programs embedding the UI
    can call this from their Ebiten Update handler.

func NewWindow(win *windowData) *windowData
    Create a new window from the default theme. Unspecified fields
    inherit their values from the theme.

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


type ItemData = itemData
    ItemData represents a widget. Set Tooltip to display a floating hint when
    hovering over the item (empty string disables it). Set LabelImage to supply
    an image label for buttons, checkboxes, radios, sliders and dropdowns.

func Overlays() []*ItemData
    Overlays returns the list of active overlays.

func (parent *ItemData) AddItem(child *ItemData)
    AddItem appends a child item to the parent item.

type ItemTypeData = itemTypeData

type StyleBools struct {
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

type StyleNumbers struct {
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
    StyleTheme controls spacing and padding used by widgets.

type StyleTheme struct {
	SliderValueGap   float32
	DropdownArrowPad float32
	TextPadding      float32

        Fillet        StyleNumbers
        Border        StyleNumbers
        BorderPad     StyleNumbers
        Filled        StyleBools
        Outlined      StyleBools
        ActiveOutline StyleBools
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

func (win *WindowData) Refresh()
    Refresh forces the window to recalculate layout, resize to its contents and
    adjust scrolling after modifying contents.
```
