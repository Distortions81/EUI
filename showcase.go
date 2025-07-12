package main

// makeShowcaseWindow creates a window demonstrating most widget types.
func makeShowcaseWindow() *windowData {
	win := NewWindow(&windowData{
		Title:    "Showcase",
		Size:     point{X: 400, Y: 420},
		Position: point{X: 8, Y: 8},
		AutoSize: true,
	})

	mainFlow := &itemData{
		ItemType: ITEM_FLOW,
		Size:     win.Size,
		FlowType: FLOW_VERTICAL,
	}
	win.addItemTo(mainFlow)

	titleText := NewText(&itemData{Text: "Demonstration of widgets", Size: point{X: 380, Y: 32}, FontSize: 10})
	mainFlow.addItemTo(titleText)

	btnImage := NewButton(&itemData{Text: "Sprite Button", ImageName: "1", Size: point{X: 100, Y: 64}, FontSize: 8})
	mainFlow.addItemTo(btnImage)
	btnText := NewButton(&itemData{Text: "Text Button", Size: point{X: 100, Y: 24}, FontSize: 8})
	mainFlow.addItemTo(btnText)

	chk := NewCheckbox(&itemData{Text: "Enable option", Size: point{X: 140, Y: 32}, FontSize: 8})
	mainFlow.addItemTo(chk)

	radioA := NewRadio(&itemData{Text: "Choice A", RadioGroup: "grp1", Size: point{X: 140, Y: 32}, FontSize: 8})
	radioB := NewRadio(&itemData{Text: "Choice B", RadioGroup: "grp1", Size: point{X: 140, Y: 32}, FontSize: 8})
	mainFlow.addItemTo(radioA)
	mainFlow.addItemTo(radioB)

	slider := NewSlider(&itemData{Size: point{X: 180, Y: 24}, MinValue: 0, MaxValue: 100, IntOnly: false, FontSize: 8})
	mainFlow.addItemTo(slider)
	intSlider := NewSlider(&itemData{Size: point{X: 180, Y: 24}, MinValue: 0, MaxValue: 10, IntOnly: true, FontSize: 8})
	mainFlow.addItemTo(intSlider)

	input := NewInput(&itemData{Size: point{X: 180, Y: 24}, FontSize: 8})
	mainFlow.addItemTo(input)

	hFlow := &itemData{
		ItemType:   ITEM_FLOW,
		FlowType:   FLOW_HORIZONTAL,
		Size:       point{X: 380, Y: 40},
		Fixed:      true,
		Scrollable: true,
	}
	mainFlow.addItemTo(hFlow)
	hFlow.addItemTo(NewButton(&itemData{Text: "One", Size: point{X: 60, Y: 24}, FontSize: 8}))
	hFlow.addItemTo(NewButton(&itemData{Text: "Two", Size: point{X: 60, Y: 24}, FontSize: 8}))
	hFlow.addItemTo(NewButton(&itemData{Text: "Three", Size: point{X: 60, Y: 24}, FontSize: 8}))
	hFlow.addItemTo(NewButton(&itemData{Text: "Four", Size: point{X: 60, Y: 24}, FontSize: 8}))

	tabFlow := &itemData{
		ItemType:   ITEM_FLOW,
		FlowType:   FLOW_VERTICAL,
		Size:       point{X: 380, Y: 120},
		Fixed:      true,
		Color:      ColorDarkGray,
		ClickColor: ColorDarkTeal,
		FontSize:   8,
		Tabs: []*itemData{
			{Name: "Tab 1", ItemType: ITEM_FLOW, FlowType: FLOW_VERTICAL},
			{Name: "Tab 2", ItemType: ITEM_FLOW, FlowType: FLOW_VERTICAL},
		},
	}
	mainFlow.addItemTo(tabFlow)
	tabFlow.Tabs[0].addItemTo(NewText(&itemData{Text: "Tab 1 content", Size: point{X: 100, Y: 32}, FontSize: 8}))
	tabFlow.Tabs[1].addItemTo(NewText(&itemData{Text: "Tab 2 content", Size: point{X: 100, Y: 32}, FontSize: 8}))

	return win
}
