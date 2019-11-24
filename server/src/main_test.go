package main

import (
	"testing"
)

func assertInt(t *testing.T, got, want int) {
	if got != want {
		t.Errorf("got %#v; want %#v", got, want)
	}
}

func assertString(t *testing.T, got, want string) {
	if got != want {
		t.Errorf("got %#v; want %#v", got, want)
	}
}
