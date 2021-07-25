package main

import (
	"strconv"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	c := 0
	a := app.New()
	w := a.NewWindow("hello")
	l := widget.NewLabel("hello fyne")
	w.SetContent(
		widget.NewVBox(
			l,
			widget.NewButton("click me", func() {
				c++
				l.SetText("count: ", +strconv.Itoa(c))
			}),
		),
	)
	w.ShowAndRun()
}
