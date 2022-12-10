package calculator

import "errors"

type Gender int

const (
	male Gender = iota
	female
	unknownGender
)

func (g Gender) String() string {
	switch g {
	case male:
		return "male"
	case female:
		return "female"
	default:
		return "unknown"
	}
}

func GenderFromString(s string) (Gender, error) {
	switch s {
	case "male":
		return male, nil
	case "female":
		return female, nil
	default:
		return unknownGender, errors.New("Unknown gender" + s)
	}
}
