package calculator

import "errors"

const (
	carbohydrateCalories = 4
	proteinCalories      = 4
	fatCalories          = 9
)

type Macro struct {
	group           string
	caloriesPerGram int
}

func (m Macro) String() string {
	return m.group
}

var (
	UnknownMacro = Macro{"", -1}
	Carbohydrate = Macro{"carbohydrate", carbohydrateCalories}
	Protein      = Macro{"protein", proteinCalories}
	Fat          = Macro{"fat", fatCalories}
)

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
