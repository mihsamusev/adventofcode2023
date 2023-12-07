package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestRangeContains(t *testing.T) {
	r := Range{2, 5}
	assert.True(t, r.In(Range{2, 5}))
	assert.True(t, r.In(Range{1, 6}))
	assert.False(t, r.In(Range{3, 5}))
    assert.False(t, r.In(Range{2, 4}))
	assert.False(t, r.In(Range{9, 10}))
}

func TestRangeUnion(t *testing.T) {
	r1 :=  Range{1, 5}
	r2 :=  Range{3, 8}
    expected := Range{1, 8} 
	assert.Equal(t, r1.Union(r2), r2.Union(r1))
	assert.Equal(t, expected, r1.Union(r2))

	empty := Range{0, 0}
	assert.Equal(t, empty, r1.Union(empty))
}

