package main

import (
	"strconv"
	"strings"
	"unicode"
)

type NumberRef struct {
	number int
	lower  int
	upper  int
}

func SymbolInRange(s string, from, to int) bool {
	if len(s) == 0 {
		return false
	}

	if from < 0 {
		from = 0
	}

	if to >= len(s) {
		to = len(s)
	}

	substr := s[from:to]
	return strings.ContainsFunc(substr, func(r rune) bool {
		return r != '.' && !unicode.IsDigit(r)
	})
}

func ParseNumberRefs(s string) []NumberRef {
	parts := make([]NumberRef, 0)
	start := 0
	end := 0

	extractorFn := func(str string, from, to int) {
		numberStr := s[from : to+1]
		number, err := strconv.Atoi(numberStr)
		if err == nil {
			bound := NumberRef{number, from, to}
			parts = append(parts, bound)
		}

	}

	collecting := false
	for i, r := range s {
		if unicode.IsDigit(r) {
			collecting = true
			end = i
		} else {
			if collecting {
				extractorFn(s, start, end)
				collecting = false
			}
			start = i + 1
			end = i + 1
		}
	}
	if collecting {
		extractorFn(s, start, end)
	}
	return parts
}

func FilterNumberRefs(lineRefs []NumberRef, lines ...string) []NumberRef {
	withSymbol := make([]NumberRef, 0)
	for _, r := range lineRefs {
		for _, line := range lines {
			found := SymbolInRange(line, r.lower-1, r.upper+2)
			if found {
				withSymbol = append(withSymbol, r)
			}
		}
	}
	return withSymbol
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

func FindDigitSliceIdx(s string, i int) SliceIdx {
	min := i
	max := i
	runes := []rune(s)
	if !unicode.IsDigit(runes[i]) {
		return SliceIdx{-1, -1}
	}

	for (min >= 0) && unicode.IsDigit(runes[min]) {
		min--
	}

	for (max < len(runes)) && unicode.IsDigit(runes[max]) {
		max++
	}

	return SliceIdx{min + 1, max}
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
func FindSliceIdxsInRange(str string, pos, radius int) []SliceIdx {
	slices := make([]SliceIdx, 0)
	from := Clamp(pos-radius, 0, len(str))
	to := Clamp(pos+radius+1, 0, len(str))

	if !DigitInRange(str, from, to) {
		return slices
	}

	for i := from; i < to; i++ {
		slice := FindDigitSliceIdx(str, i)

		if (slice.min == -1) || Contains(slice, slices) {
			continue
		}
		slices = append(slices, slice)
	}
	return slices
}

func DigitInRange(s string, from, to int) bool {
	substr := s[from:to]
	return strings.ContainsFunc(substr, unicode.IsDigit)
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
func FindGears(this, above, below string, symbol rune) [][]int {
	gears := make([][]int, 0)
	for i, r := range this {
		if r == symbol {
			found := make([]int, 0)
			numbersAbove := SlicesToNumbers(above, FindSliceIdxsInRange(above, i, 1))
			numbersThis := SlicesToNumbers(this, FindSliceIdxsInRange(this, i, 1))
			numbersBelow := SlicesToNumbers(below, FindSliceIdxsInRange(below, i, 1))
			found = append(found, numbersAbove...)
			found = append(found, numbersThis...)
			found = append(found, numbersBelow...)
			gears = append(gears, found)
		}
	}
	return gears
}

func Sum(elements []int) int {
	sum := 0
	for _, e := range elements {
		sum += e
	}
	return sum
}

func Prod(elements []int) int {
	prod := 1
	if len(elements) == 0 {
		return 0
	}
	for _, e := range elements {
		prod *= e
	}
	return prod
}
func GearPowers(gears [][]int) []int {
	powers := make([]int, len(gears))
	for i, g := range gears {
		if len(g) != 2 {
			continue
		}
		powers[i] = Prod(g)
	}
	return powers
}
