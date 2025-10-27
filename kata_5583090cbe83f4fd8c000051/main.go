package kata

type Numbers interface {
	Array() []int
}

type DigitRevers struct {
	number int
}

func DigitsRevers(n int) DigitRevers {
	return DigitRevers{number: n}
}

func (d DigitRevers) Array() []int {
	if d.number == 0 {
		return []int{0}
	}
	var nums []int
	value := d.number
	for value > 0 {
		nums = append(nums, value%10)
		value /= 10
	}
	return nums
}

func Digitize(n int) []int {
	rev := DigitsRevers(n)
	var numbers Numbers = rev
	return numbers.Array()
}
