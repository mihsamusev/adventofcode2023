package main

import (
	"fmt"
	"testing"
)

func TestSingleDigit(t *testing.T) {
	actual := ParseNumberRefs("....*4....")
	if len(actual) == 0 {
		t.Errorf("bad boi")
	}
}

func TestMoreDigits(t *testing.T) {
	actual := ParseNumberRefs("....*44...")
	if len(actual) == 0 {
		t.Errorf("bad boi")
	}
}

func TestTwoSidedSearch(t *testing.T) {
	testCases := []struct {
		line string
		pos  int
		min  int
		max  int
	}{
		{"...4.....", 4, -1, -1},
		{"...4.....", 3, 3, 4},
		{"...404...", 3, 3, 6},
		{"...404...", 4, 3, 6},
		{"...404...", 5, 3, 6},
		{"......755.", 6, 6, 9},
	}
	for _, test := range testCases {
		slice := FindDigitSliceIdx(test.line, test.pos)
		if (slice.min != test.min) || (slice.max != test.max) {
			t.Errorf("wrong min %d, max %d for searching '%s' from pos %d", slice.min, slice.max, test.line, test.pos)
		}
	}
}

func TestFindRangeSliceIdxs(t *testing.T) {
	testCases := []struct {
		line  string
		pos   int
		count int
	}{
		{"....*....", 3, 0},
		{"....*....", 4, 0},
		{"....*.2..", 4, 0},
		{"....*....", 5, 0},
		{"...4*....", 4, 1},
		{"...4*.2..", 4, 1},
		{".2.4*....", 4, 1},
		{"..44*2...", 4, 2},
		{"......755.", 6, 1},
	}
	for _, test := range testCases {
		slices := FindSliceIdxsInRange(test.line, test.pos, 1)
		fmt.Printf("Got slices: %v\n", slices)

		if len(slices) != test.count {
			t.Errorf("Got slices: %v", slices)
		}
	}
}
func TestFindGears(t *testing.T) {
	gear := '*'
	above := "......755."
	this := "...$.*...."
	below := ".664.598.."
	gears := FindGears(this, above, below, gear)
	if len(gears[0]) != 2 {
		t.Errorf("Got gears: %v", gears)
	}
}
