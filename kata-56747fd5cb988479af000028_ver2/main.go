package kata

import "fmt"

type WordWrapper struct {
	Word string
}

func GetMiddle(s string) string {
	wrapper := WordWrapper{
		Word: s,
	}

	return GetMiddleFromWrapper(wrapper)
}

func GetMiddleFromWrapper(wrapper WordWrapper) string {
	runes := []rune(wrapper.Word)
	length := len(runes)
	middle := length / 2
	if length%2 == 1 {
		return string(runes[middle])
	} else {
		return string(runes[middle-1 : middle+1])
	}
}

func main() {
	fmt.Printf("%s\n", GetMiddle("test"))
	fmt.Printf("%s\n", GetMiddle("lol"))
}
