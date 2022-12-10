package calculator

import "errors"

type ActivityLevel int

const (
	BMR ActivityLevel = iota
	Sedentary
	Light
	Moderate
	Active
	VeryActive
	unknownActivityLevel
)

func (a ActivityLevel) String() string {
	switch a {
	case BMR:
		return "Basal Metabolic Rate (BMR)"
	case Sedentary:
		return "Sedentary: little to no exercise"
	case Light:
		return "Lightly active: 1-3 times/week"
	case Moderate:
		return "Moderately active: 4-5 times/week"
	case Active:
		return "Active: 6-7 times/week"
	case VeryActive:
		return "Very active: hard exercise 6-7 times/week"
	default:
		return "unknown"
	}
}

func ActivityLevelFromString(s string) (ActivityLevel, error) {
	switch s {
	case "Basal Metabolic Rate (BMR)":
		return BMR, nil
	case "Sedentary: little to no exercise":
		return Sedentary, nil
	case "Moderately active: 4-5 times/week":
		return Moderate, nil
	case "Active: 6-7 times/week":
		return Active, nil
	case "Very active: hard exercise 6-7 times/week":
		return VeryActive, nil
	default:
		return unknownActivityLevel, errors.New("unknown activity level" + s)
	}
}
