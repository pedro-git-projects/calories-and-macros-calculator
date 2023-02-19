package gui

import (
	"calories-and-macros-calculator/src/calculator"
	"calories-and-macros-calculator/src/utils"
	"log"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// MainWindow type holds all data that will need to be presented
// on the main app window
type MainWindow struct {
	App        fyne.App
	Title      string
	Screen     fyne.Window
	Entries    map[string]*widget.Entry
	Radios     map[string]*widget.RadioGroup
	Selects    map[string]*widget.Select
	Forms      map[string]*widget.Form
	Tables     map[string]*widget.Table
	FormResult [][]string
}

// NewMainWindow returns a pointer to a window with the desired title
func NewMainWindow(application fyne.App, title string) *MainWindow {
	mw := &MainWindow{
		App:        application,
		Title:      title,
		Entries:    make(map[string]*widget.Entry),
		Radios:     make(map[string]*widget.RadioGroup),
		Selects:    make(map[string]*widget.Select),
		Forms:      make(map[string]*widget.Form),
		Tables:     make(map[string]*widget.Table),
		FormResult: make([][]string, 0),
	}
	mw.Screen = mw.App.NewWindow(title)
	mw.setEntries()
	mw.setActivityWidget()
	mw.setGenderWdiget()
	mw.constructForm()

	return mw
}

// Show and run calls ShowAndRun fyne receiver
// on the main window displaying the calculator form
func (w *MainWindow) ShowAndRun() {
	w.Screen.Resize(fyne.NewSize(500, 500))
	w.Screen.SetIcon(theme.FyneLogo())
	w.Screen.SetContent(w.Forms["calculator"])
	w.Screen.ShowAndRun()
}

// setEntries populates the entries map with
// the age height and weight entries
func (w *MainWindow) setEntries() {
	w.Entries["age"] = widget.NewEntry()
	w.Entries["height"] = widget.NewEntry()
	w.Entries["weight"] = widget.NewEntry()
}

// setActivityWidget populates the Select map
// with all activity level strings from the calculator pkg
func (w *MainWindow) setActivityWidget() {
	w.Selects["activity"] = widget.NewSelect([]string{calculator.BMRStr, calculator.SedentaryStr, calculator.LightStr, calculator.ModerateStr, calculator.ActiveStr, calculator.VeryActiveStr}, func(value string) {
	})
}

// setGenderWdiget populates the Radios map with the
// gender radio buttons and male && female options
func (w *MainWindow) setGenderWdiget() {
	w.Radios["gender"] = widget.NewRadioGroup([]string{"male", "female"}, func(value string) {
	})
}

// resultTable populates the Tables map
// with the result entry
func (w *MainWindow) resultTable() {
	w.Tables["result"] = widget.NewTable(
		func() (int, int) {
			return len(w.FormResult), len(w.FormResult[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Suggested daily caloric intake  ")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(w.FormResult[i.Row][i.Col])
		})
}

// on form submmit encapsualtes all submission logic
// including checking for error in user input
func (w *MainWindow) onFormSubmit() func() {
	return func() {
		if err := utils.IsEmpty(w.Entries["age"].Text, w.Entries["height"].Text, w.Entries["weight"].Text); err != nil {
			log.Println(err)
			return
		}

		iAge, err := strconv.Atoi(w.Entries["age"].Text)
		if err != nil {
			log.Println(err)
			return
		}

		fHeight, err := strconv.ParseFloat(w.Entries["height"].Text, 64)
		if err != nil {
			log.Println(err)
			return
		}

		fWeight, err := strconv.ParseFloat(w.Entries["weight"].Text, 64)
		if err != nil {
			log.Println(err)
			return
		}

		gGender, err := calculator.GenderFromString(w.Radios["gender"].Selected)
		if err != nil {
			log.Println(err)
			return
		}

		aActivity, err := calculator.ActivityLevelFromString(w.Selects["activity"].Selected)
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

		w.FormResult = [][]string{{"BMR", userBMR},
			{"Suggested Daily Caloric Intake", recommended},
			{"Carbohydrate", userCarbohydrate},
			{"Protein", userProtein},
			{"Fat", userFat},
		}

		w.resultTable()

		w2 := w.App.NewWindow("Calories and Macros Calculator - Results")
		w2.Resize(fyne.NewSize(500, 600))
		w2.SetIcon(theme.FyneLogo())
		w2.SetContent(w.Tables["result"])
		w2.Show()
	}
}

// constructForm combines all receivers to create the calculator form
func (w *MainWindow) constructForm() {
	w.Forms["calculator"] = &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Age in years", Widget: w.Entries["age"]},
			{Text: "Gender", Widget: w.Radios["gender"]},
			{Text: "Height in cm", Widget: w.Entries["height"]},
			{Text: "Weight in kg", Widget: w.Entries["weight"]},
			{Text: "Activity", Widget: w.Selects["activity"]}},
		OnSubmit:   w.onFormSubmit(),
		SubmitText: "calculate",
	}
}
