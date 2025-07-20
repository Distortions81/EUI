package main

// makeFlowTestWindow creates a window showcasing all flow types, nested flows,
// and pinned objects for testing layout behavior.
func makeFlowTestWindow() *windowData {
	win := NewWindow(&windowData{
		Title:     "Flow Test",
		Size:      point{X: 480, Y: 420},
		Position:  point{X: 8, Y: 8},
		AutoSize:  true,
		Open:      true,
		Resizable: true,
		Movable:   true,
	})

	mainFlow := &itemData{
		ItemType: ITEM_FLOW,
		Size:     win.Size,
		FlowType: FLOW_VERTICAL,
	}
	win.addItemTo(mainFlow)

	// Demonstrate a basic horizontal flow
	hFlow := &itemData{
		ItemType:   ITEM_FLOW,
		FlowType:   FLOW_HORIZONTAL,
		Size:       point{X: 440, Y: 32},
		Fixed:      true,
		Scrollable: true,
	}
	mainFlow.addItemTo(hFlow)
	hFlow.addItemTo(NewButton(&itemData{Text: "H1", Size: point{X: 60, Y: 24}, FontSize: 8}))
	hFlow.addItemTo(NewButton(&itemData{Text: "H2", Size: point{X: 60, Y: 24}, FontSize: 8}))
	hFlow.addItemTo(NewButton(&itemData{Text: "H3", Size: point{X: 60, Y: 24}, FontSize: 8}))

	// Demonstrate a vertical flow
	vFlow := &itemData{
		ItemType: ITEM_FLOW,
		FlowType: FLOW_VERTICAL,
		Size:     point{X: 120, Y: 72},
		Fixed:    true,
	}
	mainFlow.addItemTo(vFlow)
	vFlow.addItemTo(NewButton(&itemData{Text: "V1", Size: point{X: 80, Y: 20}, FontSize: 8}))
	vFlow.addItemTo(NewButton(&itemData{Text: "V2", Size: point{X: 80, Y: 20}, FontSize: 8}))

	// Horizontal reverse flow
	hrFlow := &itemData{
		ItemType: ITEM_FLOW,
		FlowType: FLOW_HORIZONTAL_REV,
		Size:     point{X: 440, Y: 32},
		Fixed:    true,
	}
	mainFlow.addItemTo(hrFlow)
	hrFlow.addItemTo(NewButton(&itemData{Text: "HR1", Size: point{X: 60, Y: 24}, FontSize: 8}))
	hrFlow.addItemTo(NewButton(&itemData{Text: "HR2", Size: point{X: 60, Y: 24}, FontSize: 8}))
	hrFlow.addItemTo(NewButton(&itemData{Text: "HR3", Size: point{X: 60, Y: 24}, FontSize: 8}))

	// Vertical reverse flow
	vrFlow := &itemData{
		ItemType: ITEM_FLOW,
		FlowType: FLOW_VERTICAL_REV,
		Size:     point{X: 120, Y: 72},
		Fixed:    true,
	}
	mainFlow.addItemTo(vrFlow)
	vrFlow.addItemTo(NewButton(&itemData{Text: "VR1", Size: point{X: 80, Y: 20}, FontSize: 8}))
	vrFlow.addItemTo(NewButton(&itemData{Text: "VR2", Size: point{X: 80, Y: 20}, FontSize: 8}))

	// Nested flows: horizontal flow containing a vertical flow and another horizontal flow
	nested := &itemData{
		ItemType: ITEM_FLOW,
		FlowType: FLOW_HORIZONTAL,
		Size:     point{X: 440, Y: 80},
		Fixed:    true,
	}
	mainFlow.addItemTo(nested)

	nestV := &itemData{ItemType: ITEM_FLOW, FlowType: FLOW_VERTICAL, Size: point{X: 100, Y: 60}, Fixed: true}
	nestV.addItemTo(NewButton(&itemData{Text: "NV1", Size: point{X: 60, Y: 20}, FontSize: 8}))
	nestV.addItemTo(NewButton(&itemData{Text: "NV2", Size: point{X: 60, Y: 20}, FontSize: 8}))
	nested.addItemTo(nestV)

	nestH := &itemData{ItemType: ITEM_FLOW, FlowType: FLOW_HORIZONTAL, Size: point{X: 100, Y: 32}, Fixed: true}
	nestH.addItemTo(NewButton(&itemData{Text: "NH1", Size: point{X: 40, Y: 24}, FontSize: 8}))
	nestH.addItemTo(NewButton(&itemData{Text: "NH2", Size: point{X: 40, Y: 24}, FontSize: 8}))
	nested.addItemTo(nestH)

	// Pinned buttons inside the window
	mainFlow.addItemTo(NewButton(&itemData{Text: "TopLeft", PinTo: PIN_TOP_LEFT, Size: point{X: 60, Y: 20}, FontSize: 8}))
	mainFlow.addItemTo(NewButton(&itemData{Text: "TopRight", PinTo: PIN_TOP_RIGHT, Size: point{X: 60, Y: 20}, FontSize: 8}))
	mainFlow.addItemTo(NewButton(&itemData{Text: "BottomLeft", PinTo: PIN_BOTTOM_LEFT, Size: point{X: 80, Y: 20}, FontSize: 8}))
	mainFlow.addItemTo(NewButton(&itemData{Text: "BottomRight", PinTo: PIN_BOTTOM_RIGHT, Size: point{X: 80, Y: 20}, FontSize: 8}))
	mainFlow.addItemTo(NewButton(&itemData{Text: "Center", PinTo: PIN_MID_CENTER, Size: point{X: 60, Y: 20}, FontSize: 8}))

	// Overlay flow demonstrating screen pinning
	over := &itemData{
		ItemType: ITEM_FLOW,
		FlowType: FLOW_HORIZONTAL,
		Size:     point{X: 88, Y: 32},
		Position: point{X: 4, Y: 4},
		PinTo:    PIN_TOP_LEFT,
	}
	over.addItemTo(NewButton(&itemData{Text: "Overlay", Size: point{X: 80, Y: 24}, FontSize: 8}))
	AddOverlayFlow(over)

	return win
}
