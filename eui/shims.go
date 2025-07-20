package eui

import "github.com/hajimehoshi/ebiten/v2"

// Update processes input and updates window state.
// Programs embedding the UI can call this from their Ebiten Update handler.
func Update() error {
	var g Game
	return g.Update()
}

// Draw renders the UI to the provided screen image.
// Call this from your Ebiten Draw function.
func Draw(screen *ebiten.Image) {
	var g Game
	g.Draw(screen)
}

// Layout reports the dimensions for the game's screen.
// Pass Ebiten's outside size values to this from your Layout function.
func Layout(outsideWidth, outsideHeight int) (int, int) {
	var g Game
	return g.Layout(outsideWidth, outsideHeight)
}
