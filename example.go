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
	leftText := NewText(&itemData{Text: "left panel item 1", Size: point{X: 100, Y: 32}, FontSize: 24})
	leftText2 := NewText(&itemData{Text: "left panel item 2", Size: point{X: 100, Y: 32}, FontSize: 8})
	leftText3 := NewText(&itemData{Text: "left panel item 3", Size: point{X: 100, Y: 32}, FontSize: 8})
	leftFlow.Contents = append(leftFlow.Contents, leftText)
	leftFlow.Contents = append(leftFlow.Contents, leftText2)
	leftFlow.Contents = append(leftFlow.Contents, leftText2)
	leftFlow.Contents = append(leftFlow.Contents, leftText3)
	mainFlow.Contents = append(mainFlow.Contents, leftFlow)

	rightFlow := &itemData{
		ItemType: ITEM_FLOW,
		Size:     point{X: 200, Y: 300},
		FlowType: FLOW_HORIZONTAL,
	}
	rightText := NewText(&itemData{Text: "right panel item 1", Size: point{X: 100, Y: 32}, FontSize: 8})
	rightText2 := NewText(&itemData{Text: "right panel item 2", Size: point{X: 100, Y: 32}, FontSize: 8})
	rightText3 := NewText(&itemData{Text: "right panel item 3", Size: point{X: 100, Y: 32}, FontSize: 8})
	rightFlow.Contents = append(rightFlow.Contents, rightText)
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
