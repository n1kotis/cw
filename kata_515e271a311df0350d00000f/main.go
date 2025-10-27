package kata

func SquareSum(numbers []int) int {
	sum := 0
	for _, v := range numbers {
		j := v * v

		sum += j

	}
	return sum
}
