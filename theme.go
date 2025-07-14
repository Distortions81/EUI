package main

import (
	"embed"
	"encoding/json"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

//go:embed themes/*.json
var embeddedThemes embed.FS

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
	data, err := embeddedThemes.ReadFile(filepath.Join("themes", name+".json"))
	if err != nil {
		file := filepath.Join("themes", name+".json")
		data, err = os.ReadFile(file)
		if err != nil {
			return err
		}
	}
	// Start with the compiled in defaults
	th := *baseTheme
	if err := json.Unmarshal(data, &th); err != nil {
		return err
	}
	currentTheme = &th
	currentThemeName = name
	applyThemeToAll()
	return nil
}

// listThemes returns the available theme names from the themes directory
func listThemes() ([]string, error) {
	entries, err := fs.ReadDir(embeddedThemes, "themes")
	if err != nil {
		entries, err = os.ReadDir("themes")
		if err != nil {
			return nil, err
		}
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

// applyThemeToAll updates all existing windows to use the current theme.
func applyThemeToAll() {
	if currentTheme == nil {
		return
	}
	for _, win := range windows {
		applyThemeToWindow(win)
	}
}

// applyThemeToWindow merges the current theme's window settings into the given
// window and recursively updates contained items.
func applyThemeToWindow(win *windowData) {
	if win == nil || currentTheme == nil {
		return
	}
	mergeData(win, &currentTheme.Window)
	win.Theme = currentTheme
	for _, item := range win.Contents {
		applyThemeToItem(item)
	}
}

// applyThemeToItem merges style data from the current theme based on item type
// and recursively processes child items.
func applyThemeToItem(it *itemData) {
	if it == nil || currentTheme == nil {
		return
	}
	var src *itemData
	switch it.ItemType {
	case ITEM_BUTTON:
		src = &currentTheme.Button
	case ITEM_TEXT:
		src = &currentTheme.Text
	case ITEM_CHECKBOX:
		src = &currentTheme.Checkbox
	case ITEM_RADIO:
		src = &currentTheme.Radio
	case ITEM_INPUT:
		src = &currentTheme.Input
	case ITEM_SLIDER:
		src = &currentTheme.Slider
	case ITEM_DROPDOWN:
		src = &currentTheme.Dropdown
	}
	if src != nil {
		mergeData(it, src)
	}
	for _, child := range it.Contents {
		applyThemeToItem(child)
	}
	for _, tab := range it.Tabs {
		applyThemeToItem(tab)
	}
}
