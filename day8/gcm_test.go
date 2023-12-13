package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGCM(t *testing.T) {
	numbers := []int{12, 18}
	assert.Equal(t, 36, LeastCommonMul(numbers...))

	numbers = []int{12, 18, 24}
	assert.Equal(t, 72, LeastCommonMul(numbers...))
}
