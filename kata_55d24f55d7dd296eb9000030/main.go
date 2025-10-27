package kata

type IntegerWrapper struct {
	point []int
}

// функция, которая предлагается codewars как содержащая логику решения
// из неё вызываем "Свою" функцию и возвращаем результат работы функции
func SmallestIntegerFinder(numbers []int) int {
	return FindSmallestInteger(numbers)
}

// функция, которая создает и возвращает новый экземпляр класса
func InitializeIntegerWrapper(numbers []int) IntegerWrapper {
	integerWrapper := IntegerWrapper{
		point: numbers,
	}
	return integerWrapper
}

// "Своя" функция, в которой содержится логика решения и возвращается результат.
// Функция должна иметь семантику как и у функции предлагаемой codewars
func FindSmallestInteger(numbers []int) int {
	//создаем экземпляр класса/структуры IntegerWrapper. Вызываем для этого функцию, которая создает экземпляр класса и возвращает.
	//Иными словами происходит "Инициализация экземпляра структуры". экземпляр присваивается переменной integerWrapper
	integerWrapper := InitializeIntegerWrapper(numbers)

	//получаем требуемый в задаче результат, вызываю функцию которая содержит логику задачи и аргумент типа IntegerWrapper (по сути integerWrapper это переменная-"контейнер" для входных данных)
	smallestInteger := SearchSmallestIntegerThroughWrapper(integerWrapper)

	//возвращаем результат
	return smallestInteger
}

// функция содержащая логику решения и получающая переменную-контейнер для данных
func SearchSmallestIntegerThroughWrapper(wrapper IntegerWrapper) int {
	min := wrapper.point[0]
	for _, num := range wrapper.point {
		if num < min {
			min = num
		}
	}
	return min
}
