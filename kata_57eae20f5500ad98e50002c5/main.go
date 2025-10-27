package kata

import "strings"

type StrWrapper struct {
	variable string
}

func NoSpace(word string) string {
	str := StrWrapper{
		variable: word,
	}
	return GetNoSpace(str)
}

func GetNoSpace(str StrWrapper) string {
	strng := strings.ReplaceAll(str.variable, " ", "")
	return strng
}
