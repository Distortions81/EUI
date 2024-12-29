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
	leftText2 := NewText(&itemData{Text: "left vflow 2", Size: point{X: 100, Y: 32}, FontSize: 12})
	leftFlow.Contents = append(leftFlow.Contents, leftText)
	leftFlow.Contents = append(leftFlow.Contents, leftText2)
	mainFlow.Contents = append(mainFlow.Contents, leftFlow)

	rightFlow := &itemData{
		ItemType: ITEM_FLOW,
		Size:     point{X: 200, Y: 300},
		FlowType: FLOW_HORIZONTAL,
	}
	rightText := NewText(&itemData{Text: "right vflow", Size: point{X: 100, Y: 32}, FontSize: 12})
	rightText2 := NewText(&itemData{Text: "right vflow 2", Size: point{X: 100, Y: 32}, FontSize: 12})
	rightFlow.Contents = append(rightFlow.Contents, rightText)
	rightFlow.Contents = append(rightFlow.Contents, rightText2)
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
