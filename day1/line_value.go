package main

import (
	"strings"
	"unicode"
)

type Digit struct {
	spelled string
	value   int
}

const maxSpelledLen int = 5

var pairs = [9]Digit{
	{"one", 1},
	{"two", 2},
	{"three", 3},
	{"four", 4},
	{"five", 5},
	{"six", 6},
	{"seven", 7},
	{"eight", 8},
	{"nine", 9},
}

func lineDigits(str string) []int {
	digits := make([]int, 0, 10)
	for _, rune := range str {
		if unicode.IsDigit(rune) {
			digits = append(digits, runeToValue(rune))
		}
	}
	return digits
}

func lineDigitsExtended(str string) []int {
	digits := make([]int, 0, 10)

	runes := []rune(str)

	i := 0
	step := 1
	for i < len(runes) {
		if unicode.IsDigit(runes[i]) {
			digits = append(digits, runeToValue(runes[i]))
			step = 1
		} else {
			windowEnd := i + maxSpelledLen
			if windowEnd >= len(str) {
				windowEnd = len(str)
			}
			substr := str[i:windowEnd]
			substrValue := 0
			substrValue, step = stringToValue(substr)
			
			if substrValue != 0 {
				digits = append(digits, substrValue)
				step = 3
			}
		}
		i += step
	}
	return digits
}

func calibrationValue(digits []int) int {
	numDigits := len(digits)
	if numDigits == 0 {
		return 0
	}
	first := digits[0]
	last := digits[numDigits-1]
	value := first*10 + last
	return value
}

func stringToValue(str string) (int, int) {
	width := 1
	value := 0
	for _, pair := range pairs {
		if strings.HasPrefix(str, pair.spelled) {
			return pair.value, len(pair.spelled)
		}
	}
	return value, width
}
func runeToValue(r rune) int {
	value := 0
	if unicode.IsDigit(r) {
		value = int(r - '0')
	}
	return value
}
