package common

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