package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	got := Hello()
	want := "Burgers Equation Explorer"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
