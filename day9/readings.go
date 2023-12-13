package main


func PrevReading(values []int) int {
	diff := Diff(values)
	if AllZeros(diff) {
		return First(values)
	}

	return First(values) - PrevReading(diff)
}

func NextReading(values []int) int {
	diff := Diff(values)
	if AllZeros(diff) {
		return Last(values)
	}

	return Last(values) + NextReading(diff)
}

func Diff(values []int) []int {
	diff := make([]int, len(values) - 1)
	for i := 0; i < len(diff); i++ {
		diff[i]	= values[i + 1] - values[i]
	}
	return diff
}


func AllZeros(values []int) bool {
	for _, v:= range values {
		if v != 0 {
			return false
		}
	}
	return true
}

func Last(values []int) int {
	return values[len(values) - 1]
}

func First(values []int) int {
	return values[0]
}
