package utils

import "testing"

// AssertString .
func AssertString(t *testing.T, got, want string) {
	if got != want {
		t.Errorf("got %#v; want %#v", got, want)
	}
}
