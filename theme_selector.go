package main

import "log"

func makeThemeSelector() *windowData {
	names, err := listThemes()
	if err != nil {
		log.Printf("listThemes error: %v", err)
		return nil
	}
	if currentTheme == "" {
		currentTheme = names[0]
	}
	win := NewWindow(&windowData{
		Title:     "Themes",
		PinTo:     PIN_TOP_RIGHT,
		Movable:   false,
		Resizable: false,
		Closable:  false,
		// Give the dropdown room to fully render by accounting for the
		// title bar height and the control's size.
		Size:     point{X: 192, Y: 160},
		Position: point{X: 4, Y: 4},
	})
	dd := NewDropdown(&itemData{Size: point{X: 150, Y: 24}, FontSize: 8})
	dd.Options = names
	for i, n := range names {
		if n == currentTheme {
			dd.Selected = i
			break
		}
	}
	dd.OnSelect = func(idx int) {
		currentTheme = names[idx]
		if err := LoadTheme(currentTheme); err != nil {
			log.Printf("LoadTheme error: %v", err)
		}
	}
	dd.OnHover = func(idx int) {
		if err := LoadTheme(names[idx]); err != nil {
			log.Printf("LoadTheme error: %v", err)
		}
	}
	dd.HoverIndex = -1
	win.addItemTo(dd)
	return win
}
