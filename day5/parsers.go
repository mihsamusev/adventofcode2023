package main

import (
	"errors"
	"strconv"
	"strings"
)

func ParseId(str, prefix string) (int, error) {
	result, found := strings.CutPrefix(str, prefix)
	if !found {
		return -1, nil
	}
	result = strings.TrimSpace(result)
	return strconv.Atoi(result)
}

func ParseNamedSlice(str, prefix string) ([]int, error) {
	result, found := strings.CutPrefix(str, prefix)
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

func ParseFarmingMap(str string, withHeader bool) (FarmingMap, error) {
	lines := strings.Split(str, "\n")
	if len(lines) == 0 {
		return FarmingMap{}, errors.New("rows not found")
	}

	if withHeader {
		lines = lines[1:]
	}

	lookups := make([]Lookup, 0)
	for _, line := range lines {
		n, err := ParseSlice(line)
		if err != nil {
			return FarmingMap{}, errors.New("wrong row")
		}
		if len(n) != 3 {
			return FarmingMap{}, errors.New("expected 3 numbers")
		}
		dst := n[0]
		srcRange := Range{n[1], n[1] + n[2] - 1}
		lookups = append(lookups, Lookup{dst, srcRange})
	}

	return FarmingMap{lookups}, nil
}

func ParseFarmingMaps(lines []string) ([]FarmingMap, error) {
	farmingMaps := make([]FarmingMap, 0)
	for _, b := range lines {
		farmingMap, err := ParseFarmingMap(b, true)
		if err != nil {
			panic(err)
		}
		farmingMaps = append(farmingMaps, farmingMap)
	}
	return farmingMaps, nil
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
