package main

import (
	"strconv"
	"strings"
    "errors"
)

func ParseId(str, prefix string) (int, error) {
    result, found := strings.CutPrefix(str,prefix)
    if !found {
        return -1, nil
    }
    result = strings.TrimSpace(result)
    return strconv.Atoi(result)
}

func ParseNamedSlice(str, prefix string) ([]int, error) {
    result, found := strings.CutPrefix(str,prefix)
    if !found {
        return make([]int, 0), errors.New("no prefix found")
    }
    return ParseSlice(result)

}
func ParseSlice(str string) ([]int, error) {
    trimmed := strings.Fields(str)
    slice := make([]int, 0)
    for _, t := range trimmed {
        n, err := strconv.Atoi(t)
        if err != nil {
            return slice, err
        }
        slice = append(slice, n)
    }
    return slice, nil
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

func Sum(elements []int) int {
	sum := 0
	for _, e := range elements {
		sum += e
	}
	return sum
}

