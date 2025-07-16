package main

import (
	"log"
	"reflect"
	"strings"
)

var themeEditor *windowData
var themeSelectorDropdown *itemData

func makeThemeEditor() *windowData {
	win := NewWindow(&windowData{
		Title:     "Theme Editor",
		Movable:   true,
		Resizable: true,
		Closable:  true,
		Size:      point{X: 220, Y: 400},
		Position:  point{X: 8, Y: 40},
		Open:      true,
	})
	mainFlow := &itemData{ItemType: ITEM_FLOW, Size: win.Size, FlowType: FLOW_VERTICAL, Scrollable: true}
	win.addItemTo(mainFlow)

	addEditors("", reflect.ValueOf(currentTheme).Elem(), mainFlow)

	nameInput := NewInput(&itemData{Size: point{X: 160, Y: 24}, FontSize: 8})
	mainFlow.addItemTo(nameInput)
	saveBtn := NewButton(&itemData{Text: "Save", Size: point{X: 80, Y: 24}, FontSize: 8})
	saveBtn.Action = func() {
		if err := SaveTheme(strings.TrimSpace(nameInput.Text)); err != nil {
			log.Printf("SaveTheme error: %v", err)
		} else if themeSelectorDropdown != nil {
			if names, err := listThemes(); err == nil {
				themeSelectorDropdown.Options = names
			}
		}
	}
	mainFlow.addItemTo(saveBtn)

	return win
}

func addEditors(prefix string, val reflect.Value, flow *itemData) {
	t := val.Type()
	for i := 0; i < val.NumField(); i++ {
		f := val.Field(i)
		ft := t.Field(i)
		label := prefix + ft.Name
		switch f.Kind() {
		case reflect.Struct:
			if ft.Type.Name() == "Color" {
				colPtr := f.Addr().Interface().(*Color)
				flow.addItemTo(NewText(&itemData{Text: label, FontSize: 8}))
				cw := NewColorWheel(&itemData{Size: point{X: 64, Y: 64}})
				cw.OnColorChange = func(c Color) {
					*colPtr = c
					applyThemeToAll()
				}
				flow.addItemTo(cw)
			} else {
				addEditors(label+" ", f, flow)
			}
		case reflect.Float32:
			valPtr := f.Addr().Interface().(*float32)
			sl := NewSlider(&itemData{Label: label, Size: point{X: 160, Y: 24}, MinValue: 0, MaxValue: 32, FontSize: 8})
			sl.Value = *valPtr
			sl.Action = func() {
				*valPtr = sl.Value
				applyThemeToAll()
			}
			flow.addItemTo(sl)
		case reflect.Bool:
			bptr := f.Addr().Interface().(*bool)
			cb := NewCheckbox(&itemData{Label: label, Size: point{X: 160, Y: 24}, FontSize: 8})
			cb.Checked = *bptr
			cb.Action = func() {
				*bptr = cb.Checked
				applyThemeToAll()
			}
			flow.addItemTo(cb)
		}
	}
}
