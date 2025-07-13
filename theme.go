package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// Theme groups the default style data for all widgets.
type Theme struct {
	Window   windowData `json:"Window"`
	Button   itemData   `json:"Button"`
	Text     itemData   `json:"Text"`
	Checkbox itemData   `json:"Checkbox"`
	Radio    itemData   `json:"Radio"`
	Input    itemData   `json:"Input"`
	Slider   itemData   `json:"Slider"`
}

// LoadTheme reads a theme JSON file from the themes directory
// and sets it as the default style.
func LoadTheme(name string) error {
	file := filepath.Join("themes", name+".json")
	data, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	var t Theme
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	mergeData(defaultTheme, &t.Window)
	mergeData(defaultButton, &t.Button)
	mergeData(defaultText, &t.Text)
	mergeData(defaultCheckbox, &t.Checkbox)
	mergeData(defaultRadio, &t.Radio)
	mergeData(defaultInput, &t.Input)
	mergeData(defaultSlider, &t.Slider)
	return nil
}
