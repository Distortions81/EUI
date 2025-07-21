package eui

import (
	"bytes"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

// Windows returns the list of active windows.
func Windows() []*WindowData { return windows }

// Overlays returns the list of active overlays.
func Overlays() []*ItemData { return overlays }

// SetScreenSize sets the current screen size used for layout calculations.
func SetScreenSize(w, h int) {
	screenWidth = w
	screenHeight = h
}

// ScreenSize returns the current screen size.
func ScreenSize() (int, int) { return screenWidth, screenHeight }

// SetFontSource sets the text face source used when rendering text.
func SetFontSource(src *text.GoTextFaceSource) {
	mplusFaceSource = src
	faceCache = map[float64]*text.GoTextFace{}
}

// FontSource returns the current text face source.
func FontSource() *text.GoTextFaceSource { return mplusFaceSource }

// EnsureFontSource initializes the font source from ttf data if needed.
func EnsureFontSource(ttf []byte) error {
	if mplusFaceSource != nil {
		return nil
	}
	s, err := text.NewGoTextFaceSource(bytes.NewReader(ttf))
	if err != nil {
		return err
	}
	mplusFaceSource = s
	faceCache = map[float64]*text.GoTextFace{}
	return nil
}

// AddItem appends a child item to the parent item.
func (parent *ItemData) AddItem(child *ItemData) { parent.addItemTo(child) }

// AddItem appends a child item to the window.
func (win *WindowData) AddItem(child *ItemData) { win.addItemTo(child) }

// ListThemes returns the available color theme names.
func ListThemes() ([]string, error) { return listThemes() }

// ListLayouts returns the available style theme names.
func ListLayouts() ([]string, error) { return listLayouts() }

// CurrentThemeName returns the active theme name.
func CurrentThemeName() string { return currentThemeName }

// SetCurrentThemeName updates the active theme name.
func SetCurrentThemeName(name string) { currentThemeName = name }

// CurrentLayoutName returns the active layout theme name.
func CurrentLayoutName() string { return currentLayoutName }

// SetCurrentLayoutName updates the active layout theme name.
func SetCurrentLayoutName(name string) { currentLayoutName = name }

// AccentSaturation returns the current accent color saturation value.
func AccentSaturation() float64 { return accentSaturation }
