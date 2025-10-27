package kata

func RepeatStr(repetitions int, value string) string {
	var x string
	for i := 0; i < repetitions; i++ {

		x += value
	}
	return x
}
