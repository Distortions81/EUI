# API Reference

This document provides a generated listing of every function present in the
source tree. Functions are grouped by the file that defines them along with
their full Go syntax. The project is in an early pre‑alpha state so these APIs
may change frequently. This reference can help contributors navigate the code
base and locate specific functionality quickly.

The lists below were produced by scanning all `*.go` files with `grep` and
extracting every `func` declaration. Helper methods used solely in tests are
included as well so that nothing is omitted.

## Setup and Testing

The repository includes a `scripts/setup.sh` helper that installs the system
packages required to build the Ebiten based demo. After running it once, the
project can be compiled and vetted using the standard Go tools:

```sh
./scripts/setup.sh   # install dependencies (only needed once)
go vet ./...
go build ./...
```

Running these commands prior to committing changes ensures the code remains in a
buildable state.

## color_hsv.go

```
func hsvaToRGBA(h, s, v, a float64) color.RGBA
func rgbaToHSVA(c color.RGBA) (h, s, v, a float64)
func clamp(v, min, max float64) float64
func dimColor(c Color, factor float64) Color
func (c Color) MarshalJSON() ([]byte, error)
func (c *Color) UnmarshalJSON(data []byte) error
```

## color_utils.go

```
func NewColor(r, g, b, a uint8) Color
```

## color_wheel.go

```
func ColorWheelImage(size int) *ebiten.Image
```

## geometry.go

```
func (r rect) containsPoint(p point) bool
func (r rect) getRectangle() image.Rectangle
func withinRange(a, b float32, tol float32) bool
func pointAdd(a, b point) point { return point{X: a.X + b.X, Y: a.Y + b.Y} }
func pointSub(a, b point) point { return point{X: a.X - b.X, Y: a.Y - b.Y} }
func pointMul(a, b point) point { return point{X: a.X * b.X, Y: a.Y * b.Y} }
func pointDiv(a, b point) point { return point{X: a.X / b.X, Y: a.Y / b.Y} }
func pointScaleMul(a point) point { return point{X: a.X * uiScale, Y: a.Y * uiScale} }
func pointScaleDiv(a point) point { return point{X: a.X / uiScale, Y: a.Y / uiScale} }
func unionRect(a, b rect) rect
func intersectRect(a, b rect) rect
```

## glob.go

```
func init()
```

## input.go

```
func (g *Game) Update() error
func (win *windowData) clickWindowItems(mpos point, click bool)
func clickOverlay(root *itemData, mpos point, click bool) bool
func (item *itemData) clickFlows(mpos point, click bool) bool
func (item *itemData) clickItem(mpos point, click bool) bool
func uncheckRadioGroup(parent *itemData, group string, except *itemData)
func subUncheckRadio(list []*itemData, group string, except *itemData)
func (item *itemData) setSliderValue(mpos point)
func (item *itemData) colorAt(mpos point) (Color, bool)
func scrollFlow(items []*itemData, mpos point, delta point) bool
func scrollDropdown(items []*itemData, mpos point, delta point) bool
func scrollWindow(win *windowData, delta point) bool
func dragWindowScroll(win *windowData, mpos point, vert bool)
func dropdownOpenContains(items []*itemData, mpos point) bool
func clickOpenDropdown(items []*itemData, mpos point, click bool) bool
func dropdownOpenContainsAnywhere(mpos point) bool
func closeDropdowns(items []*itemData)
func closeAllDropdowns()
```

## layout.go

```
func LoadLayout(name string) error
func applyLayoutToTheme(th *Theme)
func listLayouts() ([]string, error)
```

## main.go

```
func main()
func loadIcons() error
func subLoadIcons(parent []*itemData) error
func loadImage(name string) (*ebiten.Image, error)
func startEbiten()
func newGame() *Game
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int)
```

## overlay.go

```
func AddOverlayFlow(flow *itemData)
```

## paths.go

```
func init()
```

## render.go

```
func (g *Game) Draw(screen *ebiten.Image)
func drawOverlay(item *itemData, screen *ebiten.Image)
func (win *windowData) Draw(screen *ebiten.Image)
func (win *windowData) drawBG(screen *ebiten.Image)
func (win *windowData) drawWinTitle(screen *ebiten.Image)
func (win *windowData) drawBorder(screen *ebiten.Image)
func (win *windowData) drawScrollbars(screen *ebiten.Image)
func (win *windowData) drawItems(screen *ebiten.Image)
func (item *itemData) drawFlows(win *windowData, parent *itemData, offset point, clip rect, screen *ebiten.Image)
func (item *itemData) drawItem(parent *itemData, offset point, clip rect, screen *ebiten.Image)
func drawDropdownOptions(item *itemData, offset point, clip rect, screen *ebiten.Image)
func (win *windowData) drawDebug(screen *ebiten.Image)
func drawRoundRect(screen *ebiten.Image, rrect *roundRect)
func drawTabShape(screen *ebiten.Image, pos point, size point, col Color, fillet float32, slope float32)
func strokeTabShape(screen *ebiten.Image, pos point, size point, col Color, fillet float32, slope float32, border float32)
func drawTriangle(screen *ebiten.Image, pos point, size float32, col Color)
func drawFPS(screen *ebiten.Image)
```

## showcase.go

```
func makeShowcaseWindow() *windowData
```

## struct.go

```
func (c Color) RGBA() (r, g, b, a uint32)
func (c Color) ToRGBA() color.RGBA { return color.RGBA(c) }
```

## theme.go

```
func init()
func resolveColor(s string, colors map[string]string, seen map[string]bool) (Color, error)
func LoadTheme(name string) error
func listThemes() ([]string, error)
func SaveTheme(name string) error
func SetAccentColor(c Color)
func SetAccentSaturation(s float64)
func applyAccentColor()
func applyThemeToAll()
func copyWindowStyle(dst, src *windowData)
func applyThemeToWindow(win *windowData)
func copyItemStyle(dst, src *itemData)
func applyThemeToItem(it *itemData)
func updateColorWheels(col Color)
func updateColorWheelList(items []*itemData, col Color)
```

## theme_selector.go

```
func makeThemeSelector() *windowData
```

## util.go

```
func (win *windowData) getWinRect() rect
func (item *itemData) getItemRect(win *windowData) rect
func (parent *itemData) addItemTo(item *itemData)
func (parent *windowData) addItemTo(item *itemData)
func (win *windowData) getMainRect() rect
func (win *windowData) getTitleRect() rect
func (win *windowData) xRect() rect
func (win *windowData) dragbarRect() rect
func (win *windowData) setSize(size point) bool
func (win *windowData) getWindowPart(mpos point, click bool) dragType
func (win *windowData) titleTextWidth() point
func (win *windowData) SetTitleSize(size float32)
func SetUIScale(scale float32)
func (win *windowData) GetTitleSize() float32
func (win *windowData) GetSize() point
func (win *windowData) GetPos() point
func (item *itemData) GetSize() point
func (item *itemData) GetPos() point
func (item *itemData) bounds(offset point) rect
func (win *windowData) contentBounds() point
func (win *windowData) updateAutoSize()
func (item *itemData) contentBounds() point
func (item *itemData) resizeFlow(parentSize point)
func (win *windowData) resizeFlows()
func pixelOffset(width float32) float32
func strokeLine(dst *ebiten.Image, x0, y0, x1, y1, width float32, col color.Color, aa bool)
func strokeRect(dst *ebiten.Image, x, y, w, h, width float32, col color.Color, aa bool)
func drawFilledRect(dst *ebiten.Image, x, y, w, h float32, col color.Color, aa bool)
```

## util_test.go

```
func TestWithinRange(t *testing.T)
func TestPointOperations(t *testing.T)
func TestUnionRect(t *testing.T)
func TestMergeData(t *testing.T)
func TestPinPositions(t *testing.T)
func TestItemOverlap(t *testing.T)
func TestSetSliderValue(t *testing.T)
func sliderTrackWidth(item *itemData) float32
func TestSliderTrackLengthMatch(t *testing.T)
func TestMarkOpen(t *testing.T)
func TestSetSizeClampAndScroll(t *testing.T)
func TestFlowContentBounds(t *testing.T)
func TestPixelOffset(t *testing.T)
func roundRectKeyPoints(rrect *roundRect) []point
func TestRoundRectSymmetry(t *testing.T)
func TestRoundRectFilletClamp(t *testing.T)
func TestStrokeLineParams(t *testing.T)
func TestStrokeRectParams(t *testing.T)
```

## window.go

```
func mergeData(original interface{}, updates interface{}) interface{}
func isZeroValue(value reflect.Value) bool
func (target *windowData) AddWindow(toBack bool)
func (target *windowData) RemoveWindow()
func NewWindow(win *windowData) *windowData
func NewButton(item *itemData) *itemData
func NewCheckbox(item *itemData) *itemData
func NewRadio(item *itemData) *itemData
func NewInput(item *itemData) *itemData
func NewSlider(item *itemData) *itemData
func NewDropdown(item *itemData) *itemData
func NewColorWheel(item *itemData) *itemData
func NewText(item *itemData) *itemData
func (target *windowData) BringForward()
func (target *windowData) MarkOpen()
func (target *windowData) ToBack()
func (pin pinType) getWinPosition(win *windowData) point
func (pin pinType) getItemPosition(win *windowData, item *itemData) point
func (pin pinType) getOverlayItemPosition(item *itemData) point
func (win *windowData) getPosition() point
func (item *itemData) getPosition(win *windowData) point
func (item *itemData) getOverlayPosition() point
func (win windowData) itemOverlap(size point) (bool, bool)

To regenerate this file run:

```sh
# gather all function declarations
grep -n "^func" -n $(git ls-files "*.go") > /tmp/all_funcs.txt
# then run the helper script in scripts/ (see repository for details)
```
