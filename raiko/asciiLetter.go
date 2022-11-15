package raiko

import (
	"errors"
)

func Argum(word string) (string, error) {
	for _, letter := range word {
		if letter > 126 {
			return "", errors.New("Wrong argument syntax")
		}
	}
	return word, nil
}
