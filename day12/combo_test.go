
package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombosOne(t *testing.T) {
	combos := BuildCombos("???.###", []int{1, 1, 3})
	expected := []string{"#.#.###"}
	assert.Equal(t, expected, combos)
}

func TestCombosMany(t *testing.T) {
	combos := BuildCombos(".??..??...?##.", []int{1, 1, 3})
	expected := []string{
		".#...#....###.",
		".#....#...###.",
		"..#..#....###.",
		"..#...#...###.",
	}
	assert.Equal(t, expected, combos)
}
