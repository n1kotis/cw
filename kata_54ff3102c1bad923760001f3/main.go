package kata

import "strings"

func GetCount(str string) (sum int) {
	s1 := []string{"a", "e", "i", "o", "u"}
	for _, v := range s1 {
		sum += strings.Count(str, v)
	}
	return sum
}
