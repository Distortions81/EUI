package main

import (
	"log"
	"reflect"
)

// Merge one struct into another
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

		if !origField.CanSet() {
			continue
		}

		// Booleans need to be copied even when false so themes can
		// explicitly disable features like outlines or drag bars.
		if updField.Kind() == reflect.Bool {
			origField.Set(updField)
			continue
		}

		// Check if the update field has a non-zero value
		if !isZeroValue(updField) {
			// Set the original field to the update field's value
			origField.Set(updField)
		}
	}

	return original
}

func isZeroValue(value reflect.Value) bool {
	return reflect.DeepEqual(value.Interface(), reflect.Zero(value.Type()).Interface())
}

// Add window to window list
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

	if target.AutoSize {
		target.updateAutoSize()
		target.AutoSize = false
	}

	if currentTheme != nil {
		applyThemeToWindow(target)
	}

	if !toBack {
		windows = append(windows, target)
		activeWindow = target
	} else {
		windows = append([]*windowData{target}, windows...)
	}
}

// Remove window from window list, if found.
func (target *windowData) RemoveWindow() {
	for i, win := range windows {
		if win == target { // Compare pointers
			windows = append(windows[:i], windows[i+1:]...)
			return
		}
	}

	log.Println("Window not found")
}

// Create a new window from the default theme
func NewWindow(win *windowData) *windowData {
	if currentTheme == nil {
		currentTheme = baseTheme
	}
	newWindow := currentTheme.Window
	if win != nil {
		mergeData(&newWindow, win)
	}
	newWindow.Theme = currentTheme
	return &newWindow
}

// Create a new button from the default theme
func NewButton(item *itemData) *itemData {
	if currentTheme == nil {
		currentTheme = baseTheme
	}
	newItem := currentTheme.Button
	if item != nil {
		mergeData(&newItem, item)
	}
	return &newItem
}

// Create a new button from the default theme
func NewCheckbox(item *itemData) *itemData {
	if currentTheme == nil {
		currentTheme = baseTheme
	}
	newItem := currentTheme.Checkbox
	if item != nil {
		mergeData(&newItem, item)
	}
	return &newItem
}

// Create a new radio button from the default theme
func NewRadio(item *itemData) *itemData {
	if currentTheme == nil {
		currentTheme = baseTheme
	}
	newItem := currentTheme.Radio
	if item != nil {
		mergeData(&newItem, item)
	}
	return &newItem
}

// Create a new input box from the default theme
func NewInput(item *itemData) *itemData {
	if currentTheme == nil {
		currentTheme = baseTheme
	}
	newItem := currentTheme.Input
	if item != nil {
		mergeData(&newItem, item)
	}
	return &newItem
}

// Create a new slider from the default theme
func NewSlider(item *itemData) *itemData {
	if currentTheme == nil {
		currentTheme = baseTheme
	}
	newItem := currentTheme.Slider
	if item != nil {
		mergeData(&newItem, item)
	}
	return &newItem
}

// Create a new dropdown from the default theme
func NewDropdown(item *itemData) *itemData {
	if currentTheme == nil {
		currentTheme = baseTheme
	}
	newItem := currentTheme.Dropdown
	if item != nil {
		mergeData(&newItem, item)
	}
	return &newItem
}

// Create a new color wheel from the default theme
func NewColorWheel(item *itemData) *itemData {
	if currentTheme == nil {
		currentTheme = baseTheme
	}
	newItem := baseColorWheel
	if item != nil {
		mergeData(&newItem, item)
	}
	return &newItem
}

// Create a new textbox from the default theme
func NewText(item *itemData) *itemData {
	if currentTheme == nil {
		currentTheme = baseTheme
	}
	newItem := currentTheme.Text
	if item != nil {
		mergeData(&newItem, item)
	}
	return &newItem
}

// Bring a window to the front
func (target *windowData) BringForward() {
	for w, win := range windows {
		if win == target {
			windows = append(windows[:w], windows[w+1:]...)
			windows = append(windows, target)
			activeWindow = win
		}
	}
}

// Send a window to the back
func (target *windowData) ToBack() {
	for w, win := range windows {
		if win == target {
			windows = append(windows[:w], windows[w+1:]...)
			windows = append([]*windowData{target}, windows...)
		}
	}
}

// Get window position, considering pinned position
func (pin pinType) getWinPosition(win *windowData) point {
	switch pin {
	case PIN_TOP_LEFT:
		return win.GetPos()
	case PIN_TOP_RIGHT:
		return point{X: float32(screenWidth) - win.GetSize().X - win.GetPos().X, Y: win.GetPos().Y}
	case PIN_TOP_CENTER:
		return point{X: float32(screenWidth/2) - win.GetSize().X/2 + win.GetPos().X, Y: win.GetPos().Y}
	case PIN_MID_LEFT:
		return point{X: win.GetPos().X, Y: float32(screenHeight/2) - win.GetSize().Y/2 + win.GetPos().Y}
	case PIN_MID_CENTER:
		return point{X: float32(screenWidth/2) - win.GetSize().X/2 + win.GetPos().X, Y: float32(screenHeight/2) - (win.GetSize().Y - (win.GetTitleSize())/2) + win.GetPos().Y}
	case PIN_MID_RIGHT:
		return point{X: float32(screenWidth) - win.GetSize().X/2 - win.GetPos().X, Y: float32(screenHeight/2) - (win.GetSize().Y - (win.GetTitleSize())/2) + win.GetPos().Y}
	case PIN_BOTTOM_LEFT:
		return point{X: win.GetPos().X, Y: float32(screenHeight) - (win.GetSize().Y - (win.GetTitleSize())) - win.GetPos().Y}
	case PIN_BOTTOM_CENTER:
		return point{X: float32(screenWidth/2) - (win.GetSize().X / 2) + win.GetPos().X, Y: float32(screenHeight) - (win.GetSize().Y - (win.GetTitleSize())) - win.GetPos().Y}
	case PIN_BOTTOM_RIGHT:
		return point{X: float32(screenWidth) - win.GetSize().X - win.GetPos().X, Y: float32(screenHeight) - (win.GetSize().Y - (win.GetTitleSize())) - win.GetPos().Y}
	default:
		return win.GetPos()
	}
}

// Get item position, considering its pinned position
func (pin pinType) getItemPosition(win *windowData, item *itemData) point {
	switch pin {
	case PIN_TOP_LEFT:
		return item.Position
	case PIN_TOP_RIGHT:
		return point{
			X: item.GetPos().X + float32((win.GetSize().X)) - (item.GetSize().X) - item.GetPos().X,
			Y: item.GetPos().Y}
	case PIN_TOP_CENTER:
		return point{
			X: float32((win.GetSize().X)/2) - (item.GetSize().X)/2,
			Y: item.GetPos().Y}
	case PIN_MID_LEFT:
		return point{
			X: item.GetPos().X,
			Y: float32((win.GetSize().Y)/2) - ((item.GetSize().Y) / 2)}
	case PIN_MID_CENTER:
		return point{
			X: float32((win.GetSize().X)/2) - (item.GetSize().X)/2,
			Y: float32((win.GetSize().Y)/2) - ((item.GetSize().Y) / 2)}
	case PIN_MID_RIGHT:
		return point{
			X: float32((win.GetSize().X)) - (item.GetSize().X)/2 - item.GetPos().X,
			Y: float32((win.GetSize().Y)/2) - ((item.GetSize().Y) / 2)}
	case PIN_BOTTOM_LEFT:
		return point{
			X: item.GetPos().X,
			Y: float32((win.GetSize().Y)-(win.GetTitleSize())*2) - (item.GetSize().Y) - item.GetPos().Y}
	case PIN_BOTTOM_CENTER:
		return point{
			X: float32((win.GetSize().X)/2) - ((item.GetSize().X) / 2),
			Y: float32((win.GetSize().Y)-(win.GetTitleSize())*2) - (item.GetSize().Y) - item.GetPos().Y}
	case PIN_BOTTOM_RIGHT:
		return point{
			X: float32((win.GetSize().X)) - (item.GetSize().X) - item.GetPos().X,
			Y: float32((win.GetSize().Y)-(win.GetTitleSize())*2) - (item.GetSize().Y) - item.GetPos().Y}
	default:
		return item.GetPos()
	}
}

func (win *windowData) getPosition() point {
	pos := win.PinTo.getWinPosition(win)
	pos = pointAdd(pos, point{X: win.Margin * uiScale, Y: win.Margin * uiScale})
	return pos
}

func (item *itemData) getPosition(win *windowData) point {
	pos := item.PinTo.getItemPosition(win, item)
	pos = pointAdd(pos, point{X: item.Margin * uiScale, Y: item.Margin * uiScale})
	return pos
}

// Do the window items overlap?
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
