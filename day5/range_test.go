package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestRangeContains(t *testing.T) {
	r := Range{10, 5}

	assert.True(t, r.Contains(Range{10, 5}))
	assert.True(t, r.Contains(Range{11, 4}))
	assert.False(t, r.Contains(Range{1, 5}))
	assert.False(t, r.Contains(Range{9, 5}))
	assert.False(t, r.Contains(Range{9, 10}))
}

func TestRangeUnion(t *testing.T) {
	r1 :=  Range{10, 5}
	r2 :=  Range{12, 5}
	r1.Union(r2)
	assert.Equal(t, r1.Union(r2), r2.Union(r1))
	assert.Equal(t, Range{10, 7}, r1.Union(r2))

	empty := Range{0, 0}
	assert.Equal(t, empty, r1.Union(empty))
}

