package main

func GreatestCommonDiv(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func LeastCommonMulPair(a, b int) int {
	return a * b / GreatestCommonDiv(a, b)
}

func LeastCommonMul(numbers ...int) int {
	result := LeastCommonMulPair(numbers[0], numbers[1])
	if len(numbers) == 2 {
		return result
	}

	for _, n := range numbers[2:] {
		result = LeastCommonMulPair(result, n)
	}

	return result
}
