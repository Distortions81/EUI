package eui

import "math"

// Sinc returns sin(πx)/(πx) with the special case
// that Sinc(0) returns 1. This implementation avoids
// referencing any undefined helper functions.
func Sinc(x float64) float64 {
	if x == 0 {
		return 1
	}
	xpi := math.Pi * x
	return math.Sin(xpi) / xpi
}
