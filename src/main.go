package main

import (
	"calories-and-macros-calculator/src/gui"

	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	w := gui.NewMainWindow(a, "macro calculator")
	w.ShowAndRun()
}
