package main

import "image/color"

func makeTestWindow() *windowData {
	//Done button
	newButton := NewButton(&itemData{
		Text:     "Generate",
		Size:     point{X: 128, Y: 64},
		FontSize: 18})

	newFlow := &itemData{
		ItemType: ITEM_FLOW,
		Size:     point{X: 128 + 16, Y: 64 + 16},
		Position: point{X: 16, Y: 16},
		PinTo:    PIN_BOTTOM_RIGHT,
		Contents: []*itemData{newButton},
	}
	//Scaleup button
	newScaleup := NewButton(&itemData{
		Text:     "Scale Up +",
		Size:     point{X: 128, Y: 24},
		Position: point{X: 16, Y: 16 + 16 + 24},
		PinTo:    PIN_BOTTOM_LEFT,
		FontSize: 12})

	//Scaledown button
	newScaledown := NewButton(&itemData{
		Text:     "Scale Down -",
		Size:     point{X: 128, Y: 24},
		Position: point{X: 16, Y: 16},
		PinTo:    PIN_BOTTOM_LEFT,
		FontSize: 12})

	//Text
	newText := NewText(&itemData{
		ItemType:  ITEM_TEXT,
		Text:      "Click 'generate' to\ngenerate a new code.",
		FontSize:  18,
		Size:      point{X: 256, Y: 128},
		Position:  point{X: 16, Y: 16},
		PinTo:     PIN_TOP_LEFT,
		TextColor: color.RGBA{R: 255, G: 255, B: 255, A: 255}})

	newWindow := NewWindow(
		&windowData{
			TitleHeight: 24,
			Title:       "Test Window",
			Size:        point{X: 350, Y: 300},
			Position:    point{X: 32, Y: 32},
			Contents: []*itemData{
				newFlow, newText, newScaleup, newScaledown},
		})

	//Gen button actions
	newButton.Action = func() {
		newButton.Text = "Okay"
		newText.Text = "Secret code: 1234"
		newButton.Action = func() {
			newWindow.RemoveWindow()
		}
	}

	newScaleup.Action = func() {
		if uiScale < 8 {
			SetUIScale(uiScale + 0.1)
		}
	}

	newScaledown.Action = func() {
		if uiScale > 0.2 {
			SetUIScale(uiScale - 0.1)
		}
	}

	return newWindow
}
