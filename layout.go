package main

import (
	"embed"
	"encoding/json"
	"path/filepath"
)

//go:embed themes/layout/*.json
var embeddedLayouts embed.FS

// LayoutTheme controls spacing and padding used by widgets.
type LayoutNumbers struct {
	Window   float32
	Button   float32
	Text     float32
	Checkbox float32
	Radio    float32
	Input    float32
	Slider   float32
	Dropdown float32
	Tab      float32
}

type LayoutTheme struct {
	SliderValueGap   float32
	DropdownArrowPad float32
	TextPadding      float32

	Fillet    LayoutNumbers
	Border    LayoutNumbers
	BorderPad LayoutNumbers
}

var defaultLayout = &LayoutTheme{
	SliderValueGap:   16,
	DropdownArrowPad: 8,
	TextPadding:      4,
	Fillet: LayoutNumbers{
		Window:   4,
		Button:   8,
		Text:     0,
		Checkbox: 8,
		Radio:    8,
		Input:    4,
		Slider:   4,
		Dropdown: 4,
		Tab:      4,
	},
	Border: LayoutNumbers{
		Window:   0,
		Button:   0,
		Text:     0,
		Checkbox: 0,
		Radio:    0,
		Input:    0,
		Slider:   0,
		Dropdown: 0,
		Tab:      0,
	},
	BorderPad: LayoutNumbers{
		Window:   0,
		Button:   4,
		Text:     4,
		Checkbox: 4,
		Radio:    4,
		Input:    2,
		Slider:   2,
		Dropdown: 2,
		Tab:      2,
	},
}

var (
	currentLayout     = defaultLayout
	currentLayoutName = "Default"
)

func LoadLayout(name string) error {
	data, err := embeddedLayouts.ReadFile(filepath.Join("themes/layout", name+".json"))
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data, currentLayout); err != nil {
		return err
	}
	currentLayoutName = name
	if currentTheme != nil {
		applyLayoutToTheme(currentTheme)
		applyThemeToAll()
	}
	return nil
}

func applyLayoutToTheme(th *Theme) {
	if th == nil || currentLayout == nil {
		return
	}
	th.Window.Fillet = currentLayout.Fillet.Window
	th.Window.Border = currentLayout.Border.Window
	th.Window.BorderPad = currentLayout.BorderPad.Window

	th.Button.Fillet = currentLayout.Fillet.Button
	th.Button.Border = currentLayout.Border.Button
	th.Button.BorderPad = currentLayout.BorderPad.Button

	th.Text.Fillet = currentLayout.Fillet.Text
	th.Text.Border = currentLayout.Border.Text
	th.Text.BorderPad = currentLayout.BorderPad.Text

	th.Checkbox.Fillet = currentLayout.Fillet.Checkbox
	th.Checkbox.Border = currentLayout.Border.Checkbox
	th.Checkbox.BorderPad = currentLayout.BorderPad.Checkbox

	th.Radio.Fillet = currentLayout.Fillet.Radio
	th.Radio.Border = currentLayout.Border.Radio
	th.Radio.BorderPad = currentLayout.BorderPad.Radio

	th.Input.Fillet = currentLayout.Fillet.Input
	th.Input.Border = currentLayout.Border.Input
	th.Input.BorderPad = currentLayout.BorderPad.Input

	th.Slider.Fillet = currentLayout.Fillet.Slider
	th.Slider.Border = currentLayout.Border.Slider
	th.Slider.BorderPad = currentLayout.BorderPad.Slider

	th.Dropdown.Fillet = currentLayout.Fillet.Dropdown
	th.Dropdown.Border = currentLayout.Border.Dropdown
	th.Dropdown.BorderPad = currentLayout.BorderPad.Dropdown

	th.Tab.Fillet = currentLayout.Fillet.Tab
	th.Tab.Border = currentLayout.Border.Tab
	th.Tab.BorderPad = currentLayout.BorderPad.Tab
}
