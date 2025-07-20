package main

import eui "EUI/eui"

// makeFlowTestWindow creates a window showcasing all flow types, nested flows,
// and pinned objects for testing layout behavior.
func makeFlowTestWindow() *eui.WindowData {
	win := eui.NewWindow(&eui.WindowData{
		Title:     "Flow Test",
		Size:      eui.Point{X: 480, Y: 420},
		Position:  eui.Point{X: 8, Y: 8},
		AutoSize:  true,
		Open:      true,
		Resizable: true,
		Movable:   true,
	})

	mainFlow := &eui.ItemData{
		ItemType: eui.ITEM_FLOW,
		Size:     win.Size,
		FlowType: eui.FLOW_VERTICAL,
	}
	win.AddItem(mainFlow)

	// Demonstrate a basic horizontal flow
	hFlow := &eui.ItemData{
		ItemType:   eui.ITEM_FLOW,
		FlowType:   eui.FLOW_HORIZONTAL,
		Size:       eui.Point{X: 440, Y: 32},
		Fixed:      true,
		Scrollable: true,
	}
	mainFlow.AddItem(hFlow)
	hFlow.AddItem(eui.NewButton(&eui.ItemData{Text: "H1", Size: eui.Point{X: 60, Y: 24}, FontSize: 8}))
	hFlow.AddItem(eui.NewButton(&eui.ItemData{Text: "H2", Size: eui.Point{X: 60, Y: 24}, FontSize: 8}))
	hFlow.AddItem(eui.NewButton(&eui.ItemData{Text: "H3", Size: eui.Point{X: 60, Y: 24}, FontSize: 8}))

	// Demonstrate a vertical flow
	vFlow := &eui.ItemData{
		ItemType: eui.ITEM_FLOW,
		FlowType: eui.FLOW_VERTICAL,
		Size:     eui.Point{X: 120, Y: 72},
		Fixed:    true,
	}
	mainFlow.AddItem(vFlow)
	vFlow.AddItem(eui.NewButton(&eui.ItemData{Text: "V1", Size: eui.Point{X: 80, Y: 20}, FontSize: 8}))
	vFlow.AddItem(eui.NewButton(&eui.ItemData{Text: "V2", Size: eui.Point{X: 80, Y: 20}, FontSize: 8}))

	// Horizontal reverse flow
	hrFlow := &eui.ItemData{
		ItemType: eui.ITEM_FLOW,
		FlowType: eui.FLOW_HORIZONTAL_REV,
		Size:     eui.Point{X: 440, Y: 32},
		Fixed:    true,
	}
	mainFlow.AddItem(hrFlow)
	hrFlow.AddItem(eui.NewButton(&eui.ItemData{Text: "HR1", Size: eui.Point{X: 60, Y: 24}, FontSize: 8}))
	hrFlow.AddItem(eui.NewButton(&eui.ItemData{Text: "HR2", Size: eui.Point{X: 60, Y: 24}, FontSize: 8}))
	hrFlow.AddItem(eui.NewButton(&eui.ItemData{Text: "HR3", Size: eui.Point{X: 60, Y: 24}, FontSize: 8}))

	// Vertical reverse flow
	vrFlow := &eui.ItemData{
		ItemType: eui.ITEM_FLOW,
		FlowType: eui.FLOW_VERTICAL_REV,
		Size:     eui.Point{X: 120, Y: 72},
		Fixed:    true,
	}
	mainFlow.AddItem(vrFlow)
	vrFlow.AddItem(eui.NewButton(&eui.ItemData{Text: "VR1", Size: eui.Point{X: 80, Y: 20}, FontSize: 8}))
	vrFlow.AddItem(eui.NewButton(&eui.ItemData{Text: "VR2", Size: eui.Point{X: 80, Y: 20}, FontSize: 8}))

	// Nested flows: horizontal flow containing a vertical flow and another horizontal flow
	nested := &eui.ItemData{
		ItemType: eui.ITEM_FLOW,
		FlowType: eui.FLOW_HORIZONTAL,
		Size:     eui.Point{X: 440, Y: 80},
		Fixed:    true,
	}
	mainFlow.AddItem(nested)

	nestV := &eui.ItemData{ItemType: eui.ITEM_FLOW, FlowType: eui.FLOW_VERTICAL, Size: eui.Point{X: 100, Y: 60}, Fixed: true}
	nestV.AddItem(eui.NewButton(&eui.ItemData{Text: "NV1", Size: eui.Point{X: 60, Y: 20}, FontSize: 8}))
	nestV.AddItem(eui.NewButton(&eui.ItemData{Text: "NV2", Size: eui.Point{X: 60, Y: 20}, FontSize: 8}))

	// add another level of nesting to test deeper flows
	deepH := &eui.ItemData{ItemType: eui.ITEM_FLOW, FlowType: eui.FLOW_HORIZONTAL, Size: eui.Point{X: 80, Y: 24}, Fixed: true}
	deepH.AddItem(eui.NewButton(&eui.ItemData{Text: "DNH1", Size: eui.Point{X: 40, Y: 20}, FontSize: 8}))
	deepH.AddItem(eui.NewButton(&eui.ItemData{Text: "DNH2", Size: eui.Point{X: 40, Y: 20}, FontSize: 8}))

	deepV := &eui.ItemData{ItemType: eui.ITEM_FLOW, FlowType: eui.FLOW_VERTICAL, Size: eui.Point{X: 60, Y: 44}, Fixed: true}
	deepV.AddItem(deepH)
	deepV.AddItem(eui.NewButton(&eui.ItemData{Text: "DNV", Size: eui.Point{X: 40, Y: 20}, FontSize: 8}))

	nestV.AddItem(deepV)
	nested.AddItem(nestV)

	nestH := &eui.ItemData{ItemType: eui.ITEM_FLOW, FlowType: eui.FLOW_HORIZONTAL, Size: eui.Point{X: 100, Y: 32}, Fixed: true}
	nestH.AddItem(eui.NewButton(&eui.ItemData{Text: "NH1", Size: eui.Point{X: 40, Y: 24}, FontSize: 8}))
	nestH.AddItem(eui.NewButton(&eui.ItemData{Text: "NH2", Size: eui.Point{X: 40, Y: 24}, FontSize: 8}))
	nested.AddItem(nestH)

	// Pinned buttons inside the flow so we can test pinning behavior within flows
	mainFlow.AddItem(eui.NewButton(&eui.ItemData{Text: "TopLeft", PinTo: eui.PIN_TOP_LEFT, Size: eui.Point{X: 60, Y: 20}, FontSize: 8}))
	mainFlow.AddItem(eui.NewButton(&eui.ItemData{Text: "TopRight", PinTo: eui.PIN_TOP_RIGHT, Size: eui.Point{X: 60, Y: 20}, FontSize: 8}))
	mainFlow.AddItem(eui.NewButton(&eui.ItemData{Text: "BottomLeft", PinTo: eui.PIN_BOTTOM_LEFT, Size: eui.Point{X: 80, Y: 20}, FontSize: 8}))
	mainFlow.AddItem(eui.NewButton(&eui.ItemData{Text: "BottomRight", PinTo: eui.PIN_BOTTOM_RIGHT, Size: eui.Point{X: 80, Y: 20}, FontSize: 8}))
	mainFlow.AddItem(eui.NewButton(&eui.ItemData{Text: "Center", PinTo: eui.PIN_MID_CENTER, Size: eui.Point{X: 60, Y: 20}, FontSize: 8}))

	// Overlay flow demonstrating screen pinning
	over := &eui.ItemData{
		ItemType: eui.ITEM_FLOW,
		FlowType: eui.FLOW_HORIZONTAL,
		Size:     eui.Point{X: 88, Y: 32},
		Position: eui.Point{X: 4, Y: 4},
		PinTo:    eui.PIN_TOP_LEFT,
	}
	over.AddItem(eui.NewButton(&eui.ItemData{Text: "Overlay", Size: eui.Point{X: 80, Y: 24}, FontSize: 8}))
	eui.AddOverlayFlow(over)

	return win
}
