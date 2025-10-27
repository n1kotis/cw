package kata

func CountSheeps(numbers []bool) int {
	var x int
	for _, sheep := range numbers {
		if sheep == true {
			x++
		}
	}
	return x
}
