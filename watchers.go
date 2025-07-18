package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

var (
	themeModTime  time.Time
	layoutModTime time.Time
	modCheckTime  time.Time
)

func init() {
	modCheckTime = time.Now()
	refreshThemeMod()
	refreshLayoutMod()
}

func refreshThemeMod() {
	path := filepath.Join(os.Getenv("PWD"), "themes", "colors", currentThemeName+".json")
	if info, err := os.Stat(path); err == nil {
		themeModTime = info.ModTime()
	} else {
		themeModTime = time.Time{}
	}
}

func refreshLayoutMod() {
	path := filepath.Join(os.Getenv("PWD"), "themes", "layout", currentLayoutName+".json")
	if info, err := os.Stat(path); err == nil {
		layoutModTime = info.ModTime()
	} else {
		layoutModTime = time.Time{}
	}
}

func checkThemeLayoutMods() {
	if time.Since(modCheckTime) < 500*time.Millisecond {
		return
	}
	modCheckTime = time.Now()
	path := filepath.Join(os.Getenv("PWD"), "themes", "colors", currentThemeName+".json")
	if info, err := os.Stat(path); err == nil {
		if info.ModTime().After(themeModTime) {
			fmt.Println("Color theme reload")
			if err := LoadTheme(currentThemeName); err != nil {
				fmt.Printf("Auto reload theme error: %v\n", err)
			}
			themeModTime = info.ModTime()
		}
	} else {
		fmt.Println("Unable to stat " + currentThemeName + ": " + err.Error())
	}

	path = filepath.Join(os.Getenv("PWD"), "themes", "layout", currentLayoutName+".json")
	if info, err := os.Stat(path); err == nil {
		if info.ModTime().After(layoutModTime) {
			fmt.Println("Layout theme reload")
			if err := LoadLayout(currentLayoutName); err != nil {
				fmt.Printf("Auto reload layout error: %v\n", err)
			}
			layoutModTime = info.ModTime()
		}
	} else {
		fmt.Println("Unable to stat " + currentLayoutName + ": " + err.Error())
	}

}
