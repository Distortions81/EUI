package main

import eui "EUI/eui"

// makeShowcaseWindow creates a window demonstrating most widget types.
func makeShowcaseWindow() *eui.WindowData {
	win := eui.NewWindow(&eui.WindowData{
		Title:     "Showcase",
		Size:      eui.Point{X: 400, Y: 420},
		Position:  eui.Point{X: 8, Y: 8},
		AutoSize:  true,
		Open:      true,
		Movable:   true,
		Resizable: true,
		Closable:  true,
	})

	mainFlow := &eui.ItemData{
		ItemType: eui.ITEM_FLOW,
		Size:     win.Size,
		FlowType: eui.FLOW_VERTICAL,
	}
	win.AddItem(mainFlow)

	titleText := eui.NewText(&eui.ItemData{Text: "Demonstration of widgets", Size: eui.Point{X: 380, Y: 32}, FontSize: 10})
	mainFlow.AddItem(titleText)

	btnImage := eui.NewButton(&eui.ItemData{Text: "Sprite Button", ImageName: "1", Size: eui.Point{X: 100, Y: 64}, FontSize: 8})
	mainFlow.AddItem(btnImage)
	btnText := eui.NewButton(&eui.ItemData{Text: "Text Button", Size: eui.Point{X: 100, Y: 24}, FontSize: 8})
	mainFlow.AddItem(btnText)

	chk := eui.NewCheckbox(&eui.ItemData{Text: "Enable option", Size: eui.Point{X: 140, Y: 24}, FontSize: 8})
	mainFlow.AddItem(chk)

	radioA := eui.NewRadio(&eui.ItemData{Text: "Choice A", RadioGroup: "grp1", Size: eui.Point{X: 140, Y: 24}, FontSize: 8})
	radioB := eui.NewRadio(&eui.ItemData{Text: "Choice B", RadioGroup: "grp1", Size: eui.Point{X: 140, Y: 24}, FontSize: 8})
	mainFlow.AddItem(radioA)
	mainFlow.AddItem(radioB)

	slider := eui.NewSlider(&eui.ItemData{Label: "Float Slider", Size: eui.Point{X: 180, Y: 24}, MinValue: 0, MaxValue: 100, IntOnly: false, FontSize: 8})
	mainFlow.AddItem(slider)
	intSlider := eui.NewSlider(&eui.ItemData{Label: "Int Slider", Size: eui.Point{X: 180, Y: 24}, MinValue: 0, MaxValue: 10, IntOnly: true, FontSize: 8})
	mainFlow.AddItem(intSlider)

	input := eui.NewInput(&eui.ItemData{Label: "Text Field", Text: "", Size: eui.Point{X: 180, Y: 24}, FontSize: 8})
	mainFlow.AddItem(input)

	dropdown := eui.NewDropdown(&eui.ItemData{Label: "Select Option", Size: eui.Point{X: 180, Y: 24}, FontSize: 8})
	dropdown.Options = []string{"First", "Second", "Third", "Fourth"}
	dropdown.HoverIndex = -1
	mainFlow.AddItem(dropdown)

	hFlow := &eui.ItemData{
		ItemType:   eui.ITEM_FLOW,
		FlowType:   eui.FLOW_HORIZONTAL,
		Size:       eui.Point{X: 380, Y: 40},
		Fixed:      true,
		Scrollable: true,
	}
	mainFlow.AddItem(hFlow)
	hFlow.AddItem(eui.NewButton(&eui.ItemData{Text: "One", Size: eui.Point{X: 60, Y: 24}, FontSize: 8}))
	hFlow.AddItem(eui.NewButton(&eui.ItemData{Text: "Two", Size: eui.Point{X: 60, Y: 24}, FontSize: 8}))
	hFlow.AddItem(eui.NewButton(&eui.ItemData{Text: "Three", Size: eui.Point{X: 60, Y: 24}, FontSize: 8}))
	hFlow.AddItem(eui.NewButton(&eui.ItemData{Text: "Four", Size: eui.Point{X: 60, Y: 24}, FontSize: 8}))

	tabFlow := &eui.ItemData{
		ItemType:   eui.ITEM_FLOW,
		FlowType:   eui.FLOW_VERTICAL,
		Size:       eui.Point{X: 380, Y: 120},
		Fixed:      true,
		Color:      eui.ColorDarkGray,
		ClickColor: eui.ColorDarkTeal,
		FontSize:   8,
		Tabs: []*eui.ItemData{
			{Name: "Tab 1", ItemType: eui.ITEM_FLOW, FlowType: eui.FLOW_VERTICAL},
			{Name: "Tab 2", ItemType: eui.ITEM_FLOW, FlowType: eui.FLOW_VERTICAL},
		},
	}
	mainFlow.AddItem(tabFlow)
	tabFlow.Tabs[0].AddItem(eui.NewText(&eui.ItemData{Text: "Tab 1 content", Size: eui.Point{X: 100, Y: 32}, FontSize: 8}))
	tabFlow.Tabs[1].AddItem(eui.NewText(&eui.ItemData{Text: "Tab 2 content", Size: eui.Point{X: 100, Y: 32}, FontSize: 8}))

	return win
}
