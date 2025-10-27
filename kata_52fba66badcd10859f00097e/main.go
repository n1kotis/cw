package kata

import "regexp"

func Disemvowel(comment string) string {
	reg := regexp.MustCompile(`(a|e|i|o|u|A|E|I|O|U)`)
	return reg.ReplaceAllString(comment, "")
}
