package eui

import (
	_ "embed"
	"log"
)

//go:embed fonts/NotoSans-Regular.ttf
var defaultTTF []byte

func init() {
	if err := EnsureFontSource(defaultTTF); err != nil {
		log.Printf("default font load error: %v", err)
	}
	if err := LoadTheme(currentThemeName); err != nil {
		log.Printf("LoadTheme error: %v", err)
	}
	if err := LoadLayout(currentLayoutName); err != nil {
		log.Printf("LoadLayout error: %v", err)
	}
}
