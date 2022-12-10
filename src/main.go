package main

import (
	"calories-and-macros-calculator/src/calculator"
	"log"
	"strconv"

	"fyne.io/fyne/v2"
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
			iAge, err := strconv.Atoi(age.Text)
			fHeight, err := strconv.ParseFloat(height.Text, 64)
			fWeight, err := strconv.ParseFloat(weight.Text, 64)
			gGender, err := calculator.GenderFromString(gender.Selected)
			aActivity, err := calculator.ActivityLevelFromString(activity.Selected)
			if err != nil {
				log.Println(err)
			}

			user, err := calculator.NewUser(iAge, fHeight, fWeight, gGender, aActivity)
			if err != nil {
				log.Println(err)
			}

			recommended := strconv.FormatFloat(user.Calories, 'f', 2, 64)
			recommended = recommended + " calories"
			userBMR := strconv.FormatFloat(user.BMR, 'f', 2, 64)
			userBMR = userBMR + " calories"

			data := [][]string{{"Suggested daily caloric intake", recommended},
				{"BMR", userBMR}}

			list := widget.NewTable(
				func() (int, int) {
					return len(data), len(data[0])
				},
				func() fyne.CanvasObject {
					return widget.NewLabel("Suggested daily calorie intake")
				},
				func(i widget.TableCellID, o fyne.CanvasObject) {
					o.(*widget.Label).SetText(data[i.Row][i.Col])
				})

			w2 := a.NewWindow("Calories and Macros Calculator - Results")
			w2.Resize(fyne.NewSize(500, 500))
			w2.SetContent(list)
			w2.Show()
		},
	}

	w.Resize(fyne.NewSize(500, 500))
	w.SetContent(form)
	w.ShowAndRun()
}
