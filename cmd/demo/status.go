package main

import "github.com/Distortions81/EUI/eui"

var statusText *eui.ItemData

func setStatus(msg string) {
	if statusText != nil {
		statusText.Text = msg
		statusText.Dirty = true
	}
}
