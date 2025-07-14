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

       // Reset to the built-in defaults before applying overrides
       *defaultTheme = baseWindow
       *defaultButton = baseButton
       *defaultText = baseText
       *defaultCheckbox = baseCheckbox
       *defaultRadio = baseRadio
       *defaultInput = baseInput
       *defaultSlider = baseSlider
       *defaultDropdown = baseDropdown

       mergeData(defaultTheme, &t.Window)
       mergeData(defaultButton, &t.Button)
       mergeData(defaultText, &t.Text)
       mergeData(defaultCheckbox, &t.Checkbox)
       mergeData(defaultRadio, &t.Radio)
       mergeData(defaultInput, &t.Input)
       mergeData(defaultSlider, &t.Slider)
       mergeData(defaultDropdown, &t.Dropdown)

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
       // Merge the window with the defaultTheme values
       mergeData(win, defaultTheme)
       for _, item := range win.Contents {
               applyThemeToItem(item)
       }
}

func applyThemeToItem(it *itemData) {
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

       for _, child := range it.Contents {
               applyThemeToItem(child)
       }
       for _, tab := range it.Tabs {
               for _, sub := range tab.Contents {
                       applyThemeToItem(sub)
               }
       }
}
