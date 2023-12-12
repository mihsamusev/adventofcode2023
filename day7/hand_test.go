package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombos(t *testing.T) {
	h := Hand{[]Card("AAAAA"), 0}
	assert.Equal(t, Five, Combo(h))

	h = Hand{[]Card("AA8AA"), 0}
	assert.Equal(t, Four, Combo(h))

	h = Hand{[]Card("23332"), 0}
	assert.Equal(t, FullHouse, Combo(h))

	h = Hand{[]Card("TTT98"), 0}
	assert.Equal(t, Triplet, Combo(h))

	h = Hand{[]Card("23432"), 0}
	assert.Equal(t, TwoPair, Combo(h))

	h = Hand{[]Card("A23A4"), 0}
	assert.Equal(t, OnePair, Combo(h))

	h = Hand{[]Card("23456"), 0}
	assert.Equal(t, OneCard, Combo(h))
}

func TestComboWithJoker(t *testing.T) {
	h := Hand{[]Card("JKKK2"), 0}
	assert.Equal(t, Four, ComboWithJoker(h))

}