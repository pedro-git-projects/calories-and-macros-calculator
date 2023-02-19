package calculator

import "errors"

const (
	carbohydrateCalories = 4 // calories per gram of carb
	proteinCalories      = 4 // calories per gram of protein
	fatCalories          = 9 // calories per gram of fat
)

// The Macro strcut represents a macronutrient (carb, protein or fat)
type Macro struct {
	group           string
	caloriesPerGram int
}

// String returns the group name of the Macro
func (m Macro) String() string {
	return m.group
}

// Creating the Macro singletons
var (
	UnknownMacro = Macro{"", -1}
	Carbohydrate = Macro{"carbohydrate", carbohydrateCalories}
	Protein      = Macro{"protein", proteinCalories}
	Fat          = Macro{"fat", fatCalories}
)

// MacroFromString takes a string and converts it to
// the corresponding Macro if it is possible.
// Otherwise an error is returned
func MacroFromString(s string) (Macro, error) {
	switch s {
	case Carbohydrate.group:
		return Carbohydrate, nil
	case Protein.group:
		return Protein, nil
	case Fat.group:
		return Fat, nil
	default:
		return UnknownMacro, errors.New("unknown macro" + s)
	}
}
