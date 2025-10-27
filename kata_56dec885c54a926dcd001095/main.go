package main

func Opposite(value int) int {
	//либо все стерерть с 5 по 9 и написать тупо return -value
	if value > 0 {
		value = -value
	} else {
		value = value * -1
	}
	return value
}
