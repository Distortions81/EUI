package main

import (
	"log"
	"reflect"
)

func mergeData(original interface{}, updates interface{}) interface{} {
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

func (target *WindowData) AddWindow(toBack bool) {
	for _, win := range windows {
		if win == target {
			log.Println("Window already exists")
			return
		}
	}
	target.calcUIScale()

	if !toBack {
		windows = append(windows, target)
		activeWindow = target
	} else {
		windows = append([]*WindowData{target}, windows...)
	}
}

func (target *WindowData) RemoveWindow() {
	for i, win := range windows {
		if win == target { // Compare pointers
			windows = append(windows[:i], windows[i+1:]...)
			return
		}
	}

	log.Println("Window not found")
}

func NewWindow(win *WindowData) *WindowData {
	newWindow := *defaultTheme
	mergeData(&newWindow, win)
	return &newWindow
}

func NewButton(item *ItemData) *ItemData {
	newItem := *defaultButton
	mergeData(&newItem, item)
	return &newItem
}

func NewText(item *ItemData) *ItemData {
	newItem := *defaultText
	mergeData(&newItem, item)
	return &newItem
}

func (target *WindowData) BringForward() {
	for w, win := range windows {
		if win == target {
			windows = append(windows[:w], windows[w+1:]...)
			windows = append(windows, target)
			activeWindow = win
		}
	}
}

func (target *WindowData) ToBack() {
	for w, win := range windows {
		if win == target {
			windows = append(windows[:w], windows[w+1:]...)
			windows = append([]*WindowData{target}, windows...)
			activeWindow = win
		}
	}
}
