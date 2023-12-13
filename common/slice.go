package common

import "strconv"


type SliceIdx struct {
	min int
	max int
}

func Contains(slice SliceIdx, slices []SliceIdx) bool {
	for _, s := range slices {
		if (s.min == slice.min) && (s.max == slice.max) {
			return true
		}
	}
	return false

}

func SlicesToNumbers(str string, slices []SliceIdx) []int {
	numbers := make([]int, 0)
	for _, s := range slices {
		numberStr := str[s.min:s.max]
		number, err := strconv.Atoi(numberStr)
		if err == nil {
			numbers = append(numbers, number)
		}
	}
	return numbers
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