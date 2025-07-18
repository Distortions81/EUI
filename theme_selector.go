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
		Resizable: false,
		Closable:  false,
		PinTo:     PIN_TOP_RIGHT,
		AutoSize:  true,
		Open:      false,
	})
	mainFlow := &itemData{ItemType: ITEM_FLOW, Size: win.Size, FlowType: FLOW_VERTICAL}
	win.addItemTo(mainFlow)

	var satSlider *itemData
	layoutNames, lerr := listLayouts()
	if lerr != nil {
		log.Printf("listLayouts error: %v", lerr)
	}

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

	if len(layoutNames) > 0 {
		ldd := NewDropdown(&itemData{Size: point{X: 150, Y: 24}, FontSize: 8})
		ldd.Options = layoutNames
		for i, n := range layoutNames {
			if n == currentLayoutName {
				ldd.Selected = i
				break
			}
		}
		ldd.OnSelect = func(idx int) {
			currentLayoutName = layoutNames[idx]
			if err := LoadLayout(currentLayoutName); err != nil {
				log.Printf("LoadLayout error: %v", err)
			}
		}
		ldd.OnHover = func(idx int) {
			if err := LoadLayout(layoutNames[idx]); err != nil {
				log.Printf("LoadLayout error: %v", err)
			}
		}
		ldd.HoverIndex = -1
		mainFlow.addItemTo(ldd)
	}

	// Widen the color wheel widget slightly so that the swatch drawn to the
	// right of the wheel is within the window's bounds.
	cw := NewColorWheel(&itemData{Size: point{X: 160, Y: 128}})
	mainFlow.addItemTo(cw)

	// Match the slider width to the widened color wheel for consistent layout.
	satSlider = NewSlider(&itemData{Label: "Color Intensity", Size: point{X: 160, Y: 24}, MinValue: 0, MaxValue: 1, FontSize: 8})
	satSlider.Value = float32(accentSaturation)
	satSlider.Action = func() {
		SetAccentSaturation(float64(satSlider.Value))
	}
	mainFlow.addItemTo(satSlider)

	return win
}
