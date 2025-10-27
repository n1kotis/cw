package kata

import (
	"strconv"
)

func StringToNumber(str string) int {

	word, _ := strconv.Atoi(str)
	return word

}
