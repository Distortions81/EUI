package main

func makeTestWindow() *windowData {
	newWindow := NewWindow(
		&windowData{
			TitleHeight: 24,
			Title:       "Test Window",
			Size:        point{X: 300, Y: 300},
			Position:    point{X: 8, Y: 8},
			AutoSize:    true,
		})

	mainFlow := &itemData{
		ItemType: ITEM_FLOW,
		Size:     newWindow.Size,
		FlowType: FLOW_HORIZONTAL,
		Fixed:    false,
	}
	newWindow.addItemTo(mainFlow)

	leftFlow := &itemData{
		ItemType: ITEM_FLOW,
		Size:     point{X: 100, Y: 300},
		FlowType: FLOW_VERTICAL,
		Fixed:    true,
	}
	mainFlow.addItemTo(leftFlow)

	leftText1 := NewText(&itemData{Text: "left panel item 1", Size: point{X: 100, Y: 32}, FontSize: 8})
	leftText2 := NewText(&itemData{Text: "left panel item 2", Size: point{X: 100, Y: 32}, FontSize: 8})
	leftText3 := NewText(&itemData{Text: "left panel item 3", Size: point{X: 100, Y: 32}, FontSize: 8})
	leftButton1 := NewButton(&itemData{Text: "sprite button", Size: point{X: 64, Y: 64}, FontSize: 8, ImageName: "1"})
	leftButton2 := NewButton(&itemData{Text: "text button", Size: point{X: 64, Y: 24}, FontSize: 8})
	leftCheckbox1 := NewCheckbox(&itemData{Text: "Option 1", Size: point{X: 100, Y: 32}, FontSize: 8})
	leftRadio1 := NewRadio(&itemData{Text: "Radio A", Size: point{X: 100, Y: 32}, FontSize: 8, RadioGroup: "grp1"})
	leftRadio2 := NewRadio(&itemData{Text: "Radio B", Size: point{X: 100, Y: 32}, FontSize: 8, RadioGroup: "grp1"})
	leftSlider1 := NewSlider(&itemData{Size: point{X: 96, Y: 20}, FontSize: 8, MinValue: 0, MaxValue: 10, IntOnly: true})
	leftInput1 := NewInput(&itemData{Size: point{X: 96, Y: 20}, FontSize: 8})
	leftFlow.addItemTo(leftText1)
	leftFlow.addItemTo(leftText2)
	leftFlow.addItemTo(leftText3)
	leftFlow.addItemTo(leftButton1)
	leftFlow.addItemTo(leftButton2)
	leftFlow.addItemTo(leftCheckbox1)
	leftFlow.addItemTo(leftRadio1)
	leftFlow.addItemTo(leftRadio2)
	leftFlow.addItemTo(leftSlider1)
	leftFlow.addItemTo(leftInput1)

	rightFlow := &itemData{
		ItemType:   ITEM_FLOW,
		Size:       point{X: 200, Y: 300},
		FlowType:   FLOW_HORIZONTAL,
		Scrollable: true,
	}
	mainFlow.addItemTo(rightFlow)

	rightText1 := NewText(&itemData{Text: "right panel item 1", Size: point{X: 100, Y: 32}, FontSize: 8})
	rightText2 := NewText(&itemData{Text: "right panel item 2", Size: point{X: 100, Y: 32}, FontSize: 8})
	rightText3 := NewText(&itemData{Text: "right panel item 3", Size: point{X: 100, Y: 32}, FontSize: 8})
	rightFlow.addItemTo(rightText1)
	rightFlow.addItemTo(rightText2)
	rightFlow.addItemTo(rightText3)
	colorSel := NewColorSelector(nil)
	rightFlow.addItemTo(colorSel)

	return newWindow
}
