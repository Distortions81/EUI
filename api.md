# API Reference

This document lists the exported constants, variables, types and functions provided by the `eui` package.
The library implements a minimal retained‑mode user interface on top of the [Ebiten](https://ebiten.org/) game engine.
It is currently in a pre‑alpha state and the API may change at any time.

## Constants

- `MinWinSizeX`, `MinWinSizeY` – minimum window dimensions in pixels. Windows cannot be resized smaller than 64×64.
- `DefaultTabWidth`, `DefaultTabHeight` – dimensions used when creating tab widgets.
- `Tol` – geometry tolerance used for hit testing on window edges when resizing.

### Enumerations

`FlowType` values:
- `FLOW_HORIZONTAL`
- `FLOW_VERTICAL`
- `FLOW_HORIZONTAL_REV`
- `FLOW_VERTICAL_REV`

`AlignType` values:
- `ALIGN_NONE`
- `ALIGN_LEFT`
- `ALIGN_CENTER`
- `ALIGN_RIGHT`

`PinType` values:
- `PIN_TOP_LEFT`, `PIN_TOP_CENTER`, `PIN_TOP_RIGHT`
- `PIN_MID_LEFT`, `PIN_MID_CENTER`, `PIN_MID_RIGHT`
- `PIN_BOTTOM_LEFT`, `PIN_BOTTOM_CENTER`, `PIN_BOTTOM_RIGHT`

`DragType` values:
- `PART_NONE`, `PART_BAR`, `PART_CLOSE`
- `PART_TOP`, `PART_RIGHT`, `PART_BOTTOM`, `PART_LEFT`
- `PART_TOP_RIGHT`, `PART_BOTTOM_RIGHT`, `PART_BOTTOM_LEFT`, `PART_TOP_LEFT`
- `PART_SCROLL_V`, `PART_SCROLL_H`

`ItemTypeData` values:
- `ITEM_NONE`, `ITEM_FLOW`, `ITEM_TEXT`, `ITEM_BUTTON`, `ITEM_CHECKBOX`,
  `ITEM_RADIO`, `ITEM_INPUT`, `ITEM_SLIDER`, `ITEM_DROPDOWN`, `ITEM_COLORWHEEL`

## Variables

- `DebugMode` – when set, additional outlines are rendered for debugging.
- `DumpMode` – when set, cached images are written to `./debug` after the first
  frame is rendered and the program exits.
- `ColorWhite`, `ColorBlack`, `ColorRed`, ... – a large palette of predefined
  colors available as variables of type `Color`.

## Types

- `WindowData`
- `ItemData`
- `RoundRect`
- `Rect`
- `Point`
- `FlowType`
- `AlignType`
- `PinType`
- `DragType`
- `ItemTypeData`
- `Game`
- `Theme`
- `LayoutTheme`
- `LayoutBools`
- `LayoutNumbers`

## Functions

- `Windows() []*WindowData`
- `Overlays() []*ItemData`
- `AddOverlayFlow(flow *ItemData)`
- `DumpCachedImages() error` – write cached images to `./debug`. Requires the game to be running.
- `SetScreenSize(w, h int)`
- `ScreenSize() (int, int)`
- `SetFontSource(src *text.GoTextFaceSource)`
- `FontSource() *text.GoTextFaceSource`
- `EnsureFontSource(ttf []byte) error`
- `ListThemes() ([]string, error)` – list the bundled color theme names
- `ListLayouts() ([]string, error)` – list the bundled layout theme names
- `CurrentThemeName() string`
- `SetCurrentThemeName(name string)`
- `CurrentLayoutName() string`
- `SetCurrentLayoutName(name string)`
- `AccentSaturation() float64`
- `SetAccentColor(c Color)`
- `SetAccentSaturation(s float64)`
- `SetUIScale(scale float32)`
- `NewGame() *Game`
- `Update() error`
- `Draw(screen *ebiten.Image)`
- `Layout(outsideWidth, outsideHeight int) (int, int)`
- `(g *Game) Update() error`
- `(g *Game) Draw(screen *ebiten.Image)`
- `(g *Game) Layout(outsideWidth, outsideHeight int) (int, int)`
- `NewWindow(win *WindowData) *WindowData`
 - `NewButton(item *ItemData) (*ItemData, *EventHandler)`
 - `NewCheckbox(item *ItemData) (*ItemData, *EventHandler)`
 - `NewRadio(item *ItemData) (*ItemData, *EventHandler)`
 - `NewInput(item *ItemData) (*ItemData, *EventHandler)`
 - `NewSlider(item *ItemData) (*ItemData, *EventHandler)`
 - `NewDropdown(item *ItemData) (*ItemData, *EventHandler)`
 - `NewColorWheel(item *ItemData) (*ItemData, *EventHandler)`
 - `NewText(item *ItemData) (*ItemData, *EventHandler)`
- `ColorWheelImage(size int) *ebiten.Image`
- `LoadTheme(name string) error`
- `SaveTheme(name string) error`
- `LoadLayout(name string) error`
- `(parent *ItemData) AddItem(child *ItemData)`
- `(win *WindowData) AddItem(child *ItemData)`
- `(win *WindowData) AddWindow(toBack bool)`
- `(win *WindowData) RemoveWindow()`
- `(win *WindowData) BringForward()`
- `(win *WindowData) MarkOpen()`
- `(win *WindowData) ToBack()`
- `(win *WindowData) Draw(screen *ebiten.Image)`
- `(win *WindowData) SetTitleSize(size float32)`
- `(win *WindowData) GetTitleSize() float32`
- `(win *WindowData) GetSize() Point`
- `(win *WindowData) GetPos() Point`
- `(item *ItemData) GetSize() Point`
- `(item *ItemData) GetPos() Point`

### Example

The snippet below creates a simple window containing a button:

```go
win := eui.NewWindow(&eui.WindowData{Title: "Example", Size: eui.Point{X: 200, Y: 120}})
btn, btnEvents := eui.NewButton(&eui.ItemData{Text: "Click Me"})
win.AddItem(btn)
win.AddWindow(false)
go func() {
    for ev := range btnEvents.Events {
        if ev.Type == eui.EventClick {
            // handle click
        }
    }
}()
```

To regenerate this file run:

```sh
go doc -all EUI/eui > api.md
```
