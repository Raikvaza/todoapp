package raiko

import (
	"bufio"
	"errors"
	"os"
)

func Letters(s string) ([][]string, error) {
	var asciiLetters [][]string
	file, err := os.Open(s) // Opening the file
	if err != nil {
		return asciiLetters, errors.New("Couldn't open the banner file with the specified name")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines) // Splitting by each line of text
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text()) // appending each line of our scanner to a []string
	}

	file.Close()
	var temp []string

	start := 0
	for _, each_ln := range text {
		if start != 0 {
			temp = append(temp, each_ln) // adding 8 lines to temp
		}
		start++
		if start > 8 {
			asciiLetters = append(asciiLetters, temp) // appending 8 lined temp to our asciiLetters || [0][0] - first line of space || [0][1] - second line of space || [0][2] - third line of space || etc
			start = 0                                 // Resetting the values
			temp = []string{}
		}
	}

	return asciiLetters, nil
}
