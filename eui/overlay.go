package eui

import "log"

func AddOverlayFlow(flow *itemData) {
	for _, ov := range overlays {
		if ov == flow {
			log.Println("Overlay already exists")
			return
		}
	}
	if currentTheme != nil && flow.Theme != currentTheme {
		applyThemeToItem(flow)
	}
	overlays = append(overlays, flow)
	flow.resizeFlow(flow.GetSize())
}
