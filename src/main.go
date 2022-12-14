package main

import (
	"calories-and-macros-calculator/src/calculator"
	"calories-and-macros-calculator/src/utils"
	"fmt"
	"log"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
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
	})

	gender := widget.NewRadioGroup([]string{"male", "female"}, func(value string) {
	})

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Age in years", Widget: age},
			{Text: "Gender", Widget: gender},
			{Text: "Height in cm", Widget: height},
			{Text: "Weight in kg", Widget: weight},
			{Text: "Activity", Widget: activity}},

		OnSubmit: func() {
			err := utils.IsEmpty(age.Text, height.Text, weight.Text)
			if err != nil {
				log.Println(err)
				return
			}

			fmt.Println(age.Text)
			iAge, err := strconv.Atoi(age.Text)
			if err != nil {
				log.Println(err)
				return
			}

			fHeight, err := strconv.ParseFloat(height.Text, 64)
			if err != nil {
				log.Println(err)
				return
			}

			fWeight, err := strconv.ParseFloat(weight.Text, 64)
			if err != nil {
				log.Println(err)
				return
			}

			gGender, err := calculator.GenderFromString(gender.Selected)
			if err != nil {
				log.Println(err)
				return
			}

			aActivity, err := calculator.ActivityLevelFromString(activity.Selected)
			if err != nil {
				log.Println(err)
				return
			}

			user, err := calculator.NewUser(iAge, fHeight, fWeight, gGender, aActivity)
			if err != nil {
				log.Println(err)
				return
			}

			recommended := strconv.FormatFloat(user.Calories, 'f', 2, 64)
			recommended = recommended + " calories"
			userBMR := strconv.FormatFloat(user.BMR, 'f', 2, 64)
			userBMR = userBMR + " calories"

			userCarbohydrate := strconv.FormatFloat(user.CarbohydrateIntake, 'f', 2, 64)
			userCarbohydrate = userCarbohydrate + " grams"

			userProtein := strconv.FormatFloat(user.ProteinIntake, 'f', 2, 64)
			userProtein = userProtein + " grams"

			userFat := strconv.FormatFloat(user.FatIntake, 'f', 2, 64)
			userFat = userFat + " grams"

			data := [][]string{{"BMR", userBMR},
				{"Suggested Daily Caloric Intake", recommended},
				{"Carbohydrate", userCarbohydrate},
				{"Protein", userProtein},
				{"Fat", userFat},
			}

			list := widget.NewTable(
				func() (int, int) {
					return len(data), len(data[0])
				},
				func() fyne.CanvasObject {
					return widget.NewLabel("Suggested daily caloric intake  ")
				},
				func(i widget.TableCellID, o fyne.CanvasObject) {
					o.(*widget.Label).SetText(data[i.Row][i.Col])
				})

			w2 := a.NewWindow("Calories and Macros Calculator - Results")
			w2.Resize(fyne.NewSize(500, 600))
			w2.SetIcon(theme.FyneLogo())
			w2.SetContent(list)
			w2.Show()
		},

		SubmitText: "Calculate",
	}

	w.Resize(fyne.NewSize(500, 500))
	w.SetIcon(theme.FyneLogo())
	w.SetContent(form)
	w.ShowAndRun()
}
