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
	leftText := NewText(&itemData{Text: "left vflow", Size: point{X: 100, Y: 32}, FontSize: 12})
	leftFlow.Contents = append(leftFlow.Contents, leftText)
	mainFlow.Contents = append(mainFlow.Contents, leftFlow)

	rightFlow := &itemData{
		ItemType: ITEM_FLOW,
		Position: point{X: 100, Y: 0},
		Size:     point{X: 200, Y: 300},
		FlowType: FLOW_HORIZONTAL,
	}
	rightText := NewText(&itemData{Text: "right vflow", Size: point{X: 100, Y: 32}, FontSize: 12})
	rightFlow.Contents = append(rightFlow.Contents, rightText)
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
