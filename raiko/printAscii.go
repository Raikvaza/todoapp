package raiko

import (
	"strings"
)

func PrintAscii(word string, letters [][]string) string { // Formatting the output and storing it in a string VAR
	word = strings.ReplaceAll(word, "\n", "\\n")
	wordSplit := strings.Split(word, "\\n")
	wordSplit = checkLenWord(wordSplit)
	var result string
	for _, words := range wordSplit {
		if words == "" {
			result += "\n"
			continue
		} else {
			for i := 0; i < 8; i++ {
				for _, letter := range words {
					if letter >= 32 {
						result += letters[letter-32][i]
					}
				}
				result += "\n"
			}
		}
	}
	return result
}

func checkLenWord(words []string) []string {
	counter := 0
	for _, word := range words {
		if word == "" {
			counter++
		}
	}

	if counter == len(words) {
		return words[1:]
	}

	return words
}
