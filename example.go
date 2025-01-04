package main

func makeTestWindow() *windowData {

	mainFlow := &itemData{
		ItemType: ITEM_FLOW,
		Size:     point{X: 300, Y: 300},
		FlowType: FLOW_HORIZONTAL,
	}

	leftFlow := &itemData{
		ItemType: ITEM_FLOW,
		Size:     point{X: 100, Y: 300},
		FlowType: FLOW_VERTICAL,
	}
	leftText1 := NewText(&itemData{Text: "left panel item 1", Size: point{X: 100, Y: 32}, FontSize: 8})
	leftText2 := NewText(&itemData{Text: "left panel item 2", Size: point{X: 100, Y: 32}, FontSize: 8})
	leftText3 := NewText(&itemData{Text: "left panel item 3", Size: point{X: 100, Y: 32}, FontSize: 8})
	leftButton1 := NewButton(&itemData{Text: "test button", Size: point{X: 100, Y: 32}, FontSize: 8, ImageName: "1"})
	leftButton2 := NewButton(&itemData{Text: "test", Size: point{X: 50, Y: 32}, FontSize: 8})
	leftFlow.addItemTo(leftText1)
	leftFlow.addItemTo(leftText2)
	leftFlow.addItemTo(leftText3)
	leftFlow.addItemTo(leftButton1)
	leftFlow.addItemTo(leftButton2)
	mainFlow.addItemTo(leftFlow)

	rightFlow := &itemData{
		ItemType: ITEM_FLOW,
		Size:     point{X: 200, Y: 300},
		FlowType: FLOW_HORIZONTAL,
	}
	rightText1 := NewText(&itemData{Text: "right panel item 1", Size: point{X: 100, Y: 32}, FontSize: 8})
	rightText2 := NewText(&itemData{Text: "right panel item 2", Size: point{X: 100, Y: 32}, FontSize: 8})
	rightText3 := NewText(&itemData{Text: "right panel item 3", Size: point{X: 100, Y: 32}, FontSize: 8})
	rightFlow.addItemTo(rightText1)
	rightFlow.addItemTo(rightText2)
	rightFlow.addItemTo(rightText3)
	mainFlow.addItemTo(rightFlow)

	newWindow := NewWindow(
		&windowData{
			TitleHeight: 24,
			Title:       "Test Window",
			Size:        point{X: 300, Y: 300},
			Position:    point{X: 8, Y: 8},
			Contents:    []*itemData{mainFlow},
		})

	return newWindow
}
