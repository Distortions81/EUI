package main

import (
	"log"

	"github.com/Distortions81/EUI/eui"
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
	win := eui.NewWindow()
	win.Title = "Themes"
	win.Resizable = true
	win.Closable = false
	win.PinTo = eui.PIN_TOP_RIGHT
	win.AutoSize = true
	win.Open()
	mainFlow := &eui.ItemData{ItemType: eui.ITEM_FLOW, Size: win.Size, FlowType: eui.FLOW_VERTICAL}
	win.AddItem(mainFlow)

	var satSlider *eui.ItemData
	styleNames, serr := eui.ListStyles()
	if serr != nil {
		log.Printf("listStyles error: %v", serr)
	}

	dd, ddEvents := eui.NewDropdown()
	dd.Label = "Palette"
	dd.Size = eui.Point{X: 150, Y: 24}
	dd.FontSize = 8
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

	if len(styleNames) > 0 {
		ldd, lddEvents := eui.NewDropdown()
		ldd.Label = "Style"
		ldd.Size = eui.Point{X: 150, Y: 24}
		ldd.FontSize = 8
		ldd.Options = styleNames
		for i, n := range styleNames {
			if n == eui.CurrentStyleName() {
				ldd.Selected = i
				break
			}
		}
		ldd.HoverIndex = -1
		lddEvents.Handle = func(ev eui.UIEvent) {
			if ev.Type == eui.EventDropdownSelected {
				idx := ev.Index
				eui.SetCurrentStyleName(styleNames[idx])
				if err := eui.LoadStyle(eui.CurrentStyleName()); err != nil {
					log.Printf("eui.LoadStyle error: %v", err)
				}
			}
		}
		mainFlow.AddItem(ldd)
	}

	cw, _ := eui.NewColorWheel()
	cw.Size = eui.Point{X: 160, Y: 128}
	cw.OnColorChange = func(col eui.Color) {
		eui.SetAccentColor(col)
		if satSlider != nil {
			satSlider.Value = float32(eui.AccentSaturation())
		}
	}
	mainFlow.AddItem(cw)

	satSlider, satEvents := eui.NewSlider()
	satSlider.Label = "Color Intensity"
	satSlider.Size = eui.Point{X: 128, Y: 24}
	satSlider.MinValue = 0
	satSlider.MaxValue = 1
	satSlider.FontSize = 8
	satSlider.Value = float32(eui.AccentSaturation())
	satEvents.Handle = func(ev eui.UIEvent) {
		if ev.Type == eui.EventSliderChanged {
			eui.SetAccentSaturation(float64(ev.Value))
		}
	}
	mainFlow.AddItem(satSlider)

	return win
}
