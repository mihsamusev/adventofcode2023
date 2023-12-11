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

func TestRangeHasValue(t *testing.T) {
	r1 := Range{1, 5}
	assert.True(t, r1.HasValue(1))
	assert.True(t, r1.HasValue(2))
	assert.True(t, r1.HasValue(5))
	assert.False(t, r1.HasValue(0))
	assert.False(t, r1.HasValue(6))
}

func TestRangeOffset(t *testing.T) {
	empty := Range{}
	some := Range{1, 3}
	expected := Range{3, 5}
	assert.Equal(t, empty, empty.Offset(5))
	assert.Equal(t, expected, some.Offset(2))
}

func TestRangeRemapStart(t *testing.T) {
	some := Range{79, 92}
	expected := Range{81, 94}
	assert.Equal(t, expected, some.RemapStart(81))
}

func TestRangeUnionOk(t *testing.T) {
	r1 := Range{1, 5}
	r2 := Range{3, 8}
	expected := Range{1, 8}
	assert.Equal(t, r1.Union(r2), r2.Union(r1))
	assert.Equal(t, expected, r1.Union(r2))
}

func TestRangeUnionBorder(t *testing.T) {
	r1 := Range{1, 5}
	r2 := Range{6, 8}
	expected := Range{1, 8}
	assert.Equal(t, r1.Union(r2), r2.Union(r1))
	assert.Equal(t, expected, r1.Union(r2))
}

func TestRangeUnionEmpty(t *testing.T) {
	r1 := Range{1, 5}
	r2 := Range{7, 8}
	empty := Range{}
	assert.Equal(t, r1.Union(r2), r2.Union(r1))
	assert.Equal(t, empty, r1.Union(r2))
}

func TestIntersection(t *testing.T) {
	r1 := Range{1, 4}
	r2 := Range{3, 8}
	expected := Range{3, 4}
	assert.Equal(t, r1.Intersection(r2), r2.Intersection(r1))
	assert.Equal(t, expected, r1.Intersection(r2))
}

func TestIntersectionContanined(t *testing.T) {
	r1 := Range{1, 8}
	r2 := Range{3, 6}
	assert.Equal(t, r1.Intersection(r2), r2.Intersection(r1))
	assert.Equal(t, r2, r1.Intersection(r2))
}

func TestIntersectionNonOverlapping(t *testing.T) {
	r1 := Range{1, 4}
	r2 := Range{5, 8}
	expected := Range{}
	assert.Equal(t, r1.Intersection(r2), r2.Intersection(r1))
	assert.Equal(t, expected, r1.Intersection(r2))
}

func TestDifferenceEmpty(t *testing.T) {
	r1 := Range{1, 4}
	r2 := Range{5, 8}
	expected := Range{}
	left, right := r1.Diff(r2)
	assert.Equal(t, expected, left)
	assert.Equal(t, expected, right)
}

func TestDifferenceLeft(t *testing.T) {
	r1 := Range{1, 4}
	r2 := Range{3, 8}
	left, right := r1.Diff(r2)
	assert.Equal(t, Range{1, 2}, left)
	assert.Equal(t, Range{}, right)
}

func TestDifferenceRight(t *testing.T) {
	r1 := Range{3, 8}
	r2 := Range{1, 4}
	left, right := r1.Diff(r2)
	assert.Equal(t, Range{}, left)
	assert.Equal(t, Range{5, 8}, right)
}

func TestDifferenceTwoSided(t *testing.T) {
	r1 := Range{1, 8}
	r2 := Range{4, 6}
	left, right := r1.Diff(r2)
	assert.Equal(t, Range{1, 3}, left)
	assert.Equal(t, Range{7, 8}, right)

	left, right = r2.Diff(r1)
	assert.Equal(t, Range{}, left)
	assert.Equal(t, Range{}, right)
}
