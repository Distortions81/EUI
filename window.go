package main

import (
	"image/color"
	"log"
	"reflect"
)

var DefaultTheme = &WindowData{
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

var DefaultButton = &ItemData{
	Text:      "Button",
	ItemType:  ITEM_BUTTON,
	Size:      Point{X: 128, Y: 64},
	Position:  Point{X: 10, Y: 10},
	FontSize:  24,
	LineSpace: 1.2,
	Enabled:   true,

	Fillet: 8,
	Filled: true,

	TextColor:  color.RGBA{R: 0, G: 0, B: 0, A: 255},
	Color:      color.RGBA{R: 128, G: 128, B: 128, A: 255},
	HoverColor: color.RGBA{R: 192, G: 192, B: 192, A: 255},
	ClickColor: color.RGBA{R: 64, G: 64, B: 64, A: 255},
}

var DefaultText = &ItemData{
	Text:      "Sample text:\nThe quick brown fox\njumps over the lazy dog.",
	ItemType:  ITEM_TEXT,
	Size:      Point{X: 128, Y: 128},
	Position:  Point{X: 16, Y: 24 + 16},
	FontSize:  24,
	LineSpace: 1.2,
	Enabled:   true,
	TextColor: color.RGBA{R: 255, G: 255, B: 255, A: 255},
}

func MergeData(original interface{}, updates interface{}) interface{} {
	// Ensure both original and updates are pointers to structs
	origVal := reflect.ValueOf(original)
	updVal := reflect.ValueOf(updates)

	// Check that both are pointers to structs
	if origVal.Kind() != reflect.Ptr || updVal.Kind() != reflect.Ptr {
		panic("Both original and updates must be pointers to structs")
	}

	// Get the elements (dereference the pointers)
	origVal = origVal.Elem()
	updVal = updVal.Elem()

	// Ensure that after dereferencing, both are structs
	if origVal.Kind() != reflect.Struct || updVal.Kind() != reflect.Struct {
		panic("Both original and updates must be structs")
	}

	// Iterate through the fields of the updates struct
	for i := 0; i < updVal.NumField(); i++ {
		origField := origVal.Field(i)
		updField := updVal.Field(i)

		// Check if the update field has a non-zero value
		if !isZeroValue(updField) && origField.CanSet() {
			// Set the original field to the update field's value
			origField.Set(updField)
		}
	}

	return original
}

func isZeroValue(value reflect.Value) bool {
	return reflect.DeepEqual(value.Interface(), reflect.Zero(value.Type()).Interface())
}

func (target *WindowData) AddWindow() {
	for _, win := range Windows {
		if win == target {
			log.Println("Window already exists")
			return
		}
	}
	target.CalcUIScale()
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
	newWindow := *DefaultTheme
	MergeData(&newWindow, win)
	return &newWindow
}

func NewButton(item *ItemData) *ItemData {
	newItem := *DefaultButton
	MergeData(&newItem, item)
	return &newItem
}

func NewText(item *ItemData) *ItemData {
	newItem := *DefaultText
	MergeData(&newItem, item)
	return &newItem
}
