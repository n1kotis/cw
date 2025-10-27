package kata

import "fmt"

type WordWithMiddle struct {
	Word   string
	Middle string
}

func getWords() []WordWithMiddle {
	return []WordWithMiddle{
		{"test", " "},
		{"testing", " "},
		{"middle", " "},
		{"A", " "},
	}
}

func GetMiddle(s string) string {
	runes := []rune(s)
	lenght := len(runes)
	middle := lenght / 2
	if lenght%2 == 1 {
		return string(runes[middle])
	} else {
		return string(runes[middle-1 : middle+1])
	}
}

func findMiddle(words []WordWithMiddle) []WordWithMiddle {
	for i := range words {
		words[i].Middle = GetMiddle(words[i].Word)
	}
	return words
}

func main() {
	words := getWords()
	words = findMiddle(words)
	for _, word := range words {
		fmt.Printf("%s\n", word.Middle)
	}
}
