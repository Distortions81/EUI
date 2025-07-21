package main

import (
	eui "EUI/eui"
	"log"
)

func makeThemeSelector() *eui.WindowData {
	names, err := eui.ListThemes()
	if err != nil {
		log.Printf("listThemes error: %v", err)
		return nil
	}
	if eui.CurrentThemeName() == "" {
		eui.SetCurrentThemeName(names[0])
	}
	win := eui.NewWindow(&eui.WindowData{
		Title:     "Themes",
		Resizable: false,
		Closable:  false,
		PinTo:     eui.PIN_TOP_RIGHT,
		AutoSize:  true,
		Open:      true,
	})
	mainFlow := &eui.ItemData{ItemType: eui.ITEM_FLOW, Size: win.Size, FlowType: eui.FLOW_VERTICAL}
	win.AddItem(mainFlow)

	var satSlider *eui.ItemData
	layoutNames, lerr := eui.ListLayouts()
	if lerr != nil {
		log.Printf("listLayouts error: %v", lerr)
	}

	dd, ddEvents := eui.NewDropdown(&eui.ItemData{Label: "Palette", Size: eui.Point{X: 150, Y: 24}, FontSize: 8})
	dd.Options = names
	for i, n := range names {
		if n == eui.CurrentThemeName() {
			dd.Selected = i
			break
		}
	}
	dd.HoverIndex = -1
	ddEvents.Handle = func(ev eui.UIEvent) {
		if ev.Type == eui.EventDropdownSelected {
			idx := ev.Index
			eui.SetCurrentThemeName(names[idx])
			if err := eui.LoadTheme(eui.CurrentThemeName()); err != nil {
				log.Printf("eui.LoadTheme error: %v", err)
			}
			if satSlider != nil {
				satSlider.Value = float32(eui.AccentSaturation())
			}
		}
	}
	mainFlow.AddItem(dd)

	if len(layoutNames) > 0 {
		ldd, lddEvents := eui.NewDropdown(&eui.ItemData{Label: "Style", Size: eui.Point{X: 150, Y: 24}, FontSize: 8})
		ldd.Options = layoutNames
		for i, n := range layoutNames {
			if n == eui.CurrentLayoutName() {
				ldd.Selected = i
				break
			}
		}
		ldd.HoverIndex = -1
		lddEvents.Handle = func(ev eui.UIEvent) {
			if ev.Type == eui.EventDropdownSelected {
				idx := ev.Index
				eui.SetCurrentLayoutName(layoutNames[idx])
				if err := eui.LoadLayout(eui.CurrentLayoutName()); err != nil {
					log.Printf("eui.LoadLayout error: %v", err)
				}
			}
		}
		mainFlow.AddItem(ldd)
	}

	cw, _ := eui.NewColorWheel(&eui.ItemData{Size: eui.Point{X: 160, Y: 128}})
	cw.OnColorChange = func(col eui.Color) {
		eui.SetAccentColor(col)
		if satSlider != nil {
			satSlider.Value = float32(eui.AccentSaturation())
		}
	}
	mainFlow.AddItem(cw)

	satSlider, satEvents := eui.NewSlider(&eui.ItemData{Label: "Color Intensity", Size: eui.Point{X: 128, Y: 24}, MinValue: 0, MaxValue: 1, FontSize: 8})
	satSlider.Value = float32(eui.AccentSaturation())
	satEvents.Handle = func(ev eui.UIEvent) {
		if ev.Type == eui.EventSliderChanged {
			eui.SetAccentSaturation(float64(ev.Value))
		}
	}
	mainFlow.AddItem(satSlider)

	return win
}
