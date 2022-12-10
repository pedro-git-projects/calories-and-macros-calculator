package main

import (
	"log"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

const (
	bmr        = "Basal Metabolic Rate (BMR)"
	sedentary  = "Sedentary: little to no exercise"
	light      = "Lightly active: 1-3 times/week"
	moderate   = "Moderately active: 4-5 times/week"
	active     = "Active: 6-7 times/week"
	veryActive = "Very active: hard exercise 6-7 times/week"
)

func main() {
	a := app.New()
	w := a.NewWindow("Calories and Macros Calculator")

	age := widget.NewEntry()
	height := widget.NewEntry()
	weight := widget.NewEntry()

	activity := widget.NewSelect([]string{bmr, sedentary, light, moderate, active, veryActive}, func(value string) {
		log.Println(value)
	})

	gender := widget.NewRadioGroup([]string{"male", "female"}, func(value string) {
		log.Println(value)
	})

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Age", Widget: age},
			{Text: "Gender", Widget: gender},
			{Text: "Height", Widget: height},
			{Text: "Weight", Widget: weight},
			{Text: "Activity", Widget: activity}},

		OnSubmit: func() {
			log.Println("Form submitted", age.Text)
			log.Println("Form submitted", height.Text)
			log.Println("Form submitted", weight.Text)
			// myWindow.Close() <- use when this is not the main window
		},
	}

	w.SetContent(form)
	w.ShowAndRun()
}
