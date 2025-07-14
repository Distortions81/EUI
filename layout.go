package main

import (
	"embed"
	"encoding/json"
	"path/filepath"
)

//go:embed themes/layout/*.json
var embeddedLayouts embed.FS

// LayoutTheme controls spacing and padding used by widgets.
type LayoutTheme struct {
	SliderValueGap   float32
	DropdownArrowPad float32
	TextPadding      float32
}

var defaultLayout = &LayoutTheme{
	SliderValueGap:   16,
	DropdownArrowPad: 8,
	TextPadding:      4,
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
	return nil
}
