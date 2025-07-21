package main

import eui "EUI/eui"

var statusText *eui.ItemData

func setStatus(msg string) {
	if statusText != nil {
		statusText.Text = msg
		statusText.Dirty = true
	}
}
