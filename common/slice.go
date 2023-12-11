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