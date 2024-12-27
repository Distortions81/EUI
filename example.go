package main

import "image/color"

func makeTestWindow() *WindowData {
	//Done button
	newButton := NewButton(&ItemData{
		Text: "Generate",
		Position: Point{
			X: 300 - 128 - 16,
			Y: 300 - 24 - 64 - 16},
		Size:     Point{X: 128, Y: 64},
		FontSize: 18})

	//Scaleup button
	newScaleup := NewButton(&ItemData{
		Text: "Scale Up +",
		Position: Point{
			X: 16,
			Y: 300 - 24 - 24 - 16},
		Size:     Point{X: 128, Y: 24},
		FontSize: 12})

	//Scaledown button
	newScaledown := NewButton(&ItemData{
		Text: "Scale Down -",
		Position: Point{
			X: 16,
			Y: 300 - 24 - 24 - 24 - 16 - 16},
		Size:     Point{X: 128, Y: 24},
		FontSize: 12})

	//Text
	newText := NewText(&ItemData{
		ItemType: ITEM_TEXT,
		Text:     "Click 'generate' to\ngenerate a new code.",
		FontSize: 18,
		Position: Point{
			X: 16,
			Y: 24 + 16},
		Size:      Point{X: 256, Y: 128},
		TextColor: color.RGBA{R: 255, G: 255, B: 255, A: 255}})

	newWindow := NewWindow(
		&WindowData{
			TitleHeight: 24,
			Title:       "Test Window",
			Size:        Point{X: 300, Y: 300},
			Position:    Point{X: 32, Y: 32},
			Contents: []*ItemData{
				newButton, newText, newScaleup, newScaledown},
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
