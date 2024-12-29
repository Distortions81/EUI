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

func (target *windowData) AddWindow(toBack bool) {
	for _, win := range windows {
		if win == target {
			log.Println("Window already exists")
			return
		}
	}

	if target.PinTo != PIN_TOP_LEFT {
		target.Movable = false
	}

	if !toBack {
		windows = append(windows, target)
		activeWindow = target
	} else {
		windows = append([]*windowData{target}, windows...)
	}
}

func (target *windowData) RemoveWindow() {
	for i, win := range windows {
		if win == target { // Compare pointers
			windows = append(windows[:i], windows[i+1:]...)
			return
		}
	}

	log.Println("Window not found")
}

func NewWindow(win *windowData) *windowData {
	newWindow := *defaultTheme
	mergeData(&newWindow, win)
	return &newWindow
}

func NewButton(item *itemData) *itemData {
	newItem := *defaultButton
	mergeData(&newItem, item)
	return &newItem
}

func NewText(item *itemData) *itemData {
	newItem := *defaultText
	mergeData(&newItem, item)
	return &newItem
}

func (target *windowData) BringForward() {
	for w, win := range windows {
		if win == target {
			windows = append(windows[:w], windows[w+1:]...)
			windows = append(windows, target)
			activeWindow = win
		}
	}
}

func (target *windowData) ToBack() {
	for w, win := range windows {
		if win == target {
			windows = append(windows[:w], windows[w+1:]...)
			windows = append([]*windowData{target}, windows...)
			activeWindow = win
		}
	}
}

func (pin pinType) getWinPosition(win *windowData) point {
	switch pin {
	case PIN_TOP_LEFT:
		return win.Position
	case PIN_TOP_RIGHT:
		return point{X: float32(screenWidth) - win.Size.X, Y: 0}
	case PIN_TOP_CENTER:
		return point{X: float32(screenWidth/2) - win.Size.X/2, Y: 0}
	case PIN_MID_LEFT:
		return point{X: 0, Y: float32(screenHeight/2) - (win.Size.Y + (win.TitleHeight*uiScale)/2)}
	case PIN_MID_CENTER:
		return point{X: float32(screenWidth/2) - win.Size.X/2, Y: float32(screenHeight/2) - (win.Size.Y - (win.TitleHeight*uiScale)/2)}
	case PIN_MID_RIGHT:
		return point{X: float32(screenWidth) - win.Size.X/2, Y: float32(screenHeight/2) - (win.Size.Y - (win.TitleHeight*uiScale)/2)}
	case PIN_BOTTOM_LEFT:
		return point{X: float32(screenWidth) - win.Size.X, Y: float32(screenHeight) - (win.Size.Y - (win.TitleHeight * uiScale))}
	case PIN_BOTTOM_CENTER:
		return point{X: float32(screenWidth/2) - (win.Size.X / 2), Y: float32(screenHeight) - (win.Size.Y - (win.TitleHeight * uiScale))}
	case PIN_BOTTOM_RIGHT:
		return point{X: float32(screenWidth) - (win.Size.X), Y: float32(screenHeight) - (win.Size.Y - (win.TitleHeight * uiScale))}
	default:
		return win.Position
	}
}

func (pin pinType) getItemPosition(win *windowData, item *itemData) point {
	switch pin {
	case PIN_TOP_LEFT:
		return item.Position
	case PIN_TOP_RIGHT:
		return point{
			X: item.Position.X + float32((win.Size.X)) - (item.Size.X) - item.Position.X,
			Y: item.Position.Y}
	case PIN_TOP_CENTER:
		return point{
			X: float32((win.Size.X)/2) - (item.Size.X)/2,
			Y: item.Position.Y}
	case PIN_MID_LEFT:
		return point{
			X: item.Position.X,
			Y: float32((win.Size.Y)/2) - ((item.Size.Y) / 2)}
	case PIN_MID_CENTER:
		return point{
			X: float32((win.Size.X)/2) - (item.Size.X)/2,
			Y: float32((win.Size.Y)/2) - ((item.Size.Y) / 2)}
	case PIN_MID_RIGHT:
		return point{
			X: float32((win.Size.X)) - (item.Size.X)/2 - item.Position.X,
			Y: float32((win.Size.Y)/2) - ((item.Size.Y) / 2)}
	case PIN_BOTTOM_LEFT:
		return point{
			X: item.Position.X,
			Y: float32((win.Size.Y)-(win.TitleHeight*uiScale)*2) - (item.Size.Y) - item.Position.Y}
	case PIN_BOTTOM_CENTER:
		return point{
			X: float32((win.Size.X)/2) - ((item.Size.X) / 2),
			Y: float32((win.Size.Y)-(win.TitleHeight*uiScale)*2) - (item.Size.Y) - item.Position.Y}
	case PIN_BOTTOM_RIGHT:
		return point{
			X: float32((win.Size.X)) - (item.Size.X) - item.Position.X,
			Y: float32((win.Size.Y)-(win.TitleHeight*uiScale)*2) - (item.Size.Y) - item.Position.Y}
	default:
		return item.Position
	}
}

func (win *windowData) getPosition() point {
	return win.PinTo.getWinPosition(win)
}

func (item *itemData) getPosition(win *windowData) point {
	return item.PinTo.getItemPosition(win, item)
}

func (win windowData) itemOverlap(size point) (bool, bool) {

	rectList := []rect{}

	win.Size = size

	for _, item := range win.Contents {
		rectList = append(rectList, item.getItemRect(&win))
	}

	xc, yc := false, false
	for _, ra := range rectList {
		for _, rb := range rectList {
			if ra == rb {
				continue
			}

			if ra.containsPoint(point{X: rb.X0, Y: rb.Y0}) {
				xc = true
			}
			if ra.containsPoint(point{X: rb.X1, Y: rb.Y1}) {
				yc = true
			}
		}
	}

	return xc, yc
}
