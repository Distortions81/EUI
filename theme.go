package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
)

// applyJSON merges the JSON object data into the given struct by
// only updating fields present in the JSON. This prevents missing
// values from overwriting existing settings.
func applyJSON(target interface{}, data json.RawMessage) error {
	var obj map[string]json.RawMessage
	if err := json.Unmarshal(data, &obj); err != nil {
		return err
	}
	v := reflect.ValueOf(target).Elem()
	for name, raw := range obj {
		f := v.FieldByName(name)
		if !f.IsValid() || !f.CanSet() {
			continue
		}
		valPtr := reflect.New(f.Type())
		if err := json.Unmarshal(raw, valPtr.Interface()); err != nil {
			return err
		}
		f.Set(valPtr.Elem())
	}
	return nil
}

// LoadTheme reads a theme JSON file from the themes directory
// and sets it as the default style.
func LoadTheme(name string) error {
	file := filepath.Join("themes", name+".json")
	data, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	// Reset to the built-in defaults before applying overrides
	*defaultTheme = baseWindow
	*defaultButton = baseButton
	*defaultText = baseText
	*defaultCheckbox = baseCheckbox
	*defaultRadio = baseRadio
	*defaultInput = baseInput
	*defaultSlider = baseSlider
	*defaultDropdown = baseDropdown

	if v, ok := raw["Window"]; ok {
		if err := applyJSON(defaultTheme, v); err != nil {
			return err
		}
	}
	if v, ok := raw["Button"]; ok {
		if err := applyJSON(defaultButton, v); err != nil {
			return err
		}
	}
	if v, ok := raw["Text"]; ok {
		if err := applyJSON(defaultText, v); err != nil {
			return err
		}
	}
	if v, ok := raw["Checkbox"]; ok {
		if err := applyJSON(defaultCheckbox, v); err != nil {
			return err
		}
	}
	if v, ok := raw["Radio"]; ok {
		if err := applyJSON(defaultRadio, v); err != nil {
			return err
		}
	}
	if v, ok := raw["Input"]; ok {
		if err := applyJSON(defaultInput, v); err != nil {
			return err
		}
	}
	if v, ok := raw["Slider"]; ok {
		if err := applyJSON(defaultSlider, v); err != nil {
			return err
		}
	}
	if v, ok := raw["Dropdown"]; ok {
		if err := applyJSON(defaultDropdown, v); err != nil {
			return err
		}
	}

	// Apply new defaults to all existing windows and items so that
	// the currently displayed UI reflects the loaded theme.
	ApplyTheme()

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

// ApplyTheme updates all existing windows and items using the
// currently loaded default styles. This is used after loading a
// new theme so that the on-screen UI immediately reflects the
// new colors and margins without recreating the windows.
func ApplyTheme() {
	for _, win := range windows {
		applyThemeToWindow(win)
	}
}

func applyThemeToWindow(win *windowData) {
	// Preserve dynamic state that themes shouldn't modify
	open := win.Open
	movable := win.Movable
	resizable := win.Resizable
	closable := win.Closable

	// Merge the window with the defaultTheme values
	mergeData(win, defaultTheme)

	// Restore the preserved state
	win.Open = open
	win.Movable = movable
	win.Resizable = resizable
	win.Closable = closable

	for _, item := range win.Contents {
		applyThemeToItem(item)
	}
}

func applyThemeToItem(it *itemData) {
	// Preserve interactive state that shouldn't be affected by themes
	open := it.Open
	checked := it.Checked
	value := it.Value
	selected := it.Selected
	hoverIndex := it.HoverIndex
	scroll := it.Scroll
	focused := it.Focused

	switch it.ItemType {
	case ITEM_BUTTON:
		mergeData(it, defaultButton)
	case ITEM_TEXT:
		mergeData(it, defaultText)
	case ITEM_CHECKBOX:
		mergeData(it, defaultCheckbox)
	case ITEM_RADIO:
		mergeData(it, defaultRadio)
	case ITEM_INPUT:
		mergeData(it, defaultInput)
	case ITEM_SLIDER:
		mergeData(it, defaultSlider)
	case ITEM_DROPDOWN:
		mergeData(it, defaultDropdown)
	}

	// Restore preserved state
	it.Open = open
	it.Checked = checked
	it.Value = value
	it.Selected = selected
	it.HoverIndex = hoverIndex
	it.Scroll = scroll
	it.Focused = focused

	for _, child := range it.Contents {
		applyThemeToItem(child)
	}
	for _, tab := range it.Tabs {
		for _, sub := range tab.Contents {
			applyThemeToItem(sub)
		}
	}
}
