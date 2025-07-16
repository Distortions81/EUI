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
		// additional layout selection controls.
		Size:     point{X: 192, Y: 248},
		Position: point{X: 4, Y: 4},
		Open:     true,
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
	themeSelectorDropdown = dd

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

	cw := NewColorWheel(&itemData{Size: point{X: 128, Y: 128}})
	mainFlow.addItemTo(cw)

	satSlider = NewSlider(&itemData{Label: "Color Intensity", Size: point{X: 128, Y: 24}, MinValue: 0, MaxValue: 1, FontSize: 8})
	satSlider.Value = float32(accentSaturation)
	satSlider.Action = func() {
		SetAccentSaturation(float64(satSlider.Value))
	}
	mainFlow.addItemTo(satSlider)

	editBtn := NewButton(&itemData{Text: "Edit Theme", Size: point{X: 80, Y: 24}, FontSize: 8})
	editBtn.Action = func() {
		if themeEditor == nil {
			themeEditor = makeThemeEditor()
			themeEditor.AddWindow(false)
		} else {
			themeEditor.Open = true
			themeEditor.BringForward()
		}
	}
	mainFlow.addItemTo(editBtn)

	return win
}
