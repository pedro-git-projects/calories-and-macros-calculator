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

const (
	BMRStr        = "Basal Metabolic Rate (BMR)"
	SedentaryStr  = "Sedentary: little to no exercise"
	LightStr      = "Lightly active: 1-3 times/week"
	ModerateStr   = "Moderately active: 4-5 times/week"
	ActiveStr     = "Active: 6-7 times/week"
	VeryActiveStr = "Very active: hard exercise 6-7 times/week"
	UnknownStr    = "unknown"
)

// String converts an instance of the activity level enum
// to the corresponding string
func (a ActivityLevel) String() string {
	switch a {
	case bmr:
		return BMRStr
	case sedentary:
		return SedentaryStr
	case light:
		return SedentaryStr
	case moderate:
		return ModerateStr
	case active:
		return ActiveStr
	case veryActive:
		return VeryActiveStr
	default:
		return UnknownStr
	}
}

// ActivityLevelFromString takes a string and converts it to
// the corresponding ActivityLevel if it is possible.
// Otherwise an error is returned
func ActivityLevelFromString(s string) (ActivityLevel, error) {
	switch s {
	case BMRStr:
		return bmr, nil
	case SedentaryStr:
		return sedentary, nil
	case LightStr:
		return light, nil
	case ModerateStr:
		return moderate, nil
	case ActiveStr:
		return active, nil
	case VeryActiveStr:
		return veryActive, nil
	default:
		return unknownActivityLevel, errors.New("unknown activity level" + s)
	}
}
