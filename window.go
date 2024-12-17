package main

import (
	"image/color"
	"log"
	"reflect"
)

var DefaultTheme = WindowData{
	TitleSize:       24,
	Border:          1,
	TitleColor:      color.RGBA{R: 255, G: 255, B: 255, A: 255},
	BorderColor:     color.RGBA{R: 64, G: 64, B: 64, A: 255},
	SizeColor:       color.RGBA{R: 48, G: 48, B: 48, A: 255},
	SizeHoverColor:  color.RGBA{R: 128, G: 128, B: 128, A: 255},
	DragColor:       color.RGBA{R: 48, G: 48, B: 48, A: 255},
	DragHoverColor:  color.RGBA{R: 128, G: 128, B: 128, A: 255},
	HoverColor:      color.RGBA{R: 80, G: 80, B: 80, A: 255},
	ContentsBGColor: color.RGBA{R: 16, G: 16, B: 16, A: 255},

	Movable: true, Closable: true, Resizable: true, Open: true,
}

// MergeStructs merges the non-zero fields of updates into original.
// Uses reflection to handle all fields automatically.
func (original *WindowData) UpdateWindow(updates interface{}) interface{} {
	// Make sure both original and updates are pointers to structs
	origVal := reflect.ValueOf(original).Elem()
	updVal := reflect.ValueOf(updates).Elem()

	// Iterate through the fields of the updates struct
	for i := 0; i < updVal.NumField(); i++ {
		origField := origVal.Field(i)
		updField := updVal.Field(i)

		// Check if the update field has a non-zero value
		if !isZeroValue(updField) {
			// Set the original field to the update field's value
			origField.Set(updField)
		}
	}

	return original
}

// isZeroValue checks if a reflect.Value is the zero value for its type.
func isZeroValue(v reflect.Value) bool {
	zeroValue := reflect.Zero(v.Type())
	return reflect.DeepEqual(v.Interface(), zeroValue.Interface())
}

func (target *WindowData) AddWindow() {
	for _, win := range Windows {
		if win == target {
			log.Println("Window already exists")
			return
		}
	}
	Windows = append(Windows, target)
}

// RemoveWindow removes a *WindowData from the slice if it exists
func (target *WindowData) RemoveWindow() {
	for i, win := range Windows {
		if win == target { // Compare pointers
			Windows = append(Windows[:i], Windows[i+1:]...)
			return
		}
	}

	log.Println("Window not found")
}

func NewWindow(win *WindowData) *WindowData {
	newWindow := &DefaultTheme
	newWindow.UpdateWindow(win)
	return newWindow
}
