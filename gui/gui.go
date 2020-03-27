package gui

import (
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	"fyne.io/fyne/canvas"

	"github.com/rblack42/go-burgers/config"
)

func Create() fyne.App {
	app := app.New()
	return app
}


func Run(c *config.UiConfig) {
	var icon *canvas.Image

	app := Create()
	w := app.NewWindow(c.Title)
	w.Resize(fyne.NewSize(c.Width, c.Height))

	icon = canvas.NewImageFromFile("go-burgers.png")
	icon.FillMode = canvas.ImageFillOriginal
	icon.SetMinSize(fyne.NewSize(48, 48))

	panel := canvas.NewRectangle(color.White)
	panel.Resize(fyne.NewSize(150, 100))
	
	w.SetContent(
		widget.NewVBox(
			widget.NewHBox(
				widget.NewVBox(
					icon,
					widget.NewButton(
						"quit",
						func() {
							app.Quit()
						},
					),
				),
				panel,
			),
			widget.NewLabelWithStyle(
				"Copyright 2020 - Roie R. Black", 
				fyne.TextAlignCenter, 
				fyne.TextStyle{ Bold: true},
			),
		),
	)

	w.ShowAndRun()
}

