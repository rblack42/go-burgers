package gui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func Create() fyne.App {
	app := app.New()
	return app
}

func Run() {
	app := Create()
	w := app.NewWindow("go-burgers")
	w.SetContent(widget.NewVBox(
		widget.NewLabel("Hello"),
		widget.NewButton("Quit", func () {
			app.Quit()
		}),
	))

	w.ShowAndRun()
}

