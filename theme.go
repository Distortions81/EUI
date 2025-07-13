package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sort"
	"strings"
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
	Dropdown itemData   `json:"Dropdown"`
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
	mergeData(defaultDropdown, &t.Dropdown)
	return nil
}

// listThemes returns the available theme names from the themes directory
func listThemes() ([]string, error) {
	entries, err := os.ReadDir("themes")
	if err != nil {
		return nil, err
	}
	names := []string{}
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		name := strings.TrimSuffix(e.Name(), filepath.Ext(e.Name()))
		names = append(names, name)
	}
	sort.Strings(names)
	return names, nil
}
