package kata

import (
	"strconv"
	"strings"
)

func HighAndLow(in string) string {
	numbers := strings.Fields(in)
	low, _ := strconv.Atoi(numbers[0])
	high := low

	for _, num := range numbers {
		num1, _ := strconv.Atoi(num)
		if num1 > high {
			high = num1
		}
		if num1 < low {
			low = num1
		}
	}
	return strconv.Itoa(high) + " " + strconv.Itoa(low)
}
