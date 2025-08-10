//go:build test

package eui

import (
	"bytes"
	"log"
	"testing"
	"time"
)

func TestEmitDeliversWhenChannelFull(t *testing.T) {
	var buf bytes.Buffer
	orig := log.Writer()
	log.SetOutput(&buf)
	defer log.SetOutput(orig)

	h := &EventHandler{Events: make(chan UIEvent, 1)}
	h.Events <- UIEvent{Type: EventClick}

	done := make(chan struct{})
	go func() {
		h.Emit(UIEvent{Type: EventSliderChanged})
		close(done)
	}()

	time.Sleep(10 * time.Millisecond)
	<-h.Events

	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("emit did not complete")
	}

	ev := <-h.Events
	if ev.Type != EventSliderChanged {
		t.Fatalf("expected EventSliderChanged, got %v", ev.Type)
	}

	if buf.Len() != 0 {
		t.Fatalf("unexpected log output: %s", buf.String())
	}
}
