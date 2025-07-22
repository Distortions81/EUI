package main

import (
	eui "EUI/eui"
	"fmt"
)

// makeShowcaseWindow creates a window demonstrating most widget types.
func makeShowcaseWindow() *eui.WindowData {
	win := eui.NewWindow(&eui.WindowData{
		Title:     "Showcase",
		Size:      eui.Point{X: 400, Y: 420},
		Position:  eui.Point{X: 8, Y: 8},
		AutoSize:  true,
		Open:      true,
		Movable:   true,
		Resizable: true,
		Closable:  false,
	})

	mainFlow := &eui.ItemData{
		ItemType: eui.ITEM_FLOW,
		Size:     win.Size,
		FlowType: eui.FLOW_VERTICAL,
	}
	win.AddItem(mainFlow)

	titleText, _ := eui.NewText(&eui.ItemData{Text: "Demonstration of widgets", Size: eui.Point{X: 380, Y: 32}, FontSize: 10})
	mainFlow.AddItem(titleText)

	btnText, btnTextEvents := eui.NewButton(&eui.ItemData{Text: "Text Button", Size: eui.Point{X: 100, Y: 24}, FontSize: 8})
	btnTextEvents.Handle = func(ev eui.UIEvent) {
		if ev.Type == eui.EventClick {
			setStatus("Text Button clicked")
		}
	}
	mainFlow.AddItem(btnText)

	chk, chkEvents := eui.NewCheckbox(&eui.ItemData{Text: "Enable option", Size: eui.Point{X: 140, Y: 24}, FontSize: 8, Checked: true})
	chkEvents.Handle = func(ev eui.UIEvent) {
		if ev.Type == eui.EventCheckboxChanged {
			if ev.Checked {
				setStatus("Checkbox enabled")
			} else {
				setStatus("Checkbox disabled")
			}
		}
	}
	mainFlow.AddItem(chk)

	radioA, radioAEvents := eui.NewRadio(&eui.ItemData{Text: "Choice A", RadioGroup: "grp1", Size: eui.Point{X: 140, Y: 24}, FontSize: 8, Checked: true})
	radioAEvents.Handle = func(ev eui.UIEvent) {
		if ev.Type == eui.EventRadioSelected {
			setStatus("Selected Choice A")
		}
	}
	radioB, radioBEvents := eui.NewRadio(&eui.ItemData{Text: "Choice B", RadioGroup: "grp1", Size: eui.Point{X: 140, Y: 24}, FontSize: 8})
	radioBEvents.Handle = func(ev eui.UIEvent) {
		if ev.Type == eui.EventRadioSelected {
			setStatus("Selected Choice B")
		}
	}
	mainFlow.AddItem(radioA)
	mainFlow.AddItem(radioB)

	slider, sliderEvents := eui.NewSlider(&eui.ItemData{Label: "Float Slider", Size: eui.Point{X: 180, Y: 24}, MinValue: 0, MaxValue: 100, IntOnly: false, FontSize: 8, Value: 46.2})
	sliderEvents.Handle = func(ev eui.UIEvent) {
		if ev.Type == eui.EventSliderChanged {
			setStatus(fmt.Sprintf("Float Slider changed: %.2f", ev.Value))
		}
	}
	mainFlow.AddItem(slider)

	intSlider, intSliderEvents := eui.NewSlider(&eui.ItemData{Label: "Int Slider", Size: eui.Point{X: 180, Y: 24}, MinValue: 0, MaxValue: 10, IntOnly: true, FontSize: 8, Value: 3})
	intSliderEvents.Handle = func(ev eui.UIEvent) {
		if ev.Type == eui.EventSliderChanged {
			setStatus(fmt.Sprintf("Int Slider changed: %.0f", ev.Value))
		}
	}
	mainFlow.AddItem(intSlider)

	input, inputEvents := eui.NewInput(&eui.ItemData{Label: "Text Field", Text: "Text Text!", Size: eui.Point{X: 180, Y: 24}, FontSize: 8})
	input.Action = func() { setStatus("Text Field focused") }
	inputEvents.Handle = func(ev eui.UIEvent) {
		if ev.Type == eui.EventInputChanged {
			setStatus("Input: " + ev.Text)
		}
	}
	mainFlow.AddItem(input)

	dropdown, dropdownEvents := eui.NewDropdown(&eui.ItemData{Label: "Select Option", Size: eui.Point{X: 180, Y: 24}, FontSize: 8})
	dropdown.Options = []string{"First", "Second", "Third", "Fourth"}
	dropdown.HoverIndex = -1
	dropdownEvents.Handle = func(ev eui.UIEvent) {
		if ev.Type == eui.EventDropdownSelected {
			if ev.Index >= 0 && ev.Index < len(dropdown.Options) {
				setStatus("Dropdown selected: " + dropdown.Options[ev.Index])
			}
		}
	}
	mainFlow.AddItem(dropdown)

	hFlow := &eui.ItemData{
		ItemType:   eui.ITEM_FLOW,
		FlowType:   eui.FLOW_HORIZONTAL,
		Size:       eui.Point{X: 380, Y: 40},
		Fixed:      true,
		Scrollable: true,
	}
	mainFlow.AddItem(hFlow)
	if btn, ev := eui.NewButton(&eui.ItemData{Text: "One", Size: eui.Point{X: 60, Y: 24}, FontSize: 8}); btn != nil {
		ev.Handle = func(ev eui.UIEvent) {
			if ev.Type == eui.EventClick {
				setStatus("Button One clicked")
			}
		}
		hFlow.AddItem(btn)
	}
	if btn, ev := eui.NewButton(&eui.ItemData{Text: "Two", Size: eui.Point{X: 60, Y: 24}, FontSize: 8}); btn != nil {
		ev.Handle = func(ev eui.UIEvent) {
			if ev.Type == eui.EventClick {
				setStatus("Button Two clicked")
			}
		}
		hFlow.AddItem(btn)
	}
	if btn, ev := eui.NewButton(&eui.ItemData{Text: "Three", Size: eui.Point{X: 60, Y: 24}, FontSize: 8}); btn != nil {
		ev.Handle = func(ev eui.UIEvent) {
			if ev.Type == eui.EventClick {
				setStatus("Button Three clicked")
			}
		}
		hFlow.AddItem(btn)
	}
	if btn, ev := eui.NewButton(&eui.ItemData{Text: "Four", Size: eui.Point{X: 60, Y: 24}, FontSize: 8}); btn != nil {
		ev.Handle = func(ev eui.UIEvent) {
			if ev.Type == eui.EventClick {
				setStatus("Button Four clicked")
			}
		}
		hFlow.AddItem(btn)
	}

	tabFlow := &eui.ItemData{
		ItemType:   eui.ITEM_FLOW,
		FlowType:   eui.FLOW_VERTICAL,
		Size:       eui.Point{X: 380, Y: 120},
		Fixed:      true,
		Color:      eui.ColorDarkGray,
		ClickColor: eui.ColorDarkTeal,
		FontSize:   8,
		Tabs: []*eui.ItemData{
			{Name: "Tab 1", ItemType: eui.ITEM_FLOW, FlowType: eui.FLOW_VERTICAL},
			{Name: "Tab 2", ItemType: eui.ITEM_FLOW, FlowType: eui.FLOW_VERTICAL},
		},
	}
	mainFlow.AddItem(tabFlow)
	if txt, _ := eui.NewText(&eui.ItemData{Text: "Tab 1 content", Size: eui.Point{X: 100, Y: 32}, FontSize: 8}); txt != nil {
		tabFlow.Tabs[0].AddItem(txt)
	}
	if txt, _ := eui.NewText(&eui.ItemData{Text: "Tab 2 content", Size: eui.Point{X: 100, Y: 32}, FontSize: 8}); txt != nil {
		tabFlow.Tabs[1].AddItem(txt)
	}

	return win
}
