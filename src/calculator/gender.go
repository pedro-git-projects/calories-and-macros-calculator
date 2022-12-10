package calculator

import "errors"

type Gender int

const (
	Male Gender = iota
	Female
	unknownGender
)

func (g Gender) String() string {
	switch g {
	case Male:
		return "male"
	case Female:
		return "female"
	default:
		return "unknown"
	}
}

func GenderFromString(s string) (Gender, error) {
	switch s {
	case "male":
		return Male, nil
	case "female":
		return Female, nil
	default:
		return unknownGender, errors.New("Unknown gender" + s)
	}
}
