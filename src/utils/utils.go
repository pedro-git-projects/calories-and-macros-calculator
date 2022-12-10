package utils

import (
	"errors"
)

func IsEmpty(s ...string) error {

	for _, input := range s {
		if input == "" {
			return errors.New("empty input")
		}
	}
	return nil
}
