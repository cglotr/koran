package utils

import "testing"

// AssertInt .
func AssertInt(t *testing.T, got, want int) {
	if got != want {
		t.Errorf("got %#v; want %#v", got, want)
	}
}
