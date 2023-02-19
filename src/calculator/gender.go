package calculator

import "errors"

// Gender is an alias for the int type
// which is used to define an enum for the
// gender of the user
type Gender int

const (
	male Gender = iota
	female
	unknownGender
)

// String converts an instance of the gender enum
// to the corresponding string
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

// GenderFromString takes a string and converts it to
// the corresponding Gender if it is possible.
// Otherwise an error is returned
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
