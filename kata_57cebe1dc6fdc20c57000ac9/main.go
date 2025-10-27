package kata

import "strings"

func FindShort(s string) int {
	words := strings.Split(s, " ")

	minlength := len(words[0])
	for _, word := range words {
		if len(word) < minlength {
			minlength = len(word)
		}
	}
	return minlength
}
