package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// Theme bundles all style information for windows and widgets.
type Theme struct {
	Window   windowData
	Button   itemData
	Text     itemData
	Checkbox itemData
	Radio    itemData
	Input    itemData
	Slider   itemData
	Dropdown itemData
}

// LoadTheme reads a theme JSON file from the themes directory and
// sets it as the current theme without modifying existing windows.
func LoadTheme(name string) error {
	file := filepath.Join("themes", name+".json")
	data, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	// Start with the compiled in defaults
	th := *baseTheme
	if err := json.Unmarshal(data, &th); err != nil {
		return err
	}
	currentTheme = &th
	currentThemeName = name
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
