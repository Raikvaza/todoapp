package raiko

import (
	"errors"
)

func Run(s, b string) (string, error) {
	// Argument string
	asciiLetters, err := Argum(s) // Checking the validity of an argument
	if err != nil {
		return "", err
	}
	// Banner
	bannerInput := b
	if bannerInput != "shadow" && bannerInput != "thinkertoy" && bannerInput != "standard" { // Checking the validity of a banner syntax
		return "", errors.New("Wrong file-name of a banner")
	}
	readFile := "raiko/banner/" + bannerInput + ".txt"
	// Hash check
	hs := HashSum(readFile) // Checking the corruptness of banner files
	if hs != nil {
		return "", errors.New("Hash files are corrupted")
	}
	// Splitting a banner
	bannerLeters, err := Letters(readFile) // splitting the banner into a [][]string
	if err != nil {
		return "", err
	}
	// Returning the result after all errors have been handled
	result := PrintAscii(asciiLetters, bannerLeters)

	return result, nil
}
