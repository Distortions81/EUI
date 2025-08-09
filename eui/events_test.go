//go:build test

package eui

import (
	"bytes"
	"log"
	"testing"
)

func TestEmitLogsWhenChannelFull(t *testing.T) {
	var buf bytes.Buffer
	orig := log.Writer()
	log.SetOutput(&buf)
	defer log.SetOutput(orig)

	h := &EventHandler{Events: make(chan UIEvent, 1)}
	h.Events <- UIEvent{Type: EventClick}

	h.Emit(UIEvent{Type: EventSliderChanged})

	if buf.Len() == 0 {
		t.Fatal("expected log output when event dropped, got none")
	}
}
