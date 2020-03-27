package main

import (
	"testing"

	"fyne.io/fyne"
	"github.com/rblack42/go-burgers/gui"
	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	app := gui.Create()

	assert.Equal(t, app, fyne.CurrentApp())
}

func TestMainWindow(t *testing.T) {
	app := gui.Create()
}
