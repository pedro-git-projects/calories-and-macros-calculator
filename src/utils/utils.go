package utils

import (
	"errors"
)

// IsEmpty is a varidic function to check if an arbirtrary
// amount of strings are empty
func IsEmpty(s ...string) error {
	for _, input := range s {
		if input == "" {
			return errors.New("empty input")
		}
	}
	return nil
}
