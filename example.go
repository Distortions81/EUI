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
	leftButton2 := NewButton(nil)
	leftFlow.Contents = append(leftFlow.Contents, leftText1)
	leftFlow.Contents = append(leftFlow.Contents, leftText2)
	leftFlow.Contents = append(leftFlow.Contents, leftText3)
	leftFlow.Contents = append(leftFlow.Contents, leftButton1)
	leftFlow.Contents = append(leftFlow.Contents, leftButton2)

	mainFlow.Contents = append(mainFlow.Contents, leftFlow)

	rightFlow := &itemData{
		ItemType: ITEM_FLOW,
		Size:     point{X: 200, Y: 300},
		FlowType: FLOW_HORIZONTAL,
	}
	rightText1 := NewText(&itemData{Text: "right panel item 1", Size: point{X: 100, Y: 32}, FontSize: 8})
	rightText2 := NewText(&itemData{Text: "right panel item 2", Size: point{X: 100, Y: 32}, FontSize: 8})
	rightText3 := NewText(&itemData{Text: "right panel item 3", Size: point{X: 100, Y: 32}, FontSize: 8})
	rightFlow.Contents = append(rightFlow.Contents, rightText1)
	rightFlow.Contents = append(rightFlow.Contents, rightText2)
	rightFlow.Contents = append(rightFlow.Contents, rightText3)
	mainFlow.Contents = append(mainFlow.Contents, rightFlow)

	newWindow := NewWindow(
		&windowData{
			TitleHeight: 24,
			Title:       "Test Window",
			Size:        point{X: 300, Y: 300},
			Position:    point{X: 32, Y: 32},
			Contents:    []*itemData{mainFlow},
		})

	return newWindow
}
