package main

import "strings"

func CollapsePattern(pattern, mask string) string {
	return strings.ReplaceAll(pattern, mask, ".")
}

func Groups(pattern string) []string {
	return strings.FieldsFunc(
		pattern,
		func(r rune) bool {return r == '.'},
	)
}

func CorrectPattern(pattern string, sizes []int) bool {
	groups := Groups(pattern)
	if len(groups) != len(sizes) {
		return false
	}
	for i := range sizes {
		if strings.Count(groups[i], "#") != sizes[i] {
			return false
		}
	}
	return true
}

func CreateMask(size int) string {
	return ""	
}

func BuildCombos(pattern string, sizes []int) []string {
	result := make([]string, 0)
	return result
}