package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLookupContains(t *testing.T) {
	lookup := Lookup{50, Range{98, 99}}

	assert.False(t, lookup.Contains(97))
	assert.True(t, lookup.Contains(98))
	assert.True(t, lookup.Contains(99))
	assert.False(t, lookup.Contains(100))
}

func TestLookupConvert(t *testing.T) {
	lookup := Lookup{50, Range{98, 99}}
	dst := lookup.Convert(99)
	assert.Equal(t, 51, dst)
}

func TestFarmingMap(t *testing.T) {
	farmingMap := FarmingMap{
		[]Lookup{
			{50, Range{98, 99}},
			{52, Range{50, 97}},
		},
	}

	assert.Equal(t, 50, farmingMap.Convert(98))
	assert.Equal(t, 55, farmingMap.Convert(53))
	assert.Equal(t, 10, farmingMap.Convert(10))
}

func TestMapConvertRangeContainsEntirely(t *testing.T) {
	seedRange := Range{79, 92}
	farmingMap := FarmingMap{
		[]Lookup{
			{50, Range{98, 99}},
			{52, Range{50, 97}},
		},
	}

	ranges := farmingMap.ConvertRange(seedRange)
	expected := []Range{{81, 94}}
	assert.Equal(t, expected, ranges)
}

func TestMapConvertRangeOutsideEntirely(t *testing.T) {
	seedRange := Range{81, 94}
	farmingMap := FarmingMap{
		[]Lookup{
			{0, Range{15, 37}},
			{37, Range{52, 2}},
			{39, Range{0, 15}},
		},
	}

	ranges := farmingMap.ConvertRange(seedRange)
	expected := []Range{{81, 94}}
	assert.Equal(t, expected, ranges)
}

func TestMapConvertRangeSplitByLookups(t *testing.T) {
	seedRange := Range{74, 87}
	farmingMap := FarmingMap{
		[]Lookup{
			{45, Range{77, 99}},
			{81, Range{45, 63}},
			{68, Range{64, 76}},
		},
	}

	ranges := farmingMap.ConvertRange(seedRange)
	expected := []Range{{45, 55}, {78, 80}}
	assert.Equal(t, expected, ranges)
}
