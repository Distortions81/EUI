package main

import (
	"fmt"

	"github.com/Distortions81/EUI/eui"
)

// makeShowcaseWindow creates a window demonstrating most widget types.
func makeShowcaseWindow() *eui.WindowData {
	win := eui.NewWindow()
	win.Title = "Showcase"
	win.Size = eui.Point{X: 400, Y: 420}
	win.Position = eui.Point{X: 8, Y: 8}
	win.AutoSize = true
	win.Open = true
	win.Movable = true
	win.Resizable = true
	win.Closable = false

	mainFlow := &eui.ItemData{
		ItemType: eui.ITEM_FLOW,
		Size:     win.Size,
		FlowType: eui.FLOW_VERTICAL,
	}
	win.AddItem(mainFlow)

	titleText, _ := eui.NewText()
	titleText.Text = "Demonstration of widgets"
	titleText.Size = eui.Point{X: 380, Y: 32}
	titleText.FontSize = 10
	mainFlow.AddItem(titleText)

	btnText, btnTextEvents := eui.NewButton()
	btnText.Text = "Text Button"
	btnText.Size = eui.Point{X: 100, Y: 24}
	btnText.FontSize = 8
	btnTextEvents.Handle = func(ev eui.UIEvent) {
		if ev.Type == eui.EventClick {
			setStatus("Text Button clicked")
		}
	}
	mainFlow.AddItem(btnText)

	chk, chkEvents := eui.NewCheckbox()
	chk.Text = "Enable option"
	chk.Size = eui.Point{X: 140, Y: 24}
	chk.FontSize = 8
	chk.Checked = true
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

	radioA, radioAEvents := eui.NewRadio()
	radioA.Text = "Choice A"
	radioA.RadioGroup = "grp1"
	radioA.Size = eui.Point{X: 140, Y: 24}
	radioA.FontSize = 8
	radioA.Checked = true
	radioAEvents.Handle = func(ev eui.UIEvent) {
		if ev.Type == eui.EventRadioSelected {
			setStatus("Selected Choice A")
		}
	}
	radioB, radioBEvents := eui.NewRadio()
	radioB.Text = "Choice B"
	radioB.RadioGroup = "grp1"
	radioB.Size = eui.Point{X: 140, Y: 24}
	radioB.FontSize = 8
	radioBEvents.Handle = func(ev eui.UIEvent) {
		if ev.Type == eui.EventRadioSelected {
			setStatus("Selected Choice B")
		}
	}
	mainFlow.AddItem(radioA)
	mainFlow.AddItem(radioB)

	slider, sliderEvents := eui.NewSlider()
	slider.Label = "Float Slider"
	slider.Size = eui.Point{X: 180, Y: 24}
	slider.MinValue = 0
	slider.MaxValue = 100
	slider.IntOnly = false
	slider.FontSize = 8
	slider.Value = 46.2
	sliderEvents.Handle = func(ev eui.UIEvent) {
		if ev.Type == eui.EventSliderChanged {
			setStatus(fmt.Sprintf("Float Slider changed: %.2f", ev.Value))
		}
	}
	mainFlow.AddItem(slider)

	intSlider, intSliderEvents := eui.NewSlider()
	intSlider.Label = "Int Slider"
	intSlider.Size = eui.Point{X: 180, Y: 24}
	intSlider.MinValue = 0
	intSlider.MaxValue = 10
	intSlider.IntOnly = true
	intSlider.FontSize = 8
	intSlider.Value = 3
	intSliderEvents.Handle = func(ev eui.UIEvent) {
		if ev.Type == eui.EventSliderChanged {
			setStatus(fmt.Sprintf("Int Slider changed: %.0f", ev.Value))
		}
	}
	mainFlow.AddItem(intSlider)

	input, inputEvents := eui.NewInput()
	input.Label = "Text Field"
	input.Text = "Text Text!"
	input.Size = eui.Point{X: 180, Y: 24}
	input.FontSize = 8
	input.Action = func() { setStatus("Text Field focused") }
	inputEvents.Handle = func(ev eui.UIEvent) {
		if ev.Type == eui.EventInputChanged {
			setStatus("Input: " + ev.Text)
		}
	}
	mainFlow.AddItem(input)

	dropdown, dropdownEvents := eui.NewDropdown()
	dropdown.Label = "Select Option"
	dropdown.Size = eui.Point{X: 180, Y: 24}
	dropdown.FontSize = 8
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
	if btn, ev := eui.NewButton(); btn != nil {
		btn.Text = "One"
		btn.Size = eui.Point{X: 60, Y: 24}
		btn.FontSize = 8
		ev.Handle = func(ev eui.UIEvent) {
			if ev.Type == eui.EventClick {
				setStatus("Button One clicked")
			}
		}
		hFlow.AddItem(btn)
	}
	if btn, ev := eui.NewButton(); btn != nil {
		btn.Text = "Two"
		btn.Size = eui.Point{X: 60, Y: 24}
		btn.FontSize = 8
		ev.Handle = func(ev eui.UIEvent) {
			if ev.Type == eui.EventClick {
				setStatus("Button Two clicked")
			}
		}
		hFlow.AddItem(btn)
	}
	if btn, ev := eui.NewButton(); btn != nil {
		btn.Text = "Three"
		btn.Size = eui.Point{X: 60, Y: 24}
		btn.FontSize = 8
		ev.Handle = func(ev eui.UIEvent) {
			if ev.Type == eui.EventClick {
				setStatus("Button Three clicked")
			}
		}
		hFlow.AddItem(btn)
	}
	if btn, ev := eui.NewButton(); btn != nil {
		btn.Text = "Four"
		btn.Size = eui.Point{X: 60, Y: 24}
		btn.FontSize = 8
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
	if txt, _ := eui.NewText(); txt != nil {
		txt.Text = "Tab 1 content"
		txt.Size = eui.Point{X: 100, Y: 32}
		txt.FontSize = 8
		tabFlow.Tabs[0].AddItem(txt)
	}
	if txt, _ := eui.NewText(); txt != nil {
		txt.Text = "Tab 2 content"
		txt.Size = eui.Point{X: 100, Y: 32}
		txt.FontSize = 8
		tabFlow.Tabs[1].AddItem(txt)
	}

	return win
}
