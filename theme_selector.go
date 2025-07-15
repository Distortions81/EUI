package main

import (
	"log"
)

func makeThemeSelector() *windowData {
	names, err := listThemes()
	if err != nil {
		log.Printf("listThemes error: %v", err)
		return nil
	}
	if currentThemeName == "" {
		currentThemeName = names[0]
	}
	win := NewWindow(&windowData{
		Title:     "Themes",
		Movable:   true,
		Resizable: true,
		Closable:  true,
		// Give the dropdown room to fully render by accounting for the
		// title bar height and the control's size. Extra height is for
		// the saturation slider.
		Size:     point{X: 192, Y: 192},
		Position: point{X: 4, Y: 4},
		Open:     true,
	})
	mainFlow := &itemData{ItemType: ITEM_FLOW, Size: win.Size, FlowType: FLOW_VERTICAL}
	win.addItemTo(mainFlow)

	var satSlider *itemData

	dd := NewDropdown(&itemData{Size: point{X: 150, Y: 24}, FontSize: 8})
	dd.Options = names
	for i, n := range names {
		if n == currentThemeName {
			dd.Selected = i
			break
		}
	}
	dd.OnSelect = func(idx int) {
		currentThemeName = names[idx]
		if err := LoadTheme(currentThemeName); err != nil {
			log.Printf("LoadTheme error: %v", err)
		}
		if satSlider != nil {
			satSlider.Value = float32(accentSaturation)
		}
	}
	dd.OnHover = func(idx int) {
		if err := LoadTheme(names[idx]); err != nil {
			log.Printf("LoadTheme error: %v", err)
		}
		if satSlider != nil {
			satSlider.Value = float32(accentSaturation)
		}
	}
	dd.HoverIndex = -1
	mainFlow.addItemTo(dd)

	cw := NewColorWheel(&itemData{Size: point{X: 128, Y: 128}})
	mainFlow.addItemTo(cw)

	satSlider = NewSlider(&itemData{Size: point{X: 128, Y: 24}, MinValue: 0, MaxValue: 1, FontSize: 8})
	satSlider.Value = float32(accentSaturation)
	satSlider.Action = func() {
		SetAccentSaturation(float64(satSlider.Value))
	}
	mainFlow.addItemTo(satSlider)

	return win
}
