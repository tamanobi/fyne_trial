package main

import (
	"os"

	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func init() {
	os.Setenv("FYNE_FONT", "./fonts/mplus-1c-regular.ttf")
	os.Setenv("FYNE_FONT_MONOSPACE", "./fonts/mplus-1mn-regular.ttf")
}

func main() {
	a := app.New()
	win := a.NewWindow("Hello World")
	win.SetContent(widget.NewVBox(
		widget.NewLabel("こんにちは"),
		widget.NewButton("Quit", func() {
			a.Quit()
		}),
	))
	win.ShowAndRun()
}
