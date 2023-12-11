package common

func Min(first, second int) int {
	if first < second {
		return first
	} else {
		return second
	}
}

func Max(first, second int) int {
	if first > second {
		return first
	} else {
		return second
	}
}

func Sum(elements []int) int {
	sum := 0
	for _, e := range elements {
		sum += e
	}
	return sum
}

func Clamp(i, min, max int) int {
	if i < min {
		i = min
	}
	if i > max {
		i = max
	}
	return i
}