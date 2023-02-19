package calculator

import "errors"

// ActivityLevel is an alias for the int type
// which is used to define an enum for the
// level of activity of the user
type ActivityLevel int

const (
	bmr ActivityLevel = iota
	sedentary
	light
	moderate
	active
	veryActive
	unknownActivityLevel
)

// String converts an instance of the activity level enum
// to the corresponding string
func (a ActivityLevel) String() string {
	switch a {
	case bmr:
		return "Basal Metabolic Rate (BMR)"
	case sedentary:
		return "Sedentary: little to no exercise"
	case light:
		return "Lightly active: 1-3 times/week"
	case moderate:
		return "Moderately active: 4-5 times/week"
	case active:
		return "Active: 6-7 times/week"
	case veryActive:
		return "Very active: hard exercise 6-7 times/week"
	default:
		return "unknown"
	}
}

// ActivityLevelFromString takes a string and converts it to
// the corresponding ActivityLevel if it is possible.
// Otherwise an error is returned
func ActivityLevelFromString(s string) (ActivityLevel, error) {
	switch s {
	case "Basal Metabolic Rate (BMR)":
		return bmr, nil
	case "Sedentary: little to no exercise":
		return sedentary, nil
	case "Lightly active: 1-3 times/week":
		return light, nil
	case "Moderately active: 4-5 times/week":
		return moderate, nil
	case "Active: 6-7 times/week":
		return active, nil
	case "Very active: hard exercise 6-7 times/week":
		return veryActive, nil
	default:
		return unknownActivityLevel, errors.New("unknown activity level" + s)
	}
}
