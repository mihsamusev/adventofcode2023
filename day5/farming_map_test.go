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
