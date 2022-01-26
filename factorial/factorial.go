package factorial

func Factorial(number int) int {
	if number <= 0 {
		return 1
	}
	result := 1
	for i := number; i >= 1; i-- {
		result *= i
	}
	return result
}

func FactorialRecursive(number int) int {
	if number <= 0 {
		return 1
	} else {
		// 3 * factorial(2)
		// 3 * 2 * factorial(1)
		// 3 * 2 * 1 * factorial(0)
		// 3 * 2 * 1 * 1
		// factorial(3) * factorial(2) * factorial(1) * factorial(0)
		return number * FactorialRecursive(number-1)
	}
}

func FactorialTailRecursive(total int, number int) int {
	if number <= 0 {
		return total
	} else {
		// => factorial(1, 3)
		// => factorial(3, 2)
		// => factorial(6, 1)
		// => factorial(6, 0)
		// => 6 //
		return FactorialTailRecursive(total*number, number-1)
	}
}
