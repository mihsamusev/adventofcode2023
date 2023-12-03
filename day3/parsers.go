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
